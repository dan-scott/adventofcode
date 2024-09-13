use aoc_base::vec2::Vec2;
use aoc_base::{Day, DayResult};
use std::str::Chars;

struct Day03;

pub fn new() -> Box<dyn Day> {
    Box::new(Day03)
}

impl Day for Day03 {
    fn year(&self) -> u16 {
        2015
    }

    fn number(&self) -> u8 {
        3
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let mut map = Map::new();
        let mut reader = Reader::new(lines.first().unwrap());
        let mut loc = Vec2::zero();
        let mut ct = map.visit(loc);

        while let Some(dir) = reader.next() {
            loc.add_mut(dir);
            ct += map.visit(loc);
        }

        Ok(ct.into())
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<DayResult> {
        let mut map = Map::new();
        let mut reader = Reader::new(lines.first().unwrap());
        let mut santa = Vec2::zero();
        let mut bot = Vec2::zero();
        let mut ct = map.visit(santa);

        loop {
            let s_dir = reader.next();
            let b_dir = reader.next();

            match (s_dir, b_dir) {
                (Some(s), Some(b)) => {
                    santa.add_mut(s);
                    bot.add_mut(b);
                    ct += map.visit(santa) + map.visit(bot);
                }
                _ => break,
            }
        }

        Ok(ct.into())
    }
}

pub struct Map {
    pub cells: [[usize; 200]; 200],
}

pub enum Visit {
    FirstVisit,
    Revisit,
}

impl Map {
    pub fn new() -> Self {
        Self {
            cells: [[1; 200]; 200],
        }
    }

    pub fn visit(&mut self, pos: Vec2) -> usize {
        let x = (pos.x + 100).clamp(0, 199) as usize;
        let y = (pos.y + 100).clamp(0, 199) as usize;
        let v = self.cells[x][y];
        self.cells[x][y] &= 0;
        v
    }
}

pub struct Reader<'a> {
    line: Chars<'a>,
}

impl<'a> Reader<'a> {
    pub fn new(line: &'a str) -> Self {
        Self {
            line: line.chars().into_iter(),
        }
    }
}

impl<'a> Iterator for Reader<'a> {
    type Item = Vec2;

    fn next(&mut self) -> Option<Vec2> {
        match self.line.next() {
            None => None,
            Some(dir) => match dir {
                '^' => Some(Vec2::new(0, -1)),
                'v' => Some(Vec2::new(0, 1)),
                '<' => Some(Vec2::new(-1, 0)),
                '>' => Some(Vec2::new(1, 0)),
                _ => None,
            },
        }
    }
}
