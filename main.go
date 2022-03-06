package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func score(hand Hand, cut Card) {
	var total int
	sc := Score(hand, cut)
	sort.Sort(sc)
	fmt.Printf("Score for %s (%s)\n============\n", hand.ColorString(), cut.ColorString())
	for _, s := range sc {
		total += s.Score()
		fmt.Printf("%7s (%2d): %s\n", s.Type(), s.Score(), s.Hand().ColorString())
	}
	fmt.Printf("============\n  Total (%2d)\n", total)
}

func main() {
	var hand Hand
	var cut Card
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter hand: ")
		if scanner.Scan() {
			if hand, err = ParseHand(scanner.Text()); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			break
		}
		fmt.Print("Enter cut card: ")
		if scanner.Scan() {
			if cut, err = ParseCard(scanner.Text()); err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			break
		}
		fmt.Println()
		score(hand, cut)
		fmt.Println()
	}
}
