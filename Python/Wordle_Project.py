#WORK IN PROGRESS!!!

#!/bin/python3

#word list
#make a 5 letter array
#tell you which words are in there and which words are wrong
#give six attempts

import random
player_guess = ['-','-','-','-','-']

wordlist = ['whizz','pizza','joker']
word = []
string = random.choice(wordlist)
for letter in string:
    word.append(letter)
#chooses word from wordlist and splits into characters

Correct = 0
attempts = 0
x = 0

while x == 0:
    guess = input(f'guess a 5 letter word! {6 - attempts} attempts given')
    if len(guess) != 5:
        print("out of range")
    attempts = attempts + 1
    if attempts > 6:
        print("you lose!")
        break
    for i in range(5):
        player_guess[i] = guess[i:i+1]
    print(player_guess)
    for i in range(5):
        for j in range(5):
            if player_guess[i] == word[j]:
                print(f"{player_guess[i]} is present")
        if player_guess[i] == word[i]:
            Correct = Correct + 1
            print(f"{player_guess[i]} in position {i+1} is correct")
            if Correct == 5:
                print(f"you win! ({attempts}) attempts ")
                print(f"word was {string}")
                x = x+1
        else:
            Correct = 0



