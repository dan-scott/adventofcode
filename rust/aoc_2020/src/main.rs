use aoc_base::{runner, Day};

fn main() {
    let days: Vec<Box<dyn Day>> = vec![aoc_2020::day_01::new(), aoc_2020::day_02::new()];

    runner::run_days(days).unwrap();
}
