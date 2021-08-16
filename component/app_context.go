package component

import (
	"Tranning_food/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider}
}

func (ctx *appCtx) GetMainDBConection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
