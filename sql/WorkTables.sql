/* This script creates the tables used to generate the data files used by this app */
/*=================================================================================*/

DROP TABLE IF EXISTS LOCATIONS

CREATE TABLE LOCATIONS
(
    LocationID  INT not null IDENTITY(1,1),
    LocationCode char(3) NOT NULL PRIMARY KEY,
    [Description] varchar(100) NOT NULL,
    City varchar(100),
    [State] varchar(100),
    Country varchar(100),
    Latitude decimal(20,17),
    Longitude decimal(20,17)
);

DROP TABLE IF EXISTS ROUTES

CREATE TABLE ROUTES
(
    RouteID int NOT NULL IDENTITY(1,1),
    OriginCode char(3) NOT NULL,
    DestinationCode char(3) NOT NULL,
    ClassOfServiceID int NOT NULL,
    TransportModeID int NOT NULL,
    CostInDollars money NOT NULL,
    NumDaysTravel smallint NOT NULL,
    OriginLatitude decimal(20,17),
    OriginLongitude decimal(20,17),
    DestinationLatitude decimal(20,17),
    DestinationLongitude decimal(20,17),
    BearingToDestination decimal(20,17)
    CONSTRAINT PK_ROUTES PRIMARY KEY (OriginCode, DestinationCode, ClassOfServiceID)
);

DROP TABLE IF EXISTS TRANSPORT_MODES

CREATE TABLE TRANSPORT_MODES
(
    TransportModeID int NOT NULL IDENTITY(1,1) PRIMARY KEY,
    TransportMode varchar(100) NOT NULL
);

DROP TABLE IF EXISTS CLASS_OF_SERVICE

CREATE TABLE CLASS_OF_SERVICE
(
    ClassOfServiceID int NOT NULL IDENTITY(1,1) PRIMARY KEY,
    ClassOfService varchar(100) NOT NULL
);

