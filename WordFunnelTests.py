import unittest
import WordFunnel

class TestIt(unittest.TestCase):
    def test(self):
        self.assertEqual(WordFunnel.funnel("leave","eave"), True)


if __name__ == '__main__':
    unittest.main()