use std::fmt::Display;
use std::fs::read_to_string;

pub fn print_answer(day_num: usize, val1: &dyn Display, val2: &dyn Display) {
    println!("----- DAY {} -----", day_num);
    println!("Part 1: {}", val1);
    println!("Part 2: {}", val2);
}

// get_unsigned_int_lines converts each line to a usize
pub fn get_unsigned_int_lines(day_no: u32) -> Vec<usize> {
    let filename = format!("inputs/{}.txt", day_no);
    read_to_string(filename)
        .expect("Unable to open input")
        .lines()
        .filter_map(|l| l.parse::<usize>().ok())
        .collect()
}

// get_comma_int_list converts a single line to a list, splitting on ","
pub fn get_comma_int_list(day_no: u32) -> Vec<usize> {
    let filename = format!("inputs/{}.txt", day_no);
    read_to_string(filename)
        .expect("Unable to open input")
        .split(",")
        .filter_map(|l| l.parse::<usize>().ok())
        .collect()
}

// get_lines converts lines to a Vec<String>
pub fn get_lines(day_no: u32) -> Vec<String> {
    let filename = format!("inputs/day{}.txt", day_no);
    read_to_string(filename)
        .expect("Unable to open input")
        .lines()
        .map(str::to_string)
        .collect()
}
