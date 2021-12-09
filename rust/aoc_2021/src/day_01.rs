use crate::{runner, inputs};


struct Day01 {}

impl runner::Day for Day01 {
    fn number(&self) -> u8 {
        1
    }

    fn part_1(&self) -> Box<dyn std::fmt::Display> {
        let (total, _) = inputs::lines(2021, 1).iter().fold((0_usize, 0_usize), |(acc, prev), b| {
            let curr: usize = b.parse().unwrap();
            (if curr > prev { acc + 1 } else { acc }, curr)
        });
        Box::new(total)
    }

    fn part_2(&self) -> Box<dyn std::fmt::Display> {
        Box::new(0)
    }
}

pub fn new() -> Box<impl runner::Day> {
    return Box::new(Day01{})
}