use std::collections::HashMap;

use crate::units::{Template, Unit, UnitState};

pub struct Factory {
    mapping: HashMap<String, u32>,
    templates: HashMap<String, Template>,
}

impl Factory {
    pub fn new(mapping: HashMap<String, u32>, templates: HashMap<String, Template>) -> Self {
        Self { mapping, templates }
    }

    pub fn generate(&self, id: &str, amount: usize) -> Vec<Unit> {
        let t = &self.templates[id];
        vec![
            Unit {
                id: self.mapping[id],
                armour: t.armour,
                shield: t.shield,
                weapon: t.weapon,
            };
            amount
        ]
    }

    pub fn generate_state(&self, id: &str, amount: usize) -> Vec<UnitState> {
        let t = &self.templates[id];
        vec![
            Unit {
                id: self.mapping[id],
                armour: t.armour,
                shield: t.shield,
                weapon: t.weapon,
            }
            .into();
            amount
        ]
    }
}
