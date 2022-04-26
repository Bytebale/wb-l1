package main

import "fmt"

type mac struct{}
type client struct{}
type windows struct{}

type computer interface {
	insertUsbPort()
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (c *client) insertUsbInComputer(comp computer) {
	comp.insertUsbPort()
}

func (w *windows) insertCirclePort() {
	fmt.Println("Insert circle port into Windows machine")
}

func (m *mac) insertUsbPort() {
	fmt.Println("Insert square port into Mac machine")
}

func (w *windowsAdapter) insertUsbPort() {
	w.windowsMachine.insertCirclePort()
}

func main() {
	client := &client{}
	mac := &mac{}
	client.insertUsbInComputer(mac)
	winda := &windows{}
	adapter := &windowsAdapter{windowsMachine: winda}
	client.insertUsbInComputer(adapter)
}
