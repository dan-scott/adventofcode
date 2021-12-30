use std::time::{Duration, Instant};

use aoc_base::runner;

pub struct Day01 {}

pub fn new() -> Day01 {
    Day01 {}
}

impl runner::Day for Day01 {
    fn number(&self) -> u8 {
        1
    }

    fn part_1(&self, lines: &Vec<String>) -> (Duration, String) {
        let s = Instant::now();
        let readings: Vec<usize> = lines.iter().map(|l| l.parse().unwrap()).collect();
        let solution = readings.windows(2).filter(|w| w[0] < w[1]).count();
        let e = s.elapsed();
        (e, solution.to_string())
    }

    fn part_2(&self, lines: &Vec<String>) -> (Duration, String) {
        let s = Instant::now();
        let readings: Vec<usize> = lines.iter().map(|l| l.parse().unwrap()).collect();
        let solution = readings.windows(4).filter(|w| w[0] < w[3]).count();
        let e = s.elapsed();
        (e, solution.to_string())
    }
}

#[cfg(test)]
mod test {
    use aoc_base::runner::Day;

    use super::*;

    #[test]
    fn day_01_part_01() {
        let lines = aoc_base::inputs::lines(2021, 1);
        let expected = 1696;
        let (_, answer) = new().part_1(&lines);
        assert_eq!(answer, expected.to_string())
    }

    #[test]
    fn day_01_part_02() {
        let expected = 1737;
        let lines = aoc_base::inputs::lines(2021, 1);
        let (_, answer) = new().part_2(&lines);
        assert_eq!(answer, expected.to_string())
    }
}
