//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func IntializeService() (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil, nil
}
