mod scene;
mod years;

use crate::scene::{SceneContainer, ScenePart};
use aoc_base::inputs::lines;
use raylib::prelude::*;
use scene::{SCREEN_HEIGHT, SCREEN_WIDTH};

enum State {
    Index,
    Part(Box<dyn ScenePart>),
}

fn main() {
    let days: Vec<SceneContainer> = vec![
        years::year_2015::day_01::container(),
        years::year_2015::day_02::container(),
    ];

    let (mut rl, thread) = init()
        .size(SCREEN_WIDTH, SCREEN_HEIGHT)
        .title("Advent of Code (rust)")
        .build();

    rl.set_target_fps(60);

    let fs = rl.gui_get_style(GuiControl::DEFAULT, GuiDefaultProperty::TEXT_SIZE as i32);
    let bg_color = Color::get_color(rl.gui_get_style(
        GuiControl::DEFAULT,
        GuiDefaultProperty::BACKGROUND_COLOR as i32,
    ) as u32);

    let mut state = State::Index;

    while !rl.window_should_close() {
        if rl.is_key_pressed(KeyboardKey::KEY_BACKSPACE) {
            state = State::Index;
        }
        let mut draw = rl.begin_drawing(&thread);
        draw.clear_background(bg_color);
        match state {
            State::Index => days.iter().enumerate().for_each(|(idx, container)| {
                let txt = format!("{} Day {}", container.year, container.day);
                let width = draw.measure_text(&txt, fs) as i32;
                let top: i32 = 10 + 40 * idx as i32;
                draw.draw_text(&txt, 10, top + 5, fs, Color::BLACK);

                if draw.gui_button(rrect(width + 20, top, 50, 20), Some(rstr!("Part 1"))) {
                    let input = lines(container.year, container.day);
                    state = State::Part(container.loader.load_part_1(&input));
                }

                if draw.gui_button(rrect(width + 80, top, 50, 20), Some(rstr!("Part 2"))) {
                    let input = lines(container.year, container.day);
                    state = State::Part(container.loader.load_part_2(&input));
                }
            }),
            State::Part(ref mut scene_part) => {
                scene_part.step();
                scene_part.render(&mut draw);
            }
        }
    }
}
