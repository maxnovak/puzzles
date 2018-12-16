'''Talking clock challenge
https://old.reddit.com/r/dailyprogrammer/comments/6jr76h/20170627_challenge_321_easy_talking_clock/'''

def main():
	return speakTime("12:05")

def speakTime(time):
	spokenTime = constructHours(time) + " "
	spokenTime += constructMin(time) + " "
	spokenTime += constructMeridiem(time)
	return "It's " + spokenTime

def constructHours(time):
	return "twelve"

def constructMin(time):
	return "oh five"

def constructMeridiem(time):
	return "pm"

if __name__ == '__main__':
    print(main())