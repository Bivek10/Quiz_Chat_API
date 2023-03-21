package main

import (
	"go.uber.org/fx"

	"github.com/bivek/fmt_backend/bootstrap"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	fx.New(bootstrap.Module).Run()
}
