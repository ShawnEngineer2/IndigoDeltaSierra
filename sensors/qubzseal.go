package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
	"indigodeltasierra/datautil"
	"log/slog"
)

func QubzSealInit(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int) {
	//This routine initializes the Qubz Seal sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].QubzSeal.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].QubzSeal.SealSensorState = 1
	(*qubzMatrix)[matrixIndex].QubzSeal.SealState = 1
	(*qubzMatrix)[matrixIndex].QubzSeal.SealEventTime = "02/08/2024 09:24:15.020"

}

func QubzSealSet(qubzMatrix *[]datamodels.QubzMatrix, matrixIndex int, qubzStateDS *datamodels.QubzState, consoleLogger *slog.Logger, fileLogger *slog.Logger) {
	//This routine initializes the Qubz Seal sensor

	//Assign values to the passed sensor
	(*qubzMatrix)[matrixIndex].QubzSeal.EventState = appconstants.SENSOR_STATE_CURRENT
	(*qubzMatrix)[matrixIndex].QubzSeal.SealSensorState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SEALSTATE_SEALSENSORSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].QubzSeal.SealState = int(datautil.GetSensorStateValue(appconstants.SENSOR_DATA_POINT_SEALSTATE_SEALSTATE, qubzStateDS, consoleLogger, fileLogger))
	(*qubzMatrix)[matrixIndex].QubzSeal.SealEventTime = "02/08/2024 09:24:15.020"

}
