package main

import (
	"fmt"
	"sort"
)

func getLarge(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// set two int for number of cake being sold in shop and number of cake to eat
	var n, m int
	fmt.Scan(&n, &m)
	// set beauty[0], taste[1] and popularity[2] for all cakes being sold in shop
	scores := make([][]int, n)
	for i := 0; i < n; i++ {
		scores[i] = make([]int, 3)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&scores[i][0], &scores[i][1], &scores[i][2])
	}
	// we have 8 ways of chosing maxvalue
	// Max or Min of Beauty or Taste or Popularity
	var max int
	// bit全探索
	for bit := uint(0); bit < (1 << 3); bit++ {
		// store scores for each bit
		var tmpScores []int
		// store value of this bit
		var tmpMax int
		for i := 0; i < n; i++ {
			// a score of each cake in each bit
			eachScore := 0
			for j := uint(0); j < 3; j++ {
				if bit&(1<<j) == 1<<j {
					// if jth digit is 1
					eachScore += scores[i][j]
				} else {
					//if jth digit is 0
					eachScore -= scores[i][j]
				}
			}
			// append a score of a cake in that bit
			tmpScores = append(tmpScores, eachScore)
		}
		//sort cakes by their scores in descending order
		sort.Sort(sort.Reverse(sort.IntSlice(tmpScores)))
		// fetch scores from top to bottom till mth element and add them all to get the maximum value of that bit
		for i := 0; i < m; i++ {
			tmpMax += tmpScores[i]
		}
		// check this temporary max exceeds max that we held till this timepoint
		max = getLarge(max, tmpMax)
	}
	fmt.Printf("%d", max)
}
