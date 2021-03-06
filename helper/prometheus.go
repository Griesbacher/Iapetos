package helper

import "github.com/prometheus/client_golang/prometheus"

//CopyLabels copies prometheus labels
func CopyLabels(in prometheus.Labels) (out prometheus.Labels) {
	out = prometheus.Labels{}
	for k, v := range in {
		out[k] = v
	}
	return out
}
