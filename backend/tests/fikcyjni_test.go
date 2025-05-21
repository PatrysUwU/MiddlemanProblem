package tests

import (
	"backend/controllers"
	"reflect"
	"testing"
)

func TestFikcyjni1(t *testing.T) {
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

func TestFikcyjni2(t *testing.T) {
	odbiorcy := []controllers.Odbiorca{
		{Przychod: 12, Popyt: 30},
		{Przychod: 13, Popyt: 30},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 45, Koszt: 6},
		{Podaz: 25, Koszt: 7},
	}

	tabelaKosztow := [][]int{
		{-1, 3},
		{2, 1},
	}

	expected := [][]int{
		{-1, 3, 0},
		{2, 1, 0},
		{0, 0, 0},
	}

	odbiorcy, dostawcy, result := controllers.DodajFikcyjnych(odbiorcy, dostawcy, tabelaKosztow)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("zle tabela")
	}
	if odbiorcy[2].Popyt != 70 {
		t.Errorf("zle odbiorca")
	}
	if dostawcy[2].Podaz != 60 {
		t.Errorf("zle dostawca")
	}
}
