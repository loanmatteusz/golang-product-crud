package middlewares

import (
	"backend/internal/custom_errors"
	"backend/internal/dtos/helpers"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			authHeader := ctx.Request().Header.Get("Authorization")
			if authHeader == "" {
				return sendUnauthorized(ctx)
			}

			tokenSplited := strings.Split(authHeader, " ")
			if len(tokenSplited) != 2 || tokenSplited[0] != "Bearer" {
				return sendUnauthorized(ctx)
			}

			tokenStr := tokenSplited[1]
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil || !token.Valid {
				return sendUnauthorized(ctx)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return sendUnauthorized(ctx)
			}

			userID, ok := claims["user_id"].(string)
			if !ok {
				return sendUnauthorized(ctx)
			}

			if claims["type"] == "refresh" {
				return sendUnauthorized(ctx)
			}

			ctx.Set("user_id", userID)

			return next(ctx)
		}
	}
}

func sendUnauthorized(ctx echo.Context) error {
	return helpers.SendError(ctx, http.StatusUnauthorized, custom_errors.CodeInvalidToken, custom_errors.ErrInvalidToken)
}
