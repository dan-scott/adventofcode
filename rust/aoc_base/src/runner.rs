use crate::{inputs, Day};
use anyhow::{Context, Ok, Result};
use std::time::{Duration, Instant};

pub fn run_days(days: Vec<Box<dyn Day>>) -> Result<()> {
    let mut total_duration = Duration::default();
    for day in days {
        let lines = inputs::lines(day.year(), day.number());

        println!("Day {}", day.number());
        print!("\tSolving part one...");
        let s = Instant::now();
        let p1_solution = day
            .part_1(&lines)
            .context(format!("Failed on day {} part 1", day.number()))?;
        let p1_time = Instant::now().duration_since(s);
        println!(" done in {:<15?}{}", p1_time, p1_solution);

        print!("\tSolving part two...");
        let s = Instant::now();
        let p2_solution = day
            .part_2(&lines)
            .context(format!("Failed on day {} part 2", day.number()))?;
        let p2_time = Instant::now().duration_since(s);
        println!(" done in {:<15?}{}\n", p2_time, p2_solution);

        total_duration += p1_time + p2_time;
    }

    println!("Finished in {:?}", total_duration);

    Ok(())
}
