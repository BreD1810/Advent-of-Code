use aoc::{get_lines, print_answer};

pub fn run() {
    let input = get_lines(1);
    let answer1 = part1(input.clone());
    let answer2 = part2(input);
    print_answer(1, &answer1, &answer2);
}

fn part1(input: Vec<String>) -> usize {
    let mut left = input
        .iter()
        .filter_map(|l| l.split_whitespace().next().unwrap().parse::<usize>().ok())
        .collect::<Vec<usize>>();
    left.sort();

    let mut right = input
        .iter()
        .filter_map(|l| {
            l.split_whitespace()
                .skip(1)
                .next()
                .unwrap()
                .parse::<usize>()
                .ok()
        })
        .collect::<Vec<usize>>();
    right.sort();

    left.iter()
        .zip(right.iter())
        .map(|(l, r)| usize::abs_diff(*l, *r))
        .sum()
}

fn part2(input: Vec<String>) -> usize {
    let left = input
        .iter()
        .filter_map(|l| l.split_whitespace().next().unwrap().parse::<usize>().ok())
        .collect::<Vec<usize>>();

    let right = input
        .iter()
        .filter_map(|l| {
            l.split_whitespace()
                .skip(1)
                .next()
                .unwrap()
                .parse::<usize>()
                .ok()
        })
        .collect::<Vec<usize>>();

    left.iter()
        .map(|l| l * right.iter().filter(|r| *r == l).count())
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn day1_example1() {
        let input = vec![
            String::from("3   4"),
            String::from("4   3"),
            String::from("2   5"),
            String::from("1   3"),
            String::from("3   9"),
            String::from("3   3"),
        ];
        let res = part1(input);
        assert_eq!(11, res);
    }

    #[test]
    fn day1_actual1() {
        let input = get_lines(1);
        let res = part1(input);
        assert_eq!(1879048, res);
    }

    #[test]
    fn day1_example2() {
        let input = vec![
            String::from("3   4"),
            String::from("4   3"),
            String::from("2   5"),
            String::from("1   3"),
            String::from("3   9"),
            String::from("3   3"),
        ];
        let res = part2(input);
        assert_eq!(31, res);
    }

    #[test]
    fn day1_actual2() {
        let input = get_lines(1);
        let res = part2(input);
        assert_eq!(21024792, res);
    }
}
