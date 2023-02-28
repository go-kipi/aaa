package repository

import (
	"github.com/go-kipi/aaa/domain/repository"
)

type aaaRepository struct {
}

func NewAaaRepository() repository.AaaRepository {
	return aaaRepository{}
}
