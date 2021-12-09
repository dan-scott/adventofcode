use std::{fmt::Display, time::Instant};

pub trait Day {
    fn number(&self) -> u8;
    fn part_1(&self) -> Box<dyn Display>;
    fn part_2(&self) -> Box<dyn Display>;
}

pub fn run_days(days: Vec<Box<impl Day>>) {
    let runner_start = Instant::now();

    days.iter().for_each(|day| {
        println!("Day {}", day.number());
        print!("\tSolving part one...");
        let part_1_start = Instant::now();
        let p1 = day.part_1();
        println!(" done in {:?}\t{}", part_1_start.elapsed(), p1);

        print!("\tSolving part two...");
        let part_2_start = Instant::now();
        let p1 = day.part_2();
        println!(" done in {:?}\t{}\n", part_2_start.elapsed(), p1);
    });

    println!("Finished in {:?}", runner_start.elapsed());
}
