package appconstants

// This package provides constants used throughout this app

// System constants
const EXIT_FAILURE int = 125
const EXIT_CLEAN int = 0

// Data file paths
const CONFIG_FILE_PATH string = "./config.dat"
const LOCATIONS_FILE_PATH string = "./locations.dat"
const ROUTES_FILE_PATH string = "./routes.dat"
const CLASS_OF_SERVICE_FILE_PATH string = "./classofservice.dat"
const QUBZ_NAME_FILE_PATH string = "./qubz.dat"
const SENSOR_TYPE_FILE_PATH string = "./sensor_types.dat"
const SHIPMENT_TYPES_FILE_PATH string = "./shipment_types.dat"
const TRANSPORT_MODE_FILE_PATH string = "./transportmodes.dat"

// App constants

const STARTUP_MSG string = "System Starting"
const SIMULATION_COMPLETE_MSG string = "Simulation Complete"

// Simulator constants
const DEFAULT_SVC_WAIT int = 10

const DATA_POINT_TYPE_RANGED int = 1
const DATA_POINT_TYPE_BOOLEAN int = 2
const DATA_POINT_TYPE_CALCULATED int = 3
const DATA_POINT_TYPE_INFORMATIONAL int = 4
