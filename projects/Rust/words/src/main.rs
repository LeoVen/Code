use std::{collections::BTreeMap, fmt::Write};

static MIN_LETTER_SIZE: usize = 2;
static LETTERS: &[char; 27] = &[
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
    't', 'u', 'v', 'w', 'x', 'y', 'z', '-',
];

#[derive(Default, Debug)]
struct Stats {
    pub total_appearances: usize,
    pub first_letter: usize,
    pub last_letter: usize,
    pub letters_before: BTreeMap<char, usize>,
    pub letters_after: BTreeMap<char, usize>,
}

fn new_empty_map() -> BTreeMap<char, usize> {
    let mut map = BTreeMap::new();

    for letter in LETTERS {
        map.insert(*letter, 0);
    }

    map
}

fn main() {
    let result = std::fs::read_to_string("./names.txt").expect("failed to read words.txt");

    let words = result.split('\n').collect::<Vec<&str>>();

    let mut map: BTreeMap<char, Stats> = BTreeMap::new();

    for letter in LETTERS {
        map.insert(*letter, Stats{
            total_appearances: 0,
            first_letter: 0,
            last_letter: 0,
            letters_before: new_empty_map(),
            letters_after: new_empty_map(),
        });
    }

    let mut longest = String::new();

    for word in words {
        if word.len() < MIN_LETTER_SIZE {
            continue;
        }

        if word.len() > longest.len() {
            longest = word.to_string();
        }

        let word = word.to_lowercase();

        let mut first_letter = true;
        let mut letter_before = ' ';

        for char in word.chars() {
            if let Some(curr_stat) = map.get_mut(&char) {
                curr_stat.total_appearances += 1;

                if first_letter {
                    curr_stat.first_letter += 1;
                    first_letter = false
                } else {
                    if let Some(num) = curr_stat.letters_before.get_mut(&letter_before) {
                        *num += 1;
                    }

                    if let Some(stat_before) = map.get_mut(&letter_before) {
                        if let Some (num) = stat_before.letters_after.get_mut(&char) {
                            *num += 1;
                        }
                    }
                }

                letter_before = char;
            }
        }

        if let Some(stat_last) = map.get_mut(&letter_before) {
            stat_last.last_letter += 1;
        }
    }

    let mut content = String::new();

    _ = content.write_str("char,total,first,last\n");

    for (letter, stats) in map {
        content += &format!("{},{},{},{}\n", letter, stats.total_appearances, stats.first_letter, stats.last_letter);
        println!("Letter {} has\n{:?}\n\n\n", letter, stats);
    }

    _ = std::fs::write("./output.csv", content);

    println!("Longest word is {}", &longest);

}
