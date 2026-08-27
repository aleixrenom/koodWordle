# WeatherStation instructions

Optimize payload sizes for incoming weather station telemetry[cite: 1].

## The Situation

You're working on a custom server which accepts incoming data from weather stations[cite: 1]. You're building the part of the system which transforms incoming messages from weather stations into a format which can be used by other parts of the system[cite: 1].

Those weather stations send their meteorological data every minute, but they only send the data that has changed in the last 60 seconds to keep message payloads as small as possible[cite: 1]. Luckily, those weather stations send a full snapshot every 10 minutes, so there shouldn't be too many gaps in the data if the odd message goes missing[cite: 1].

Your infrastructure is considered as critical national infrastructure, and your senior architect has instructed you not to use any fancy tools for remembering the previous state of the weather station[cite: 1].

You decide to tackle this problem head on with nothing but raw code[cite: 1].

## Functional Requirements

Your objective is to implement a component that can remember each weather station's previous state, so that you can broadcast its full known state when a partial update message arrives[cite: 1]. You decide to make a test to deal with a single weather station, before scaling your solution to deal with all weather stations[cite: 1].

### Overview

Your program will:

1. **Read** lines from standard input until `exit` is entered[cite: 1].
2. **Parse** each line as one of:
   - **Data update**: `<id>,<value>`[cite: 1]
   - **Query**: `get`[cite: 1]
   - **Reset**: `clear`[cite: 1]
   - **Termination**: `exit`[cite: 1]
3. **Maintain** an internal state of nine meteorological sensors[cite: 1].
4. **Print** the full state on `get`[cite: 1].

No menus or prompts should be printed—each line of input triggers exactly one action[cite: 1].

### Program Initialisation

When your program starts, it should print:

```text
--- Weather Station ---
```[cite: 1]

### Data Points

Weather stations send their data as comma separated value strings, and each data point has a numerical code[cite: 1]. The rest of your system works with keys, and so your service must do some conversions[cite: 1].

For historical reasons, some IDs are no longer in use, but some old weather stations may still send data which is no longer interesting to your system[cite: 1]. You should ignore any codes which are no longer in use[cite: 1].

You can use the following table to link data point IDs with their internal strings:

| ID | Key |
|---|---|
| 1 | `airTemp` |
| 2 | `airPressure` |
| 7 | `precipitation` |
| 11 | `windSpeed` |
| 12 | `windDirection` |
| 13 | `humidity` |
| 14 | `dewPoint` |
| 15 | `soilMoisture` |
| 22 | `cloudCover` |[cite: 1]

The values are formatted as floats[cite: 1].

The initial state is that all data points are `NULL`[cite: 1].

### Messages

Messages are received as CSV strings[cite: 1]. For example if `windSpeed` and `humidity` have changed since the last message, the following string would be received:

```text
11,15.5
13,32.3
```[cite: 1]

There are occasions where there is no data to report for a given sensor, in these cases `NULL` values will be sent[cite: 1]. For example, if there is no wind speed, the wind direction cannot be known[cite: 1].

```text
11,0
12,NULL
```[cite: 1]

The inputs will always be valid[cite: 1]. The CSV string will never be malformed[cite: 1]. The first column will always be a positive integer[cite: 1]. The second column will always be a float or `NULL`[cite: 1].

If no data has changed, no message will be sent[cite: 1]. So we'll never need to deal with an empty string[cite: 1]. A full snapshot will still be sent every 10 minutes, so it is possible to detect if a weather station has *probably* gone offline[cite: 1]. Detecting that is a problem for some other service[cite: 1].

### Return State

The state is always reported in full, ordered by ID[cite: 1]. If no state is known about a data point, `NULL` is displayed[cite: 1].

```text
airTemp:21.6
airPressure:31.0
precipitation:0.4
windSpeed:0.0
windDirection:NULL
humidity:9.1
dewPoint:12.3
soilMoisture:33.2
cloudCover:1008.0
```[cite: 1]

### Technical Implementation

Your program must read commands from standard input (stdin) and write responses to standard output (stdout)[cite: 1].

No menus or prompts should be printed—each line of input triggers exactly one action[cite: 1].

When your program starts, it should print:

```text
--- Weather Station ---
```[cite: 1]

Process input lines as follows:

- **`id,value`** — Updates a sensor value[cite: 1].
  - Example: `11,15.5` sets `windSpeed` to `15.5`[cite: 1].
  - Example: `12,NULL` sets `windDirection` to missing[cite: 1].
  - Invalid lines or IDs outside the table are ignored[cite: 1].
- **`get`** — Prints all sensor values, one per line, in ascending ID order[cite: 1].
- **`clear`** — Resets all fields back to missing values (`NULL` when printed)[cite: 1].
- **`exit`** — Prints `Exiting...` and exits the program[cite: 1].

`main.go` needs to be located at the **root** of your `weatherStation` repository[cite: 1]. This is the entry point of the application, maintaining the weather station state and processing input commands[cite: 1]. If you use packages, your **Go module name** should match your repository name: `weatherStation`[cite: 1].

While the location of `main.go` is a strict requirement, you're free to implement the program's internal logic in any way you choose, as long as it produces the expected standard output (`stdout`)[cite: 1].

## Usage

### Getting Initial State

```text
/weatherStation $ go run .
--- Weather Station ---
get
airTemp:NULL
airPressure:NULL
precipitation:NULL
windSpeed:NULL
windDirection:NULL
humidity:NULL
dewPoint:NULL
soilMoisture:NULL
cloudCover:NULL
```[cite: 1]

### Updating States

```text
11,15.5
13,32.3
get
airTemp:NULL
airPressure:NULL
precipitation:NULL
windSpeed:15.5
windDirection:NULL
humidity:32.3
dewPoint:NULL
soilMoisture:NULL
cloudCover:NULL
```[cite: 1]

### Clearing State

```text
clear
get
airTemp:NULL
airPressure:NULL
precipitation:NULL
windSpeed:NULL
windDirection:NULL
humidity:NULL
dewPoint:NULL
soilMoisture:NULL
cloudCover:NULL
```[cite: 1]

### Exiting

```text
exit
Exiting...
```[cite: 1]

## Useful Links

- [Bufio Package (for buffered I/O)](https://pkg.go.dev/bufio)[cite: 1]
- [Go Structs, Methods and Receivers](https://dev.to/jpoly1219/structs-methods-and-receivers-in-go-5g4f)[cite: 1]
