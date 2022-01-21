package svc

import (
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/internal/config"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"moul.io/zapgorm2"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	zapLogger := zapgorm2.New(zap.L())
	zapLogger.SetAsDefault()
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: zapLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Account{})
	return &ServiceContext{Config: c, Db: db}
}
