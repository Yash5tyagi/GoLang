package auth

import (
	"net/http"
)

type TokenValues struct {
	Session string
}
type Service interface {
	Authorize(next http.Handler) http.Handler
}

type service struct {
	//logger            logging.Service
	//config            *config.Config

}

func NewService() *service {
	return &service{}
}
