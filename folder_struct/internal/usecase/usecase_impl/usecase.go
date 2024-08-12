package usecase_impl

import (
	"krriya/internal/core/core_impl"
	"krriya/pkg/logging"
)

type UsecaseImpl struct {
	Logs logging.Service
}

func NewUsecase(logs logging.Service) core_impl.Usecase {
	return &UsecaseImpl{

		Logs: logs,
	}
}
