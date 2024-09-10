use crate::scene::{SceneContainer, ScenePart, ScenePartLoader};
use raylib::color::Color;
use raylib::prelude::*;

pub(crate) fn container() -> SceneContainer {
    SceneContainer::of(2015, 1, Box::new(Loader {}))
}

struct Loader;

impl ScenePartLoader for Loader {
    fn load_part_1(&self, lines: &[String]) -> Box<dyn ScenePart> {
        Box::new(Part1::new(lines))
    }

    fn load_part_2(&self, lines: &[String]) -> Box<dyn ScenePart> {
        Box::new(Part2::new(lines))
    }
}

struct Part1 {
    floor: isize,
    cursor: usize,
    input: String,
}

impl Part1 {
    fn new(lines: &[String]) -> Self {
        Self {
            input: (*lines.first().unwrap()).clone(),
            cursor: 0,
            floor: 0,
        }
    }
}

impl ScenePart for Part1 {
    fn step(&mut self) {
        for _ in 0..5 {
            if self.cursor >= self.input.len() {
                return;
            }
            match self.input.chars().nth(self.cursor) {
                Some('(') => self.floor += 1,
                Some(')') => self.floor -= 1,
                _ => {}
            };
            self.cursor += 1;
        }
    }

    fn render(&self, draw: &mut RaylibDrawHandle) {
        draw.draw_text(&format!("Floor {}", self.floor), 10, 10, 20, Color::BLACK);
        draw.draw_text(
            &format!("Cursor {}/{}", self.cursor, self.input.len()),
            10,
            40,
            20,
            Color::BLACK,
        );
    }
}

struct Part2 {
    floor: isize,
    cursor: usize,
    input: String,
}
impl Part2 {
    fn new(lines: &[String]) -> Self {
        Self {
            input: (*lines.first().unwrap()).clone(),
            cursor: 0,
            floor: 0,
        }
    }
}

impl ScenePart for Part2 {
    fn step(&mut self) {
        for _ in 0..5 {
            if self.floor < 0 {
                return;
            }
            match self.input.chars().nth(self.cursor) {
                Some('(') => self.floor += 1,
                Some(')') => self.floor -= 1,
                _ => {}
            };
            self.cursor += 1;
        }
    }

    fn render(&self, draw: &mut RaylibDrawHandle) {
        draw.draw_text(&format!("Floor {}", self.floor), 10, 10, 20, Color::BLACK);
        draw.draw_text(
            &format!("Cursor {}/{}", self.cursor, self.input.len()),
            10,
            40,
            20,
            Color::BLACK,
        );
    }
}
