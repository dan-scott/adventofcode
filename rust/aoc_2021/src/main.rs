mod inputs;
mod runner;
mod day_01;

fn main() {
    let days = vec![
        day_01::new(),
    ];

    runner::run_days(days);
}
