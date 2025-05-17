package webres

// return the data or if the operation is completed otherwise will block
func (w *WebResource) Await() ([]byte, error) {

	if w.isCompleted {
		return w.cached, w.chachedErr
	}

	select {
	case data := <-w.channel:
		return data, nil
	case err := <-w.errorChannel:
		return nil, err

	case <-w.complete:
		return w.cached, w.chachedErr
	}
}

func (w *WebResource) Then(then func(data []byte)) {
	if w.isCompleted {
		then(w.cached)
		return
	}
	go func() {
		select {
		case data := <-w.channel:
			then(data)
		case <-w.complete:
			then(w.cached)
		}
	}()
}

func (w *WebResource) ThenError(then func(err error)) {
	if w.isCompleted {
		then(w.chachedErr)
		return
	}
	go func() {
		select {
		case err := <-w.errorChannel:
			then(err)
		case <-w.complete:
			then(w.chachedErr)
		}
	}()
}

func (w *WebResource) Finally(then func()) {
	if w.isCompleted {
		then()
		return
	}
	go func() {
		select {
		case <-w.complete:
			then()
		}
	}()
}

func (w *WebResource) GetChannel() chan []byte {
	return w.channel
}
func (w *WebResource) GetErrorChannel() chan error {
	return w.errorChannel
}
func (w *WebResource) GetCompleteChannel() chan bool {
	return w.complete
}
