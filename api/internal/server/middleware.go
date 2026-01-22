package server

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

func useAuth(jwtSecret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(user.JwtCustomClaims)
		},
	})
}

func requireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Get("user").(*jwt.Token)
			claims := userID.Claims.(*user.JwtCustomClaims)

			if claims.Role != role {
				return response.Error(c, http.StatusForbidden, "You do not have permission to perform this action")
			}

			return next(c)
		}
	}
}

func serverLogger(s *Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			res := c.Response()
			req := c.Request()
			rid := res.Header().Get(echo.HeaderXRequestID)

			logger := s.app.Log().With().
				Str("method", req.Method).
				Str("path", req.URL.Path).
				Str("remote_ip", c.RealIP()).
				Str("user_agent", req.UserAgent()).
				Int("status", res.Status).
				Str("request_id", rid).
				Dur("latency", stop.Sub(start)).
				Logger()

			switch {
			case res.Status >= 500:
				logger.Error().Msg("request completed")
			case res.Status >= 400:
				logger.Warn().Msg("request completed")
			default:
				logger.Info().Msg("request completed")
			}

			return err
		}
	}
}
