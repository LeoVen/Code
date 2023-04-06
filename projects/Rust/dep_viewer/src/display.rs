use comfy_table::{Cell, CellAlignment, Table};

use crate::fetch::DepInfo;

pub fn display_table(data: Vec<DepInfo>) {
    let mut table = Table::new();

    table.set_header(vec![
        "Type",
        "Name",
        "Version",
        "Max Version",
        "Downloads",
        "Last Update",
    ]);

    for dep in data {
        table.add_row(vec![
            Cell::new(dep.set.dep_type),
            Cell::new(dep.set.name),
            Cell::new(dep.set.version).set_alignment(CellAlignment::Center),
            Cell::new(dep.max_version).set_alignment(CellAlignment::Center),
            Cell::new(dep.downloads).set_alignment(CellAlignment::Right),
            Cell::new(dep.last_update),
        ]);
    }

    println!("{table}");
}
