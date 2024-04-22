package main

import "context"

type HandlerFunc func(context.Context)

type HandlersChain []HandlerFunc

type RoutesGroup struct {
	Handlers HandlersChain
	basePath string
}

func main() {

}
