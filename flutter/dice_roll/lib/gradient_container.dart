import 'package:flutter/material.dart';

import 'package:dice_roll/dice_roller.dart';

class GradientContainer extends StatelessWidget {
  const GradientContainer({
    super.key,
    required this.firstColor,
    required this.secondColor,
  });

  const GradientContainer.purple({super.key})
      : firstColor = Colors.purple,
        secondColor = Colors.deepPurple;

  final Color firstColor;
  final Color secondColor;

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: BoxDecoration(
        gradient: LinearGradient(
          begin: Alignment.topLeft,
          end: Alignment.bottomRight,
          colors: [
            firstColor,
            secondColor,
          ],
        ),
      ),
      child: const Center(
        child: Center(
          child: DiceRoller(),
        ),
      ),
    );
  }
}
