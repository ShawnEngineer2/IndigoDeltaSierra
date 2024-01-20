package validators

import (
	"indigodeltasierra/datamodels"
	"log/slog"
)

func ValidateConfig(configToValidate datamodels.Config, consoleLogger slog.Logger) int {

	//This routine validates that the contents of the Config struct are within acceptable parameters

	validationErrorCount := 0

	//Validate Qubz Count
	if configToValidate.QubzCount < 0 {
		consoleLogger.Error("QubzCount value cannot be less than 1")
		validationErrorCount++

	}

	//Validate Event Interval
	if configToValidate.EventInterval < 10 {
		consoleLogger.Error("EventInterval value cannot be less than 10")
		validationErrorCount++
	}

	//Validate Event Cycle Count
	if configToValidate.EventCycleCount < 1 {
		consoleLogger.Error("EventCycleCount value cannot be less than 1")
		validationErrorCount++
	}

	//Validate Email string
	if len(configToValidate.EmailAddress) == 0 {
		consoleLogger.Error("Email Address cannot be blank")
		validationErrorCount++
	}

	//Validate Queue Endpoint value
	if len(configToValidate.QueueEndpoint) == 0 {
		consoleLogger.Error("Message Queue Endpoint address cannot be blank")
		validationErrorCount++
	}

	//Validate Queue Topic value
	if len(configToValidate.QueueTopic) == 0 {
		consoleLogger.Error("Message Queue Topic name cannot be blank")
		validationErrorCount++
	}

	//Validate Log File Location
	if len(configToValidate.LogLocation) == 0 {
		consoleLogger.Error("Message Queue Topic value cannot be blank")
		validationErrorCount++
	}

	//Return the number of errors found
	return validationErrorCount

}
