'''https://old.reddit.com/r/dailyprogrammer/comments/8xbxi9/20180709_challenge_365_easy_uparrow_notation/'''

def main():
	return arrow(2,4)

def arrow(base_number, iterations, number_of_arrows = 1):
	if number_of_arrows == 1:
		return base_number ** iterations
	result = base_number
	for item in range(iterations - 1):
		result = arrow(base_number, result, number_of_arrows -1)
	print(result)
	return result

if __name__ == '__main__':
	main()