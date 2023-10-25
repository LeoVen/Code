use std::{
    fmt::Display,
    process::{ExitCode, Termination},
};

#[derive(Debug)]
pub struct AkkadiaError {
    pub code: u8,
    pub message: String,
}

impl Display for AkkadiaError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "error {}: {}", self.code, self.message)
    }
}

impl Termination for AkkadiaError {
    fn report(self) -> std::process::ExitCode {
        eprintln!("{}", self);
        ExitCode::from(self.code)
    }
}

impl From<std::io::Error> for AkkadiaError {
    fn from(value: std::io::Error) -> Self {
        AkkadiaError {
            message: format!("std::io::Error {}: {}", value.kind(), value),
            code: 1,
        }
    }
}

impl From<&str> for AkkadiaError {
    fn from(value: &str) -> Self {
        AkkadiaError {
            code: 2,
            message: format!("string Error: {value}"),
        }
    }
}

impl From<String> for AkkadiaError {
    fn from(value: String) -> Self {
        AkkadiaError {
            code: 2,
            message: format!("string Error: {value}"),
        }
    }
}

impl From<toml::de::Error> for AkkadiaError {
    fn from(value: toml::de::Error) -> Self {
        AkkadiaError {
            code: 3,
            message: format!("toml deserialize error: {value}"),
        }
    }
}

impl From<crates_io_api::Error> for AkkadiaError {
    fn from(value: crates_io_api::Error) -> Self {
        AkkadiaError {
            code: 4,
            message: format!("crates-io error: {value}"),
        }
    }
}
