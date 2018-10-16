'''https://old.reddit.com/r/dailyprogrammer/comments/8xbxi9/20180709_challenge_365_easy_uparrow_notation/'''

def main():
	return arrow(2,4)

def arrow(base_number, iterations):
	result = 1
	for count in range(iterations):
		result = result * base_number
	print(result)
	return result

if __name__ == '__main__':
	main()