#!/usr/bin/python3

import random
from ASGlossary import questionlist, answerlist

number = random.randint(1,164)
know = []
no_know = []

def question():
  question = print(f"What is the definition of {questionlist[number]}?")
  MCQ = [['a',random.choice(questionlist)],
         ['b',random.choice(questionlist)],
         ['c',random.choice(questionlist)],
         ['d',random.choice(questionlist)]]
def casual():
  question()

def competitive():
  print("competitive")

def choose():
  while True:
    choice = input("""
                   Would you like to practice flashcards (inf lives) or go for a high score?                       (3iives).
                   [a] casual
                   [b] competitive
                   [exit] exit
                   """)
    if choice == 'a':
       casual()
       break
    elif choice == 'b':
       competitive()
       break
    elif choice == 'exit':
      return
choose()
