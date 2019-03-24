# Gojek parking lot problem

## Usage guide
1. Run "setup" script in bin folder.
```
./bin/setup
```
2. This will place executable file in ./target folder which can be run like:
```
./bin/parking_lot
or
./bin/parking_lot "INPUT_FILE"
```
3. Additionally you can run test from project root directory as:
```
go test ./... -v
```
----

## Development reference
1. This app does not have any dependency other than core go. Basic structure of app is described here.
2. Entities used are:
    Interfaces:
        
        Vehicle
    Structs:
        
        Car,
        ParkingLot,
        Slot
    Enums:

        Color
    Packages:

        parkingLot,
        main
3. "parkingLot.appService.go" contains all the application layer services, other files deal with domain.
4. Info comments are written above individual functions in the files.

---