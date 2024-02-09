package sensors

import (
	"indigodeltasierra/appconstants"
	"indigodeltasierra/datamodels"
)

const altimeterState_DPCode int = 44
const altitude_ship_DPCode int = 47
const altitude_plane_DPCode int = 48
const altitude_truck_DPCode int = 45
const altitude_falcon_DPCode int = 49
const altitude_dragon_DPCode int = 50
const altitude_train_DPCode int = 46

func AltimeterInit(qubzMatrix *[]datamodels.QubzMatrix) {
	//This routine initializes the Altimeter sensor

	var altitudeValue float64 = 0

	for i := range *qubzMatrix {

		//Generate an appropriate value based on the passed transport mode
		switch (*qubzMatrix)[i].TransportMode {
		case appconstants.TRANSPORT_MODE_DRAGON:
			altitudeValue = 15
		case appconstants.TRANSPORT_MODE_FALCON:
			altitudeValue = 15
		case appconstants.TRANSPORT_MODE_PLANE:
			altitudeValue = 15
		case appconstants.TRANSPORT_MODE_SHIP:
			altitudeValue = 15
		case appconstants.TRANSPORT_MODE_TRAIN:
			altitudeValue = 15
		case appconstants.TRANSPORT_MODE_TRUCK:
			altitudeValue = 15
		default:
			altitudeValue = 0
		}

		//Assign values to the passed sensor
		(*qubzMatrix)[i].Altimeter.AltimeterState = 1
		(*qubzMatrix)[i].Altimeter.Altitude = altitudeValue
		(*qubzMatrix)[i].Altimeter.EventState = appconstants.SENSOR_STATE_CURRENT
	}

}
