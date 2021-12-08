use std::{
    fs::File,
    io::{BufRead, BufReader},
    path::Path,
};

trait InputParser {}

pub fn lines(year: u16, day: u8) -> Vec<String> {
    let root_path = match std::env::var_os("ADVENT_OF_CODE_ROOT") {
        Some(v) => v.into_string().unwrap(),
        None => "../../".to_string(),
    };
    let file = format!("{}/inputs/{}/{}.txt", root_path, year, day);
    let input_file = Path::new(file.as_str());
    let read = File::open(input_file).expect("Failed to open the thing");
    let reader = BufReader::new(read);
    return reader
        .lines()
        .enumerate()
        .map(|(_, r)| r.unwrap())
        .collect();
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_loads() {
        let loaded_lines = lines(2021, 1);
        println!("{:?}", loaded_lines);
        assert_ne!(loaded_lines.len(), 0)
    }
}
