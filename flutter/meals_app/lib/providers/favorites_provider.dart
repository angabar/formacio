import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:meals_app/models/meal.dart';

class FavoriteMealsNotifier extends StateNotifier<List<Meal>> {
  FavoriteMealsNotifier() : super([]);

  bool toggleMealFavoriteStatus(Meal meal) {
    final mealIsFavourite = state.contains(meal);

    // LESSON: El estado es inmutable, por lo que no podemos hacer cosas como
    // add o remove
    if (mealIsFavourite) {
      state = state
          .where((Meal mealToFilter) => mealToFilter.id != meal.id)
          .toList();
      return false;
    } else {
      state = [...state, meal];
      return true;
    }
  }
}

final favoriteMealsProvider =
    StateNotifierProvider<FavoriteMealsNotifier, List<Meal>>(
        (StateNotifierProviderRef ref) {
  return FavoriteMealsNotifier();
});
