use crate::{inputs, Day};
use anyhow::{Context, Ok, Result};
use std::time::Duration;

pub fn run_days(days: Vec<Box<dyn Day>>) -> Result<()> {
    let mut total_duration = Duration::default();
    for day in days {
        let lines = inputs::lines(day.year(), day.number());
        println!("Day {}", day.number());
        print!("\tSolving part one...");
        let (p1_time, p1_solution) = day
            .part_1(&lines)
            .context(format!("Failed on day {} part 1", day.number()))?;
        println!(" done in {:?}\t{}", p1_time, p1_solution);

        print!("\tSolving part two...");
        let (p2_time, p2_solution) = day
            .part_2(&lines)
            .context(format!("Failed on day {} part 2", day.number()))?;
        println!(" done in {:?}\t{}\n", p2_time, p2_solution);

        total_duration += p1_time + p2_time;
    }

    println!("Finished in {:?}", total_duration);

    Ok(())
}
