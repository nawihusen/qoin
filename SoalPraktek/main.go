package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DiceGame(player int, Dice int) {
	rand.Seed(time.Now().UTC().UnixNano())
	var bigger int
	var winner []int
	Score := map[int]int{}
	TotalDice := map[int]int{}
	Play := map[int]bool{}
	DiceResults := map[int][]int{}
	diceValue := []int{1, 2, 3, 4, 5, 6}

	for i := 1; i <= player; i++ {
		Score[i] = 0
		TotalDice[i] = Dice
		DiceResults[i] = []int{}
		Play[i] = true
	}

	looping := 0

outer:
	for {

		for i := 1; i <= player; i++ {
			if Play[i] {
				new := []int{}
				for j := 1; j <= TotalDice[i]; j++ {
					temp := diceValue[rand.Intn(6)]
					new = append(new, temp)
				}
				DiceResults[i] = new
			}
		}

		// Evaluasi
		for i := 1; i <= player; i++ {
			if Play[i] {
				for j := 0; j < len(DiceResults[i]); j++ {
					if DiceResults[i][j] == 6 {
						Score[i] += 1
						TotalDice[i] -= 1
					} else if DiceResults[i][j] == 1 {
						if player == i {
						inside:
							for k := 1; k <= player; k++ {
								if Play[k] {
									TotalDice[k] += 1
									TotalDice[i] -= 1
									break inside
								} else {
									TotalDice[i+1] += 1
									TotalDice[i] -= 1
									break inside
								}
							}
						} else {
						inside2:
							for k := i + 1; ; k++ {
								if k > player {
									k = 1
								}

								if Play[k] {
									TotalDice[k] += 1
									TotalDice[i] -= 1
									break inside2
								}
							}
						}
					}
				}
			}
		}

		for i := 1; i <= player; i++ {
			if TotalDice[i] == 0 {
				Play[i] = false
				DiceResults[i] = []int{}
			}
		}

		looping += 1
		fmt.Println("looping ke", looping)
		fmt.Print("\n=====================================================\n\n")

		var temp int
		for i := 1; i <= player; i++ {
			if Play[i] {
				temp += 1
			}
		}

		if temp < 2 {
			break outer
		}

	}

	for i := 1; i <= player; i++ {
		if Score[i] > bigger {
			bigger = Score[i]
		}
	}

	for i := 1; i <= player; i++ {
		if Score[i] == bigger {
			winner = append(winner, i)
		}
	}

	msg := fmt.Sprintf("Game Di Menangkan Oleh %d dengan score %d", winner, bigger)

	fmt.Println(msg)
}

func main() {
	DiceGame(3, 7)
}
