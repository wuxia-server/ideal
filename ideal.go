package ideal

import (
	"ideal/app"
	"ideal/module"
)

func CreateApp(opts ...module.Option) module.App {
	return app.NewApp(opts...)
}
