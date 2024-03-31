# This script builds the IOT simulator for various target platforms

# Clear the existing output files
echo "Clearing previous files ..."
rm ./bin/*

# Build for Apple Silicon
echo "Building for Apple Silicon ..."
GOOS=darwin GOARCH=arm64 go build -o ./bin/iotsim-arm64 .

# Build for Linux on Intel
echo "Building for Linux Intel ..."
GOOS=linux GOARCH=amd64 go build -tags musl -o ./bin/iotsim-amd64 .


# Build the Mac Intel version
echo "Building for Apple Intel ..."
GOOS=darwin GOARCH=amd64 go build -o ./bin/iotsim-amd64 .

# Build the Windows version
echo "Building for Windows Intel ..."
GOOS=windows GOARCH=amd64 go build -o ./bin/iotsim-win.exe .

# Copy the config file
echo "Copying config file ..."
cp ./config.dat ./bin/config.dat

# Copy the data files
echo "Copying data files ..."
cp ./classofservice.dat ./bin/classofservice.dat
cp ./locations.dat ./bin/locations.dat
cp ./qubz.dat ./bin/qubz.dat
cp ./routes.dat ./bin/routes.dat
cp ./sensor_ranges.dat ./bin/sensor_ranges.dat
cp ./sensor_types.dat ./bin/sensor_types.dat
cp ./shipment_types.dat ./bin/shipment_types.dat
cp ./transportmodes.dat ./bin/transportmodes.dat
cp ./exceptions.dat ./bin/exceptions.dat

echo "Done"