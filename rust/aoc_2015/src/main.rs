use aoc_base::runner::run_days;
use aoc_base::Day;

mod day_01;
mod day_02;

fn main() {
    let days: Vec<Box<dyn Day>> = vec![day_01::Day01::new(), day_02::Day02::new()];

    run_days(days).unwrap();
}
