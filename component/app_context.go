package component

import (
	"Tranning_food/component/otpProvider"
	"Tranning_food/component/uploadprovider"
	"Tranning_food/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.PubSub
	OtpProvider() otpProvider.OtpProvider
}

type appCtx struct {
	db          *gorm.DB
	upProvider  uploadprovider.UploadProvider
	secretKey   string
	pb          pubsub.PubSub
	otpProvider otpProvider.OtpProvider
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, pb pubsub.PubSub, otpProvider otpProvider.OtpProvider) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey, pb: pb, otpProvider: otpProvider}
}

func (ctx *appCtx) GetMainDBConection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) GetPubSub() pubsub.PubSub { return ctx.pb }

func (ctx *appCtx) OtpProvider() otpProvider.OtpProvider {
	return ctx.otpProvider
}
