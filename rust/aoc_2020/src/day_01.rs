use aoc_base::{Day, DayResult};
use itertools::Itertools;

pub struct Day01 {}

pub fn new() -> Box<dyn Day> {
    Box::new(Day01 {})
}

impl Day for Day01 {
    fn year(&self) -> u16 {
        2020
    }

    fn number(&self) -> u8 {
        1
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let (l, r) = lines
            .iter()
            .map(|l| l.parse::<i64>())
            .collect::<Result<Vec<_>, _>>()?
            .into_iter()
            .tuple_combinations()
            .find(|(a, b)| a + b == 2020)
            .expect("No combination found");

        Ok((l * r).into())
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let (a, b, c) = lines
            .iter()
            .map(|l| l.parse::<i64>())
            .collect::<Result<Vec<_>, _>>()?
            .into_iter()
            .tuple_combinations()
            .find(|(a, b, c)| a + b + c == 2020)
            .expect("No combination found");

        Ok((a * b * c).into())
    }
}

#[cfg(test)]
mod test {

    use super::*;

    #[test]
    fn day_01_part_1() {
        let lines = aoc_base::inputs::lines(2020, 1);
        let expected = 211899;
        let (_, answer) = new().part_1(&lines).unwrap();
        assert_eq!(answer, expected.to_string())
    }

    #[test]
    fn day_01_part_2() {
        let lines = aoc_base::inputs::lines(2020, 1);
        let expected = 275765682;
        let (_, answer) = new().part_2(&lines).unwrap();
        assert_eq!(answer, expected.to_string())
    }
}
