use aoc_base::{runner, Day};

fn main() {
    let days: Vec<Box<dyn Day>> = vec![
        aoc_2016::day_01::new(),
        aoc_2016::day_02::new(),
        aoc_2016::day_03::new(),
    ];

    runner::run_days(days).unwrap();
}
