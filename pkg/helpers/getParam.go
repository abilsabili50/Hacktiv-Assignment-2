package helpers

import (
	"strconv"

	"github.com/abilsabili50/Hacktiv-Assignment-2.git/pkg/errs"
	"github.com/gin-gonic/gin"
)

func GetParam(ctx *gin.Context, key string) (int, errs.MessageErr) {
	value := ctx.Param(key)

	param, err := strconv.Atoi(value)
	if err != nil {
		return 0, errs.BadRequest("invalid uri parameter")
	}

	return param, nil
}
