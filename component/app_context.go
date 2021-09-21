package component

import (
	"Tranning_food/component/uploadprovider"
	"Tranning_food/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.PubSub
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
	pb         pubsub.PubSub
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, pb pubsub.PubSub) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey, pb: pb}
}

func (ctx *appCtx) GetMainDBConection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}
func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) GetPubSub() pubsub.PubSub { return ctx.pb }
