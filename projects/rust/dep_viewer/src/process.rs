use std::{fs, path::PathBuf};

use log::info;

use crate::custom_error::AkkadiaError;

#[derive(Debug)]
pub struct ResultSet {
    pub name: String,
    pub version: String,
    pub dep_type: &'static str,
}

pub fn process(path: String) -> Result<Vec<ResultSet>, AkkadiaError> {
    let path = fs::canonicalize(PathBuf::from(path))?;
    let meta = fs::symlink_metadata(&path)?;

    let file_path = if meta.is_dir() {
        solve_file_in_dir(path)?
    } else {
        path
    };

    info!(
        "solved path to file {}",
        file_path.to_str().unwrap_or("<non utf8 path>")
    );

    let root = fs::read_to_string(file_path)?.parse::<toml::Value>()?;

    get_deps(root)
}

fn solve_file_in_dir(path: PathBuf) -> Result<PathBuf, AkkadiaError> {
    for entry in fs::read_dir(&path)? {
        let entry = entry?;

        let meta = entry.metadata()?;

        if !meta.is_file() {
            continue;
        }

        let file_name = entry
            .file_name()
            .into_string()
            .map_err(|_| "file name is non utf8")?;

        if file_name == "Cargo.toml" {
            return Ok(entry.path());
        }
    }

    let msg = format!(
        "could not find toml file in folder {}",
        &path.to_str().unwrap_or("<non utf8 path>")
    );

    Err(msg.as_str().into())
}

fn get_deps(root: toml::Value) -> Result<Vec<ResultSet>, AkkadiaError> {
    static DEPS_TABLES: [&str; 3] = ["dependencies", "dev-dependencies", "build-dependencies"];

    if !root.is_table() {
        return Err(
            format!("toml root is not table (it is {})", root.type_str())
                .as_str()
                .into(),
        );
    }

    let table = root.as_table().unwrap();
    let mut result = vec![];

    for deps in DEPS_TABLES {
        if let Some(dep_table) = table.get(deps) {
            if !dep_table.is_table() {
                return Err(format!(
                    "dependency {} is not in table format (it is {})",
                    deps,
                    dep_table.type_str()
                )
                .as_str()
                .into());
            }

            let dep_table = dep_table.as_table().unwrap();

            for (k, v) in dep_table {
                match v {
                    // foo-bar = "0.1.0"
                    toml::Value::String(version) => result.push(ResultSet {
                        name: k.clone(),
                        version: version.clone(),
                        dep_type: deps,
                    }),
                    // foo-bar = { version = "0.1.0" }
                    toml::Value::Table(table) => {
                        if let Some(version) = table.get("version") {
                            if version.is_str() {
                                result.push(ResultSet {
                                    name: k.clone(),
                                    version: version.as_str().unwrap().to_string(),
                                    dep_type: deps,
                                })
                            } else {
                                return Err(format!(
                                    "unexpected version format for {} as type {}",
                                    &k,
                                    version.type_str()
                                )
                                .as_str()
                                .into());
                            }
                        } else {
                            return Err(format!("could not find version for {}", &k)
                                .as_str()
                                .into());
                        }
                    }
                    _ => {
                        return Err(format!(
                            "unexpected version format for {} as type {}",
                            &k,
                            v.type_str()
                        )
                        .as_str()
                        .into())
                    }
                }
            }
        }
    }

    Ok(result)
}
