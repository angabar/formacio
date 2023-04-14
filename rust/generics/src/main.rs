use std::cmp::PartialOrd;

#[derive(Debug)]
struct Point<T> {
    x: T,
    y: T,
}

// Ojo en las implementaciones, tendremos que usar el tipo genérico dos veces
impl<T> Point<T> {
    fn x(&self) -> &T {
        &self.x
    }
}

// Podemos también crear métodos que solo pertenezcan a un tipo concreto de
// dato, como en este ejemplo el i32, aquellos objetos creados como f64 no
// tendrán acceso al mismo
impl Point<i32> {
    fn sum_points_number(&self) -> i32 {
        &self.x + &self.y
    }
}

#[derive(Debug)]
struct OtherPoint<T, U> {
    x: T,
    y: U,
}

fn main() {
    let number_list = vec![34, 59, 12, 1];

    // Ojo a la hora de usar los structs genéricos, si son del mismo tipo no
    // podremos poner valores diferentes
    let integer_point = Point { x: 2, y: 1 };
    let floating_point = Point { x: 2.1, y: 1.3 };

    // En este caso si, porque OtherPoint admite dos tipos diferentes de datos,
    // uno para cada punto
    let other_point = OtherPoint { x: 2.1, y: 3 };

    print!("{:?}", integer_point.x);
    print!("{:?}", floating_point.y);
    print!("{:?}", floating_point.x());
    print!("{:?}", integer_point.sum_points_number());

    print!("{:?}", other_point.x);
    print!("{:?}", other_point.y);

    println!("The largest number is {}", largest(&number_list));
}

// Cuando definimos un generico tenemos que pensar en los tipos de datos que va
// a recibir este, si son elementos comparables (como los usado en el for)
// tenemos que implementar el PartialOrd
fn largest<T: PartialOrd>(list: &[T]) -> &T {
    let mut largest = &list[0];

    for number in list {
        if number > largest {
            largest = number;
        }
    }

    largest
}
