package setting

import (
	"blog-service/pkg/logger"
	"github.com/spf13/viper"
)

type Setting struct {
	vp     *viper.Viper
	Logger *logger.Logger
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp, }, nil
}
