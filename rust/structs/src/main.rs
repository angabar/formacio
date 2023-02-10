struct Rectangle {
    x: i32,
    y: i32,
}

impl Rectangle {
    fn defatult(x: i32, y: i32) -> Self {
        Self { x, y }
    }

    fn area(&self) -> i32 {
        &self.x * &self.y
    }
}

fn main() {
    // let my_rectangle = Rectangle { x: 22, y: 12 };

    let my_rectangle = Rectangle::defatult(12, 22);

    crate::helpers::text::getNames();

    println!("Area of the rectangle {}", my_rectangle.area());
}
