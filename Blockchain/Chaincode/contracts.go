package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	layout = "2006-01-02"
)

// AVD: Autonomous Vehicle Data Block
type AVD struct {
	ID                      string        `json:"id"`
	LicensePlate            string        `json:"licensePlate"`
	VehicleModel            string        `json:"vehicleModel"`
	VinNumber               string        `json:"vinNumber"`
	DataTransmission        string        `json:"dataTransmission"`
	SurroundingVehicleData  []AVD         `json:"nearbyVehicles"`
}

// HeroesServiceChaincode implementation of Chaincode
type HeroesServiceChaincode struct {
}

// Init of the chaincode
// This function is called only one when the chaincode is instantiated.
// So the goal is to prepare the ledger to handle future requests.
func (t *HeroesServiceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### HeroesServiceChaincode Init ###########")

	// Get the function and arguments from the request
	function, _ := stub.GetFunctionAndParameters()

	// Check if the request is the init function
	if function != "init" {
		return shim.Error("Unknown function call")
	}

	var avd AVD
	avd.ID = "placeholder"
	avd.LicensePlate = "placeholder"
	avd.VehicleModel = "placeholder"
	avd.VinNumber = "placeholder"
	avd.DataTransmission = nil
	avd.SurroundingVehicleData = nil

	behr, err := json.Marshal(avd)
	if err != nil {
		return shim.Error("error marshalling AVD to Json")
	}

	// Put in the ledger the key/value hello/world
	err = stub.PutState("hello", behr)
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return a successful message
	return shim.Success(nil)
}

// Invoke of the chaincode
// All future requests named invoke will arrive here.
func (t *HeroesServiceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### HeroesServiceChaincode Invoke ###########")

	// Get the function and arguments from the request
	function, args := stub.GetFunctionAndParameters()

	// Handle different functions
	switch function {
	case "createAVD":
		return createAVD(stub, args)
	case "getAVD":
		return getAVD(stub, args)
	case "updateAVD":
		return updateAVD(stub, args)
	}

	// Check whether it is an invoke request
	if function != "invoke" {
		return shim.Error("Unknown function call")
	}

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// In order to manage multiple type of request, we will check the first argument.
	// Here we have one possible argument: query (every query request will read in the ledger without modification)
	if args[0] == "query" {
		return t.query(stub, args)
	}

	// The update argument will manage all update in the ledger
	if args[0] == "invoke" {
		return t.invoke(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

// query
// Every readonly functions in the ledger will be here
func (t *HeroesServiceChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### HeroesServiceChaincode query ###########")

	// Check whether the number of arguments is sufficient
	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	// Like the Invoke function, we manage multiple type of query requests with the second argument.
	// We also have only one possible argument: hello
	if args[1] == "hello" {

		// Get the state of the value matching the key hello in the ledger
		state, err := stub.GetState("hello")
		if err != nil {
			return shim.Error("Failed to get state of hello")
		}

		// Return this value in response
		return shim.Success(state)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown query action, check the second argument.")
}

// invoke
// Every functions that read and write in the ledger will be here
func (t *HeroesServiceChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("########### HeroesServiceChaincode invoke ###########")

	if len(args) < 2 {
		return shim.Error("The number of arguments is insufficient.")
	}

	avdID := args[1] //stub.GetTxID
	value := args[2]
	// Check if the ledger key is "hello" and process if it is the case. Otherwise it returns an error.
	if avdID == "hello" && len(args) == 3 {

		// Add random suffix to the value
		value = value + strconv.Itoa(time.Now().Nanosecond())
		// Write the new value in the ledger
		err := stub.PutState(avdID, []byte(value))
		if err != nil {
			return shim.Error("Failed to update state of hello")
		}

		// Notify listeners that an event "eventInvoke" have been executed (check line 19 in the file invoke.go)
		err = stub.SetEvent("eventInvoke", []byte{})
		if err != nil {
			return shim.Error(err.Error())
		}

		// Return this value in response
		return shim.Success(nil)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown invoke action, check the second argument.")
}

// ==========================================================================================
//	createAVD - create a donor-recipient pair of data transmission from an AV
// ==========================================================================================
func createAVD(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	fmt.Println("running the function createPair()")

	if len(args) != 4 {
		return shim.Error("Wrong input")
	}

	var avd AVD
	avdID := stub.GetTxID()
	avd.ID = avdID
	avd.LicensePlate = args[0]
	avd.VehicleModel = args[1]
	avd.VinNumber = args[2]
	avd.DataTransmission, err = args[3]
	avd.SurroundingVehicleData = nil

	if err != nil {
		return shim.Error("Error parsing data")
	}

	jsonAVD, err := json.Marshal(ehr)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("Error marshalling to JSON")
	}

	err = stub.PutState(avdID, jsonAVD)
	if err != nil {
		return shim.Error("createAVD() : Error writing to state")
	}

	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(ehrID))
}

// ==========================================================================================
// getAVD : query to get an AVD by its key
// ==========================================================================================
func getAVD(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var avdID string
	var err error

	if len(args) != 1 {
		return shim.Error("Wrong input")
	}
	avdID = args[0]
	valAsbytes, err := stub.GetState(avdID)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	} else if valAsbytes == nil {
		fmt.Println("AVD does not exist")
		return shim.Error("AVD does not exist")
	}

	return shim.Success(valAsbytes)
}

// ==========================================================================================
// updateAVD : get an AVD by its key and add new data to the block
// ==========================================================================================
func updateAVD(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var avdID string
	var err error
	var avd *AVD

	if len(args) != 3 { // ehrID, DrID, comment
		return shim.Error("Wrong input")
	}
	avdID = args[0]
	avd, err = getAVDbyID(stub, avdID)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	}
	if avd == nil {
		fmt.Println("Error reading state : AVD is nil")
		return shim.Error("nil avd")
	}
	err = avd.addAppointment(args[1], args[2])
	if avd != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	}

	jsonAVD, err := json.Marshal(avd)
	if err != nil {
		fmt.Println(err.Error())
		return shim.Error("error marshalling json" + err.Error())
	}

	err = stub.PutState(avdID, jsonAVD)
	if err != nil {
		return shim.Error("updateAVD() : Error put state")
	}

	// Notify listeners that an event "eventInvoke" has been executed
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(jsonAVD)
}

// ==========================================================================================
// add data to AVD
// ==========================================================================================
func (avd *AVD) addData(location string, new_data string) error {

	_now := time.Now()
	yyyy := _now.Year()
	MM := _now.Month()
	dd := _now.Day()
	hh := _now.Hour()
	mm := _now.Minute()
	_now = time.Date(yyyy, MM, dd, hh, mm, 0, 0, time.UTC)
	avd.DataTransmission = append(avd.DataTransmission, {location, _now, new_data})
	return nil
}

// ==========================================================================================
// getAVDbyID : get the AVD object by ID - Auxiliary function
// ==========================================================================================
func getAVDbyID(stub shim.ChaincodeStubInterface, ID string) (*AVD, error) {
	valAsbytes, err := stub.GetState(ID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else if valAsbytes == nil {
		return nil, errors.New("AVD does not exist")
	}

	var avd AVD
	err = json.Unmarshal(valAsbytes, &ehr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("Error unmarshalling JSON")
	}

	return &avd, nil
}

func main() {
	// Start the chaincode and make it ready for futures requests
	err := shim.Start(new(HeroesServiceChaincode))
	if err != nil {
		fmt.Printf("Error starting Heroes Service chaincode: %s", err)
	}
}
