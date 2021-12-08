mod day_01;
mod inputs;
mod runner;

fn main() {
    let days = vec![day_01::new()];

    runner::run_days(days);
}
