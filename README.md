# Indigo Delta Sierra
Indigo Delta Sierra is an IOT Data Simulator originally based on an idea to simulate events generated by LeMonde Qubz. This repo is a POC (proof-of-concept) for this idea. It's a true MVP and presently contains some massive violations of common coding practice. Please give us a few more iterations before forking or otherwise utilizing this research.

## Kafka Configuration
Before using Kafka (_or Kafka-compatible alternatives like Redpanda_) as an output target for the IOT Data Simulator, a Kafka queue must be provisioned. The queue must be configured as follows:

* The name of the topic must be **_qubz_**
* The topic must have **13** partitions

Retention / deletion policies may be set to defaults

## IOT Simulator Installation
Pre-built compilations of the IOT Data Simulator are found in the *bin* folder of this repo. These are built for the following platforms:

* **Apple Silicon :** iotsim-arm64
* **MacOS on Intel :** iotsim-amd64
* **Windows :** iotsim-wim.exe

**To install on your local system, do the following:**
1. On your local system, create a folder named *iotsim*
2. Copy the contents of the *bin* folder in the repo to your *iotsim* folder
3. In your *iotsim* folder, locate the config.dat file and configure it as described under *IOT Simulator Configuration* below
4. _If your system is a MacOS sytem_, open a terminal console, navigate to the *iotsim* folder and use the following command to make the program file executable: chmod 755 _filename_

## IOT Simulator Configuration
After opening the config.dat file located in your *iotsim* folder, modify the following parameters as described below:

| Parameter Name | Parameter Value |
| -------------- | --------------- |
| qubzCount | Number of Qubz Devices to be simulated. Valid values are a range from 1 to 100001. Default value is 30. |
| eventCycleCount | The number of processing "cycles" performed. Each cycle will produce a set of Current and Previous Sensor Values for each simulated Qubz Device |
| outputChannel | Where Qubz events will routed after being created by the simulator. Valid values are "segmentio", "kafka", and "filesystem". Default is "segmentio", which uses a pure Golang kafka client to route events into a Kafka queue |
| queueEndpoint | The host IP and Port of the Kafka cluster to produce events to. If running Redpanda in a local Docker this should be set to _127.0.0.1:19092_ |
| queueTopic | The Kafka topic to produce to. Default value is _qubz_ |
| logLocation | The location on the local filesystem where logs are written to. Default is to a folder named "logs" in the same directory where IOT Data Simulator is running |

**Note that this version of the IOT Data Simulator has a memory management bug** when emitting events to Kafka. This limits the number of Qubz units that can be simulated. If you experience issues with the program abending during transmission of events to Kafka, lower the value of _qubzCount_ and / or _eventCycleCount_ as needed until the program runs to completion

## Running the IOT Data Simulator
To run the IOT Data Simulator do the following actions:
* **Ensure the Kafka endpoint is up and running.** If you are running Kafka or Redpanda in Docker, ensure that both Docker and the container are running.
* **If running on MacOS with Apple Silicon :** Open a terminal window, navigate to the _iotsim_ folder, and enter _./iotsim_arm64_ , then click on Menu Option 1 and press _Enter_
* **If running on MacOS with Intel :** Open a terminal window, navigate to the _iotsim_ folder, and enter _./iotsim_amd64_ , then click on Menu Option 1 and press _Enter_
* **If running on Windows with PowerShell :** Open a terminal window, navigate to the _iotsim_ folder, and enter _./iotsim_win_ , then click on Menu Option 1 and press _Enter_
* **If running on Windows from File Explorer:** Open File Explorer, navigate to the _iotsim_ folder, and doubleclick the _./iotsim_win_ file, then click on Menu Option 1 and press _Enter_
 
