use aoc_base::Day;
use std::{str::FromStr, time::Instant};

enum Instruction {
    Forward(usize),
    Down(usize),
    Up(usize),
}

impl FromStr for Instruction {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut pts = s.split(" ");
        let i = pts.next().unwrap();
        let v: usize = pts.next().unwrap().parse().unwrap();
        match i.chars().nth(0) {
            Some('f') => Ok(Instruction::Forward(v)),
            Some('d') => Ok(Instruction::Down(v)),
            Some('u') => Ok(Instruction::Up(v)),
            _ => unreachable!(),
        }
    }
}

pub struct Day02 {}

impl Day for Day02 {
    fn year(&self) -> u16 {
        2021
    }

    fn number(&self) -> u8 {
        2
    }

    fn part_1(&self, lines: &Vec<String>) -> anyhow::Result<(std::time::Duration, String)> {
        let s = Instant::now();

        let mut h = 0;
        let mut v = 0;
        let ins: Vec<Instruction> = lines.iter().map(|l| l.parse().unwrap()).collect();

        for i in ins {
            match i {
                Instruction::Up(m) => v -= m,
                Instruction::Down(m) => v += m,
                Instruction::Forward(m) => h += m,
            }
        }
        let e = s.elapsed();
        let a = (h * v).to_string();
        anyhow::Ok((e, a))
    }

    fn part_2(&self, lines: &Vec<String>) -> anyhow::Result<(std::time::Duration, String)> {
        let s = Instant::now();

        let mut h = 0;
        let mut v = 0;
        let mut ah = 0;
        let ins: Vec<Instruction> = lines.iter().map(|l| l.parse().unwrap()).collect();

        for i in ins {
            match i {
                Instruction::Up(m) => ah -= m,
                Instruction::Down(m) => ah += m,
                Instruction::Forward(m) => {
                    h += ah * m;
                    v += m;
                }
            }
        }
        let e = s.elapsed();
        let a = (h * v).to_string();
        anyhow::Ok((e, a))
    }
}

pub fn new() -> Day02 {
    Day02 {}
}
