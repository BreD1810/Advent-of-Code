use crate::{get_unsigned_lines, print_answer};

pub fn run() {
    let lines = get_unsigned_lines(1, 1);
    let a1: u64 = lines.iter().map(|input| fuel_from_mass(input)).sum();
    let a2: u64 = lines
        .iter()
        .map(|input| recursive_fuel_from_mass(input))
        .sum();
    print_answer(1, &a1, &a2);
}

fn fuel_from_mass(mass: &u64) -> u64 {
    let r: u64 = mass / 3;
    if r >= 2 {
        r - 2
    } else {
        0u64
    }
}

fn recursive_fuel_from_mass(mass: &u64) -> u64 {
    let mut fuel_counts: Vec<u64> = Vec::new();
    let mut fuel = fuel_from_mass(mass);
    fuel_counts.push(fuel);
    loop {
        fuel = fuel_from_mass(&fuel);
        if fuel > 0 {
            fuel_counts.push(fuel);
        } else {
            break;
        }
    }
    fuel_counts.iter().sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    // Part 1
    #[test]
    fn mass_12() {
        assert_eq!(2, fuel_from_mass(&12));
    }

    #[test]
    fn mass_14() {
        assert_eq!(2, fuel_from_mass(&14));
    }

    #[test]
    fn mass_1969() {
        assert_eq!(654, fuel_from_mass(&1969));
    }

    #[test]
    fn mass_100756() {
        assert_eq!(33583, fuel_from_mass(&100756));
    }

    // Part 2
    #[test]
    fn recursive_mass_14() {
        assert_eq!(2, recursive_fuel_from_mass(&14));
    }

    #[test]
    fn recursive_mass_1969() {
        assert_eq!(966, recursive_fuel_from_mass(&1969));
    }

    #[test]
    fn recursive_mass_100756() {
        assert_eq!(50346, recursive_fuel_from_mass(&100756));
    }
}
