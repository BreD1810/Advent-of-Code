use std::env;
use std::process::exit;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() == 1 {
        eprintln!("Please specify a day number.");
        exit(1);
    }
    let day: u8 = args[1].parse::<u8>().expect("Invalid day number");

    match day {
        1 => aoc_2019::days::day1::run(),
        2 => aoc_2019::days::day2::run(),
        3 => aoc_2019::days::day3::run(),
        _ => eprintln!("Provided day not available."),
    };
}
