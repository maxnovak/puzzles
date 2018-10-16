import unittest
import UpArrowNotation

class TestIt(unittest.TestCase):
	def test(self):
		self.assertEqual(UpArrowNotation.arrow(2, 4), 16)

if __name__ == '__main__':
	unittest.main()