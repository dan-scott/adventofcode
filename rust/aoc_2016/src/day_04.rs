use aoc_base::{Day, DayResult};
use itertools::Itertools;
use std::str::FromStr;

pub struct Day04;

pub fn new() -> Box<dyn Day> {
    Box::new(Day04)
}

#[derive(Debug)]
struct Room {
    name: String,
    serial: isize,
}

const A_ORD: usize = 'a' as usize;

impl FromStr for Room {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut counts = [0; 26];
        let chars = s.chars().collect_vec();
        let mut idx = 0;
        let mut c = chars[idx];
        while !c.is_numeric() {
            if c.is_numeric() {
                break;
            }
            if c.is_alphabetic() {
                counts[c as usize - A_ORD] += 1;
            }
            idx += 1;
            c = chars[idx];
        }

        let top = *counts
            .iter()
            .filter(|c| **c > 0)
            .sorted_by_key(|c| -(*c))
            .take(5)
            .last()
            .unwrap();

        let mut serial = 0;
        while c != '[' {
            serial *= 10;
            serial += c as isize - '0' as isize;
            idx += 1;
            c = chars[idx];
        }

        idx += 1;
        c = chars[idx];
        for _ in 0..4 {
            let nc = chars[idx + 1];
            let a = counts[c as usize - A_ORD];
            let b = counts[nc as usize - A_ORD];
            if a < top || b < top {
                return Err(());
            }

            if a < b {
                return Err(());
            }
            if a == b && c > nc {
                return Err(());
            }
            idx += 1;
            c = chars[idx];
        }

        Ok(Room {
            name: s.to_string(),
            serial,
        })
    }
}

impl Day for Day04 {
    fn year(&self) -> u16 {
        2016
    }

    fn number(&self) -> u8 {
        4
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<aoc_base::DayResult> {
        let mut sum: isize = 0;
        for line in lines {
            match Room::from_str(line) {
                Ok(room) => {
                    sum += room.serial;
                }
                _ => {}
            }
        }
        Ok(sum.into())
    }

    fn part_2(&self, _lines: &[String]) -> anyhow::Result<aoc_base::DayResult> {
        Ok(DayResult::Todo)
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    pub fn day_04_part_1() {
        let input = vec![
            "aaaaa-bbb-z-y-x-123[abxyz]".to_string(),
            "a-b-c-d-e-f-g-h-987[abcde]".to_string(),
            "not-a-real-room-404[oarel]".to_string(),
            "totally-real-room-200[decoy]".to_string(),
        ];
        let day = new();
        let actual = day.part_1(&input).unwrap();
        assert_eq!(DayResult::Int(1514), actual);
    }
}
