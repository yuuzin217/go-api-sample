package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseItemID(ctx *gin.Context) (uint64, error) {
	return strconv.ParseUint(ctx.Param("id"), 10, 64)
}
