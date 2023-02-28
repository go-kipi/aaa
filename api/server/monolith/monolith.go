/* This file is auto-generated and should not be modified */

package monolith

import (
	"context"
	"github.com/go-kipi/aaa/api"
	"github.com/go-kipi/aaa/domain/application"
	"github.com/orchestd/dependencybundler/interfaces/validations"
	. "github.com/orchestd/servicereply"
)

type AaaInterface struct {
	aaaApp     application.AaaApp
	validation validations.Validations
}

func NewAaaInterface(aaaApp application.AaaApp, validation validations.Validations) AaaInterface {
	return AaaInterface{aaaApp: aaaApp, validation: validation}
}

func (i AaaInterface) NewPath(c context.Context, req api.NewPathRequest) ServiceReply {
	if vErr := i.validation.Validate(req); vErr != nil && !vErr.IsSuccess() {
		return vErr
	}
	err := i.aaaApp.NewPath(c)
	return err
}
