pub mod days;

use std::fmt::Display;
use std::fs::read_to_string;

fn print_answer(day_num: usize, val1: &dyn Display, val2: &dyn Display) {
    println!("----- DAY {} -----", day_num);
    println!("Part 1: {}", val1);
    println!("Part 2: {}", val2);
}

fn get_unsigned_lines(day_no: u32, task_no: u32) -> Vec<u64> {
    let filename = format!("inputs/{}-{}.txt", day_no, task_no);
    read_to_string(filename)
        .expect("Unable to open input")
        .split("\n")
        .filter_map(|l| l.parse::<u64>().ok())
        .collect()
}

fn get_comma_int_list(day_no: u32, task_no: u32) -> Vec<usize> {
    let filename = format!("inputs/{}-{}.txt", day_no, task_no);
    read_to_string(filename)
        .expect("Unable to open input")
        .split(",")
        .filter_map(|l| l.parse::<usize>().ok())
        .collect()
}

fn get_lines(day_no: u32, task_no: u32) -> Vec<String> {
    let filename = format!("inputs/{}-{}.txt", day_no, task_no);
    read_to_string(filename)
        .expect("Unable to open input")
        .split("\n")
        .map(str::to_string)
        .collect()
}
