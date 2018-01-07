package logging

import (
	"context"
	"github.com/grtl/mysql-operator/controller"
)

// LogEvents logs events from controllers event channel.
func LogEvents(ctx context.Context, clusterController controller.ClusterController) {
loop:
	for {
		select {
		case e := <-clusterController.GetEventsChan():
			logClusterEvent(e)
		case <-ctx.Done():
			break loop
		}
	}
}
