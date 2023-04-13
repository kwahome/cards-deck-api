package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/kwahome/cards-deck-api/internal/api/v1/handlers"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
)

func RegisterRoutes(router gin.IRouter) {

	deckService := service.CreateDeckService()
	createDeckHandler := handlers.NewCreateDeckHandler(deckService)
	getDeckHandler := handlers.NewGetDeckHandler(deckService)
	drawCardsHandler := handlers.NewDrawCardsHandler(deckService)

	v1 := router.Group("api/v1")
	v1.POST("/decks", createDeckHandler.CreateDeck)
	v1.GET("/decks/:id", getDeckHandler.OpenDeck)
	v1.GET("/decks/:id/draw", drawCardsHandler.DrawCards)
}
