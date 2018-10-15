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

def bonus(word):
    word_file = open("WordList.txt", "r")
    word_list = word_file.read()
    found_words = []

    for test_word in word_list.split():
        if funnel(word, test_word):
            found_words.append(test_word)

    return found_words

if __name__ == '__main__':
    print(main())