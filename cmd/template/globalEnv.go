package template

func GlobalEnvTemplate() []byte {
	return []byte(`
PORT=8080
APP_ENV=local
`)
}
