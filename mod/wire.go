// +build wireinject

package main

func initApplication(ctx context.Context) all.Repository{
	panic(wire.Build(
		
	))
}