import unittest
from UpArrowNotation import arrow

class TestIt(unittest.TestCase):
	def testSingleArrow(self):
		self.assertEqual(arrow(2, 4), 16)
		self.assertEqual(arrow(2, 3), 8)
	def testDoubleArrow(self):
		self.assertEqual(arrow(2, arrow(2, 4)), 65536)
	def testTripleArrow(self):
		self.assertEqual(arrow(2, arrow(2, arrow(2, 3))), 65536)

if __name__ == '__main__':
	unittest.main()