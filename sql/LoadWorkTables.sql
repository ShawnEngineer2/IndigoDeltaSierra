/* This script loads the Work Tables from Bronze-type tables and calc */

/* Create the CLASS_OF_SERVICE table */
TRUNCATE TABLE CLASS_OF_SERVICE;

INSERT INTO CLASS_OF_SERVICE(ClassOfService)
SELECT DISTINCT ClassOfService
FROM ROUTES_LOAD;

/* Create the TRANSPORT_MODES table */
TRUNCATE TABLE TRANSPORT_MODES;

INSERT INTO TRANSPORT_MODES(TransportMode)
SELECT DISTINCT TransportMode
FROM ROUTES_LOAD;

/* Load the LOCATIONS table */
TRUNCATE TABLE LOCATIONS;

INSERT INTO LOCATIONS (LocationCode, [Description], City, [State], Country,
                       Latitude, Longitude)
SELECT  LocationCode,
        [Description],
        City,
        [State],
        Country,
        Latitude,
        Longitude
FROM LOCATIONS_LOAD;

/* Load the ROUTES table */
TRUNCATE TABLE ROUTES;

INSERT INTO ROUTES (OriginCode, DestinationCode, ClassOfServiceID, TransportModeID, CostInDollars,
                    NumDaysTravel, OriginLatitude, OriginLongitude, DestinationLatitude, DestinationLongitude,
                    BearingToDestination)
SELECT  A.OriginCode,
        A.DestinationCode,
        B.ClassOfServiceID,
        C.TransportModeID,
        A.CostInDollars,
        A.NumDaysTravel,
        D_ORIGIN.Latitude,
        D_ORIGIN.Longitude,
        D_DESTINATION.Latitude,
        D_DESTINATION.Longitude,
        A.DestinationBearing
FROM ROUTES_LOAD A INNER JOIN CLASS_OF_SERVICE B 
ON A.ClassOfService = B.ClassOfService
INNER JOIN TRANSPORT_MODES C 
ON A.TransportMode = C.TransportMode
INNER JOIN LOCATIONS D_ORIGIN
ON A.OriginCode=D_ORIGIN.LocationCode
INNER JOIN LOCATIONS D_DESTINATION
ON A.DestinationCode=D_DESTINATION.LocationCode;

/* Create A Work Table for calculating bearing */
/*---------------------------------------------*/
DROP TABLE IF EXISTS #BEARING_WORK;

CREATE TABLE #BEARING_WORK
(
    OriginCode                      char(3) NOT NULL,
    DestinationCode                 char(3) NOT NULL,
    OriginLatitude_Degrees          decimal(20,17) NOT NULL,
    OriginLongitude_Degrees         decimal(20,17) NOT NULL,
    OriginLatitude_Radians          decimal(20,17) NULL,
    OriginLongitude_Radians         decimal(20,17) NULL,
    DestinationLatitude_Degrees     decimal(20,17) NOT NULL,
    DestinationLongitude_Degrees    decimal(20,17) NOT NULL,
    DestinationLatitude_Radians     decimal(20,17) NULL,
    DestinationLongitude_Radians    decimal(20,17) NULL,
    Heading_Y                       decimal(20,17) NULL,
    Heading_X                       decimal(20,17) NULL,
    Bearing_Radians                 decimal(20,17) NULL,
    Bearing_Degrees                 decimal(20,17) NULL 
    CONSTRAINT PK_TMP_BEARINGWORK PRIMARY KEY(OriginCode, DestinationCode)
);

/* Load the Bearing Work Table */
/*-----------------------------*/
INSERT INTO #BEARING_WORK (OriginCode, DestinationCode, OriginLatitude_Degrees, OriginLongitude_Degrees,
                           DestinationLatitude_Degrees, DestinationLongitude_Degrees)
SELECT  OriginCode,
        DestinationCode,
        MAX(OriginLatitude),
        MAX(OriginLongitude),
        MAX(DestinationLatitude),
        MAX(DestinationLongitude)
FROM ROUTES
GROUP BY OriginCode,
         DestinationCode;

/* Convert Origin and Destination Lat/Long into Radians */
/*------------------------------------------------------*/
DECLARE @CONST_DEGREES DECIMAL(20,17) = 180
DECLARE @CONST_CIRCLE_DEGREES DECIMAL(20,17) = 360

UPDATE #BEARING_WORK
SET OriginLatitude_Radians = OriginLatitude_Degrees * (PI() / @CONST_DEGREES),
    OriginLongitude_Radians = OriginLongitude_Degrees * (PI() / @CONST_DEGREES),
    DestinationLatitude_Radians = DestinationLatitude_Degrees * (PI() / @CONST_DEGREES),
    DestinationLongitude_Radians = DestinationLongitude_Degrees * (PI() / @CONST_DEGREES)

/* Calculate the heading X and Y to the destination */
/*--------------------------------------------------*/
UPDATE #BEARING_WORK
SET Heading_Y = SIN(DestinationLongitude_Radians - OriginLongitude_Radians) * COS(DestinationLatitude_Radians),
    Heading_X = COS(OriginLatitude_Radians) * SIN(DestinationLatitude_Radians) - SIN(OriginLatitude_Radians) * COS(DestinationLatitude_Radians) * COS(DestinationLongitude_Radians - OriginLongitude_Radians);

/* Calculate the Bearing in Radians */
/*----------------------------------*/
UPDATE #BEARING_WORK
SET Bearing_Radians = ATN2(Heading_Y, Heading_X)

/* Convert Bearing in Radians to Degrees and Normalize to Compass bearing */
/*------------------------------------------------------------------------*/
UPDATE #BEARING_WORK
SET Bearing_Degrees = Bearing_Radians * (@CONST_DEGREES / PI())

/* Normalize Bearing Degrees to Compass bearing */
/*------------------------------------------------------------------------*/
UPDATE #BEARING_WORK
SET Bearing_Degrees = (Bearing_Degrees + @CONST_CIRCLE_DEGREES) % @CONST_CIRCLE_DEGREES

/* Update the ROUTES table with the results of calculations */
/*----------------------------------------------------------*/
UPDATE A 
SET A.BearingToDestination = B.Bearing_Degrees
FROM ROUTES A INNER JOIN #BEARING_WORK B 
ON (A.OriginCode = B.OriginCode) AND (A.DestinationCode = B.DestinationCode)

SELECT * FROM #BEARING_WORK

