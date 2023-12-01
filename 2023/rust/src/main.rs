use std::{env, process};

mod days;

fn main() {
    let args: Vec<String> = env::args().collect();

    let day_number = match args.get(1) {
        Some(n) => n.parse().unwrap_or_default(),
        None => {
            eprintln!("no day number supplied");
            process::exit(1);
        }
    };

    match day_number {
        1 => days::day1::run(),
        0 => {
            let arg_val = args.get(1).unwrap();
            eprintln!("day number '{arg_val}' not recognised");
            process::exit(1)
        }
        _ => {
            eprintln!("day number {day_number} not recognised");
            process::exit(1);
        }
    }
}
