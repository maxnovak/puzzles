import unittest
import WordFunnel

class TestIt(unittest.TestCase):
    def test_eave(self):
        self.assertEqual(WordFunnel.funnel("leave","eave"), True)
    def test_rest(self):
        self.assertEqual(WordFunnel.funnel("reset","rest"), True)
    def test_dragon(self):
        self.assertEqual(WordFunnel.funnel("dragoon","dragon"), True)
    def test_leave(self):
        self.assertEqual(WordFunnel.funnel("eave","leave"), False)
    def test_lets(self):
        self.assertEqual(WordFunnel.funnel("sleet","lets"), False)
    def test_ski(self):
        self.assertEqual(WordFunnel.funnel("skiff","ski"), False)


if __name__ == '__main__':
    unittest.main()