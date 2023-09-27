// Una constante siempre a de tener tipo y debe ser calculada en tiempo de
// compilacion
const IVA_PERCENT: i8 = 16;

fn main() {
    let mut x: i32 = 5;
    println!("The value of x is: {x}");
    x = 6;
    println!("The value of x is: {x}");

    println!("The value of IVA is: {IVA_PERCENT}");

    // Podemos hacer re-asignación de variables re-usando let
    let b: i8 = 2;
    let b: i8 = b + 3;

    println!("b = {b}");

    // En Rust existen dos tipos de valores, scalar y compound

    // Debemos tipar todas las variables, en algunos casos esto se hace de manera
    // implicita, pero en otros, como en la conversión de tipos ha de hacerse de
    // manera explicita
    let guess: i8 = "42".parse().expect("Not a number");
    println!("guess: {guess}");

    // Scalar: eneteros, decimales, booleans y carácteres

    // El tipo puede ir junto con el valor y este separarse con _ si es muy
    // largo
    let num = 1_000i16;
    println!("guess: {num}");

    // El integer overflow es manejado en rust de la siguiente manera, si es
    // build normal el overflow es tratado con un panic, si es con --release no,
    // en este caso se produce un desbordamiento y el valor que sobrepasa
    // empezara a contar de 0

    // Caso u8: 256 -> 0, 257 -> 1...

    let character: char = 'a';
    println!("guess: {character}");

    // Compound: tuplas y arrays

    // Las tuplas son agrupaciones fijas de tamaño que pueden ser de diferente
    // tipo
    let tup: (i16, f32, char) = (500, 22.4, 'a');

    // Puede accederse mediante indice
    let second_value: f32 = tup.1;
    println!("guess: {second_value}");

    // O mediante destrucción de parámetros
    let (x, y, z) = tup;
    println!("guess: {x}");
    println!("guess: {y}");
    println!("guess: {z}");

    // Una tupla vacía se llama unit y se representa con paréntesis vacíos () se
    // usa para indicar que una entidad no devuelve nada (caso void?)

    // El otro tipo de compound es el array, este es un conjunto de valores de
    // tamaño y tipo fijo (esto último difiere de las tuplas)
    let list: [i8; 5] = [1, 2, 3, 4, 5];
    let first_element = list[0];
    println!("guess: {first_element}");

    // Los arrays permiten setearse con valores por defecto un número
    // determinado de veces sustituyendo el tipo por el valor
    let prefered_list: [i8; 5] = [3; 5];
    let last_element = prefered_list[4];
    println!("guess: {last_element}");
}
