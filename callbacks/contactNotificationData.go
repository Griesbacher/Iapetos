package callbacks

import (
	"unsafe"

	"github.com/ConSol/go-neb-wrapper/neb"
	"github.com/ConSol/go-neb-wrapper/neb/structs"
	"github.com/griesbacher/Iapetos/prom"
	"github.com/prometheus/client_golang/prometheus"
)

//ContactNotificationCheckData is a callback for neb.ContactNotificationData
func ContactNotificationCheckData(callbackType int, data unsafe.Pointer) int {
	if callbackType != neb.ContactNotificationData {
		return neb.Error
	}

	contactNotification := structs.CastContactNotificationCheck(data)

	identifier := prometheus.Labels{
		prom.HostName:           contactNotification.HostName,
		prom.ServiceDescription: contactNotification.ServiceDescription,
		prom.AckAuthor:          contactNotification.AckAuthor,
		prom.ContactName:        contactNotification.ContactName,
		prom.Reason:             structs.CastNotificationReasonToString(contactNotification.ReasonType),
		prom.Type:               structs.CastNotificationTypeToString(contactNotification.NotificationType),
	}

	if contactNotification.Type == neb.ContactNotificationStart {
		prom.ContactNotificationStart.With(identifier).Inc()
	} else if contactNotification.Type == neb.ContactNotificationEnd {
		prom.ContactNotificationEnd.With(identifier).Inc()
	}

	return neb.Ok
}
