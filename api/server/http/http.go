/* This file is auto-generated and should not be modified */

package http

import (
	"github.com/go-kipi/aaa/api"
	"github.com/go-kipi/aaa/api/server/monolith"
	"github.com/orchestd/dependencybundler/interfaces/configuration"
	"github.com/orchestd/dependencybundler/interfaces/transport"
)

func InitHandlers(router transport.IRouter, m monolith.AaaInterface, c configuration.Config) {
	router.POST(api.NewPathMethod, transport.HandleFunc(m.NewPath))
}
