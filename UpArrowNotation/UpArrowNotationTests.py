import unittest
import UpArrowNotation

class TestIt(unittest.TestCase):
	def testSingleArrow(self):
		self.assertEqual(UpArrowNotation.arrow(2, 4), 16)
	def testDoubleArrow(self):
		self.assertEqual(UpArrowNotation.arrow(2, UpArrowNotation.arrow(2, 4)), 65536)

if __name__ == '__main__':
	unittest.main()