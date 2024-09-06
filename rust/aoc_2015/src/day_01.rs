use anyhow::Error;
use aoc_base::{Day, DayResult};

pub struct Day01 {}

impl Day01 {
    pub(crate) fn new() -> Box<dyn Day> {
        Box::new(Self {})
    }
}

impl Day for Day01 {
    fn year(&self) -> u16 {
        2015
    }

    fn number(&self) -> u8 {
        1
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        if let Some(line) = lines.first() {
            let floor: isize = line
                .chars()
                .map(|c| match c {
                    '(' => 1,
                    ')' => -1,
                    _ => 0,
                })
                .sum();
            Ok(floor.into())
        } else {
            Err(Error::msg("Expected one line of input"))
        }
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        if let Some(line) = lines.first() {
            let idx = line
                .chars()
                .scan(0, |acc, ch| {
                    if ch == '(' {
                        *acc += 1;
                    } else if ch == ')' {
                        *acc -= 1;
                    }

                    if *acc < 0 {
                        None
                    } else {
                        Some(*acc)
                    }
                })
                .count();

            Ok((idx + 1).into())
        } else {
            Err(Error::msg("Expected one line of input"))
        }
    }
}
