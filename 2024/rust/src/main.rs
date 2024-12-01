use days::day1;
use std::{env, process};

mod days;

fn main() {
    // let args = env::args().collect<Vec<String>>();
    let args: Vec<String> = env::args().collect();

    let day_number = match args.get(1) {
        Some(n) => n.parse().unwrap(),
        None => {
            eprintln!("no day number supplied");
            process::exit(1);
        }
    };

    match day_number {
        1 => day1::run(),
        _ => {
            eprintln!("day number {day_number} not recognised");
            process::exit(1);
        }
    }
}
