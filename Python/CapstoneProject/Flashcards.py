#!/usr/bin/python3

import random
from ASGlossary import questionlist
from ASAnswer import answerlist
from casual import casual
from competitive import competitive

def choose():
  while True:
    choice = input("""
                   Would you like to practice flashcards (inf lives) or go for a high score? (3 lives).
                   [a] practice
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
