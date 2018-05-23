package FactoryMethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	factory := new(FileLoggerFactory)
	logger := factory.createLogger()
	logger.writeLog()

}



