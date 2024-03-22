package http_v1

import (
	"github.com/gin-gonic/gin"
)

func getFilters(ctx *gin.Context) map[string]any {
	filter := make(map[string]interface{})
	if ctx.Query("first_name") != "" {
		filter["first_name"] = ctx.Query("first_name")
	}
	if ctx.Query("second_name") != "" {
		filter["second_name"] = ctx.Query("second_name")
	}
	if ctx.Query("middle_name") != "" {
		filter["middle_name"] = ctx.Query("middle_name")
	}
	if ctx.Query("phone_number") != "" {
		filter["phone_number"] = ctx.Query("phone_number")
	}
	return filter
}
