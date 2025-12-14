fn main() {
    let value_a: Option<i32> = Some(1);
    let value_b: Option<i32> = None;

    match (value_a, value_b) {
        (Some(v), None) | (None, Some(v)) => println!("{}", v),
        _ => println!("other"),
    }
}
