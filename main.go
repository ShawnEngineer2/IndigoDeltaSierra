package main

import (
	"bufio"
	"flag"
	"fmt"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/svcclient"
	"indigodeltasierra/sysfile"
	"indigodeltasierra/testutil"
	"indigodeltasierra/validators"
	"io"
	"log/slog"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

const EXIT_FAILURE int = 125
const EXIT_CLEAN int = 0

func main() {

	//Define the "silent" flag. If the "silent" flag was not sent in on the command line, show a menu and process based on the input
	silentPtr := flag.Bool("silent", false, "When specified, causes the simulator to start without displaying the user menu")

	if *silentPtr {
		fmt.Print("Starting IOT event simulation ...")
		startSimulation()
	} else {
		//Run interactively
		runInteractive()
	}
}

func RotatingLog(path string) io.Writer {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    1,     // In MB before rotating the file
		MaxAge:     1,     // In days before deleting the file
		MaxBackups: 5,     // Maximum number of backups to keep track of
		Compress:   false, // Compress the rotated log files, false by default.
	}
}

func generateTestEvents() {
	//This function runs the routines in the testutil package to create
	//example JSON files of the events emitted by this simulator

	fmt.Println("Generating Test Event Files ...")
	fmt.Print("\n")

	fmt.Println("Altimeter event ...")
	testutil.AltimeterTestEvent("./output/events/altimeter.json")

	fmt.Println("Battery event ...")
	testutil.BatteryTestEvent("./output/events/battery.json")

	fmt.Println("Compute event ...")
	testutil.ComputeTestEvent("./output/events/compute.json")

	fmt.Println("Geiger Counter event ...")
	testutil.GeigerTestEvent("./output/events/geiger.json")

	fmt.Println("GPS event ...")
	testutil.GPSTestEvent("./output/events/gps.json")

	fmt.Println("Gyroscopic event ...")
	testutil.GyroscopicTestEvent("./output/events/gyro.json")

	fmt.Println("Lock event ...")
	testutil.LockTestEvent("./output/events/lock.json")

	fmt.Println("Motion Sensor event ...")
	testutil.MotionTestEvent("./output/events/motion.json")

	fmt.Println("Qubz Seal event ...")
	testutil.QubzSealTestEvent("./output/events/seal.json")

	fmt.Println("Radio event ...")
	testutil.RadioTestEvent("./output/events/radio.json")

	fmt.Println("Temperature and Barometric event ...")
	testutil.TempBarometricTestEvent("./output/events/tempbarometric.json")

	fmt.Println("Spectrometer event ...")
	testutil.SpectrometerTestEvent("./output/events/spectrometer.json")

	fmt.Print("\n")
	fmt.Println("Test Event File Generation Complete")

}

func runInteractive() {
	//Call the menu routine and wait for user input
	showMenu()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		inputChar, _, _ := reader.ReadRune()

		switch inputChar {
		case '1':
			startSimulation()
			fmt.Println("")
			os.Exit(EXIT_CLEAN)

		case '2':
			generateTestEvents()
			fmt.Println("")
			os.Exit(EXIT_CLEAN)

		case '9':
			fmt.Println("")
			fmt.Println("Program Exit - Good-bye")
			fmt.Println("")
			os.Exit(EXIT_CLEAN)
		}
	}

}

func showMenu() {
	//Prints the menu to the screen
	fmt.Println("")
	fmt.Println("===============================")
	fmt.Println("IOT EVENT SIMULATOR")
	fmt.Println("===============================")
	fmt.Println("")
	fmt.Println("This application simulates random sensor events from a number of virtual IOT devices.")
	fmt.Println("Please choose an option from the menu below:")
	fmt.Println("")
	fmt.Println("1 : Begin Simulation")
	fmt.Println("2 : Generate Example Events")
	fmt.Println("9 : Exit")
	fmt.Println("")
}

