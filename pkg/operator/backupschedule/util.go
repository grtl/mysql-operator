package backupschedule

import (
	"fmt"
	"text/template"
)

// FuncMap can be used to execute templates with the helper functions from the
// Backup Shcedule operator fail.
var FuncMap = template.FuncMap{
	"CronJobName": CronJobName,
	"PVCName":     PVCName,
}

// CronJobName returns a name for a cron job associated with the given
// scheduleName.
func CronJobName(scheduleName string) string {
	return fmt.Sprintf("%s-job", scheduleName)
}

// PVCName returns a name for a PVC associated with the given backupName.
func PVCName(scheduleName string) string {
	return scheduleName
}
