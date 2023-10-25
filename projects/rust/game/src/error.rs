use std::{
    fmt::Display,
    process::{ExitCode, Termination},
};

use ron::error::SpannedError;

#[derive(Debug)]
pub enum LibError {
    AlgorithmError(String),
    ConfigError(String),
}

impl Display for LibError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "LibError {:?}", self)
    }
}

impl Termination for LibError {
    fn report(self) -> std::process::ExitCode {
        eprintln!("{}", self);
        match self {
            LibError::AlgorithmError(_) => ExitCode::from(1),
            LibError::ConfigError(_) => ExitCode::from(2),
        }
    }
}

impl From<std::io::Error> for LibError {
    fn from(value: std::io::Error) -> Self {
        Self::ConfigError(format!("[config] io error: {}", value))
    }
}

impl From<SpannedError> for LibError {
    fn from(value: SpannedError) -> Self {
        Self::ConfigError(format!("[config] parse error: {}", value))
    }
}
