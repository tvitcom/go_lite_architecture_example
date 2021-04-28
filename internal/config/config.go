package config

import (
	"my.localhost/others/bookstore/internal/models"
)

type ConfigService struct {
	db      models.IBooks
	mongodb models.IProfile
}
