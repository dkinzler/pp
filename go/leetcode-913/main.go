package main

import "fmt"

func main() {
	graph := [][]int{{2, 3}, {3, 4}, {0, 4}, {0, 1}, {1, 2}}
	fmt.Println(catMouseGame(graph))
	graph = [][]int{{2, 3, 4}, {2, 4}, {0, 1, 4}, {0, 4}, {0, 1, 2, 3}}
	fmt.Println(catMouseGame(graph))
}

func catMouseGame(graph [][]int) int {
	// can represent a state as the two nodes the cat and mouse are on
	// and 0/1 depending on whose turn it is
	// initial state is (1,2,0)
	// a terminal state is (0,x,1) or (x,x, 0/1)
	// there are <= 50*50*2 = 5k states
	n := len(graph)
	hasZeroNb := make([]bool, n)
	for i := 0; i < n; i++ {
		for _, x := range graph[i] {
			if x == 0 {
				hasZeroNb[i] = true
				break
			}
		}
	}

	// counts for each state the number of neighbors where mouse/cat wins
	m := map[S]int{}
	c := map[S]int{}

	mouseWin := map[S]bool{}
	catWin := map[S]bool{}

	q := make([]S, 0)
	for i := 1; i < n; i++ {
		s := S{0, i, 1}
		q = append(q, s)
		mouseWin[s] = true
		s = S{i, i, 0}
		q = append(q, s)
		catWin[s] = true
		s = S{i, i, 1}
		q = append(q, s)
		catWin[s] = true
	}

	seen := make(map[S]bool)
	for _, s := range q {
		seen[s] = true
	}

	for len(q) > 0 {
		curr := q[0]
		var node int
		var ns S
		if curr.turn == 1 {
			node = curr.mouse
			ns.turn = 0
			ns.cat = curr.cat
		} else {
			node = curr.cat
			ns.turn = 1
			ns.mouse = curr.mouse
		}
		_, mouseWins := mouseWin[curr]
		q = q[1:]
		for _, nb := range graph[node] {
			if ns.turn == 0 {
				ns.mouse = nb
			} else {
				ns.cat = nb
			}
			if ns.mouse == 0 || ns.cat == 0 || ns.mouse == ns.cat || seen[ns] {
				continue
			}
			if mouseWins {
				m[ns] += 1
			} else {
				c[ns] += 1
			}
			if ns.turn == 0 {
				if mouseWins {
					mouseWin[ns] = true
					q = append(q, ns)
					seen[ns] = true
				} else {
					// mouse's turn
					// so neighbors are all those were cat has to move
					// (x, ns.cat==curr.cat, 1), 0 is allowed
					if c[ns] == len(graph[nb]) {
						catWin[ns] = true
						q = append(q, ns)
						seen[ns] = true
					}
				}
			} else {
				if mouseWins {
					// cats's turn
					// so neighbors are all those were mouse has to move
					// (ns.mouse==curr.mouse, x, 1), 0 is not allowed
					target := len(graph[nb])
					if hasZeroNb[nb] {
						target -= 1
					}
					if m[ns] == target {
						mouseWin[ns] = true
						q = append(q, ns)
						seen[ns] = true
					}
				} else {
					catWin[ns] = true
					q = append(q, ns)
					seen[ns] = true
				}
			}
		}
	}

	initialState := S{1, 2, 0}
	if _, ok := mouseWin[initialState]; ok {
		return 1
	} else if _, ok := catWin[initialState]; ok {
		return 2
	}

	return 0
}

type S struct {
	mouse int
	cat   int
	turn  int
}
