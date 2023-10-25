use std::collections::HashMap;

use game::{
    battle::Battle,
    error::LibError,
    factory::Factory,
    parse_ron_file,
    units::{RapidFireConfig, Template},
};
use log::{info, trace};

fn main() -> Result<(), LibError> {
    dotenv::dotenv().ok();
    env_logger::init();

    info!("starting");

    let templates = parse_ron_file::<Vec<Template>>("templates.ron")?;
    let rapidfire = parse_ron_file::<RapidFireConfig>("rapidfire.ron")?;

    trace!("LOADED\n{:?}\n{:?}", &templates, rapidfire);

    let (mapping, units, rapid_fire) = game::create_instances(&templates, &rapidfire);

    trace!("UNITS\n{:?}", units);
    trace!("RAPIDFIRE\n{:?}", rapid_fire);

    let mut attackers = vec![];
    let mut defenders = vec![];
    let factory = Factory::new(
        mapping,
        templates.into_iter().map(|t| (t.id.clone(), t)).collect(),
    );

    let quantities =
        parse_ron_file::<(HashMap<String, usize>, HashMap<String, usize>)>("battlefield.ron")?;

    for (k, v) in &quantities.0 {
        attackers.extend(factory.generate_state(&k, *v));
    }
    for (k, v) in &quantities.1 {
        defenders.extend(factory.generate_state(&k, *v));
    }

    println!(
        "ATTACKERS\n{:?}\n\nDEFENDERS\n{:?}\n",
        &quantities.0, &quantities.1
    );

    println!("BATTLE START");

    let mut battle = Battle::new(rapid_fire, attackers, defenders);

    for _ in 0..11 {
        battle.round();
        battle.filter_out();

        if battle.ended() {
            break;
        }
    }

    battle.summary();

    Ok(())
}
