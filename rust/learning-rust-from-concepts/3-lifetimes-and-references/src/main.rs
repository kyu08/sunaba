// fn return_hello() -> &String {}
fn return_hello() -> &String {
    let s = "hi".to_string();
    &s
}

fn main() {
    let s = return_hello();
    println!("{s}");
}
