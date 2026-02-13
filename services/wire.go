package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserService,
	NewS3Service,
)
