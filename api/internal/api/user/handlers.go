package user

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/pkg/passwords"
	"github.com/nixoncode/skillflow/pkg/response"
)

func (uh *UserHandler) Register(c echo.Context) error {
	uh.app.Log().Info().Msg("Register endpoint hit")
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to bind register request")
		return response.BadRequest(c, "Invalid request payload, check documentation")
	}

	if err := req.Validate(); err != nil {
		uh.app.Log().Warn().Err(err).Msg("Validation failed for register request")
		return response.ValidationError(c, err)
	}

	// can be moved to validation
	emailExists, err := uh.EmailExists(req.Email)
	if err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to check existing email")
		return response.InternalServerError(c, "Failed to process request")
	}
	if emailExists {
		return response.Error(c, http.StatusConflict, "Email already registered")
	}

	newUser := &User{
		Email:     req.Email,
		Role:      req.Role,
		CreatedAt: time.Now(),
	}
	hashedPassword, err := passwords.HashPassword(req.Password)
	if err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to hash password")
		return response.InternalServerError(c, "Failed to process password")
	}
	newUser.PasswordHash = hashedPassword

	err = uh.CreateUser(newUser)
	if err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to create user")
		return response.InternalServerError(c, "Failed to create user")
	}

	return response.Ok(c, "Registration successful", newUser)
}

func (uh *UserHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to bind login request")
		return response.BadRequest(c, "Invalid request payload, check documentation")
	}

	if err := req.Validate(); err != nil {
		uh.app.Log().Warn().Err(err).Msg("Validation failed for login request")
		return response.ValidationError(c, err)
	}

	user, err := uh.GetUserByEmail(req.Email)
	if err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to fetch user by email")
		return response.InternalServerError(c, "Failed to process request")
	}
	if user == nil || !passwords.CheckPasswordHash(req.Password, user.PasswordHash) {
		return response.Error(c, http.StatusUnauthorized, "Invalid email or password")
	}

	token, err := GenerateToken(user.ID, user.Role, uh.app.Config().JWT.SecretKey, uh.app.Config().JWT.ExpirationMins)
	if err != nil {
		uh.app.Log().Error().Err(err).Msg("Failed to generate JWT token")
		return response.InternalServerError(c, "Failed to generate authentication token")
	}

	return response.Ok(c, "Login successful", map[string]string{"token": token})
}

func (uh *UserHandler) Logout(c echo.Context) error {
	// Placeholder implementation
	return response.Ok(c, "Logout successful", nil)
}

func (uh *UserHandler) Profile(c echo.Context) error {
	user := User{
		ID:    12345,
		Email: "me@example.com",
	}
	return response.Ok(c, "User profile fetched successfully", user)
}
