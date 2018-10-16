'''https://old.reddit.com/r/dailyprogrammer/comments/8s0cy1/20180618_challenge_364_easy_create_a_dice_roller/'''
from random import *
import datetime

def main():
	'''roll two six sided dice'''
	print(roll(2,6))

def roll(number, sides):
	result = []
	seed(datetime.datetime.now())
	for number in range(number):
		result.append(randrange(sides))

	return result

if __name__ == '__main__':
    main()