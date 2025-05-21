package tests

import (
	"backend/controllers"
	"reflect"
	"testing"
)

func TestOptymalizacja(t *testing.T) {
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
	expected := [][]int{
		{10, 0, 10, 0},
		{0, 28, 0, 2},
		{0, 0, 17, 48},
	}
	result := controllers.Optymalizuj(planTransportu, tabelaKosztow)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("zle czat zleee")
	}
}
