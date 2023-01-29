# SIMULATED BODY AREA NETWORK (BAN)
Scalable Computing Project 4

The purpose of this project is to implement, validate and demonstrate a simple scalable communications protocol that employs Peer to Peer modalities of the type that might be found in Body Area Networking communication scenarios.

We propose a Wireless Body Area Network implemented for real time monitoring of patients in a hospital. Each patient is equipped with a number of sensors sending vital information to the sink attached to their body, which in turn, is capable of sending the data to a nearby edge node; the edge node being a server within the hospital. The edge node is where a doctor can monitor the status of all his patients using a dashboard.

The proposed network has three major components:

1. Sink
2. Sensors
3. Edge node


**Below are the guidelines to get the network simulation up and running for simulation:**

# Sink

**"sink" is a simulator program for a sink, it is written in Golang**

* directory: sink

* how to build:

  * install Golang

  * ```shell
    cd sink && go build -o sink main.go
    ```

* how to run

  ```shell
  ./sink -h
  Usage of ./sink:
    -local string
      	local addr (default "0.0.0.0:9090")
    -pda string
      	personal digital assistant address (default "127.0.0.1:9091")
  ```

# Sensors

**"sensor" is a simulator program for sensor, it is written in Golang.**

**you have to start sink first, then start all the sensors and give them the right sink address**

sensor receives json files (dummy) as dataset, you can use following files as dataset for sensors

```shell
sensorBloodAlcohol.json
sensorBloodPressure.json
sensorBodyOxygen.json
sensorBreathingRate.json
sensorInsulin.json
sensorPacemaker.json
sensorTemprature.json
```

* directory: sensors

* how to build: 

  * install Golang

  * ```shell
    cd sensors && go build -o sensor cmd/main.go 
    ```

* how to run:

  ```shell
  ./sensor -h 
  Usage of ./censor:
    -addr string
      	listen ip address (default "127.0.0.1")
    -dataset string
      	dataset file (default "data.txt")
    -duration int
      	working duration(seconds) (default 5)
    -id string
      	sensor id (default "1")
    -interval int
      	interval between working(seconds) (default 10)
    -port string
      	listen port (default "1234")
    -sink string
      	sink address (default "127.0.0.1:9090")
    -x int
      	coordinate x (default 1)
    -y int
      	coordinate y (default 1)
    -z int
      	coordinate z (default 1)
  
  ```
  
 # For Windows Systems
 
 **There are two approaches:**
 
 **A. Directly on the system using cmd**
 1. Intsalling Go in the system.
 2. Starting the sink and builiding it via command prompt using GO libraries.
 3. After that starting sensors so as to connect to the sync.
 
 Note: There might be IP address issues in this approach. Second approach using Ubuntu was therefore used to avoid such issues.
 
 **B. Using Virtualbox**
 
 1. Install virtualbox.
 2. Download a Ubuntu iso image.
 3. Create a VM using iso image.
 4. Check the IP address(WiFi) of the main machine(i.e Windows) and provide it to the Edge node.
 5. Install tmux in the VM using
  * ```shell
    sudo apt install tmux 
    ```
 
 6. Install git in the VM
  * ```shell
    sudo apt install git 
    ```
    
 7. Cloning this repository using git clone (---)
 8. Go to the clone directory using cd
 9. Running the shell script using 
  * ```shell
    sh start.sh
    ``` 
 
 10. This script will autmatically turn on sink node as well as the sensors.
 
**Above steps should get the sensor running and send data to sink. This simulation is run on one single system, the system representing one patient**
# Steps to transfer data from Sink to Edge node and display dashboard.

1. Open Final_Sending_Code.py and Final_Receiving_Code.py and update the hostname, port number.
2. Run the receiver code at edge node and sending code at sink node.
3. Open the Microsoft PowerBI dashboard and click update. The dashboard gets updated with new data.
  

  
