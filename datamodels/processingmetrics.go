package datamodels

// This struct is used to capture processing metrics for display at the end of the simulation run
type ProcessingMetrics struct {
	EstimatedEventCount int
	ActualEventCount    int
	Sev01ExceptionCount int
	Sev02ExceptionCount int
	Sev03ExceptionCount int
	TotalQubzCount      int
	TotalExceptionCount int
}
