package testutil

import (
	"fmt"
)

func GenerateTestEvents() {
	//This function runs the routines in the testutil package to create
	//example JSON files of the events emitted by this simulator

	fmt.Println("Generating Test Event Files ...")
	fmt.Print("\n")

	fmt.Println("Altimeter event ...")
	AltimeterTestEvent("./output/events/altimeter.json")

	fmt.Println("Battery event ...")
	BatteryTestEvent("./output/events/battery.json")

	fmt.Println("Compute event ...")
	ComputeTestEvent("./output/events/compute.json")

	fmt.Println("Geiger Counter event ...")
	GeigerTestEvent("./output/events/geiger.json")

	fmt.Println("GPS event ...")
	GPSTestEvent("./output/events/gps.json")

	fmt.Println("Gyroscopic event ...")
	GyroscopicTestEvent("./output/events/gyro.json")

	fmt.Println("Lock event ...")
	LockTestEvent("./output/events/lock.json")

	fmt.Println("Motion Sensor event ...")
	MotionTestEvent("./output/events/motion.json")

	fmt.Println("Qubz Seal event ...")
	QubzSealTestEvent("./output/events/seal.json")

	fmt.Println("Radio event ...")
	RadioTestEvent("./output/events/radio.json")

	fmt.Println("Temperature and Barometric event ...")
	TempBarometricTestEvent("./output/events/tempbarometric.json")

	fmt.Println("Spectrometer event ...")
	SpectrometerTestEvent("./output/events/spectrometer.json")

	fmt.Println("Internal Fire event ...")
	FireTestEvent("./output/events/fire.json")

	fmt.Print("\n")
	fmt.Println("Test Event File Generation Complete")

}
