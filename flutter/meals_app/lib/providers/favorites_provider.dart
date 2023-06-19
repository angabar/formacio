import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:meals_app/models/meal.dart';

class FavoriteMealsNotifier extends StateNotifier<List<Meal>> {
  FavoriteMealsNotifier() : super([]);

  void toggleMealFavoriteStatus(Meal meal) {
    final mealIsFavourite = state.contains(meal);

    // El estado es inmutable, por lo que no podemos hacer cosas como add o
    // remove
    if (mealIsFavourite) {
      state = state
          .where((Meal mealToFilter) => mealToFilter.id != meal.id)
          .toList();
    } else {
      state = [...state, meal];
    }
  }
}

final favoriteMealsProvider =
    StateNotifierProvider<FavoriteMealsNotifier, List<Meal>>(
        (StateNotifierProviderRef ref) {
  return FavoriteMealsNotifier();
});
