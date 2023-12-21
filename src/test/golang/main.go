package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/web-app-template/modules/wat"
)

func main() {
	m := wat.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
