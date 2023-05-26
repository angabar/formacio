import 'package:flutter/material.dart';

class QuestionsSummary extends StatelessWidget {
  const QuestionsSummary({
    super.key,
    required this.selectedAnswers,
  });

  final List<Map<String, Object>> selectedAnswers;

  @override
  Widget build(BuildContext context) {
    return Column(
      children: selectedAnswers.map((singleSelectedAnswer) {
        return Row(
          children: [
            Text(singleSelectedAnswer['question_index'].toString()),
            Column(
              children: [
                Text(singleSelectedAnswer['question'].toString()),
                const SizedBox(height: 5),
                Text(singleSelectedAnswer['user_answer'].toString()),
                Text(singleSelectedAnswer['correct_answer'].toString()),
              ],
            ),
          ],
        );
      }).toList(),
    );
  }
}
