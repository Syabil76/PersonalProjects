#!/usr/bin/python3

import random
from ASGlossary import questionlist, answerlist

Score = 0

def question():
  no_knowq = []
  no_knowa = []
  ans =[]
  # while True:
  number = random.randint(1,164)
  question = print(f"What is the definition of {questionlist[number]}?")
  ans.append(answerlist[number])
  for i in range(3):
    randomnum = random.randint(1,164)
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
    else:
      print("input error, did you make a typo? (case sensitive)")
  if outputans == answerlist[number]:
    print("Correct!")
    Score = Score + 1
    print(Score)
  else:
    choose2 = input("try again (a) or exit? (b)")
    while True:
      if choose2 == 'a':
        no_knowq.append(questionlist[number])
        no_knowa.append(answerlist[number])
        if len(no_knowq) > 10:
          print(f"""You seem like you're getting these questions wrong...
                {no_knowq}
                """)
          break
        elif choose2 == 'b':
          print(f'goodbye, finished with a score of {Score}')
          return
        else:
          print("input error, did you make a typo? (case sensitive)")

def competitive():
  print("competitive")

def choose():
  while True:
    choice = input("""
                   Would you like to practice flashcards (inf lives) or go for a high score? (3 lives).
                   [a] casual
                   [b] competitive
                   [exit] exit
                   """)
    if choice == 'a':
       question()
       break
    elif choice == 'b':
       competitive()
       break
    elif choice == 'exit':
      return
choose()
