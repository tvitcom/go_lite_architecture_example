package config

import (
	"my.localhost/others/bookstore/internal/model"
)

type ConfigService struct {
	db      model.IBooks
	mongodb model.IProfile
}
