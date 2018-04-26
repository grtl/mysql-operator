package backupinstance

import (
	"fmt"
	"text/template"

	"github.com/grtl/mysql-operator/pkg/operator/backupschedule"
)

// FuncMap can be used to execute templates with the helper functions from the
// Backup Instance operator.
var FuncMap = template.FuncMap{
	"PVCName":       backupschedule.PVCName,
	"JobCreateName": JobCreateName,
	"JobDeleteName": JobDeleteName,
}

// JobCreateName returns a "Create job" name for a given backup.
func JobCreateName(backupName string) string {
	return fmt.Sprintf("%s-create", backupName)
}

// JobDeleteName returns a "Delete job" name for a given backup.
func JobDeleteName(backupName string) string {
	return fmt.Sprintf("%s-delete", backupName)
}
