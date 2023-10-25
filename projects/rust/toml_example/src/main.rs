use std::{fs, path::PathBuf};

use toml::Value;

fn show_value(value: &Value, indent: usize) {
    let pfx = "  ".repeat(indent);
    print!("{}", pfx);
    match value {
        Value::String(string) => {
            println!("[string] {}", string);
        }
        Value::Integer(integer) => {
            println!("[integer] {}", integer);
        }
        Value::Float(float) => {
            println!("[float] {}", float);
        }
        Value::Boolean(boolean) => {
            println!("[boolean] {}", boolean);
        }
        Value::Datetime(datetime) => {
            println!("[datetime] {}", datetime);
        }
        Value::Array(array) => {
            println!("[array]");
            for v in array.iter() {
                show_value(v, indent + 1);
            }
        }
        Value::Table(table) => {
            println!("[table]");
            for (k, v) in table.iter() {
                println!("{} > {}", pfx, k);
                show_value(v, indent + 1);
            }
        }
    }
}

fn main() {
    let mut args = std::env::args().skip(1);

    let path = args.next().unwrap();
    let path = fs::canonicalize(PathBuf::from(path)).unwrap();
    let content = fs::read_to_string(path).unwrap();

    let value = content.parse::<Value>().unwrap();
    show_value(&value, 0);
}
