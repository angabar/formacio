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

    for (int i = 0; i < selectedAnswers.length; i++) {
      summary.add({
        'question_index': i + 1,
        'question': questions[i].text,
        'user_answer': selectedAnswers[i],
        'correct_answer': questions[i].answers[0],
      });
    }

    return summary;
  }

  @override
  Widget build(BuildContext context) {
    final List<Map<String, Object>> selectedAnswers = getSummaryData();
    final int totalQuestions = questions.length;
    final int correctQuestions = selectedAnswers
        .where((data) => data['user_answer'] == data['correct_answer'])
        .length;

    return SizedBox(
      width: double.infinity,
      child: Container(
        margin: const EdgeInsets.all(40),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              'You answered $correctQuestions out of $totalQuestions questions correctly',
            ),
            Container(
              margin: const EdgeInsets.only(top: 30),
              child: QuestionsSummary(selectedAnswers: selectedAnswers),
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
