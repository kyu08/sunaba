trait PrintHello {
    fn print_hello(&self) {
        println!("Hello!!")
    }
}

struct Test1;
struct Test2;

impl PrintHello for Test1 {}

impl PrintHello for Test2 {
    fn print_hello(&self) {
        println!("Hello from not default impl")
    }
}

fn main() {
    let test1 = Test1;
    test1.print_hello();

    let test2 = Test2;
    test2.print_hello();
}
