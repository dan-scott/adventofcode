use aoc_base::Day;

pub struct Day03;

pub fn new() -> Box<dyn Day> {
    Box::new(Day03)
}

impl Day for Day03 {
    fn year(&self) -> u16 {
        2016
    }

    fn number(&self) -> u8 {
        3
    }

    fn part_1(&self, lines: &[String]) -> anyhow::Result<aoc_base::DayResult> {
        Ok(lines
            .iter()
            .map(parse)
            .filter(is_valid_triangle)
            .count()
            .into())
    }

    fn part_2(&self, lines: &[String]) -> anyhow::Result<aoc_base::DayResult> {
        let mut col_a = vec![-1 as isize; lines.len()];
        let mut col_b = vec![-1 as isize; lines.len()];
        let mut col_c = vec![-1 as isize; lines.len()];
        for (i, line) in lines.iter().enumerate() {
            (col_a[i], col_b[i], col_c[i]) = parse(line);
        }
        col_a.append(&mut col_b);
        col_a.append(&mut col_c);

        Ok(col_a
            .chunks(3)
            .map(|c| (c[0], c[1], c[2]))
            .filter(is_valid_triangle)
            .count()
            .into())
    }
}

fn is_valid_triangle((a, b, c): &(isize, isize, isize)) -> bool {
    if a > b && a > c {
        b + c > *a
    } else if b > a && b > c {
        a + c > *b
    } else {
        a + b > *c
    }
}

fn parse(line: &String) -> (isize, isize, isize) {
    let (a, rest) = line.split_at(5);
    let (b, c) = rest.split_at(5);
    let an = a.trim().parse::<isize>().unwrap();
    let bn = b.trim().parse::<isize>().unwrap();
    let cn = c.trim().parse::<isize>().unwrap();
    (an, bn, cn)
}
#[cfg(test)]
mod test {
    use aoc_base::DayResult;

    use super::*;

    #[test]
    pub fn day_03_part_1() {
        let input = &["5 10 25".to_string()];
        let day = new();

        let actual = day.part_1(input).unwrap();

        assert_eq!(DayResult::Int(0), actual);
    }

    #[test]
    pub fn day_03_part_2() {
        let line: Vec<String> = "  101  301  501
  102  302  502
  103  303  503
  201  401  601
  202  402  602
  203  403  603"
            .split("\n")
            .map(|s| s.to_string())
            .collect();

        let day = new();
        let actual = day.part_2(&line).unwrap();
        assert_eq!(DayResult::Int(6), actual);
    }
}
