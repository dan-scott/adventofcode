use anyhow::anyhow;
use aoc_base::DayResult::Skip;
use aoc_base::{Day, DayResult};
use crypto::digest::Digest;
use crypto::md5::Md5;

struct Day04;

pub fn new() -> Box<dyn Day> {
    Box::new(Day04)
}

impl Day for Day04 {
    fn year(&self) -> u16 {
        2015
    }

    fn number(&self) -> u8 {
        4
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let mut hasher = Md5::new();

        let key = lines.first().unwrap().as_bytes();
        let mut output = [0; 16]; // An MD5 is 16 bytes
        for i in 0..u64::MAX {
            hasher.input(key);
            hasher.input(i.to_string().as_bytes());

            hasher.result(&mut output);

            let first_five = output[0] | output[1] | (output[2] >> 4);
            if first_five == 0 {
                return Ok(i.into());
            }
            hasher.reset();
        }

        Err(anyhow!("fail"))
    }

    fn part_2(&self, _lines: &[String]) -> anyhow::Result<DayResult> {
        Skip.into()
        // let mut hasher = Md5::new();
        //
        // let key = _lines.first().unwrap().as_bytes();
        // let mut output = [0; 16]; // An MD5 is 16 bytes
        // for i in 0..u64::MAX {
        //     hasher.input(key);
        //     hasher.input(i.to_string().as_bytes());
        //
        //     hasher.result(&mut output);
        //
        //     let first_six = output[0] | output[1] | output[2];
        //     if first_six == 0 {
        //         return Ok(i.into());
        //     }
        //     hasher.reset();
        // }
        //
        // Err(anyhow!("fail"))
    }
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn day_04_part_01() {
        let input = vec!["abcdef".to_string()];
        let day = Day04;
        assert_eq!(day.part_1(&input).unwrap(), DayResult::Int(609043));
    }
}
