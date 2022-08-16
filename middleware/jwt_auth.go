package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/services"
)

func AuthorizeJWT(isForAdminOnly bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER = "Bearer "
		header := ctx.GetHeader("Authorization")
		rawToken := header[len(BEARER):]

		authServ, err := services.NewJsonWebTokenService()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   true,
				"message": "Unable to validate authorization token due to a server error. Please submit a support ticket.",
			})
			return
		}

		claims, err := authServ.ValidateAndParseToken(rawToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		if isForAdminOnly && !claims.IsAdmin {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "User does not have permission to perform this action.",
			})

		}
		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
}
