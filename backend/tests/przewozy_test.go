package tests

import (
	"backend/controllers"
	"reflect"
	"testing"
)

func TestPrzewozy1(t *testing.T) {
	odbiorcy := []controllers.Odbiorca{
		{Przychod: 30, Popyt: 10},
		{Przychod: 25, Popyt: 28},
		{Przychod: 30, Popyt: 27},
		{Przychod: 0, Popyt: 50},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 20, Koszt: 10},
		{Podaz: 30, Koszt: 12},
		{Podaz: 65, Koszt: 0},
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

	result := controllers.TabelaPrzewozow(odbiorcy, dostawcy, tabelaKosztow)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("zle prewozy")
	}
}

func TestPrzewozy2(t *testing.T) {
	odbiorcy := []controllers.Odbiorca{
		{Przychod: 12, Popyt: 30},
		{Przychod: 13, Popyt: 30},
		{Przychod: 0, Popyt: 70},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 45, Koszt: 6},
		{Podaz: 25, Koszt: 7},
		{Podaz: 60, Koszt: 0},
	}

	tabelaKosztow := [][]int{
		{-1, 3, 0},
		{2, 1, 0},
		{0, 0, 0},
	}

	expected := [][]int{
		{0, 30, 15},
		{25, 0, 0},
		{5, 0, 55},
	}

	result := controllers.TabelaPrzewozow(odbiorcy, dostawcy, tabelaKosztow)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("zle prewozy")
	}
}
