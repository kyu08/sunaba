use std::{
    fs::File,
    io::{BufRead, BufReader},
};

fn main() {
    let file = File::open("src/main.rs").unwrap();
    let reader = BufReader::new(file);
    let start = 1;
    let end = 100;

    let source_lines: Vec<_> = reader
        .lines()
        .skip(start - 1)
        .take(end - start + 1)
        .collect();
    for line in &source_lines {
        println!("{}", line.as_ref().unwrap());
    }
}
