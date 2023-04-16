// Un trait es el equivalente a una interfaz en otro lenguaje de programación,
// se encarga de definir los métodos y atributos que deberán tener las
// implementaciones

use std::fmt::Display;

// La aplicación de los trait tiene una restricción, tan solo podemos aplicar
// aquellos que sean locales a nuestro Crate, no podemos añadir un trait como
// Display al modelo Vec
pub trait Summary {
    fn summarize(&self) -> String;

    // Los traits pueden tener valores por defecto para no tener que obligar el
    // desarrollador a implementarlos todos
    fn summarize_def(&self) -> String {
        String::from("Read more...")
    }
}

pub struct NewsArticle {
    pub headline: String,
    pub location: String,
    pub author: String,
}

impl Summary for NewsArticle {
    fn summarize(&self) -> String {
        format!("{}", self.author)
    }
}

fn main() {
    let news_common = NewsArticle {
        headline: String::from("Test"),
        location: String::from("City"),
        author: String::from("Test author"),
    };

    print!("{}", news_common.summarize());
    print!("{}", news_common.summarize_def());

    notify(&news_common)
}

// También podemos usar los traits en los parámetros de una función para que de
// esta manera aseguremos antes que estos tendrán implementados los métodos del
// trait correspondiente, se le puede pasar cualquier cosa que implemente el
// trait
pub fn notify(item: &impl Summary) {
    println!("{}", item.summarize())
}

// Cuando la función que creamos tenga parámetos genéricos y queremos que
// implementen un trait concreto, tenemos que usar los trait bounds
pub fn notify2<T: Summary>(item: &T) {
    println!("{}", item.summarize())
}

// Si lo que queremos es que implementen varios traits al mismo tiempo tenemos
// que usar el operador +
pub fn notify3(item: &(impl Summary + Display)) {
    println!("{}", item.summarize())
}

// Y para los genéricos
pub fn notify4<T: Summary + Display>(item: &T) {
    println!("{}", item.summarize())
}
