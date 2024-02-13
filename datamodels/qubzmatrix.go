package datamodels

//This struct is used to implement the Qubz Matrix - this is the centerpiece of data tracking in this app

type QubzMatrix struct {
	QubzID                int
	QubzName              string
	RouteAssignment       int
	ShipmentType          int
	TransportMode         int
	ExceptionAssignment   int
	ExceptionSeverity     int
	ExceptionIntermittent bool
	ExceptionInterval     int //Note: Only applies if Exception is Intermitten. If Exception Interval is -1, the error occurs randomly (0 /1 type stuff)
	Altimeter             AltimeterReading
	Battery               BatteryReading
	Compute               ComputeReading
	Fire                  FireReading
	Geiger                GeigerReading
	GPS                   GPSReading
	Gyro                  GyroscopeReading
	Lock                  LockReading
	Motion                MotionReading
	QubzSeal              SealReading
	Radio                 RadioReading
	Spectrometer          SpectrometerReading
	TempBarometrics       TempBarometricsReading
}
