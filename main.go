package ideal

import (
	"github.com/wuxia-server/ideal/app"
	"github.com/wuxia-server/ideal/module"
)

func CreateApp(opts ...module.Option) module.App {
	return app.NewApp(opts...)
}
