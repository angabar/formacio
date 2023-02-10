use std::{cmp::Ordering, io};

use rand::Rng;

fn main() {
    println!("Guess the number!");

    let mut tries = 0;
    let random_number = rand::thread_rng().gen_range(1..101);

    loop {
        println!("Please, input your guess");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("error reading line");

        let guess: i32 = match guess.trim().parse() {
            Ok(number) => number,
            Err(_) => continue,
        };

        tries = tries + 1;

        match guess.cmp(&random_number) {
            Ordering::Less => println!("Number is greatter"),
            Ordering::Greater => println!("Number is less"),
            Ordering::Equal => {
                println!("You win! Need {} tries", tries);
                break;
            }
        }
    }
}
