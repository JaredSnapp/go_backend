package HTTPHandler

import (
	"strings"

	"github.com/JaredSnapp/go_backend/internal/config"
)

type Handler struct {
	Conf              *config.Config
	corsAllowedOrigin []string
	PersonsService    PersonsService
}

func NewHandler(conf *config.Config, ps PersonsService) *Handler {

	corsOrigins := strings.Split(conf.CORSAllowedOrigin, ",")
	return &Handler{
		Conf:              conf,
		corsAllowedOrigin: corsOrigins,
		PersonsService:    ps,
	}
}
