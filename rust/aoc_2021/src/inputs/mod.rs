use std::{
    fs::File,
    io::{BufRead, BufReader},
    path::Path,
};

trait InputParser {}

fn open_file(year: u16, day: u8) -> BufReader<File> {
    let root_path = match std::env::var_os("ADVENT_OF_CODE_ROOT") {
        Some(v) => v.into_string().unwrap(),
        None => "../../".to_string(),
    };
    let file = format!("{}/inputs/{}/{}.txt", root_path, year, day);
    let input_file = Path::new(file.as_str());
    let read = File::open(input_file).expect("Failed to open the thing");
    return BufReader::new(read);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_loads_lines() {
        let loaded_lines = lines(2021, 1);
        assert_ne!(loaded_lines.len(), 0)
    }
}
