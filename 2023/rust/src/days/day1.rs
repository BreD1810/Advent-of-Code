use aoc_2023::{get_lines, print_answer};

pub fn run() {
    let input = get_lines(1);
    let answer_1 = part_1(input.clone());
    let answer_2 = part_2(input.clone());
    print_answer(1, &answer_1, &answer_2);
}

fn part_1(input: Vec<String>) -> u32 {
    input
        .iter()
        .map(|line| {
            line.chars()
                .filter_map(|c| c.to_digit(10))
                .collect::<Vec<u32>>()
        })
        .map(|vec| 10 * vec.first().unwrap() + vec.last().unwrap())
        .sum()
}

fn part_2(input: Vec<String>) -> u32 {
    input
        .iter()
        .map(|line| {
            line.replace("one", "one1one")
                .replace("two", "two2two")
                .replace("three", "three3three")
                .replace("four", "four4four")
                .replace("five", "five5five")
                .replace("six", "six6six")
                .replace("seven", "seven7seven")
                .replace("eight", "eight8eight")
                .replace("nine", "nine9nine")
        })
        .map(|line| {
            line.chars()
                .filter_map(|c| c.to_digit(10))
                .collect::<Vec<u32>>()
        })
        .map(|vec| 10 * vec.first().unwrap() + vec.last().unwrap())
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn example_1() {
        let input = vec![
            String::from("1abc2"),
            String::from("pqr3stu8vwx"),
            String::from("a1b2c3d4e5f"),
            String::from("treb7uchet"),
        ];
        let res = part_1(input);
        assert_eq!(142, res);
    }

    #[test]
    fn actual_1() {
        let input = get_lines(1);
        let res = part_1(input);
        assert_eq!(55108, res);
    }

    #[test]
    fn example_2() {
        let input = vec![
            String::from("two1nine"),
            String::from("eightwothree"),
            String::from("abcone2threexyz"),
            String::from("xtwone3four"),
            String::from("4nineeightseven2"),
            String::from("zoneight234"),
            String::from("7pqrstsixteen"),
        ];
        let res = part_2(input);
        assert_eq!(281, res);
    }

    #[test]
    fn actual_2() {
        let input = get_lines(1);
        let res = part_2(input);
        assert_eq!(56324, res);
    }
}
