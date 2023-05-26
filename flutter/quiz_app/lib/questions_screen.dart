import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:quiz_app/answer_button.dart';
import 'package:quiz_app/data/questions.dart';
import 'package:quiz_app/models/quiz_question.dart';

class QuestionsScreen extends StatefulWidget {
  const QuestionsScreen({
    super.key,
    required this.chooseAnswer,
  });

  final void Function(String answer) chooseAnswer;

  @override
  State<QuestionsScreen> createState() {
    return _QuestionsScreenState();
  }
}

class _QuestionsScreenState extends State<QuestionsScreen> {
  int questionIndex = 0;

  void answerQuestion(String answer) {
    widget.chooseAnswer(answer);

    setState(() {
      questionIndex++;
    });
  }

  @override
  Widget build(BuildContext context) {
    QuizQuestion currentQuestion = questions[questionIndex];

    return SizedBox(
      width: double.infinity,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Container(
            margin: const EdgeInsets.symmetric(horizontal: 40),
            child: Text(
              currentQuestion.text,
              textAlign: TextAlign.center,
              style: GoogleFonts.lato(
                color: Colors.white,
                fontSize: 24,
              ),
            ),
          ),
          Container(
            margin: const EdgeInsets.only(
              top: 30,
              left: 40,
              right: 40,
            ),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children:
                  currentQuestion.getShuffledAnswers().map((singleAnswer) {
                return AnswerButton(
                  text: singleAnswer,
                  onPress: () {
                    answerQuestion(singleAnswer);
                  },
                );
              }).toList(),
            ),
          ),
        ],
      ),
    );
  }
}
