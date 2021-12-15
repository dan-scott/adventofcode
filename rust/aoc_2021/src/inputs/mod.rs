use std::{fs::File, io::BufReader};

trait InputParser {}

pub fn open_file(year: u16, day: u8) -> BufReader<File> {
    let root_path = match std::env::var_os("ADVENT_OF_CODE_ROOT") {
        Some(v) => v.into_string().unwrap(),
        None => "../../".to_string(),
    };
    let file = format!("{}/inputs/{}/{}.txt", root_path, year, day);
    let read = File::open(file).expect("Failed to open the thing");
    return BufReader::new(read);
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::io::BufRead;

    #[test]
    fn it_opens_a_file() {
        let strs: Vec<String> = open_file(2021, 1).lines().map(|l| l.unwrap()).collect();
        assert_ne!(strs.len(), 0)
    }
}
