class Product {
  final String id;
  final String title;
  final String description;
  final double price;
  final String imageUrl;

  // isFavourite no es final porque queremos que se pueda modificar despues de
  // haberlo creado
  bool isFavourite;

  Product({
    required this.id,
    required this.title,
    required this.description,
    required this.price,
    required this.imageUrl,
    this.isFavourite = false,
  });
}
