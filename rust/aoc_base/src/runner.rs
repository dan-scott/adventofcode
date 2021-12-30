use crate::inputs;
use std::time::Duration;

pub trait Day {
    fn number(&self) -> u8;
    fn part_1(&self, lines: &Vec<String>) -> (Duration, String);
    fn part_2(&self, lines: &Vec<String>) -> (Duration, String);
}

pub fn run_days(days: Vec<Box<dyn Day>>) {
    let d = days.iter().fold(Duration::default(), |t, day| {
        let lines = inputs::lines(2021, day.number());
        println!("Day {}", day.number());
        print!("\tSolving part one...");
        let p1 = day.part_1(&lines);
        println!(" done in {:?}\t{}", p1.0, p1.1);

        print!("\tSolving part two...");
        let p2 = day.part_2(&lines);
        println!(" done in {:?}\t{}\n", p2.0, p2.1);

        t + p1.0 + p2.0
    });

    println!("Finished in {:?}", d);
}
