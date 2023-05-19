import 'package:dice_roll/main_title.dart';
import 'package:flutter/material.dart';

class GradientContainer extends StatelessWidget {
  const GradientContainer({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        gradient: LinearGradient(
          begin: Alignment.topLeft,
          end: Alignment.bottomRight,
          colors: [
            Colors.deepPurple,
            Colors.purple,
          ],
        ),
      ),
      child: const Center(
        child: MainTitle(
          text: "Hello world",
        ),
      ),
    );
  }
}
