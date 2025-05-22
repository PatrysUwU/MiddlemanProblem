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
		{Przychod: 30, Popyt: 10},
		{Przychod: 25, Popyt: 28},
		{Przychod: 30, Popyt: 27},
	}

	dostawcy := []controllers.Dostawca{
		{Podaz: 20, Koszt: 10},
		{Podaz: 30, Koszt: 12},
	}

	tabelaKosztow := [][]int{
		{8, 14, 17},
		{12, 9, 19},
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
	zyskCalkowity, przychod, koszt := controllers.ObliczZysk(odbiorcy, dostawcy, tabelaZyskowJednostkowych, result)
	fmt.Println(zyskCalkowity)
	fmt.Println(przychod)
	fmt.Println(koszt)
}
