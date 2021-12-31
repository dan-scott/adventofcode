use std::time::Duration;

pub mod inputs;
pub mod runner;

pub trait Day {
    fn year(&self) -> u16;
    fn number(&self) -> u8;
    fn part_1(&self, lines: &Vec<String>) -> anyhow::Result<(Duration, String)>;
    fn part_2(&self, lines: &Vec<String>) -> anyhow::Result<(Duration, String)>;
}
