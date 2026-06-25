//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func IntializeService() *SimpleService {
	wire.Build(
		NewSimpleRepository,
		NewSimpleService,
	)
	return nil
}
