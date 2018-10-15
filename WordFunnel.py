'''This program will solve the problem:
https://old.reddit.com/r/dailyprogrammer/comments/98ufvz/20180820_challenge_366_easy_word_funnel_1/
'''

def main():
    return funnel("leave", "eave")

def funnel(word, contains):
    for index in range(len(word)):
        testWord = word[:index] + word[index+1:]
        if testWord == contains:
            return True

    return False

if __name__ == '__main__':
    print(main())