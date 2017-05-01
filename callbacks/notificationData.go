package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

//NotificationCheckData is a callback for neb.NotificationData
func NotificationCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.NotificationData {
		return neb.Error
	}

	notification := structs.CastNotificationCheck(data)

	identifier := prometheus.Labels{
		prom.HostName:           notification.HostName,
		prom.ServiceDescription: notification.ServiceDescription,
		prom.AckAuthor:          notification.AckAuthor,
		prom.Type:               structs.CastNotificationTypeToString(notification.NotificationType),
		prom.Reason:             structs.CastNotificationReasonToString(notification.ReasonType),
	}

	if notification.Type == neb.NotificationStart {
		prom.NotificationStart.With(identifier).Inc()
		prom.NotificationContactsNotified.With(identifier).Add(float64(notification.ContactsNotified))
	} else if notification.Type == neb.NotificationEnd {
		prom.NotificationEnd.With(identifier).Inc()
	}

	return neb.Ok
}
