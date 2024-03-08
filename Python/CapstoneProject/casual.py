#!/usr/bin/python3

import random
from ASGlossary import questionlist
from ASAnswer import answerlist

def casual():
    ans =[]
    answ = 0
    streak = 0
    encouragement = ['You are on fire!','Keep going!','spectacular!','Supercalifragilisticexpialidocious work!','You work at google or something?']
    while True:
        number = random.randint(0,len(questionlist))
        print("type 'exit' to exit")
        question = print(f"{'':>10s}What is the definition of {questionlist[number]}?\n")
        ans.append(answerlist[number])
        for i in range(3):
            randomnum = random.randint(0,len(questionlist)-1)
            ans.append(answerlist[randomnum])
            random.shuffle(ans)
        while True:
            inputans =input(f"""
            Pick answer
            A : {ans[0]}
            B : {ans[1]}
            C : {ans[2]}
            D : {ans[3]}
            """)
            if inputans == 'A':
                outputans = ans[0]
                break
            elif inputans == 'B':
                outputans = ans[1]
                break
            elif inputans == 'C':
                outputans = ans[2]
                break
            elif inputans == 'D':
                outputans = ans[3]
                break
            elif inputans == 'exit':
                print(f"finished with a streak of {streak} and {answ} flashcards answered")
                return
            else:
                print("input error, did you make a typo? (case sensitive)")
        if outputans == answerlist[number]:
            print("Correct!")
            answ = answ+1
            ans.clear()
            streak = streak +1
            if streak%5 == 0:
                print(f"{random.choice(encouragement):>16s}, you're on a streak of {streak}!")
        else:
            print("no worries!")
            answ = answ+1
            streak = 0
