package user

import (
	"github.com/itpark/market/auth/internal/service"
)

type UserHandler struct {
	Service *service.UserService
}
