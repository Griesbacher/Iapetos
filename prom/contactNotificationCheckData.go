package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemContactNotification = "contact_notification"
	//ContactName is the string "contact_name"
	ContactName = "contact_name"
)

var contactNotificationLabelNames = []string{HostName, ServiceDescription, AckAuthor, ContactName, Type, Reason}

//ContactNotificationStart is a Prometheus counter vector
var ContactNotificationStart = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemContactNotification,
		Name:      "start_total",
		Help:      "Notification started",
	},
	contactNotificationLabelNames,
)

//ContactNotificationEnd is a Prometheus counter vector
var ContactNotificationEnd = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemContactNotification,
		Name:      "end_total",
		Help:      "Notification ended",
	},
	contactNotificationLabelNames,
)

func initContactNotificationCheckData() {
	prometheus.MustRegister(ContactNotificationStart)
	prometheus.MustRegister(ContactNotificationEnd)
}
