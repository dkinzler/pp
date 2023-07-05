package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	// can match players[i] <= trainers[j]
	// max number of matches, each player and trainer can only be used once
	// first sort players and trainers
	// lets go through players 1 by 1 and match
	// suppose we have already matched some
	// we want to match with the lowest trainer possible
	// could it ever make sense to not match a player?
	// if we instead matched a higher player we might need a higher trainer as well
	// and thus the original trainer would go unmatched
	// so greedy strategy should work here to give the optimal result
	sort.Ints(players)
	sort.Ints(trainers)

	result := 0
	j := 0
	for _, p := range players {
		for j < len(trainers) && trainers[j] < p {
			j += 1
		}
		if j < len(trainers) {
			result += 1
			j += 1
		} else {
			break
		}
	}
	return result
}
