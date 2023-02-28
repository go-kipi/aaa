package main

import (
	"github.com/go-kipi/aaa/api/server/monolith"
	"github.com/go-kipi/aaa/application/defaultApp"
)

func deps() []interface{} {
	return append(internalDeps(), externalDeps()...)
}

func internalDeps() []interface{} {
	return []interface{}{
		defaultApp.NewAaaApp,
		monolith.NewAaaInterface,
	}
}

func externalDeps() []interface{} {
	return []interface{}{}
}
