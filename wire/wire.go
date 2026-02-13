//go:build wireinject
// +build wireinject

package wire

import (
	"s3-poc/config"
	"s3-poc/data"
	"s3-poc/handlers"
	"s3-poc/services"

	"github.com/google/wire"
)

func InitializeUserHandler() (*handlers.UserHandler, error) {
	wire.Build(
		config.ProviderSet,
		data.ProviderSet,
		services.ProviderSet,
		handlers.ProviderSet,
	)
	return &handlers.UserHandler{}, nil
}

func InitializeS3Handler() (*handlers.S3Handler, error) {
	wire.Build(
		config.ProviderSet,
		data.ProviderSet,
		services.ProviderSet,
		handlers.ProviderSet,
	)
	return &handlers.S3Handler{}, nil
}
