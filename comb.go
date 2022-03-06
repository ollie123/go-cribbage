package main

import "errors"

var ErrBreak = errors.New("break")

func Combs(n, r int, f func([]int) error) error {
	var i int
	comb := make([]int, r)
	for i = 0; i < r; i++ {
		comb[i] = i
	}
	if err := f(comb); err != nil {
		if err == ErrBreak {
			err = nil
		}
		return err
	}
	for {
		done := true
		for i = r - 1; i >= 0; i-- {
			if comb[i] != i+n-r {
				done = false
				break
			}
		}
		if done {
			return nil
		}
		comb[i]++
		for j := i + 1; j < r; j++ {
			comb[j] = comb[j-1] + 1
		}
		if err := f(comb); err != nil {
			if err == ErrBreak {
				err = nil
			}
			return err
		}
	}
}
