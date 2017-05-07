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
	FirstName    	string
	LastName		string
	VehicleNumber	string
	Make			string
	Model			string
	ManYear         string
	RegNo			string
	RegState		string
	ECC				string
	Status 			string
	MetroInsurance 	string
	AvonInsurance  	string
	BharatiInsurance	string
	Rating 			string

}

func main() {
	err := shim.Start(new(Policy))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}


func (t *Policy) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	msg="In side Init"
	return []byte(msg), nil
}

func (t *Policy) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)
	// Handle different functions
	
	fmt.Println("invoke did not find func: " + function)
  if function == "updatePolicy" {
        return t.updatePolicy(stub, args)
    }else if function=="createPolicy" {
			 return t.createPolicy(stub, args)		
	}
	return nil, errors.New("Received unknown function invocation: " + function)
	
	
}

// Query is our entry point for queries
func (t *Policy) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	

	// Handle different functions
	 if function == "getPolicy" {   
        	return t.getPolicy(stub, args)
    }else if function == "getAllPolicies" {   
        	return t.getAllPolicies(stub, args)
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
	var key,key2,key3,key4,key5, jsonResp string
	var err error
	var msg string

	key = args[0]
	key2= args[1]
	key3=args[2]
	key4=args[3]
	key5=args[4]
	
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	src_json:=[]byte(valAsbytes)
u := PolicyDetails{}
	json.Unmarshal(src_json, &u)
	//ponits are hardcoded as of now can be made dynamic by getting value from args
	if key2== "status" {
			u.Status=key3
				
	
	}
	if key4=="MetroInsurance" {
					u.MetroInsurance=key5
					
	}else if key4=="AvonInsurance"{
		u.AvonInsurance=key5
	}else if key4=="BharatiInsurance"{
		u.BharatiInsurance	=key5
	}	
	json_byte, err:=json.Marshal(u);
if err != nil {
						msg="Some error occured"
		panic(err)
	}

	err = stub.PutState(key, json_byte)
if err != nil {
		msg="UnSuccesful"
			return []byte(msg), err
			panic(err)
	}
	msg="Success"
	
	

	return []byte(msg), nil
}

func (t *Policy) createPolicy(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var msg, jsonResp string
	if len(args) <= 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting more then 1")
	}

u:=PolicyDetails{}
u.FirstName=args[0]
u.LastName=args[1]
u.VehicleNumber=args[2]
u.Make=args[3]
u.Model=args[4]
u.ManYear=args[5]
u.RegNo=args[6]
u.RegState=args[7]
u.ECC=args[8]
u.Status=args[9]
u.MetroInsurance=args[10]
u.AvonInsurance=args[11]
u.BharatiInsurance=args[12]
u.Rating=args[13]


//for checking duplicacy
 status:=t.checkDuplicacy(stub,u.VehicleNumber)
 
if status {
	jsonResp = "{\"Error\":\"Vehicle id is already there" + u.VehicleNumber + "\"}"
		return []byte(jsonResp), errors.New(jsonResp)}


//end for duplicacy



//for all policyIds
policy_no, dt:=stub.GetState("policyIds")
if dt != nil {
		jsonResp = "{\"Error\":\"Failed to get state for key policyId}"
		return nil, errors.New(jsonResp)
	}
if policy_no != nil{

	policy:=string(policy_no)
	
	stub.PutState("policyIds",[]byte(policy+ ","+u.VehicleNumber))

}else {
	stub.PutState("policyIds",[]byte(u.VehicleNumber))
}

//end of get all policyIds
json_byte, err:=json.Marshal(u);
	
	err = stub.PutState(u.VehicleNumber, json_byte)
	if err != nil {
		msg="UnSuccesful"
			return []byte(msg), err
			panic(err)
	}
	msg="Success"
	return []byte(msg), nil
}

func (t *Policy) getAllPolicies(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		var jsonResp string
	policy_no, dt:=stub.GetState("policyIds")
if dt != nil {
		jsonResp = "{\"Error\":\"Failed to get state for key policyId}"
		return nil, errors.New(jsonResp)
		panic(dt)
	}
	

	return policy_no, nil
}

func (t *Policy) checkDuplicacy(stub shim.ChaincodeStubInterface, args string) (bool) {
		
		var res bool
	policy, dt:=stub.GetState(args)
	if dt !=nil {
		res =false
		
		panic(dt)
	}
	if policy !=nil {
		res=true
	
	}else{
	res= false
	}
	
	return res
}



