package cyclic

//Stoppable is a struct with one chan in, which will be used to signal an go routine to stop
type Stoppable struct {
	stop chan bool
}

//Stop will send true to the channel and waits for an response
func (s Stoppable) Stop() {
	if s.stop == nil {
		return
	}
	s.stop <- true
	<-s.stop
}
