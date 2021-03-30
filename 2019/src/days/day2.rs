use crate::{print_answer, get_comma_int_list};

pub fn run() {
    let mut lines = get_comma_int_list(2, 1);
    let mut lines2 = lines.clone();
    lines[1] = 12;
    lines[2] = 2;
    let a1 = execute_intcode(lines)[0];

    let a2 = search_noun_verb(&mut lines2, 19690720);
    print_answer(2, &a1, &a2);
}

fn search_noun_verb(initial_memory: &mut Vec<usize>, target: usize) -> usize {
    for noun in 0..99 {
        for verb in 0..99 {
            let mut memory = initial_memory.clone();
            memory[1] = noun;
            memory[2] = verb;

            let result = execute_intcode(memory);
            if result[0] == target {
                return 100 * noun + verb;
            }
        }
    }
    0
}

fn execute_intcode(input_lines: Vec<usize>) -> Vec<usize> {
    let mut lines = input_lines.clone();
    let mut index = 0;
    loop {
        match lines[index] {
            1 => perform_addition(&mut lines, index),
            2 => perform_multiplication(&mut lines, index),
            99 => break,
            _ => panic!("Encountered unexpected opcode.")
        }
        index += 4;
    }
    lines
}

fn get_instruction_parameters(lines: Vec<usize>, index: usize) -> (usize, usize, usize) {
    let pos1 = lines[index + 1] as usize;
    let pos2 = lines[index + 2] as usize;
    let pos3 = lines[index + 3] as usize;
    (pos1, pos2, pos3)
}

fn perform_addition(lines: &mut Vec<usize>, index: usize) {
    let (pos1, pos2, pos3) = get_instruction_parameters(lines.to_vec(), index);
    lines[pos3] = lines[pos1] + lines[pos2];
}

fn perform_multiplication(lines: &mut Vec<usize>, index: usize) {
    let (pos1, pos2, pos3) = get_instruction_parameters(lines.to_vec(), index);
    lines[pos3] = lines[pos1] * lines[pos2];
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn detailed_example() {
        let input = vec![1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50];
        let expected = vec![3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50];
        assert_eq!(expected, execute_intcode(input));
    }

    #[test]
    fn t1_ex1() {
        let input = vec![1, 0, 0, 0, 99];
        let expected = vec![2, 0, 0, 0, 99];
        assert_eq!(expected, execute_intcode(input));
    }

    #[test]
    fn t1_ex2() {
        let input = vec![2, 3, 0, 3, 99];
        let expected = vec![2, 3, 0, 6, 99];
        assert_eq!(expected, execute_intcode(input));
    }

    #[test]
    fn t1_ex3() {
        let input = vec![2, 4, 4, 5, 99, 0];
        let expected = vec![2, 4, 4, 5, 99, 9801];
        assert_eq!(expected, execute_intcode(input));
    }

    #[test]
    fn t1_ex4() {
        let input = vec![1, 1, 1, 4, 99, 5, 6, 0, 99];
        let expected = vec![30, 1, 1, 4, 2, 5, 6, 0, 99];
        assert_eq!(expected, execute_intcode(input));
    }
}
