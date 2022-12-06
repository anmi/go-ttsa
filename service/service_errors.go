package service

import (
	api "github.com/anmi/go-ttsa/api"
)

type ServiceErrorsType struct {
}

func (s ServiceErrorsType) Unauthorized() *api.ErrorUnauthorized {
	return &api.ErrorUnauthorized{
		Message: "Unauthorized",
	}
}

func (s ServiceErrorsType) Forbidden() *api.ErrorForbidden {
	return &api.ErrorForbidden{
		Message: "Forbidden",
	}
}

func (s ServiceErrorsType) Circular() *api.ErrorCircular {
	return &api.ErrorCircular{
		Message: "Circular",
	}
}

var ServiceErrors ServiceErrorsType = ServiceErrorsType{}
