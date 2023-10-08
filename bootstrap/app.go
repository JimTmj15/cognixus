package bootstrap

import "github.com/jmoiron/sqlx"

type Application struct {
	Env *Env
	Db  *sqlx.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Db = NewDb(app.Env)
	return *app
}
