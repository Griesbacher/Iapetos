package prom

import (
	"net"
	"net/http"

	"io"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace_core   = "core"
	subsystem_events = "events"
)

func handleMainPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `
<HTML>
	<HEAD><TITLE>Iapetos</TITLE></HEAD>
<BODY>
	<h2>Hi, this is the landingpage of Iapetos.</h2>
	<p>The data you are looking for: <a href="/metrics">metrics</a></p>
</BODY>
</HTML>`)
}

func InitPrometheus(address string) (net.Listener, error) {
	var err error
	prometheusListener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", prometheus.Handler())
		mux.HandleFunc("/", handleMainPage)
		http.Serve(prometheusListener, mux)
	}()
	initHostCheckData()
	initServiceCheckData()
	return prometheusListener, nil
}
