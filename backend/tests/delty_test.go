package tests

import (
	"backend/controllers"
	"testing"
)

func TestDelta1(t *testing.T) {
	planTransportu := [][]int{
		{10, 0, 10, 0},
		{0, 28, 0, 2},
		{0, 0, 17, 48},
	}

	tabelaKosztow := [][]int{
		{12, 1, 3, 0},
		{6, 4, -1, 0},
		{0, 0, 0, 0},
	}

	alpha := []int{3, 0, 0}
	beta := []int{9, 4, 0, 0}

	delta, pos, maxdelta, improved := controllers.WyznaczDelty(planTransportu, tabelaKosztow, alpha, beta)
	t.Log(delta)
	t.Log(pos)
	t.Log(maxdelta)
	t.Log(improved)
}

func TestDelta2(t *testing.T) {
	planTransportu := [][]int{
		{10, 0, 10, 0},
		{0, 28, 2, 0},
		{0, 0, 15, 50},
	}

	tabelaKosztow := [][]int{
		{12, 1, 3, 0},
		{6, 4, -1, 0},
		{0, 0, 0, 0},
	}

	alpha := []int{3, -1, 0}
	beta := []int{9, 5, 0, 0}

	delta, pos, maxdelta, improved := controllers.WyznaczDelty(planTransportu, tabelaKosztow, alpha, beta)
	t.Log(delta)
	t.Log(pos)
	t.Log(maxdelta)
	t.Log(improved)
}
