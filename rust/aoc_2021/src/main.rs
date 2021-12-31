use aoc_base::{runner, Day};

fn main() {
    let days: Vec<Box<dyn Day>> = vec![
        Box::new(aoc_2021::day_01::new()),
        Box::new(aoc_2021::day_02::new()),
        Box::new(aoc_2021::day_03::new()),
    ];

    runner::run_days(days).unwrap();
}
