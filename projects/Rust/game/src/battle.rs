use std::collections::HashMap;

use crate::units::{RapidFire, UnitState};

pub struct Battle {
    pub attackers: BattleSide,
    pub defenders: BattleSide,

    pub attacker_stats: Vec<Stat>,
    pub defender_stats: Vec<Stat>,
}

impl Battle {
    pub fn new(
        rapid_fire: RapidFire,
        attackers: Vec<UnitState>,
        defenders: Vec<UnitState>,
    ) -> Self {
        Self {
            attackers: BattleSide::new(attackers, rapid_fire.clone()),
            defenders: BattleSide::new(defenders, rapid_fire),
            attacker_stats: vec![],
            defender_stats: vec![],
        }
    }

    pub fn ended(&self) -> bool {
        self.attackers.is_empty() || self.defenders.is_empty()
    }

    pub fn filter_out(&mut self) {
        self.attackers.filter_out();
        self.defenders.filter_out();
    }

    pub fn count_units(&self) {
        // TODO
    }

    pub fn summary(&self) {
        // TODO
        // output it to a file?
        println!("SUMMARY");

        let mut attacker: HashMap<u32, usize> = HashMap::new();
        let mut defender: HashMap<u32, usize> = HashMap::new();

        self.attackers.units.iter().for_each(|a| {
            let _ = attacker.entry(a.id).and_modify(|e| *e += 1).or_insert(1);
        });
        self.defenders.units.iter().for_each(|d| {
            let _ = defender.entry(d.id).and_modify(|e| *e += 1).or_insert(1);
        });

        println!("{attacker:?}\n{defender:?}");

        let defender_stats = self
            .defender_stats
            .iter()
            .fold(Stat::default(), |acc, x| acc.merge(x));
        let attacker_stats = self
            .attacker_stats
            .iter()
            .fold(Stat::default(), |acc, x| acc.merge(x));

        println!(
            "ATTACKERS\n{:?}\n\nDEFENDERS\n{:?}\n",
            attacker_stats, defender_stats
        )
    }

    pub fn round(&mut self) {
        self.attacker_stats
            .push(self.attackers.round(&mut self.defenders));
        self.defender_stats
            .push(self.defenders.round(&mut self.attackers));
    }
}

#[derive(Default, Debug)]
pub struct Stat {
    pub absorbtion: u64,
    pub explosion: u64,
    pub rapid_fire: u64,

    // [shield, hull]
    pub damage_dealt: [u64; 2],
}

impl Stat {
    pub fn merge(mut self, other: &Stat) -> Self {
        self.absorbtion += other.absorbtion;
        self.explosion += other.explosion;
        self.rapid_fire += other.rapid_fire;
        self.damage_dealt[0] += other.damage_dealt[0];
        self.damage_dealt[1] += other.damage_dealt[1];

        self
    }
}

pub struct BattleSide {
    pub rapid_fire: RapidFire,
    pub units: Vec<UnitState>,
    pub rng: fastrand::Rng,
}

impl BattleSide {
    pub fn new(units: Vec<UnitState>, rapid_fire: RapidFire) -> Self {
        Self {
            rapid_fire,
            units,
            rng: fastrand::Rng::new(),
        }
    }

    pub fn is_empty(&self) -> bool {
        self.units.is_empty()
    }

    pub fn filter_out(&mut self) {
        self.units.retain(|unit| unit.armour > 0);
    }

    pub fn get_random(&mut self) -> &mut UnitState {
        let idx = self.rng.usize(0..self.units.len());

        &mut self.units[idx]
    }

    // Steps
    // 1. Check if the attack bounces
    // 2. Do damage
    // 3. Check if defender explodes
    // 4. Check for rapid fire
    pub fn attack(&self, stat: &mut Stat, attacker: &UnitState, defender: &mut UnitState) -> bool {
        // 1.
        if (attacker.weapon as f64) < defender.shield as f64 * 0.01 {
            stat.absorbtion += 1;
            return false;
        }

        // 2.
        if attacker.weapon > defender.shield {
            let mut armour_damage = attacker.weapon - defender.shield;

            if armour_damage > defender.armour {
                armour_damage = defender.armour;
            }

            stat.damage_dealt[0] += defender.shield as u64;
            stat.damage_dealt[1] += armour_damage as u64;

            defender.armour -= armour_damage;
            defender.shield = 0;
        } else {
            defender.shield -= attacker.weapon;

            stat.damage_dealt[0] += attacker.weapon as u64;
        }

        // 3.
        let pct_armour = defender.armour as f64 / defender.init_armour as f64;
        if pct_armour < 0.7 {
            let rng = self.rng.f64();

            if rng < pct_armour {
                stat.explosion += 1;
                defender.armour = 0;
            }
        }

        // 4.
        let rf = self.rapid_fire[attacker.id as usize][defender.id as usize];
        if rf > 0 {
            let rf = rf as f64;
            let chance = (rf - 1.0) / rf;
            let rng = self.rng.f64();

            if rng < chance {
                stat.rapid_fire += 1;
                return true;
            }
        }

        false
    }

    // For every unit, choose a random target and attack it until rapid fire fails
    pub fn round(&self, targets: &mut BattleSide) -> Stat {
        let mut round_stat = Stat::default();

        for attacker in self.units.iter() {
            let mut target = targets.get_random();

            while self.attack(&mut round_stat, attacker, target) {
                target = targets.get_random();
            }
        }

        round_stat
    }
}
