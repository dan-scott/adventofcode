use aoc_base::runner::Day;
use criterion::{criterion_group, criterion_main, Criterion};

macro_rules! benchmark {
    ($day: ident, $part: ident) => {
        paste::item! {
            fn [< benchmark_ $day _ $part >](c: &mut Criterion) {
                let d = aoc_2021::$day::new();
                let lines = aoc_base::inputs::lines(2021, d.number());
                c.bench_function(stringify!($day, $part), |b| b.iter(|| d.$part(&lines)));
            }
        }
    };
}

macro_rules! benchmarks {
    ($($day: ident),+) => {
        $(
            benchmark!{$day, part_1}
            benchmark!{$day, part_2}
        )+
        paste::item! {
            criterion_group!(benches, $([< benchmark_ $day _ part_1 >],[< benchmark_ $day _ part_2 >],)+);
        }
        criterion_main!(benches);
    };
}

benchmarks! {day_01, day_02, day_03}
