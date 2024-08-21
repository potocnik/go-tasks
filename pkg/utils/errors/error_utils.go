package error

import logger "tasks/pkg/logging"

func Check(err error) {
	if err != nil {
		logger.Error("Managed error", err)
		panic(err)
	}
}

func CheckWithMessage(err error, message string) {
	if err != nil {
		logger.Error(message, err)
		panic(err)
	}
}