func startSimulation() {
	//This function configures the app and starts the simulation
	//Allocate constants
	const CONFIG_FILE_PATH string = "./config.dat"
	const STARTUP_MSG string = "System Starting"
	const SHUTDOWN_MSG string = "Run Complete ... Shutting Down"
	const DEFAULT_SVC_WAIT int = 10
	var max_qubz_count int = 0

	//Set up console logger
	consoleLogger := slog.Default()
	consoleLogger.Info(STARTUP_MSG)

	//Check for Config file
	consoleLogger.Info("Checking for Config file ... ")

	if !sysfile.FileExists(CONFIG_FILE_PATH) {
		consoleLogger.Error("Cannot find Config file in boot directory ... startup terminated!")
		os.Exit(EXIT_FAILURE)
	}

	//Load Config file
	consoleLogger.Info("Found Config file ... begin load ...")

	config := datamodels.Config{}

	if !sysfile.LoadFileToStruct(CONFIG_FILE_PATH, &config) {
		consoleLogger.Error("Could not load Config file ... startup terminated")
		os.Exit(EXIT_FAILURE)
	} else {
		consoleLogger.Info("Config values loaded")
		consoleLogger.Info("Validating Config ...")

		errCount := validators.ValidateConfig(config, *consoleLogger)

		if errCount > 0 {
			consoleLogger.Error(fmt.Sprintf("%s errors identified in loaded config. Startup terminated ... please correct and start again", fmt.Sprint(errCount)))
			os.Exit(EXIT_FAILURE)
		} else {
			consoleLogger.Info("Config values validated. Loaded values are:")
			consoleLogger.Info(fmt.Sprintf("%+v", config))
		}

	}

	//Setup logfile logger
	consoleLogger.Info("Configuring File Logger to output path " + config.LogLocation + " ...")
	fileLogger := slog.New(slog.NewJSONHandler(RotatingLog(config.LogLocation), nil))
	fileLogger.Info(STARTUP_MSG)
	fileLogger.Info("Config Values for this run ...")
	fileLogger.Info(fmt.Sprintf("%+v", config))

	//Load Qubz Names and IDs from the Qubz Name File
	qubznames := []datamodels.Qubz{}

	if !sysfile.LoadFileToStruct(config.QubzNameFile, &qubznames) {
		consoleLogger.Error(fmt.Sprintf("Could not load Qubz Names from file %s ... startup terminated", config.QubzNameFile))
		os.Exit(EXIT_FAILURE)
	} else {
		max_qubz_count = len(qubznames)
		consoleLogger.Info(fmt.Sprintf("Qubz Names loaded from file %s ... %d names loaded", config.QubzNameFile, max_qubz_count))
	}

	//Use the loaded names list to source a random set of Qubz names
	qubzmatrix := make([]datamodels.QubzMatrix, config.QubzCount)

	if config.QubzCount == 0 || config.QubzCount == max_qubz_count {
		//Just load'em all into the Qubz Matrix
		fileLogger.Info("Loading All Qubz Names ...")

		qubzmatrix = make([]datamodels.QubzMatrix, max_qubz_count)

		for i, x := range qubznames {
			qubzmatrix[i].QubzID = x.QubzID
			qubzmatrix[i].QubzName = x.QubzName
		}

		fmt.Println(qubzmatrix[10])

	} else if config.QubzCount < max_qubz_count {
		//Grab a random set of numbers of QubzCount between 1 and max qubz count and load those names
		fileLogger.Info("Loading Random Qubz Names ...")

		fmt.Println(svcclient.GetRandomNumbers(config.QubzCount, 0, (max_qubz_count - 1), fileLogger, config.EmailAddress))

	} else if config.QubzCount > max_qubz_count {
		//Too many Qubz requested - throw an error
		consoleLogger.Error(fmt.Sprintf("Not enough names in %s file (%d) to satisfy requested number of Qubz (%d) ... startup terminated", config.QubzNameFile, len(qubznames), config.QubzCount))
		os.Exit(EXIT_FAILURE)
	}

	//Cooling off wait so we don't overload the random number generator service
	consoleLogger.Info("Service Cooloff Wait ...")
	time.Sleep(time.Duration(DEFAULT_SVC_WAIT) * time.Second)

	quotaExceeded, err := svcclient.CheckQuotaExceeded(fileLogger, config.EmailAddress)

	if err != nil {
		msg := "Service Abend: Error checking Random Service Quota: " + err.Error()
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		os.Exit(EXIT_FAILURE)
	} else if quotaExceeded {
		msg := "Service Abend: Random Service Quota exceeded"
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		os.Exit(EXIT_FAILURE)
	}

	//Load the route table into a struct

	//Randomly assign routes to the Qubz in the Qubz Matrix

	//Cool off wait for random number generator service - then check quota
	consoleLogger.Info("Service Cooloff Wait ...")
	time.Sleep(time.Duration(DEFAULT_SVC_WAIT) * time.Second)

	quotaExceeded, err = svcclient.CheckQuotaExceeded(fileLogger, config.EmailAddress)

	if err != nil {
		msg := "Service Abend: Error checking Random Service Quota: " + err.Error()
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		os.Exit(EXIT_FAILURE)
	} else if quotaExceeded {
		msg := "Service Abend: Random Service Quota exceeded"
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		os.Exit(EXIT_FAILURE)
	}

	//Load the Event Types file into a struct

	quotaExceeded, err = svcclient.CheckQuotaExceeded(fileLogger, config.EmailAddress)

	if err != nil {
		msg := "Service Abend: Error checking Random Service Quota: " + err.Error()
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		os.Exit(EXIT_FAILURE)
	} else if quotaExceeded {
		msg := "Service Abend: Random Service Quota exceeded"
		consoleLogger.Error(msg)
		fileLogger.Error(msg)
		os.Exit(EXIT_FAILURE)
	}

	//Exit with a CLEAN (no errors) code
	fileLogger.Info(SHUTDOWN_MSG)
	consoleLogger.Info(SHUTDOWN_MSG)
	os.Exit(EXIT_CLEAN)
}
