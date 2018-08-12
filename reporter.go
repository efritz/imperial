package imperial

type Reporter interface {
	Report(name string, value int, configs ...ConfigFunc)
	Shutdown()
}
