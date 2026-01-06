package global

import (
	"github.com/shaxiaozz/sangfor-ad-exporter/config"
	"go.uber.org/zap"
)

var (
	Config         *config.App
	Logger         *zap.Logger
	SangforAdToken string
)
