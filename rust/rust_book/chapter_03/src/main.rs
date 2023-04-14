fn main() {
    let my_tuple = (1, 2, 'a');

    let my_array_1 = [3; 5];
    let my_array_2: [i32; 5] = [0, 2, 4, 5, 8];

    println!("{}", my_tuple.0);
    println!("{}", my_array_1[4]);
    println!("{}", my_array_2[3]);
    println!("{}", give_me_five());

    let mut counter = 0

    una de las caracter

    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };
}

fn give_me_five() -> i32 {
    5
}
