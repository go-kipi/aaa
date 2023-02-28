/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/go-kipi/aaa/api"
	"github.com/go-kipi/aaa/api/server/monolith"
	. "github.com/orchestd/servicereply"
)

type AaaMonolithClient struct {
	MonolithInterface monolith.AaaInterface
}

func NewAaaMonolithClient(monolithInterface monolith.AaaInterface) api.AaaApi {
	return AaaMonolithClient{MonolithInterface: monolithInterface}
}

func (p AaaMonolithClient) NewPath(c context.Context, req api.NewPathRequest) ServiceReply {
	return p.MonolithInterface.NewPath(c, req)
}
