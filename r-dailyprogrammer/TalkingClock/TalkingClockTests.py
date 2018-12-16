import unittest
import TalkingClock

class TestIt(unittest.TestCase):
	def test_1205(self):
		self.assertEqual(TalkingClock.speakTime("12:05"), "It's twelve oh five pm")

if __name__ == '__main__':
    unittest.main()