package data

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserRepository,
	NewS3Repository,
)
