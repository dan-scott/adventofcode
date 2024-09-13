use aoc_2015::{day_01, day_02, day_03};
use aoc_base::runner::run_days;
use aoc_base::Day;

fn main() {
    let days: Vec<Box<dyn Day>> = vec![day_01::new(), day_02::new(), day_03::new()];

    run_days(days).unwrap();
}
