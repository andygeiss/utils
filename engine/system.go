package engine

// System implements a specific aspect of the solution.
type System interface {
	Error() (err error)
	Process(stopCh chan bool)
	Setup()
	Teardown()
}
