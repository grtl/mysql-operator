package backup

import (
	"fmt"
	"text/template"
)

// FuncMap can be used to execute templates with the helper functions from the
// backup operator util.
var FuncMap = template.FuncMap{
	"CronJobName": CronJobName,
	"PVCName":     PVCName,
}

// CronJobName returns a name for a cron job associated with the given
// backupName.
func CronJobName(backupName string) string {
	return fmt.Sprintf("%s-job", backupName)
}

// PVCName returns a name for a PVC associated with the given backupName.
func PVCName(backupName string) string {
	return backupName
}
