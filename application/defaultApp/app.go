package defaultApp

import (
	"github.com/go-kipi/aaa/domain/application"
)

type aaaApp struct {
}

func NewAaaApp() application.AaaApp {
	return &aaaApp{}
}
