use std::io::BufRead;

use crate::{inputs, runner};

struct Day01 {}

impl runner::Day for Day01 {
    fn number(&self) -> u8 {
        1
    }

    fn part_1(&self) -> Box<dyn std::fmt::Display> {
        let readings: Vec<usize> = inputs::open_file(2021, 1)
            .lines()
            .map(|l| l.unwrap().parse().unwrap())
            .collect();
        let mut total = 0;
        for i in 0..readings.len() - 1 {
            if readings[i] < readings[i + 1] {
                total += 1;
            }
        }
        Box::new(total)
    }

    fn part_2(&self) -> Box<dyn std::fmt::Display> {
        let readings: Vec<usize> = inputs::open_file(2021, 1)
            .lines()
            .map(|l| l.unwrap().parse().unwrap())
            .collect();
        let total = readings.windows(4).filter(|w| w[0] < w[3]).count();
        Box::new(total)
    }
}

pub fn new() -> Box<impl runner::Day> {
    return Box::new(Day01 {});
}

#[cfg(test)]
mod test {
    use crate::runner::Day;

    use super::*;

    #[test]
    fn day_01_part_01() {
        let expected = 1696;
        let answer = new().part_1();
        assert_eq!(answer.to_string(), expected.to_string())
    }

    #[test]
    fn day_01_part_02() {
        let expected = 1737;
        let answer = new().part_2();
        assert_eq!(answer.to_string(), expected.to_string())
    }
}
