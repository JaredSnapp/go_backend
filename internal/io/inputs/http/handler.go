package HTTPHandler

import (
	"strings"

	"github.com/JaredSnapp/go_backend/internal/config"
)

type Handler struct {
	Conf              *config.Config
	corsAllowedOrigin []string
	PersonsService    PersonsService
	GoalsService      GoalsService
	ActionService     ActionService
	RecruiterService  RecruiterService
}

func NewHandler(conf *config.Config, ps PersonsService, gs GoalsService, as ActionService, r RecruiterService) *Handler {

	corsOrigins := strings.Split(conf.CORSAllowedOrigin, ",")
	return &Handler{
		Conf:              conf,
		corsAllowedOrigin: corsOrigins,
		PersonsService:    ps,
		GoalsService:      gs,
		ActionService:     as,
		RecruiterService:  r,
	}
}
