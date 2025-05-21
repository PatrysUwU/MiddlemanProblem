package controllers

import "fmt"

type Dostawca struct {
	Podaz int
	Koszt int
}
type Odbiorca struct {
	Popyt    int
	Przychod int
}

func Macierz_zyskow_j(odbiorcy []Odbiorca, dostawcy []Dostawca, kosztyTransportu [][]int) [][]int {
	result := make([][]int, len(dostawcy))
	for i := range result {
		result[i] = make([]int, len(odbiorcy))
		for j := range result[i] {
			zysk := odbiorcy[j].Przychod - dostawcy[i].Koszt - kosztyTransportu[i][j]
			result[i][j] = zysk
		}
	}
	return result
}

func DodajFikcyjnych(odbiorcy []Odbiorca, dostawcy []Dostawca, tabelaZyskow [][]int) ([]Odbiorca, []Dostawca, [][]int) {
	sumPopyt := 0
	sumPodaz := 0
	for _, v := range dostawcy {
		sumPopyt += v.Podaz
	}
	for _, v := range odbiorcy {
		sumPodaz += v.Popyt
	}

	odbiorcy = append(odbiorcy, Odbiorca{Popyt: sumPopyt, Przychod: 0})
	dostawcy = append(dostawcy, Dostawca{Podaz: sumPodaz, Koszt: 0})

	tabelaZyskow = append(tabelaZyskow, make([]int, len(odbiorcy)-1))
	for i := range tabelaZyskow {
		tabelaZyskow[i] = append(tabelaZyskow[i], 0)
	}
	return odbiorcy, dostawcy, tabelaZyskow
}

type para struct {
	A int
	B int
}

func TabelaPrzewozow(odbiorcy []Odbiorca, dostawcy []Dostawca, tabelaKosztow [][]int) [][]int {
	var seen []para
	result := make([][]int, len(dostawcy))
	for i := range result {
		result[i] = make([]int, len(odbiorcy))
	}

	for {
		// Sprawdź czy koniec: wszystkie popyty i podaż wyzerowane
		done := true
		for _, o := range odbiorcy {
			if o.Popyt > 0 {
				done = false
				break
			}
		}
		for _, d := range dostawcy {
			if d.Podaz > 0 {
				done = false
				break
			}
		}
		if done {
			break
		}

		// Znajdź maksymalny zysk
		_, pos := findMax(tabelaKosztow, seen)

		i := pos.A
		j := pos.B

		if odbiorcy[j].Popyt == 0 || dostawcy[i].Podaz == 0 {
			seen = append(seen, pos)
			continue
		}

		q := min(odbiorcy[j].Popyt, dostawcy[i].Podaz)

		result[i][j] = q
		dostawcy[i].Podaz -= q
		odbiorcy[j].Popyt -= q

		if dostawcy[i].Podaz == 0 || odbiorcy[j].Popyt == 0 {
			seen = append(seen, pos)
		}
	}

	return result
}

func findMax(tabela [][]int, skip []para) (int, para) {
	tempTabela := copyTable(tabela)
	temp := -999999
	indexes := para{0, 0}
	for _, v := range skip {
		tempTabela[v.A][v.B] = -999999
	}

	for i, row := range tempTabela {
		for j := range row {
			if tempTabela[i][j] > temp {
				temp = tempTabela[i][j]
				indexes.A = i
				indexes.B = j
			}
		}
	}
	return temp, indexes
}

func copyTable(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func ZmienneDualne(plan, koszty [][]int) ([]int, []int) {
	m := len(plan)
	n := len(plan[0])

	alpha := make([]int, m)
	beta := make([]int, n)

	for i := range alpha {
		alpha[i] = -999
	}
	for j := range beta {
		beta[j] = -999
	}

	alpha[m-1] = 0

	zmiana := true
	for zmiana {
		zmiana = false
		for i := range m {
			for j := range n {
				if plan[i][j] > 0 {
					if alpha[i] != -999 && beta[j] == -999 {
						beta[j] = koszty[i][j] - alpha[i]
						zmiana = true
					} else if beta[j] != -999 && alpha[i] == -999 {
						alpha[i] = koszty[i][j] - beta[j]
						zmiana = true
					}
				}
			}
		}
	}

	return alpha, beta
}

func WyznaczDelty(plan [][]int, koszty [][]int, alpha, beta []int) ([][]int, para, int, bool) {
	n := len(plan)
	m := len(plan[0])
	delta := make([][]int, n)
	for i := range delta {
		delta[i] = make([]int, m)
	}

	maxDelta := -1
	maxDeltaPos := para{-1, -1}
	improved := false

	for i := range n {
		for j := range m {
			if plan[i][j] == 0 {
				d := koszty[i][j] - alpha[i] - beta[j]
				delta[i][j] = d
				if d > maxDelta {
					maxDelta = d
					maxDeltaPos = para{i, j}
				}
				if d > 0 {
					improved = true
				}
			}
		}
	}
	return delta, maxDeltaPos, maxDelta, improved
}

func ZnajdzCykl(plan [][]int, start para) []para {
	i := start.A
	j := start.B

	for y := 0; y < len(plan[0]); y++ {
		if y == j || plan[i][y] == 0 {
			continue
		}
		for x := 0; x < len(plan); x++ {
			if x == i || plan[x][y] == 0 {
				continue
			}
			if plan[x][j] > 0 {
				// Mamy cykl: start -> (i, y) -> (x, y) -> (x, j) -> start
				return []para{
					{i, j}, // +
					{i, y}, // -
					{x, y}, // +
					{x, j}, // -
				}
			}
		}
	}
	return nil
}

func PoprawPlan(plan [][]int, cykl []para) {
	min := 999999999999 // max int
	// Pola z minusem to: 1, 3, 5, ...
	for i := 1; i < len(cykl); i += 2 {
		p := cykl[i]
		if plan[p.A][p.B] < min {
			min = plan[p.A][p.B]
		}
	}

	// Aktualizacja
	for i, p := range cykl {
		if i%2 == 0 {
			plan[p.A][p.B] += min
		} else {
			plan[p.A][p.B] -= min
		}
	}
}

func Optymalizuj(plan [][]int, koszty [][]int) [][]int {
	for {
		alpha, beta := ZmienneDualne(plan, koszty)
		_, najLepszaPara, _, poprawa := WyznaczDelty(plan, koszty, alpha, beta)

		if !poprawa {
			break
		}

		cykl := ZnajdzCykl(plan, najLepszaPara)
		if cykl == nil {
			fmt.Println("Nie znaleziono cyklu – błąd")
			break
		}
		PoprawPlan(plan, cykl)

	}
	return plan
}
