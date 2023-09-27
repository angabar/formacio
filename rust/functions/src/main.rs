fn main() {
    println!("Hello, world!");

    another_function(5, 'a');
}

fn another_function(x: i8, unit_label: char) {
    println!("Another function! {x}{unit_label}");
}
