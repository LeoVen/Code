use serde::Deserialize;
use std::collections::HashMap;

pub type RapidFireConfig = HashMap<String, HashMap<String, u32>>;
pub type RapidFire = Vec<Vec<u32>>;

#[derive(Debug, Deserialize)]
pub struct Template {
    pub id: String,
    pub name: String,

    pub armour: u32,
    pub shield: u32,
    pub weapon: u32,
}

#[derive(Debug, Clone, Copy)]
pub struct Unit {
    pub id: u32,
    pub armour: u32,
    pub shield: u32,
    pub weapon: u32,
}

#[derive(Debug, Clone, Copy)]
pub struct UnitState {
    pub id: u32,
    pub armour: u32,
    pub shield: u32,
    pub weapon: u32,

    pub init_armour: u32,
    // pub init_shield: u32,
    // pub init_weapon: u32,
}

impl From<Unit> for UnitState {
    fn from(value: Unit) -> Self {
        Self {
            id: value.id,
            armour: value.armour,
            shield: value.shield,
            weapon: value.weapon,
            init_armour: value.armour,
            // init_shield: value.shield,
            // init_weapon: value.weapon,
        }
    }
}
