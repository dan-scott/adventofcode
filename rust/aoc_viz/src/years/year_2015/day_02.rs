use anyhow::{Context, Result};
use nom::{bytes::complete::tag, character::complete::digit1, sequence::tuple};
use raylib::prelude::*;

use crate::scene::{SceneContainer, ScenePart, ScenePartLoader};
use std::mem::swap;

const STEP: usize = 60;
const SCALE: i32 = 4;
const SPACE: i32 = 3;

pub fn container() -> SceneContainer {
    SceneContainer::of(2015, 2, Box::new(Loader))
}

struct Loader;

impl ScenePartLoader for Loader {
    fn load_part_1(&self, lines: &[String]) -> Box<dyn ScenePart> {
        Box::new(Part1::new(lines))
    }

    fn load_part_2(&self, lines: &[String]) -> Box<dyn ScenePart> {
        todo!()
    }
}

struct Dims(i32, i32, i32);

impl Dims {
    fn of(input: &str) -> Result<Self> {
        let (_, (a, _, b, _, c)) = tuple((
            digit1,
            tag("x"),
            digit1,
            tag("x"),
            digit1, //
        ))(input)
        .map_err(|e: nom::Err<nom::error::Error<&str>>| e.to_owned())
        .context("Failed to nom or whatever")?;

        Ok(Self::new(a.parse()?, b.parse()?, c.parse()?).into())
    }

    fn new(a: i32, b: i32, c: i32) -> Self {
        let mut d = Self(a * SCALE, b * SCALE, c * SCALE);
        if d.1 > d.2 {
            swap(&mut d.1, &mut d.2);
        }
        if d.0 > d.1 {
            swap(&mut d.0, &mut d.1);
        }
        if d.1 > d.2 {
            swap(&mut d.1, &mut d.2);
        }
        d
    }
}

struct Part1 {
    index: usize,
    ct: usize,
    dims: Vec<Dims>,
}

impl Part1 {
    fn new(input: &[String]) -> Self {
        let dims: Result<Vec<Dims>> = input.iter().map(|line| Dims::of(line)).collect();

        Self {
            index: 0,
            ct: 0,
            dims: dims.unwrap(),
        }
    }
}

impl ScenePart for Part1 {
    fn step(&mut self) {
        self.ct += 1;

        if self.dims.len() <= self.index + STEP {
            return;
        }

        if self.ct % STEP == 0 {
            self.index += 60;
        }
    }

    fn render(&self, draw: &mut raylib::prelude::RaylibDrawHandle) {
        if let Some(dims) = self.dims.get(self.index) {
            draw.draw_rectangle(30, 30, dims.0, dims.1, Color::GREEN);
            draw.draw_rectangle(
                30 + dims.0 + SPACE + dims.2 + SPACE,
                30,
                dims.0,
                dims.1,
                Color::GREEN,
            );

            draw.draw_rectangle(30 + dims.0 + SPACE, 30, dims.2, dims.1, Color::BLUE);
            draw.draw_rectangle(
                30 + dims.0 + SPACE,
                30 + dims.1 + SPACE + dims.0 + SPACE,
                dims.2,
                dims.1,
                Color::BLUE,
            );

            draw.draw_rectangle(
                30 + dims.0 + SPACE,
                30 + dims.1 + SPACE,
                dims.2,
                dims.0,
                Color::RED,
            );
            draw.draw_rectangle(
                30 + dims.0 + SPACE,
                30 + 2 * (dims.1 + SPACE) + dims.0 + SPACE,
                dims.2,
                dims.0,
                Color::RED,
            );
        }
    }
}
