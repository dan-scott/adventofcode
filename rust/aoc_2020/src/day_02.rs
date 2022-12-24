use aoc_base::{Day, DayResult};

struct Day02 {}

pub fn new() -> Box<dyn Day> {
    Box::new(Day02 {})
}

impl Day for Day02 {
    fn year(&self) -> u16 {
        2020
    }

    fn number(&self) -> u8 {
        2
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        Ok(lines
            .iter()
            .map(parse_line)
            .map(Result::unwrap)
            .filter(|(policy, password)| policy.is_valid_range(password))
            .count()
            .into())
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        Ok(lines
            .iter()
            .map(parse_line)
            .map(Result::unwrap)
            .filter(|(policy, password)| policy.is_valid_pos(password))
            .count()
            .into())
    }
}

fn parse_line(s: &String) -> anyhow::Result<(PasswordPolicy, &str)> {
    peg::parser! {
        grammar parser() for str {
            rule number() -> usize
                = n:$(['0'..='9']+) { n.parse::<usize>().unwrap() - 1 }

            rule nums() -> [usize; 2]
                = min:number() "-" max:number() { [min, max] }

            rule byte() -> u8
                = letter:$['a'..='z'] { letter.as_bytes()[0] }

            rule password() -> &'input str
                = letters:$([_]*) { letters }

            pub(in crate::day_02) rule line() -> (PasswordPolicy, &'input str)
                = nums:nums() " " byte:byte() ": " password:password() {
                    (PasswordPolicy { nums, byte }, password)
                }
        }
    }

    Ok(parser::line(s)?)
}

#[derive(PartialEq, Debug)]
struct PasswordPolicy {
    byte: u8,
    nums: [usize; 2],
}

impl PasswordPolicy {
    fn is_valid_range(&self, password: &str) -> bool {
        (self.nums[0] + 1..=self.nums[1] + 1).contains(
            &password
                .as_bytes()
                .iter()
                .copied()
                .filter(|&b| b == self.byte)
                .count(),
        )
    }

    fn is_valid_pos(&self, password: &str) -> bool {
        self.nums
            .iter()
            .copied()
            .filter(|&index| password.as_bytes()[index] == self.byte)
            .count()
            == 1
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn day_02_part_1() {
        let lines = aoc_base::inputs::lines(2020, 2);
        let expected = 393;
        let res = new().part_1(&lines);
        match res {
            Ok(answer) => assert_eq!(answer, expected.into()),
            Err(er) => assert_eq!(er.to_string(), ""),
        }
    }

    #[test]
    fn day_02_part_2() {
        let lines = aoc_base::inputs::lines(2020, 2);
        let expected = 690;
        let answer = new().part_2(&lines).unwrap();
        assert_eq!(answer, expected.into())
    }
}
