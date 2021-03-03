package main

import (
	"github.com/vedhavyas/go-wasm"
)

func main() {
	b, err := wasm.BridgeFromFile("test", "../hello-wasm/hello-wasm", nil)
	if err != nil {
		panic(err)
	}

	/*
		ctx, canc := context.WithCancel(context.Background())
		defer canc()
		init := make(chan error)
		go b.Run(ctx, init)
		if err := <-init; err != nil {
			panic(err)
		}
	*/
	run, _ := b.GetMain()
	run(0, 0)
}
