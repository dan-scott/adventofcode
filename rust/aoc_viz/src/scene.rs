use raylib::drawing::RaylibDrawHandle;

pub const SCREEN_WIDTH: i32 = 800;
pub const SCREEN_HEIGHT: i32 = 450;

pub struct SceneContainer {
    pub year: u16,
    pub day: u8,
    pub loader: Box<dyn ScenePartLoader>,
}

impl SceneContainer {
    pub fn of(year: u16, day: u8, loader: Box<dyn ScenePartLoader>) -> Self {
        Self { year, day, loader }
    }
}

pub trait ScenePartLoader {
    fn load_part_1(&self, lines: &[String]) -> Box<dyn ScenePart>;
    fn load_part_2(&self, lines: &[String]) -> Box<dyn ScenePart>;
}

pub trait ScenePart {
    fn step(&mut self);
    fn render(&self, draw: &mut RaylibDrawHandle);
}
