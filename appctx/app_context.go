package appctx

import "gorm.io/gorm"

// AppContext app ctx
type AppContext interface {
	GetDBConnection() *gorm.DB
}
type appContext struct {
	db *gorm.DB
}

// NewAppContext ctx
func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (ctx *appContext) GetDBConnection() *gorm.DB {
	return ctx.db.Session(
		&gorm.Session{
			NewDB: true,
		},
	)
}
