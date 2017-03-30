package prom

import "github.com/prometheus/client_golang/prometheus"

const (
	subsystemNotification = "notification"
	AckAuthor             = "ack_author"
	Reason                = "reason"
)

var notificationLabelNames = []string{HostName, ServiceDescription, AckAuthor, Type, Reason}

//NotificationStarted is a Prometheus counter vector
var NotificationStart = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNotification,
		Name:      "start_total",
		Help:      "Notification started",
	},
	notificationLabelNames,
)

//NotificationEnded is a Prometheus counter vector
var NotificationEnd = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Namespace: namespaceCore,
		Subsystem: subsystemNotification,
		Name:      "end_total",
		Help:      "Notification ended",
	},
	notificationLabelNames,
)

//NotificationReason is a Prometheus counter vector
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
