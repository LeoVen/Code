use std::process::{ExitCode, Termination};

#[derive(Debug)]
pub struct AkkadiaError {
    pub code: u8,
    pub message: String,
}

impl Termination for AkkadiaError {
    fn report(self) -> std::process::ExitCode {
        eprintln!("Akkadia Error: {}", self.message);
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
            message: format!("String Error: {value}"),
        }
    }
}

impl From<toml::de::Error> for AkkadiaError {
    fn from(value: toml::de::Error) -> Self {
        AkkadiaError {
            code: 3,
            message: format!("Toml deserialize error: {value}"),
        }
    }
}
