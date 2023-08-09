package adapter

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {
	w.windowsMachine.insertInCirclePort()
}
