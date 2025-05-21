package controllers

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
	temp := -999999
	indexes := para{0, 0}
	for _, v := range skip {
		tabela[v.A][v.B] = -999999
	}

	for i, row := range tabela {
		for j := range row {
			if tabela[i][j] > temp {
				temp = tabela[i][j]
				indexes.A = i
				indexes.B = j
			}
		}
	}
	return temp, indexes
}
