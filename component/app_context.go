package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (ctx *appCtx) GetMainDBConection() *gorm.DB {
	return ctx.db
}
