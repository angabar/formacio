import 'package:test/test.dart';
import 'package:test_tutorial/counter.dart';

void main() {
  group("Counter", () {
    test("Counter should work", () {
      final counter = Counter();
      counter.increment();
      expect(counter.value, 1);
    });

    test("Counter should work", () {
      final counter = Counter();
      counter.decrement();
      expect(counter.value, -1);
    });
  });
}
