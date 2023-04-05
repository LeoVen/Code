use std::vec;

use crates_io_api::SyncClient;

use crate::{custom_error::AkkadiaError, process::ResultSet};

pub struct DepInfo {
    pub set: ResultSet,
}

pub fn fetch(set: Vec<ResultSet>) -> Result<Vec<DepInfo>, AkkadiaError> {
    let result = vec![];

    let client = SyncClient::new(
        "Akkadia (github.com/LeoVen)",
        std::time::Duration::from_millis(1000),
    )
    .map_err(|err| AkkadiaError::from(&*err.to_string()))?;

    Ok(result)
}
