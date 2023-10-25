package models

import "go.uber.org/fx"

var Module = fx.Module(
	"models",
	fx.Provide(
		NewMigration,
	),
)
