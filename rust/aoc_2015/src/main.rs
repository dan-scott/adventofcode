use aoc_base::runner::run_days;
use aoc_base::Day;

mod day_01;

fn main() {
    let days: Vec<Box<dyn Day>> = vec![
        day_01::Day01::new(),
    ];

    run_days(days).unwrap();
}
