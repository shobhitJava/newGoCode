package main

import (

	"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

type Policy struct{
	
}
type PolicyDetails struct {
	FirstName    string
	LastName	string
	VehicleNumber	string
	Make	string
	Model	string
	RegNo			string
	RegState		string
	ECC		string
}

func main() {
	err := shim.Start(new(Policy))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


func (t *Policy) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	fmt.Print("inside init method" + function)
	if len(args) <= 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting more then 1")
	}

u:=PolicyDetails{}
u.FirstName=args[0]
u.LastName=args[1]
u.VehicleNumber=args[2]
u.Make=args[3]
u.Model=args[4]
//u.RegNo=args[5]	
//u.RegState=args[6]
//u.ECC=args[7]


json_byte, err:=json.Marshal(u);
	//hardcoded the key since not using the DB
	err = stub.PutState(u.VehicleNumber, json_byte)
	if err != nil {
		msg="UnSuccesful"
			return []byte(msg), err
	}
	msg="Success"
	return []byte(msg), nil
}

func (t *Policy) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)
	// Handle different functions
	
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
	
	
}

// Query is our entry point for queries
func (t *Policy) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	 if function == "getPolicy" {
        	return t.getPolicy(stub, args)
    } else if function == "updatePolicy" {
        return t.updatePolicy(stub, args)
    } 
	
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}


// read - query function to read key/value pair
func (t *Policy) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	return valAsbytes, nil
}

func (t *Policy) getPolicy(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	return valAsbytes, nil
}

func (t *Policy) updatePolicy(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key,key2, jsonResp string
	var err error
	var msg string
	if len(args) <= 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	key2= args[1]
	fmt.Println(key2+" is the new key2")
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	src_json:=[]byte(valAsbytes)
u := PolicyDetails{}
	json.Unmarshal(src_json, &u)
	//ponits are hardcoded as of now can be made dynamic by getting value from args
	if key2== "firstName" {
			u.FirstName=key2
			
	
	}else if key2=="lastName"	{
					u.LastName=key2
					
	}	
	json_byte, err:=json.Marshal(u);
if err != nil {
						msg="Some error occured"
		panic(err)
	}

	err = stub.PutState(key, json_byte)
	

	return []byte(msg), nil
}


