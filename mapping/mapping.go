package mapping

import (
	"fmt"
	"math"
)

func DegreesToRadians(degrees float64) float64 {
	//rads = deg * pi/180
	return degrees * (math.Pi / 180)
}

func RadiansToDegrees(radians float64) float64 {
	//degrees = rads * pi/180
	return radians * (180 / math.Pi)
}

func FindIntermediatePoint(startLat float64, startLong float64, destLat float64, destLong float64, distanceDegrees float64) (float64, float64) {
	//This routine finds intermediate points between two points at the specified distance from the starting Lat / Long

	/*
				public static void normalizedPoint(double lat1, double lon1,
		                                       double lat2, double lon2,
		                                       double dist)
		    {
		        double constant = Math.PI / 180;
		        double angular = dist / 6371;
		        double a = Math.Sin(0 * angular) / Math.Sin(angular);
		        double b = Math.Sin(1 * angular) / Math.Sin(angular);
		        double x = a * Math.Cos(lat1* constant) * Math.Cos(lon1* constant) +
		                   b * Math.Cos(lat2* constant) * Math.Cos(lon2* constant);
		        double y = a * Math.Cos(lat1* constant) * Math.Sin(lon1* constant) +
		                   b * Math.Cos(lat2* constant) * Math.Sin(lon2* constant);
		        double z = a * Math.Sin(lat1* constant) + b * Math.Sin(lat2* constant);
		        double lat3 = Math.Atan2(z, Math.Sqrt(x * x + y * y));
		        double lon3 = Math.Atan2(y, x);
		        Console.WriteLine(lat3 / constant + " " + lon3 / constant);
		    }

			Calling:

		normalizedPoint(47.20761, 27.02185, 47.20754, 27.02177, 1);
		I get the output:

		47.20754 27.02177
	*/
	fmt.Println("Not Implemented")
	return 0, 0
}

func FindMidpoint(startLat float64, startLong float64, destLat float64, destLong float64) (float64, float64) {
	//This routine locates the midpoint between two Lat/Long pairs
	/*
			double dLon = Math.toRadians(lon2 - lon1);

		    //convert to radians
		    lat1 = Math.toRadians(lat1);
		    lat2 = Math.toRadians(lat2);
		    lon1 = Math.toRadians(lon1);

		    double Bx = Math.cos(lat2) * Math.cos(dLon);
		    double By = Math.cos(lat2) * Math.sin(dLon);
		    double lat3 = Math.atan2(Math.sin(lat1) + Math.sin(lat2), Math.sqrt((Math.cos(lat1) + Bx) * (Math.cos(lat1) + Bx) + By * By));
		    double lon3 = lon1 + Math.atan2(By, Math.cos(lat1) + Bx);

		    //print out in degrees
		    System.out.println(Math.toDegrees(lat3) + " " + Math.toDegrees(lon3));
	*/

	var longDiff float64 = DegreesToRadians(destLong - startLong)

	var startLat_Rads float64 = DegreesToRadians(startLat)
	var destLat_Rads float64 = DegreesToRadians(destLat)
	var startLong_Rads float64 = DegreesToRadians(startLong)

	var Bx float64 = math.Cos(destLat_Rads) * math.Cos(longDiff)
	var By float64 = math.Cos(destLat_Rads) * math.Sin(longDiff)
	var midpointLat float64 = math.Atan2(math.Sin(startLat_Rads)+math.Sin(destLat_Rads), math.Sqrt((math.Cos(startLat_Rads)+Bx)*(math.Cos(startLat_Rads)+Bx)+By*By))
	var midpointLong float64 = startLong_Rads + math.Atan2(By, math.Cos(startLat_Rads)+Bx)

	var midpointLat_Degs float64 = RadiansToDegrees(midpointLat)
	var midpointLong_Degs float64 = RadiansToDegrees(midpointLong)

	return midpointLat_Degs, midpointLong_Degs
}

func DecimalToDMS(decimalCoordinate float64) (int, int, int) {
	//Converts coordinates in decimal format to Degrees / Minutes / Seconds
	//const latlong float64 = -97.5296526046655

	var degrees float64 = 0
	var latlongdecimal float64 = 0
	var minutes float64 = 0
	var minutes_decimal float64 = 0
	var seconds_decimal float64 = 0
	var seconds float64 = 0
	var crosscheck float64 = 0

	//Get degress and separate the integer portion from the decimal portion
	degrees = math.Trunc(decimalCoordinate)
	latlongdecimal = decimalCoordinate - degrees

	//Calculate minutes and get minutes decimal
	minutes_decimal = latlongdecimal * 60
	minutes = math.Trunc(minutes_decimal)

	//Calculate seconds
	seconds_decimal = (decimalCoordinate - degrees - (minutes / 60)) * 3600
	seconds = math.Round(seconds_decimal)

	fmt.Println(decimalCoordinate)
	fmt.Println(degrees, minutes, seconds)

	//Cross check to verify
	crosscheck = degrees + (minutes / 60) + (seconds_decimal / 3600)
	fmt.Println(crosscheck)

	if decimalCoordinate == crosscheck {
		fmt.Println("Winner!")
	} else {
		fmt.Println("Loser")
	}

	return int(degrees), int(minutes), int(seconds)

}
