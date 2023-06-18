import 'package:flutter/material.dart';
import 'package:meals_app/data/dummy_data.dart';
import 'package:meals_app/models/category.dart';
import 'package:meals_app/models/meal.dart';
import 'package:meals_app/screens/filters.dart';
import 'package:meals_app/screens/meals.dart';
import 'package:meals_app/screens/tabs.dart';
import 'package:meals_app/widgets/category_grid_item.dart';
import 'package:meals_app/widgets/main_drawer.dart';

const kInitialFilters = {
  Filter.glutenFree: false,
  Filter.lactoseFree: false,
  Filter.vegetarian: false,
  Filter.vegan: false,
};

class CategoriesScreen extends StatefulWidget {
  const CategoriesScreen({
    super.key,
    required this.onToggleFavourite,
  });

  final void Function(Meal meal) onToggleFavourite;

  @override
  State<CategoriesScreen> createState() => _CategoriesScreenState();
}

class _CategoriesScreenState extends State<CategoriesScreen> {
  Map<Filter, bool> _selectedFilters = kInitialFilters;

  void _selectCategory(BuildContext context, Category category) {
    List<Meal> categoryMeals = dummyMeals
        .where((Meal meal) => meal.categories.contains(category.id))
        .toList();

    Navigator.push(
      context,
      MaterialPageRoute(
        builder: (context) => MealsScreen(
          title: category.title,
          meals: categoryMeals,
          onToggleFavourite: widget.onToggleFavourite,
        ),
      ),
    );
  }

  void _setScreen(BuildContext context, String identifier) async {
    Navigator.of(context).pop();

    if (identifier == 'filters') {
      // Aqui se guardara lo que se envie desde los filtros al hacer el pop
      final result = await Navigator.push<Map<Filter, bool>>(
        context,
        MaterialPageRoute(
          builder: (context) => const FiltersScreen(),
        ),
      );

      setState(() {
        _selectedFilters = result ?? kInitialFilters;
      });
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
      body: GridView(
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
    );
  }
}
