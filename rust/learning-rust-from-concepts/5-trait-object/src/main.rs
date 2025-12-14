fn generic_return_type() -> Box<dyn std::fmt::Display> {
    Box::new(1)
}

fn main() {
    println!("{}", generic_return_type());
}
