package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemNotification = "notification"
	//AckAuthor is the string "ack_author"
	AckAuthor = "ack_author"
	//Reason is the string "reason"
	Reason = "reason"
)

var notificationLabelNames = []string{HostName, ServiceDescription, AckAuthor, Type, Reason}

//NotificationStart is a Prometheus counter vector
var NotificationStart = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNotification,
		Name:      "start_total",
		Help:      "Notification started",
	},
	notificationLabelNames,
)

//NotificationEnd is a Prometheus counter vector
var NotificationEnd = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNotification,
		Name:      "end_total",
		Help:      "Notification ended",
	},
	notificationLabelNames,
)

//NotificationContactsNotified is a Prometheus counter vector
var NotificationContactsNotified = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNotification,
		Name:      "contacts_notified_total",
		Help:      "Contacts notified",
	},
	notificationLabelNames,
)

func initNotificationCheckData() {
	prometheus.MustRegister(NotificationStart)
	prometheus.MustRegister(NotificationEnd)
	prometheus.MustRegister(NotificationContactsNotified)
}
