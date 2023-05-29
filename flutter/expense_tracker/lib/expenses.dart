import 'package:expense_tracker/models/expense.dart';
import 'package:flutter/material.dart';

class Expenses extends StatefulWidget {
  const Expenses({super.key});

  @override
  State<Expenses> createState() {
    return _ExpensesState();
  }
}

class _ExpensesState extends State<Expenses> {
  final List<Expense> _registeredExpenses = [
    Expense(
      title: 'Asmorsar',
      amount: 4.5,
      date: DateTime.now(),
      category: Category.food,
    ),
    Expense(
      title: 'Sine',
      amount: 10,
      date: DateTime.now(),
      category: Category.leisure,
    ),
    Expense(
      title: 'Tren',
      amount: 5,
      date: DateTime.now(),
      category: Category.work,
    ),
  ];

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Column(
        children: [
          Text('Charts'),
          Text('List'),
        ],
      ),
    );
  }
}
