package main

import (
	"bufio"
	"fmt"
	"github.com/frida/frida-go/frida"
	"os"
)

func main() {
	cluster, err := frida.NewEndpointParameters(&frida.EParams{
		Address: "0.0.0.0",
		Port:    27052,
	})

	control, err := frida.NewEndpointParameters(&frida.EParams{
		Address: "0.0.0.0",
		Port:    27042,
	})

	portal := frida.NewPortal(cluster, control)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating portal: %v\n", err)
		os.Exit(1)
	}

	portal.On("node-connected", func(connId int, addr *frida.Address) {
		fmt.Printf("Got connection from: %s\n", addr)
	})

	portal.On("node_joined", func(connId uint, app *frida.Application) {
		fmt.Printf("[*] Node joined with app: %s\n", app.Name())
	})

	if err := portal.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting portal: %v\n", err)
		os.Exit(1)
	}
	defer portal.Stop()

	fmt.Println("Portal started")

	r := bufio.NewReader(os.Stdin)
	r.ReadLine()
}
