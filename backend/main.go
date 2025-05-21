package main

import (
	"backend/controllers"
	"fmt"
)

// import "github.com/gin-gonic/gin"
func main() {
	//	router := gin.Default()
	//	router.GET("/ping", func(c *gin.Context) {
	//		c.JSON(200, gin.H{
	//			"message": "pong",
	//		})
	//	})
	//	router.Run()

	odbiorcy := []controllers.Odbiorca{
		{Przychod: 12, Popyt: 30},
		{Przychod: 13, Popyt: 30},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 45, Koszt: 6},
		{Podaz: 25, Koszt: 7},
	}

	tabelaKosztow := [][]int{
		{7, 4},
		{3, 5},
	}
	tabelaZyskowJednostkowych := controllers.Macierz_zyskow_j(odbiorcy, dostawcy, tabelaKosztow)
	fmt.Println(tabelaZyskowJednostkowych)
	odbiorcy, dostawcy, tabelaZyskowJednostkowych = controllers.DodajFikcyjnych(odbiorcy, dostawcy, tabelaZyskowJednostkowych)
	fmt.Println(tabelaZyskowJednostkowych)
	tabelaTrasportu := controllers.TabelaPrzewozow(odbiorcy, dostawcy, tabelaZyskowJednostkowych)
	fmt.Println(tabelaTrasportu)
	fmt.Println(controllers.ZmienneDualne(tabelaTrasportu, tabelaZyskowJednostkowych))
	result := controllers.Optymalizuj(tabelaTrasportu, tabelaZyskowJednostkowych)
	fmt.Println(result)
}
