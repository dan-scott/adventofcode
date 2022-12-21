use anyhow::{Error, Ok, Result};
use aoc_base::vec2::Vec2;
use aoc_base::Day;
use std::collections::HashMap;

pub struct Day02;

pub fn new() -> Box<dyn Day> {
    Box::new(Day02)
}

impl Day for Day02 {
    fn year(&self) -> u16 {
        2016
    }

    fn number(&self) -> u8 {
        2
    }

    fn part_1(&self, lines: &[String]) -> Result<aoc_base::DayResult> {
        let mut current = 4;
        let mut code = "".to_string();
        for line in lines.iter() {
            for c in line.chars() {
                current = match c {
                    'U' => (current - 3).max(current % 3),
                    'D' => (current + 3).min(6 + current % 3),
                    'L' => (current - 1).max(current - current % 3),
                    'R' => (current + 1).min(current - current % 3 + 2),
                    _ => return Err(Error::msg("Unrecognised direction")),
                }
            }
            code.push_str(&(current + 1).to_string())
        }
        Ok(code.into())
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<aoc_base::DayResult> {
        let map = HashMap::from([
            (Vec2::new(2, 0), "1"),
            (Vec2::new(1, 1), "2"),
            (Vec2::new(2, 1), "3"),
            (Vec2::new(3, 1), "4"),
            (Vec2::new(0, 2), "5"),
            (Vec2::new(1, 2), "6"),
            (Vec2::new(2, 2), "7"),
            (Vec2::new(3, 2), "8"),
            (Vec2::new(4, 2), "9"),
            (Vec2::new(1, 3), "A"),
            (Vec2::new(2, 3), "B"),
            (Vec2::new(3, 3), "C"),
            (Vec2::new(2, 4), "D"),
        ]);
        let u = Vec2::new(0, -1);
        let d = Vec2::new(0, 1);
        let l = Vec2::new(-1, 0);
        let r = Vec2::new(1, 0);
        let mut loc = Vec2::new(0, 2);
        let mut code = "".to_string();
        for line in lines.iter() {
            for c in line.chars() {
                let nl = match c {
                    'U' => loc.add(u),
                    'D' => loc.add(d),
                    'L' => loc.add(l),
                    'R' => loc.add(r),
                    _ => return Err(Error::msg("Unrecognised direction")),
                };
                if map.contains_key(&nl) {
                    loc = nl;
                }
            }
            code.push_str(map.get(&loc).unwrap())
        }

        Ok(code.into())
    }
}

#[cfg(test)]
mod test {
    use aoc_base::DayResult;

    use super::*;

    #[test]
    pub fn day_02_part_1() {
        let input: Vec<String> = "ULL
RRDDD
LURDL
UUUUD"
            .lines()
            .map(|l| l.to_string())
            .collect();

        let day = new();

        let actual = day.part_1(&input).unwrap();
        assert_eq!(DayResult::Str("1985".to_string()), actual);
    }

    #[test]
    pub fn day_02_part_2() {
        let input: Vec<String> = "ULL
RRDDD
LURDL
UUUUD"
            .lines()
            .map(|l| l.to_string())
            .collect();

        let day = new();

        let actual = day.part_2(&input).unwrap();
        assert_eq!(DayResult::Str("5DB3".to_string()), actual);
    }
}
