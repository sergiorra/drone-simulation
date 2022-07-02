<div align="center">
  <h1>Drone Simulation</h1>
  <blockquote>Simulation of two drones that are flying around London and report on traffic conditions while receiving the coordinates of the following positions to be reported.</blockquote>
</div>

## 📜 Documentation

### Scenario
There are two automatic drones that fly around London and report on traffic conditions. When a drone flies over a tube station, it assesses what the traffic condition is like in the area, and reports on it.

### How it works
- There is one dispatcher and two drones. 
- Each drone is "moving" independently on different processes. 
- The dispatcher sends the coordinates to each drone detailing where the drone's next position should be. 
- The dispatcher is the responsible for terminating the program.
- When the drone receives a new coordinate, it moves, checks if there is a tube station in the area, and if so, reports on the traffic conditions there.

### Notes
- The simulation finishes @ 08:10, where the drones will receive a "SHUTDOWN" signal.
- The two drones have IDs 6043 and 5937. 
- There is a file containing their lat/lon points for their routes. The CSV file layout is "droneId, latitude, longitude, time".
- There is also a file with the lat/lon points for London tube stations with the layout "station, lat, lon".
- Traffic reports have the following format:
  - Drone ID
  - Time
  - Speed
  - Conditions of Traffic (HEAVY, LIGHT, MODERATE). This can be chosen randomly.

### Remarks
- Assume that the drones follow a straight line between each point, travelling at constant speed.
- Disregard the fact that the start time is not in sync. The dispatcher can start pumping data as soon as it has read the files.
- A nearby station should be less than 350 meters from the drone's position.
- There should be a constraint on each drone to have limited memory, so they can only consume ten points at a time.