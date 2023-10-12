package appctx

import (
	"food-delivery/component/uploadprovider"
	"food-delivery/pubsub"
	"food-delivery/skio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
}

type appCtx struct {
	db         *gorm.DB
	upProvider uploadprovider.UploadProvider
	secretKey  string
	ps         pubsub.Pubsub
	rtEngine   skio.RealtimeEngine
}

func NewAppContext(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string, ps pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, upProvider: upProvider, secretKey: secretKey, ps: ps}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.upProvider
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}

func (ctx *appCtx) GetRealtimeEngine() skio.RealtimeEngine {
	return ctx.rtEngine
}

func (ctx *appCtx) SetRealtimeEngine(rt skio.RealtimeEngine) {
	ctx.rtEngine = rt
}
