package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Log       *logrus.Logger
	DB        *gorm.DB
	JWTSecret = []byte("your-secret-key")
)
