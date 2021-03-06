import unittest
from UpArrowNotation import arrow

class TestIt(unittest.TestCase):
	def testSingleArrow(self):
		self.assertEqual(arrow(2, 4), 16)
		self.assertEqual(arrow(2, 3), 8)
		self.assertEqual(arrow(1, 0), 1)

	def testDoubleArrow(self):
		self.assertEqual(arrow(2, 4, 2), 65536)
		self.assertEqual(arrow(1, 0, 2), 1)
	def testTripleArrow(self):
		self.assertEqual(arrow(2, 3, 3), 65536)
		self.assertEqual(arrow(-1, 3, 3), -1)

if __name__ == '__main__':
	unittest.main()