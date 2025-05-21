package tests

import (
	"backend/controllers"
	"reflect"
	"testing"
)

func TestMacierzZyskow1(t *testing.T) {
	odbiorcy := []controllers.Odbiorca{
		{Przychod: 30, Popyt: 10},
		{Przychod: 25, Popyt: 28},
		{Przychod: 30, Popyt: 27},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 20, Koszt: 10},
		{Podaz: 30, Koszt: 12},
	}

	kosztyTransportu := [][]int{
		{8, 14, 17},
		{12, 9, 19},
	}

	expected := [][]int{
		{12, 1, 3},
		{6, 4, -1},
	}

	result := controllers.Macierz_zyskow_j(odbiorcy, dostawcy, kosztyTransportu)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("zle")
	}
}

func TestMacierzZyskow2(t *testing.T) {
	odbiorcy := []controllers.Odbiorca{
		{Przychod: 12, Popyt: 30},
		{Przychod: 13, Popyt: 30},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 45, Koszt: 6},
		{Podaz: 25, Koszt: 7},
	}

	kosztyTransportu := [][]int{
		{7, 4},
		{3, 5},
	}

	expected := [][]int{
		{-1, 3},
		{2, 1},
	}

	result := controllers.Macierz_zyskow_j(odbiorcy, dostawcy, kosztyTransportu)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("zle")
	}
}
