package ideal

import (
	"ideal/app"
	"ideal/module"
)

func CreateApp(opts ...module.Option) module.App {
	app := app.NewApp(opts...)
	return app
}
