class QuizQuestion {
  const QuizQuestion({
    required this.text,
    required this.answers,
  });

  final String text;
  final List<String> answers;

  List<String> getShuffledAnswers() {
    final List<String> shufflededAnswers = List.of(answers);
    shufflededAnswers.shuffle();
    return shufflededAnswers;
  }
}
