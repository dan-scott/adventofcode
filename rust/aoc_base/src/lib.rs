use std::fmt::Display;

pub mod inputs;
pub mod runner;
pub mod vec2;

#[derive(Debug, PartialEq, Eq)]
pub enum DayResult {
    Int(i128),
    Str(String),
}

macro_rules! day_result_from {
    ($t:ty) => {
        impl From<$t> for DayResult {
            fn from(value: $t) -> Self {
                DayResult::Int(value as i128)
            }
        }
    };
}

day_result_from!(i128);
day_result_from!(i64);
day_result_from!(i32);
day_result_from!(isize);
day_result_from!(usize);

impl From<String> for DayResult {
    fn from(value: String) -> Self {
        DayResult::Str(value)
    }
}

impl From<&str> for DayResult {
    fn from(value: &str) -> Self {
        DayResult::Str(value.to_string())
    }
}

impl Display for DayResult {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            DayResult::Int(i) => write!(f, "{}", i),
            DayResult::Str(s) => write!(f, "{}", s),
        }
    }
}

pub trait Day {
    fn year(&self) -> u16;
    fn number(&self) -> u8;
    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult>;
    fn part_2(&self, lines: &[String]) -> anyhow::Result<DayResult>;
}
