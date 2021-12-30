use aoc_base::runner;

fn main() {
    let days: Vec<Box<dyn runner::Day>> = vec![
        Box::new(aoc_2021::day_01::new()),
        Box::new(aoc_2021::day_02::new()),
        Box::new(aoc_2021::day_03::new()),
    ];

    runner::run_days(days);
}
