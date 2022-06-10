package main

import (
	"log"
	"net/http"

	"purchase/gapi"
	"purchase/util"

	"github.com/gin-gonic/gin"
)

func main() {
	runGin()
}

func runGin() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	router := gin.Default()
	router.POST("/purchase", purchase)
	router.Run(config.PurchaseHttpAddress)
}

type purchaseRequest struct {
	UserId int64 `form:"userid" binding:"required,min=1"`
	BookId int64 `form:"bookid" binding:"required,min=1"`
	Amount int32 `form:"amount" binding:"required,min=1"`
}

func purchase(ctx *gin.Context) {
	var req purchaseRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := gapi.Grpc_createPurchase(req.UserId, req.BookId, req.Amount)
	if err != nil {
		log.Fatalf("Failed to create order because: %v\n", err)
	}
	ctx.JSON(http.StatusOK, order)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
