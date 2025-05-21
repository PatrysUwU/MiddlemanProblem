package tests

import (
	"backend/controllers"
	"reflect"
	"testing"
)

func TestFikcyjni(t *testing.T) {
	odbiorcy := []controllers.Odbiorca{
		{Przychod: 30, Popyt: 10},
		{Przychod: 25, Popyt: 28},
		{Przychod: 30, Popyt: 27},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 20, Koszt: 10},
		{Podaz: 30, Koszt: 12},
	}

	tabelaKosztow := [][]int{
		{12, 1, 3},
		{6, 4, -1},
	}

	expected := [][]int{
		{12, 1, 3, 0},
		{6, 4, -1, 0},
		{0, 0, 0, 0},
	}
	odbiorcy, dostawcy, result := controllers.DodajFikcyjnych(odbiorcy, dostawcy, tabelaKosztow)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("zle tabela")
	}
	if odbiorcy[3].Popyt != 50 {
		t.Errorf("zle odbiorca")
	}
	if dostawcy[2].Podaz != 65 {
		t.Errorf("zle dostawca")
	}
}
