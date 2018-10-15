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
    def test_bonus_dragoon(self):
        self.assertEqual(WordFunnel.bonus("dragoon"), ["dragon"])
    def test_bonus_boats(self):
        self.assertEqual(WordFunnel.bonus("boats"), ["bats", "boas", "boat", "bots", "oats"])
    def test_bonus_affidavit(self):
        self.assertEqual(WordFunnel.bonus("affidavit"), [])

if __name__ == '__main__':
    unittest.main()