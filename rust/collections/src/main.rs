mod helpers;

use std::collections::HashMap;

use helpers::math;

fn main() {
    let _result = math::sum(2, 2);

    let my_vector = vec![1, 2, 3];

    println!("{}", &my_vector[1]);

    match my_vector.get(100) {
        Some(element) => println!("{}", element),
        None => println!("There is no element with tthat index"),
    }

    for element in &my_vector {
        println!("element in for: {}", element);
    }

    let mut scores: HashMap<String, i32> = HashMap::new();

    scores.insert(String::from("test"), 22);

    match scores.get("test") {
        Some(value) => println!("{}", value),
        None => println!("No value"),
    }

    let word = "my word is my word";

    let counter = HashMap::new();

    for w in word.split_whitespace() {
        let &mut valueRef: &i32 = counter.entry(w);
        *valueRef += 1;
    }
}
