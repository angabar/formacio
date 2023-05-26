import 'package:flutter/material.dart';

class ResultsScreen extends StatelessWidget {
  const ResultsScreen({
    super.key,
    required this.selectedAnswers,
  });

  final List<String> selectedAnswers;

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: double.infinity,
      child: Container(
        margin: const EdgeInsets.all(40),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('You answered X out of Y questions correctly'),
            Container(
              margin: const EdgeInsets.only(top: 30),
              child: const Text('List of answered questions'),
            ),
            Container(
              margin: const EdgeInsets.only(top: 30),
              child: TextButton(
                child: const Text('RestarQuiz'),
                onPressed: () {},
              ),
            )
          ],
        ),
      ),
    );
  }
}
