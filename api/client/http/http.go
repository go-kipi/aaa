/* This file is auto-generated and should not be modified */

package http

import (
	"context"
	"github.com/go-kipi/aaa/api"
	"github.com/orchestd/dependencybundler/interfaces/transport"
	. "github.com/orchestd/servicereply"
)

const serviceName = "aaa"

type aaaHTTPClient struct {
	client transport.HttpClient
}

func NewAaaApiHTTPClient(client transport.HttpClient) api.AaaApi {
	return aaaHTTPClient{client: client}
}

func (h aaaHTTPClient) NewPath(c context.Context, req api.NewPathRequest) ServiceReply {
	if reply := h.client.Call(c, req, serviceName, api.NewPathMethod, nil, nil); !reply.IsSuccess() {
		return reply
	}
	return nil
}
