package adapter

import "testing"

func TestAdapter(t *testing.T) {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)

	windowsMachine := &windows{}
	windowsAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.insertSquareUsbInComputer(windowsAdapter)
}
