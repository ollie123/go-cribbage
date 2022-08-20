## go-cribbage
Terminal-based score calculator for the Cribbage card game. Scores are calculated assuming a standard two-player game as described here: https://en.wikipedia.org/wiki/Rules_of_cribbage

Example usage:
```
$ ./go-cribbage 
Enter hand: 6D JH 4H 7C
Enter cut card: 5H

Score for 6♦, J♥, 4♥, 7♣ (5♥)
============
Fifteen ( 2): 5♥, J♥
Fifteen ( 2): 4♥, 5♥, 6♦
    Run ( 4): 4♥, 5♥, 6♦, 7♣
   Nobs ( 1): J♥
============
  Total ( 9)
```