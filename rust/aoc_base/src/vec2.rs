use std::cmp::{max, min};

#[derive(Clone, Copy, Hash, PartialEq, Eq, Debug)]
pub struct Vec2 {
    pub x: isize,
    pub y: isize,
}

impl Vec2 {
    pub fn zero() -> Self {
        Vec2 { x: 0, y: 0 }
    }

    pub fn new(x: isize, y: isize) -> Self {
        Vec2 { x, y }
    }

    pub fn add(&self, other: Self) -> Self {
        Vec2 {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }

    pub fn mult(&self, mag: isize) -> Self {
        Vec2 {
            x: self.x * mag,
            y: self.y * mag,
        }
    }

    pub fn rot_left(&self) -> Self {
        Vec2 {
            x: -self.y,
            y: self.x,
        }
    }

    pub fn rot_right(&self) -> Self {
        Vec2 {
            x: self.y,
            y: -self.x,
        }
    }

    pub fn manhattan_dist(&self, other: Self) -> isize {
        (self.x - other.x).abs() + (self.y - other.y).abs()
    }
}

impl Vec2 {
    pub fn add_mut(&mut self, other: Self) {
        self.x += other.x;
        self.y += other.y;
    }

    pub fn min_of_mut(&mut self, other: Self) {
        self.x = min(self.x, other.x);
        self.y = min(self.y, other.y);
    }

    pub fn max_of_mut(&mut self, other: Self) {
        self.x = max(self.x, other.x);
        self.y = max(self.y, other.y);
    }
}
