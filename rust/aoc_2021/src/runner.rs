use crate::inputs;
use std::time::{Duration, Instant};

pub trait Day {
    fn number(&self) -> u8;
    fn part_1(&self, lines: &Vec<String>) -> (Duration, String);
    fn part_2(&self, lines: &Vec<String>) -> (Duration, String);
}

pub fn run_days(days: Vec<Box<impl Day>>) {
    let runner_start = Instant::now();

    days.iter().for_each(|day| {
        let lines = inputs::lines(2021, day.number());
        println!("Day {}", day.number());
        print!("\tSolving part one...");
        let p1 = day.part_1(&lines);
        println!(" done in {:?}\t{}", p1.0, p1.1);

        print!("\tSolving part two...");
        let p2 = day.part_2(&lines);
        println!(" done in {:?}\t{}\n", p2.0, p2.1);
    });

    println!("Finished in {:?}", runner_start.elapsed());
}
