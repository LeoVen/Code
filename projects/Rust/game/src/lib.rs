pub mod battle;
pub mod error;
pub mod factory;
pub mod units;

use std::{
    collections::{HashMap, HashSet},
    fs,
};

use error::LibError;
use log::{info, trace};
use units::{RapidFire, RapidFireConfig, Template, Unit};

pub fn parse_ron_file<'a, T>(path: &str) -> Result<T, LibError>
where
    T: serde::de::DeserializeOwned,
{
    let data = fs::read_to_string(path)?;
    let parsed = ron::from_str(&data)?;

    trace!("loaded {}", path);

    Ok(parsed)
}

pub fn create_instances(
    templates: &[Template],
    rapidfire_config: &RapidFireConfig,
) -> (HashMap<String, u32>, Vec<Unit>, RapidFire) {
    let unique_ids = templates.iter().fold(HashSet::new(), |mut set, val| {
        set.insert(val.id.clone());
        set
    });

    let dim = unique_ids.len();

    info!("found {} unique IDs", dim);
    trace!("IDs found: {:?}", unique_ids);

    let mut counter = 0;
    let mapping = unique_ids
        .into_iter()
        .map(|id| {
            let result = (id, counter);
            counter += 1;
            result
        })
        .collect::<HashMap<String, u32>>();

    trace!("mapping {:?}", mapping);

    let instances = templates
        .iter()
        .map(|t| Unit {
            armour: t.armour,
            shield: t.shield,
            weapon: t.weapon,
            id: mapping[&t.id],
        })
        .collect();

    // TODO optimize this to a continuous array via chunks and &mut references
    let mut rapid_fire = vec![vec![0; dim]; dim];
    for (id, rf) in rapidfire_config.iter() {
        let idx = mapping[&*id] as usize;
        for (target_id, val) in rf.iter() {
            let target = mapping[&*target_id] as usize;
            rapid_fire[idx][target] = *val;
        }
    }

    (mapping, instances, rapid_fire)
}
