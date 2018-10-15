import unittest
import WordFunnel

class TestIt(unittest.TestCase):
    def test(self):
        self.assertEqual(WordFunnel.main(), "hello world")


if __name__ == '__main__':
    unittest.main()