use anyhow::{Context, Result};
use aoc_base::{Day, DayResult};
use nom::bytes::complete::tag;
use nom::character::complete::digit1;
use nom::sequence::tuple;

pub struct Day02 {}

impl Day02 {
    pub fn new() -> Box<dyn Day> {
        Box::new(Self {})
    }
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
    let (_, (a, _, b, _, c)) = tuple((
        digit1,
        tag("x"),
        digit1,
        tag("x"),
        digit1, //
    ))(input)
    .map_err(|e: nom::Err<nom::error::Error<&str>>| e.to_owned())
    .context("Failed to nom or whatever")?;

    let mut vals = [
        a.parse::<usize>()?,
        b.parse::<usize>()?,
        c.parse::<usize>()?,
    ];
    vals.sort();
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
