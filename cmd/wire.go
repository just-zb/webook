//go:build wireinject

package cmd

import "github.com/google/wire"

func initApplication() {
	panic(wire.Build())
}
