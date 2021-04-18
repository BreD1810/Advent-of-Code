use crate::{print_answer, get_lines};
use std::str::FromStr;
use std::collections::HashSet;
use std::iter::FromIterator;
use std::hash::{Hash, Hasher};

enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl FromStr for Direction {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "U" => Ok(Self::Up),
            "R" => Ok(Self::Right),
            "D" => Ok(Self::Down),
            "L" => Ok(Self::Left),
            _ => Err(()),
        }
    }
}

impl Direction {
    fn changes(&self) ->(isize, isize) {
        match self {
            Direction::Up => (0, 1),
            Direction::Right => (1, 0),
            Direction::Down => (0, -1),
            Direction::Left => (-1, 0),
        }
    }
}

struct Instruction {
    dir: Direction,
    size: isize,
}

#[derive(Clone, Eq)]
struct Point {
    x: i64,
    y: i64,
    steps: usize,
}

impl PartialEq for Point {
    fn eq(&self, other: &Self) -> bool {
        self.x == other.x && self.y == other.y
    }
}

impl Hash for Point {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.x.hash(state);
        self.y.hash(state);
    }
}

impl Point {
    fn manhatten_distance(&self) -> u64 {
        (self.x.abs() + self.y.abs()) as u64
    }
}

pub fn run() {
    let lines = get_lines(3, 1);
    let ins_1 = get_instructions(&lines[0]);
    let ins_1 = get_points(&ins_1);
    let ins_2 = get_instructions(&lines[1]);
    let ins_2 = get_points(&ins_2);

    let a0 = task_1(&ins_1, &ins_2);
    let a1 = task_2(&ins_1, &ins_2);
    print_answer(3, &a0, &a1);
}

fn get_instructions(input: &str) -> Vec<Instruction> {
    input.split(",")
        .map(|s| Instruction {
            dir: s[0..1].parse::<Direction>().expect("error parsing direction"),
            size: s[1..].parse::<isize>().expect("error parsing size"),
        })
        .collect()
}

fn get_points(ins: &Vec<Instruction>) -> Vec<Point> {
    let mut points = vec![];
    let mut p = Point{
        x: 0,
        y: 0,
        steps: 0,
    };

    for i in ins {
        let mut cur_points = generate_points(&p, i);
        p = cur_points.last().expect("No last point").clone();
        points.append(&mut cur_points);
    }

    points
}

fn generate_points(origin: &Point, i: &Instruction) -> Vec<Point> {
    let (x_change, y_change) = i.dir.changes();
    let mut p = origin.clone();
    let mut points = vec![];

    for _ in 0..i.size {
        p = Point {
            x: p.x + x_change as i64,
            y: p.y + y_change as i64,
            steps: p.steps + 1,
        };
        points.push(p.clone());
    }

    points
}

fn task_1(wire1: &Vec<Point>, wire2: &Vec<Point>) -> u64 {
    let wire1_set: HashSet<Point> = HashSet::from_iter(wire1.clone().into_iter());
    let wire2_set: HashSet<Point> = HashSet::from_iter(wire2.clone().into_iter());

    wire1_set
        .intersection(&wire2_set)
        .map(|p| p.manhatten_distance())
        .min()
        .unwrap_or(0)
}

fn task_2(wire1: &Vec<Point>, wire2: &Vec<Point>) -> u64 {
    let wire1_set: HashSet<Point> = HashSet::from_iter(wire1.clone().into_iter());
    let wire2_set: HashSet<Point> = HashSet::from_iter(wire2.clone().into_iter());

    let intersections = wire1_set
        .intersection(&wire2_set)
        .collect::<Vec<_>>();

    intersections.iter()
        .map(|i| {
            let p1 = wire1_set.get(i).expect("couldn't find point on wire");
            let p2 = wire2_set.get(i).expect("couldn't find point on wire");
            (p1.steps + p2.steps) as u64
        })
    .min()
    .unwrap_or(0)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn detailed_example() {
        let wire1 = get_instructions(&String::from("R8,U5,L5,D3"));
        let wire1 = get_points(&wire1);
        let wire2 = get_instructions(&String::from("U7,R6,D4,L4"));
        let wire2 = get_points(&wire2);

        assert_eq!(6, task_1(&wire1, &wire2));
    }

    #[test]
    fn t1_ex1() {
        let wire1 = get_instructions(&String::from("R75,D30,R83,U83,L12,D49,R71,U7,L72"));
        let wire1 = get_points(&wire1);
        let wire2 = get_instructions(&String::from("U62,R66,U55,R34,D71,R55,D58,R83"));
        let wire2 = get_points(&wire2);
        assert_eq!(159, task_1(&wire1, &wire2));
    }

    #[test]
    fn t1_ex2() {
        let wire1 = get_instructions(&String::from("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"));
        let wire1 = get_points(&wire1);
        let wire2 = get_instructions(&String::from("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"));
        let wire2 = get_points(&wire2);

        assert_eq!(135, task_1(&wire1, &wire2));
    }

    #[test]
    fn t2_ex1() {
        let wire1 = get_instructions(&String::from("R8,U5,L5,D3"));
        let wire1 = get_points(&wire1);
        let wire2 = get_instructions(&String::from("U7,R6,D4,L4"));
        let wire2 = get_points(&wire2);

        assert_eq!(30, task_2(&wire1, &wire2));
    }

    #[test]
    fn t2_ex2() {
        let wire1 = get_instructions(&String::from("R75,D30,R83,U83,L12,D49,R71,U7,L72"));
        let wire1 = get_points(&wire1);
        let wire2 = get_instructions(&String::from("U62,R66,U55,R34,D71,R55,D58,R83"));
        let wire2 = get_points(&wire2);
        assert_eq!(610, task_2(&wire1, &wire2));
    }

    #[test]
    fn t2_ex3() {
        let wire1 = get_instructions(&String::from("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"));
        let wire1 = get_points(&wire1);
        let wire2 = get_instructions(&String::from("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"));
        let wire2 = get_points(&wire2);

        assert_eq!(410, task_2(&wire1, &wire2));
    }
}
