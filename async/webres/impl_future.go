package webres

// return the data or if the operation is completed otherwise will block
func (w *WebResource) Await() ([]byte, error) {

	if w.isCompleted {
		return w.cached, w.chachedErr
	}
	//block until the operation is completed
	<-w.completeChannel

	return w.cached, w.chachedErr
}

func (w *WebResource) Then(thenFn func(data []byte)) {
	go func() {

		data, err := w.Await()
		if err != nil {
			return
		}
		thenFn(data)
	}()

}

func (w *WebResource) ThenError(thenErrFn func(err error)) {
	go func() {
		_, err := w.Await()
		if err != nil {
			thenErrFn(err)
			return
		}
	}()
}

func (w *WebResource) Finally(finallyFn func()) {
	go func() {
		_, err := w.Await()
		if err != nil {
			return
		}
	
		finallyFn()
	}()
}

func (w *WebResource) GetChannel() *chan []byte {
	return &w.channel
}
func (w *WebResource) GetErrorChannel() *chan error {
	return &w.errorChannel
}
func (w *WebResource) GetCompleteChannel() *chan bool {
	return &w.completeChannel
}
