use std::{thread, time::Duration};

fn main() {
    let simulated_user_specified_value = 100;
    let simulated_random_number = 3;

    generate_workout(simulated_user_specified_value, simulated_random_number);
}

fn generate_workout(intensity: u32, random_number: u32) {
    let expensive_result = |num| {
        println!("calculating slowly...");
        thread::sleep(Duration::from_secs(2));
        num
    };
    let mut cacher = Cacher::new(expensive_result);

    if intensity < 25 {
        println!("Today, do {} pushups!", cacher.value(intensity));

        println!("Next, do {} situps!", cacher.value(intensity));
    } else {
        if random_number == 3 {
            println!("Take a break today");
        } else {
            println!("Today, run for {} minutes", cacher.value(intensity));
        }
    }
}

struct Cacher<T>
where
    T: Fn(u32) -> u32,
{
    calculation: T,
    value: Option<u32>,
}

impl<T> Cacher<T>
where
    T: Fn(u32) -> u32,
{
    fn new(calculation: T) -> Cacher<T> {
        Cacher {
            calculation,
            value: None,
        }
    }

    fn value(&mut self, arg: u32) -> u32 {
        match self.value {
            Some(v) => v,
            None => {
                let value = (self.calculation)(arg);
                self.value = Some(value);
                value
            }
        }
    }
}
