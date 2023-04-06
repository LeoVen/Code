use std::vec;

use crates_io_api::SyncClient;
use log::{info, warn};

use crate::{custom_error::AkkadiaError, process::ResultSet};

#[derive(Debug)]
pub struct DepInfo {
    pub set: ResultSet,
    pub downloads: u64,
    pub max_version: String,
    pub last_update: String,
}

pub fn fetch(set: Vec<ResultSet>) -> Result<Vec<DepInfo>, AkkadiaError> {
    let mut result = vec![];

    let client = SyncClient::new(
        "Akkadia (github.com/LeoVen)",
        std::time::Duration::from_millis(1000),
    )
    .map_err(|err| AkkadiaError::from(err.to_string()))?;

    for data in set {
        let name = &data.name;

        let response = client.get_crate(name)?;

        if response.crate_data.name != data.name {
            warn!(
                "crate names don't match: {} - {}",
                response.crate_data.name, data.name
            );
        }

        info!("fetched {}", response.crate_data.name);

        result.push(DepInfo {
            set: data,
            downloads: response.crate_data.downloads,
            max_version: response.crate_data.max_version,
            last_update: response.crate_data.updated_at.to_string(),
        })
    }

    Ok(result)
}
