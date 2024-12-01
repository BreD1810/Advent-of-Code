use std::{fmt::Display, fs::read_to_string};

pub fn print_answer(day_no: u32, val1: &dyn Display, val2: &dyn Display) {
    println!("----- DAY {} -----", day_no);
    println!("Part 1: {}", val1);
    println!("Part 2: {}", val2);
}

// get_lines gives a string for every line in the input file
pub fn get_lines(day_no: u32) -> Vec<String> {
    let filename = format!("inputs/day{day_no}.txt");
    read_to_string(filename)
        .expect("Unable to open input")
        .lines()
        .map(str::to_string)
        .collect()
}
