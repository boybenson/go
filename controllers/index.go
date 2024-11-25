package controllers

import (
	"example/go/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetPoems(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message":"All Poems",
	})
}

func HandleCreatePoem(ctx *gin.Context) {
	var body types.Poem

	if err := ctx.BindJSON(&body); err != nil {

	}

	fmt.Println(body.Genre)
}

func HandleGetPoem(ctx *gin.Context) {
	println(ctx.Params)
}

func HandleDeletePoem(ctx *gin.Context) {
	println(ctx.Params)
}