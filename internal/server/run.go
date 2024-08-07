package server

import "test/internal/server/builder"

func Run() {
	server := new(builder.Server)
	director := builder.ServerDirector{
		Builder: &builder.ServerBuilder{Product: server},
	}

	director.Construct()
}
