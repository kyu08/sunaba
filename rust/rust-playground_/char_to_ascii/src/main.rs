fn main() {
    let str = "abc";
    for char in str.chars() {
        let ascii_code = char as u8;
        println!("{}({})", char, ascii_code);
    }
}
