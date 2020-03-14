package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/core/service/gamesrv"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/handlers/gamehdl"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/internal/repositories/gamesrepo"
	"github.com/matiasvarela/minesweeper-hex-arch-sample/pkg/uidgen"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehdl.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)
	router.PUT("/games/:id", gamesHandler.RevealCell)

	router.Run(":8080")
}