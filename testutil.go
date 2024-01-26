package main

import (
	"encoding/json"
	"fmt"
	"indigodeltasierra/datamodels"
	"os"
)

func testjson() {
	//This routine is used to test the event JSON formats
	fmt.Println("Creating JSON Representation of Event ...")
	fmt.Println("Opening Output File and Writing JSON ...")

	f, err := os.Create("./output/events/altimeter.json")
	check(err)

	defer f.Close()

	altimeterdata := datamodels.AltimeterReading{}

	altimeterdata.AltimeterState = 1
	altimeterdata.Altitude = 2000
	altimeterdata.EventState = 0

	jsondata := datamodels.AltimeterEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.AltimeterReading, 2)

	jsondata.SensorData[0].AltimeterState = 1
	jsondata.SensorData[0].Altitude = 2000
	jsondata.SensorData[0].EventState = 0

	jsondata.SensorData[1].AltimeterState = 1
	jsondata.SensorData[1].Altitude = 2330
	jsondata.SensorData[1].EventState = 1

	jsonbytes, err := json.Marshal(jsondata)

	check(err)

	_, err = f.WriteString(string(jsonbytes))

	check(err)

	fmt.Println("Done")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func testAltimeter() {
	//This routine is used to test the altimeter JSON
	fmt.Println("Creating JSON Representation of Altimeter Event ...")
	fmt.Println("Opening Output File and Writing JSON ...")

	f, err := os.Create("./output/events/altimeter.json")
	check(err)

	defer f.Close()

	altimeterdata := datamodels.AltimeterReading{}

	altimeterdata.AltimeterState = 1
	altimeterdata.Altitude = 2000
	altimeterdata.EventState = 0

	jsondata := datamodels.AltimeterEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.AltimeterReading, 2)

	jsondata.SensorData[0].AltimeterState = 1
	jsondata.SensorData[0].Altitude = 2000
	jsondata.SensorData[0].EventState = 0

	jsondata.SensorData[1].AltimeterState = 1
	jsondata.SensorData[1].Altitude = 2330
	jsondata.SensorData[1].EventState = 1

	jsonbytes, err := json.Marshal(jsondata)

	check(err)

	_, err = f.WriteString(string(jsonbytes))

	check(err)

	fmt.Println("Done")
}

func testBattery() {
	//This routine is used to test the battery JSON
	const sensorName string = "Battery"

	fmt.Println("Creating JSON Representation of " + sensorName + " Event ...")
	fmt.Println("Opening Output File and Writing JSON ...")

	f, err := os.Create("./output/events/" + sensorName + ".json")
	check(err)

	defer f.Close()

	jsondata := datamodels.BatteryEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.BatteryReading, 2)

	jsondata.SensorData[0].BatterySensorState = 1
	jsondata.SensorData[0].BatteryTemperature = 15
	jsondata.SensorData[0].CapacitanceGelActualVolume = 300
	jsondata.SensorData[0].CapacitanceGelExpectedVolume = 300
	jsondata.SensorData[0].DrainRate = 134
	jsondata.SensorData[0].MaxAmpHours = 50000
	jsondata.SensorData[0].RemainingAmpHours = 34000
	jsondata.SensorData[0].EventState = 0

	jsondata.SensorData[1].BatterySensorState = 1
	jsondata.SensorData[1].BatteryTemperature = 15
	jsondata.SensorData[1].CapacitanceGelActualVolume = 300
	jsondata.SensorData[1].CapacitanceGelExpectedVolume = 300
	jsondata.SensorData[1].DrainRate = 134
	jsondata.SensorData[1].MaxAmpHours = 50000
	jsondata.SensorData[1].RemainingAmpHours = 34000
	jsondata.SensorData[1].EventState = 1

	jsonbytes, err := json.Marshal(jsondata)

	check(err)

	_, err = f.WriteString(string(jsonbytes))

	check(err)

	fmt.Println("Done")
}

func testCompute() {
	//This routine is used to test the Compute JSON
	const sensorName string = "Compute"

	fmt.Println("Creating JSON Representation of " + sensorName + " Event ...")
	fmt.Println("Opening Output File and Writing JSON ...")

	f, err := os.Create("./output/events/" + sensorName + ".json")
	check(err)

	defer f.Close()

	jsondata := datamodels.ComputeEvent{}

	jsondata.EventHeader.EventTimestamp = "XYZ"
	jsondata.EventHeader.QubzId = "SunnyRainyJellyfish"
	jsondata.EventHeader.RouteAssignment = 3
	jsondata.EventHeader.ShipmentType = 2

	jsondata.SensorData = make([]datamodels.ComputeReading, 2)

	jsondata.SensorData[0].EventState = 0
	jsondata.SensorData[0].AmountMemoryUtilized = 2400
	jsondata.SensorData[0].CPUUtilization = 57
	jsondata.SensorData[0].ComputeFirmwareState = 1
	jsondata.SensorData[0].ComputeFirmwareVersion = 17
	jsondata.SensorData[0].ComputeSensorState = 1
	jsondata.SensorData[0].NumCPUCores = 16
	jsondata.SensorData[0].OSTypeVersion = 2
	jsondata.SensorData[0].RemainingDiskStorage = 2300
	jsondata.SensorData[0].TotalAmountOfMemory = 3600
	jsondata.SensorData[0].TotalDiskStorage = 4000

	jsondata.SensorData[1].EventState = 1
	jsondata.SensorData[1].AmountMemoryUtilized = 2400
	jsondata.SensorData[1].CPUUtilization = 57
	jsondata.SensorData[1].ComputeFirmwareState = 1
	jsondata.SensorData[1].ComputeFirmwareVersion = 17
	jsondata.SensorData[1].ComputeSensorState = 1
	jsondata.SensorData[1].NumCPUCores = 16
	jsondata.SensorData[1].OSTypeVersion = 2
	jsondata.SensorData[1].RemainingDiskStorage = 2300
	jsondata.SensorData[1].TotalAmountOfMemory = 3600
	jsondata.SensorData[1].TotalDiskStorage = 4000

	jsonbytes, err := json.Marshal(jsondata)

	check(err)

	_, err = f.WriteString(string(jsonbytes))

	check(err)

	fmt.Println("Done")
}
