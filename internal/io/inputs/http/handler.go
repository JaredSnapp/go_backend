package HTTPHandler

import (
	"strings"

	"github.com/JaredSnapp/go_backend/internal/config"
)

type Handler struct {
	Conf              *config.Config
	corsAllowedOrigin []string
	PersonsService    PersonsService
	RecruiterService  RecruiterService
	RichTextService   RichTextService
}

func NewHandler(conf *config.Config, ps PersonsService, r RecruiterService, rt RichTextService) *Handler {

	corsOrigins := strings.Split(conf.CORSAllowedOrigin, ",")
	return &Handler{
		Conf:              conf,
		corsAllowedOrigin: corsOrigins,
		PersonsService:    ps,
		// GoalsService:      gs,
		// ActionService:     as,
		RecruiterService: r,
		RichTextService:  rt,
	}
}
