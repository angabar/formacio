import 'package:flutter/material.dart';
import 'package:quiz_app/data/questions.dart';
import 'package:quiz_app/questions_summary.dart';

class ResultsScreen extends StatelessWidget {
  const ResultsScreen({
    super.key,
    required this.selectedAnswers,
  });

  final List<String> selectedAnswers;

  List<Map<String, Object>> getSummaryData() {
    final List<Map<String, Object>> summary = [];

    for (int i = 1; i <= selectedAnswers.length; i++) {
      summary.add({
        'question_index': i,
        'question': questions[i].text,
        'user_answer': selectedAnswers[i],
        'correct_answer': questions[i].answers[0],
      });
    }

    return summary;
  }

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
              child: QuestionsSummary(selectedAnswers: getSummaryData()),
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
