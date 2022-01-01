use anyhow::{Ok, Result};
use aoc_base::Day;
use std::time::Instant;

pub struct Day03 {}

pub fn new() -> Day03 {
    Day03 {}
}

impl Day for Day03 {
    fn year(&self) -> u16 {
        2021
    }

    fn number(&self) -> u8 {
        3
    }

    fn part_1(&self, lines: &[String]) -> Result<(std::time::Duration, String)> {
        let s = Instant::now();
        let zero_count = count_zeros(lines);
        let mut gamma = 0;
        let mut epsilon = 0;
        let half = lines.len() / 2;
        let mut val = 1 << 12;
        for i in zero_count {
            val >>= 1;
            if i >= half {
                gamma |= val;
            } else {
                epsilon |= val;
            }
        }
        let sln = epsilon * gamma;
        let e = s.elapsed();
        Ok((e, sln.to_string()))
    }

    fn part_2(&self, lines: &[String]) -> Result<(std::time::Duration, String)> {
        let s = Instant::now();
        let mut o2_list = lines.to_vec();
        let mut co2_list = lines.to_vec();
        let mut i = 0;
        while o2_list.len() > 1 && co2_list.len() > 1 {
            if o2_list.len() > 1 {
                let mc = most_common_at(i, &o2_list);
                o2_list.retain(|s| s.chars().nth(i).unwrap() == mc);
            }
            if co2_list.len() > 1 {
                let lc = least_common_at(i, &co2_list);
                co2_list.retain(|s| s.chars().nth(i).unwrap() == lc);
            }
            i += 1;
        }
        let o2 = o2_list
            .first()
            .map(|f| usize::from_str_radix(f, 2).unwrap())
            .unwrap();
        let co2 = co2_list
            .first()
            .map(|l| usize::from_str_radix(l, 2).unwrap())
            .unwrap();

        let sln = (o2 * co2).to_string();
        let e = s.elapsed();
        Ok((e, sln))
    }
}

fn count_zeros(lines: &[String]) -> [usize; 12] {
    let mut zero_count: [usize; 12] = [0; 12];
    for line in lines {
        line.bytes()
            .enumerate()
            .for_each(|(i, c)| zero_count[i] += if c == '0' as u8 { 1 } else { 0 })
    }
    zero_count
}

fn most_common_at(i: usize, lines: &Vec<String>) -> char {
    let half = lines.len() / 2;
    let ct = count_zeros_at(i, lines);
    if ct > half {
        '0'
    } else {
        '1'
    }
}

fn least_common_at(i: usize, lines: &Vec<String>) -> char {
    let half = lines.len() / 2;
    let ct = count_zeros_at(i, lines);
    if ct > half {
        '1'
    } else {
        '0'
    }
}

fn count_zeros_at(i: usize, lines: &Vec<String>) -> usize {
    lines
        .iter()
        .filter(|f| f.chars().nth(i).unwrap() == '0')
        .count()
}
