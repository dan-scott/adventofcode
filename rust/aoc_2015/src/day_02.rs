use anyhow::Result;
use aoc_base::{Day, DayResult};

struct Day02;

pub fn new() -> Box<dyn Day> {
    Box::new(Day02 {})
}

impl Day for Day02 {
    fn year(&self) -> u16 {
        2015
    }

    fn number(&self) -> u8 {
        2
    }

    fn part_1(&self, lines: &[String]) -> Result<DayResult> {
        let sizes: Result<Vec<usize>> = lines.iter().map(|line| get_sq_ft(&line)).collect();
        Ok(sizes?.iter().sum::<usize>().into())
    }

    fn part_2(&self, lines: &[String]) -> Result<DayResult> {
        let sizes: Result<Vec<usize>> = lines.iter().map(|line| get_ribbon_ft(&line)).collect();
        Ok(sizes?.iter().sum::<usize>().into())
    }
}

fn get_ribbon_ft(input: &str) -> Result<usize> {
    let [a, b, c] = parse_line(input)?;
    Ok(a * b * c + 2 * (a + b))
}

fn get_sq_ft(input: &str) -> Result<usize> {
    let [a, b, c] = parse_line(input)?;
    let (s1, s2, s3) = (a * b, a * c, b * c);
    Ok(2 * (s1 + s2 + s3) + s1)
}

fn parse_line(input: &str) -> Result<[usize; 3]> {
    let mut chars = input.split("x");

    let mut vals = [
        chars.next().unwrap().parse::<usize>()?,
        chars.next().unwrap().parse::<usize>()?,
        chars.next().unwrap().parse::<usize>()?,
    ];
    if vals[1] > vals[2] {
        vals.swap(1, 2);
    }
    if vals[0] > vals[1] {
        vals.swap(0, 1);
    }
    if vals[1] > vals[2] {
        vals.swap(1, 2);
    }
    // vals.sort();
    Ok(vals)
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_get_sq_ft() -> Result<()> {
        assert_eq!(get_sq_ft("2x3x4")?, 58);
        Ok(())
    }

    #[test]
    fn test_get_ribbon_ft() -> Result<()> {
        assert_eq!(get_ribbon_ft("2x3x4")?, 34);
        Ok(())
    }
}
