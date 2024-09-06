use std::{
    fs::File,
    io::{BufRead, BufReader},
};

pub fn open_file(year: u16, day: u8) -> BufReader<File> {
    let root_path = match std::env::var_os("ADVENT_OF_CODE_ROOT") {
        Some(v) => v.into_string().unwrap(),
        None => "../../".to_string(),
    };
    let file = format!("{}/inputs/{}/{}.txt", root_path, year, day);
    let read = File::open(file).expect("Failed to open the thing");

    BufReader::new(read)
}

pub fn lines(year: u16, day: u8) -> Vec<String> {
    open_file(year, day).lines().map(|l| l.unwrap()).collect()
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::io::BufRead;

    #[test]
    fn it_opens_a_file() {
        assert_ne!(open_file(2021, 1).lines().count(), 0)
    }
}
