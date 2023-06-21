import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:meals_app/data/dummy_data.dart';
import 'package:meals_app/models/category.dart';
import 'package:meals_app/models/meal.dart';
import 'package:meals_app/providers/filters_provider.dart';
import 'package:meals_app/providers/meals_provider.dart';
import 'package:meals_app/screens/filters.dart';
import 'package:meals_app/screens/meals.dart';
import 'package:meals_app/screens/tabs.dart';
import 'package:meals_app/widgets/category_grid_item.dart';
import 'package:meals_app/widgets/main_drawer.dart';

class CategoriesScreen extends ConsumerStatefulWidget {
  const CategoriesScreen({
    super.key,
  });

  @override
  ConsumerState<CategoriesScreen> createState() => _CategoriesScreenState();
}

class _CategoriesScreenState extends ConsumerState<CategoriesScreen>
    with SingleTickerProviderStateMixin {
  // LESSON: late determina que la variable tendra un valor definido en el
  // momento de utilizarse, pero no en el momento de declararse
  late AnimationController _animationController;

  @override
  void initState() {
    super.initState();

    _animationController = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 300),
      lowerBound: 0,
      upperBound: 1,
    );

    _animationController.forward();
  }

  @override
  void dispose() {
    // LESSON: Todos los controladores de animaciones deben cerrarse al terminar
    // el widget
    _animationController.dispose();
    super.dispose();
  }

  void _selectCategory(BuildContext context, Category category) {
    final List<Meal> meals = ref.watch(mealsProvider);
    final Map<Filter, bool> activeFilters = ref.watch(filtersProvider);

    List<Meal> categoryMeals = meals.where((Meal meal) {
      if (!meal.categories.contains(category.id)) {
        return false;
      }
      if (activeFilters[Filter.glutenFree]! && !meal.isGlutenFree) {
        return false;
      }
      if (activeFilters[Filter.lactoseFree]! && !meal.isLactoseFree) {
        return false;
      }
      if (activeFilters[Filter.vegetarian]! && !meal.isVegetarian) {
        return false;
      }
      if (activeFilters[Filter.vegan]! && !meal.isVegan) {
        return false;
      }

      return true;
    }).toList();

    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => MealsScreen(
          title: category.title,
          meals: categoryMeals,
        ),
      ),
    );
  }

  void _setScreen(BuildContext context, String identifier) async {
    Navigator.of(context).pop();

    if (identifier == 'filters') {
      // Aqui se guardara lo que se envie desde los filtros al hacer el pop
      // final result = await Navigator.push<Map<Filter, bool>>(
      //   context,
      //   MaterialPageRoute(
      //     builder: (context) => const FiltersScreen(),
      //   ),
      // );

      Navigator.push<Map<Filter, bool>>(
        context,
        MaterialPageRoute(
          builder: (context) => const FiltersScreen(),
        ),
      );
    } else {
      Navigator.push(
        context,
        MaterialPageRoute(
          builder: (context) => const TabsScreen(),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Pick your category'),
      ),
      drawer: MainDrawer(onSelectScreen: _setScreen),
      body: AnimatedBuilder(
        animation: _animationController,
        child: GridView(
          padding: const EdgeInsets.all(24),
          gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
            crossAxisCount: 2,
            childAspectRatio: 3 / 2,
            crossAxisSpacing: 20,
            mainAxisSpacing: 20,
          ),
          children: [
            for (final category in availableCategories)
              CategoryGridItem(
                category: category,
                onSelectCategory: _selectCategory,
              ),
          ],
        ),
        builder: (BuildContext context, Widget? child) => Padding(
          padding: EdgeInsets.only(
            top: 100 - _animationController.value * 100,
          ),
          child: child,
        ),
      ),
    );
  }
}
