package appconstants

// This package provides constants used throughout this app

// System constants
const EXIT_FAILURE int = 125
const EXIT_CLEAN int = 0

// App constants
const CONFIG_FILE_PATH string = "./config.dat"

// Message constants
const STARTUP_MSG string = "System Starting"
const SIMULATION_COMPLETE_MSG string = "Simulation Complete"
const SIMULATION_FAILED_MSG string = "Simulation Failed - Please see log for details!"

// Simulator constants
const DEFAULT_SVC_WAIT int = 10

// Data Point Type constants
const DATA_POINT_TYPE_RANGED int = 1
const DATA_POINT_TYPE_BOOLEAN int = 2
const DATA_POINT_TYPE_CALCULATED int = 3
const DATA_POINT_TYPE_INFORMATIONAL int = 4

// Sensor constants
const SENSOR_STATE_PREVIOUS int = 0
const SENSOR_STATE_CURRENT int = 1

// Tranportation Mode Constants
const TRANSPORT_MODE_SHIP int = 1
const TRANSPORT_MODE_PLANE int = 2
const TRANSPORT_MODE_TRUCK int = 3
const TRANSPORT_MODE_FALCON int = 4
const TRANSPORT_MODE_DRAGON int = 5
const TRANSPORT_MODE_TRAIN int = 6
