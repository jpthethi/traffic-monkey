# Traffic Monkey

As Vehicles are getting more connectivity and location awareness, the information can be well leveraged to make the traffic lights on an intersection smarter. This application is a traffic generator to test and validate the algorithms that needs to be developed for smart traffic lights.

  - Generates the load of vehicles crossing a n road crossection
  - Provides UDP endpoints where the vehicles location stream will arrive
  - Provides a Websocket on which these events will arrive.
  - UI based on the events to visualize the traffic flow
  - APIs for controlling the traffic signals and Vehicle movement.


# How to use Traffic Monkey

 Clone the repository, get the dependencies, build the project and run the generated exec

```sh
$ git clone <this repo>
$ go get
$ go build
$ ./traffic-monkey
```

## UDP port for listening to events
The events come on the UDP port 6000. you can listen to these events using nc or hook up with your stream ingestion platform

On a saperate terminal / machine
```sh
$ nc -ul localhost 6000
```

## Websocket
The events also come on the websocket service /ws

## Webpage / Visualization
A d3 based visualization of the model is shown in real time from http://localhost:3060

Screenshot:
![Web Page Screenshot](./screenshot "Screenshot")

The page provides intereactivity to toggle signal state from red to green for all roads. The position of the vehicles is updated based on events on websocket

## APIs
| API | Purpose|
|-|-|
|/stats|shows wait time of all vehicles because of signal stoppage|
|/roads|reusable test steps that can be refered in the test yml file|
|/roads/{rd}/stop|Set the signal on road as Red. Vehicles will stop at intersection|
|/roads/{rd}/go|Set the signal on road as Go. Vehicles will start moving if present at intersection|
|/roads/{rd}/vehicles|Count of vehicles on a road|
|/roads/{rd}/vehicle/{veh}|Details of a vehicle on a road|
|/roads/{rd}/vehicle/{veh}/stop|Stop the vehicle|
|/roads/{rd}/vehicle/{veh}/go|Let the vehicle go (from stop state)|


# Running in Docker container

Building the container
```
$ docker build -t traffic-monkey .
```

Running the container
```
$ docker run --net host --name tr -d traffic-monkey
```

pl note we are running in host network as UDP traffic can run in this mode only

# Feedback
Please submit your suggestions / bugs via Github. You are welcome to fork and submit back your pull requests
