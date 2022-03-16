import unittest
import text_helper


class TestTextHelper(unittest.TestCase):
    def test_one_word(self):
        text = "python"
        result = text_helper.cap_text(text)

        self.assertEqual(result, "Python")

    def test_multiple_words(self):
        text = "monty python"
        result = text_helper.cap_text(text)

        self.assertEqual(result, "Monty Python")


if __name__ == "__main__":
    unittest.main()
