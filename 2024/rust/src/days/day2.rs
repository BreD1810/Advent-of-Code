use aoc::{get_lines, print_answer};

pub fn run() {
    let input = get_lines(1);
    let answer1 = part1(input.clone());
    let answer2 = part2(input);
    print_answer(2, &answer1, &answer2);
}

fn part1(input: Vec<String>) -> usize {
    input
        .iter()
        .map(|l| {
            l.split_whitespace()
                .filter_map(|v| v.parse::<i32>().ok())
                .collect::<Vec<i32>>()
        })
        .filter(|v| is_safe(v))
        .count()
}

fn is_safe(report: &[i32]) -> bool {
    let increasing = report[0] < report[1];
    report.windows(2).all(|w| {
        let l = w[0];
        let r = w[1];

        let diff = if increasing { r - l } else { l - r };
        if diff > 3 || diff < 1 {
            return false;
        }
        true
    })
}

fn part2(input: Vec<String>) -> usize {
    input
        .iter()
        .map(|l| {
            l.split_whitespace()
                .filter_map(|v| v.parse::<i32>().ok())
                .collect::<Vec<i32>>()
        })
        .filter(|report| {
            if is_safe(report) {
                return true;
            }

            report.iter().enumerate().any(|(n, _)| {
                let mut r = Vec::clone(report)[0..n].to_vec();
                let mut r2 = Vec::clone(report)[n + 1..].to_vec();
                r.append(&mut r2);
                is_safe(&r)
            })
        })
        .count()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn day2_example1() {
        let input = vec![
            String::from("7 6 4 2 1"),
            String::from("1 2 7 8 9"),
            String::from("9 7 6 2 1"),
            String::from("1 3 2 4 5"),
            String::from("8 6 4 4 1"),
            String::from("1 3 6 7 9"),
        ];
        let res = part1(input);
        assert_eq!(2, res);
    }

    #[test]
    fn day2_actual1() {
        let input = get_lines(2);
        let res = part1(input);
        assert_eq!(624, res);
    }

    #[test]
    fn day2_example2() {
        let input = vec![
            String::from("7 6 4 2 1"),
            String::from("1 2 7 8 9"),
            String::from("9 7 6 2 1"),
            String::from("1 3 2 4 5"),
            String::from("8 6 4 4 1"),
            String::from("1 3 6 7 9"),
        ];
        let res = part2(input);
        assert_eq!(4, res);
    }

    #[test]
    fn day2_actual2() {
        let input = get_lines(2);
        let res = part2(input);
        assert_eq!(658, res);
    }
}
