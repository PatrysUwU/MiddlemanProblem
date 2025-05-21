package tests

import (
	"backend/controllers"
	"reflect"
	"testing"
)

func TestDualne1(t *testing.T) {
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

	expectedAlpha := []int{3, -1, 0}
	expectedBeta := []int{9, 5, 0, 0}

	alpha, beta := controllers.ZmienneDualne(planTransportu, tabelaKosztow)
	if !reflect.DeepEqual(alpha, expectedAlpha) {
		t.Errorf("zle alfa")
	}
	if !reflect.DeepEqual(beta, expectedBeta) {
		t.Errorf("zle beta")
	}
}

func TestDualne2(t *testing.T) {
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

	expectedAlpha := []int{3, 0, 0}
	expectedBeta := []int{9, 4, 0, 0}

	alpha, beta := controllers.ZmienneDualne(planTransportu, tabelaKosztow)
	if !reflect.DeepEqual(alpha, expectedAlpha) {
		t.Errorf("zle alfa")
	}
	if !reflect.DeepEqual(beta, expectedBeta) {
		t.Errorf("zle beta")
	}
}

func TestDualne3(t *testing.T) {
	planTransportu := [][]int{
		{0, 30, 15},
		{25, 0, 0},
		{5, 0, 55},
	}

	tabelaKosztow := [][]int{
		{-1, 3, 0},
		{2, 1, 0},
		{0, 0, 0},
	}

	expectedAlpha := []int{0, 2, 0}
	expectedBeta := []int{0, 3, 0}

	alpha, beta := controllers.ZmienneDualne(planTransportu, tabelaKosztow)
	if !reflect.DeepEqual(alpha, expectedAlpha) {
		t.Errorf("zle alfa")
	}
	if !reflect.DeepEqual(beta, expectedBeta) {
		t.Errorf("zle beta")
	}
}
