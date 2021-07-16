use crate::print_answer;

pub fn run() {
    let a1 = part_1(359282, 820401);
    let a2 = part_2(359282, 820401);
    print_answer(4, &a1, &a2);
}

fn part_1(start: usize, end: usize) -> usize {
    (start..end + 1)
        .filter(|i| check_valid_password(&i))
        .count()
}

fn part_2(start: usize, end: usize) -> usize {
    (start..end + 1)
        .filter(|i| check_valid_password_2(&i))
        .count()
}

fn check_valid_password(password: &usize) -> bool {
    check_adjacent_equal(&password) && check_ascending_digits(&password)
}

fn check_valid_password_2(password: &usize) -> bool {
    check_only_2_adjacent(&password) && check_ascending_digits(&password)
}

fn check_adjacent_equal(val: &usize) -> bool {
    let digits: Vec<_> = val
        .to_string()
        .chars()
        .map(|d| d.to_digit(10).unwrap())
        .collect();

    let equal_adjacent: Vec<_> = digits.windows(2).filter(|l| l[0] == l[1]).collect();

    equal_adjacent.len() > 0
}

fn check_only_2_adjacent(val: &usize) -> bool {
    let digits: Vec<_> = val
        .to_string()
        .chars()
        .map(|d| d.to_digit(10).unwrap())
        .collect();

    let adjacent_matches: Vec<_> = digits.windows(2).map(|l| l[0] == l[1]).collect();

    adjacent_matches
        .split(|r| !r)
        .into_iter()
        .filter(|l| l.len() == 1)
        .count()
        > 0
    // .fold(true, |a, l| l.len() % 2 == 1 && a)
}

fn check_ascending_digits(val: &usize) -> bool {
    let digits: Vec<_> = val
        .to_string()
        .chars()
        .map(|d| d.to_digit(10).unwrap())
        .collect();

    let adjacent_ascending: Vec<_> = digits.windows(2).filter(|l| l[0] <= l[1]).collect();

    adjacent_ascending.len() == digits.len() - 1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part_1_example_1() {
        let value = 111111;
        assert!(check_adjacent_equal(&value));
        assert!(check_ascending_digits(&value));
        assert!(check_valid_password(&value));
    }

    #[test]
    fn part_1_example_2() {
        let value = 223450;
        assert!(check_adjacent_equal(&value));
        assert!(!check_ascending_digits(&value));
        assert!(!check_valid_password(&value))
    }

    #[test]
    fn part_1_example_3() {
        let value = 123789;
        assert!(!check_adjacent_equal(&value));
        assert!(check_ascending_digits(&value));
        assert!(!check_valid_password(&value));
    }

    #[test]
    fn part_2_example_1() {
        let value = 112233;
        assert!(check_ascending_digits(&value));
        assert!(check_only_2_adjacent(&value));
        assert!(check_valid_password_2(&value));
    }

    #[test]
    fn part_2_example_2() {
        let value = 123444;
        assert!(check_ascending_digits(&value));
        assert!(!check_only_2_adjacent(&value));
        assert!(!check_valid_password_2(&value))
    }

    #[test]
    fn part_2_example_3() {
        let value = 111122;
        assert!(check_ascending_digits(&value));
        assert!(check_only_2_adjacent(&value));
        assert!(check_valid_password_2(&value));
    }
}
