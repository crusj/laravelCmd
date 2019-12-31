package init

import (
	logger "github.com/crusj/logger"
)

func init() {
	logger.SetLogger(`{"Console": {"level": "DEBG","color":true}}`)
	logger.OpenShortLog()
}
