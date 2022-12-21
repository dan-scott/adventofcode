use std::{collections::HashSet, str::FromStr};

use anyhow::Error;
use aoc_base::{vec2::Vec2, Day, DayResult};

pub struct Day01 {}

pub fn new() -> Box<dyn Day> {
    Box::new(Day01 {})
}

enum Dir {
    Left(isize),
    Right(isize),
}

impl FromStr for Dir {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.len() < 2 {
            return Err(());
        }
        let val = s[1..].parse::<isize>().unwrap();
        if s.chars().nth(0) == Some('L') {
            Ok(Dir::Left(val))
        } else {
            Ok(Dir::Right(val))
        }
    }
}

impl Day for Day01 {
    fn year(&self) -> u16 {
        2016
    }

    fn number(&self) -> u8 {
        1
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let (_, loc) = lines[0].split(", ").map(Dir::from_str).fold(
            (Vec2::new(0, 1), Vec2::zero()),
            |(heading, location), dir| {
                let dist: isize;
                let nh: Vec2;
                (nh, dist) = match dir.unwrap() {
                    Dir::Left(d) => (heading.rot_left(), d),
                    Dir::Right(d) => (heading.rot_right(), d),
                };
                (nh, location.add(nh.mult(dist)))
            },
        );
        Ok(loc.manhattan_dist(Vec2::zero()).into())
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let mut loc_set = HashSet::new();
        let mut heading = Vec2::new(0, 1);
        let mut loc = Vec2::zero();
        loc_set.insert(loc);
        for dir in lines[0].split(", ").map(Dir::from_str) {
            let dist: isize;
            (heading, dist) = match dir.unwrap() {
                Dir::Left(d) => (heading.rot_left(), d),
                Dir::Right(d) => (heading.rot_right(), d),
            };
            for _ in 0..dist {
                loc = loc.add(heading);
                if loc_set.contains(&loc) {
                    return Ok(loc.manhattan_dist(Vec2::zero()).into());
                } else {
                    loc_set.insert(loc);
                }
            }
        }
        Err(Error::msg("not found"))
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn day_01_part_1() {
        let day = new();
        let cases = vec![("R2, L3", 5), ("R2, R2, R2", 2), ("R5, L5, R5, R3", 12)];
        for (input, expected) in cases {
            let lines = [input.to_string()];
            let actual = day.part_1(&lines).unwrap();
            assert_eq!(DayResult::Int(expected), actual);
        }
    }

    #[test]
    fn day_01_part_2() {
        let day = new();
        let input = ["R8, R4, R4, R8".to_string()];
        let actual = day.part_2(&input).unwrap();
        assert_eq!(DayResult::Int(4), actual)
    }
}
