package cmd

import (
	"github.com/crusj/logger"
	"testing"
)

func TestYaPi(t *testing.T) {
	logger.OpenShortLog()
	parser := NewYaPiPath("../test_Route.php","//route开始","//route结束","route_start","route_end")
	actions := NewActions("../api.json", parser)
	actions.AnalyzeGroup()
	actions.PrintRoutes()
	actions.InsertRoute()

}
