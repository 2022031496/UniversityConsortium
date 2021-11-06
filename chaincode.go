package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)
type Chaincode struct {
}
type Teacher struct {
	Name string ‘json:“name”’
	University string ‘json:“university”’
	Course string ‘json:“course”’
}
type Student struct {
	Name string ‘json:“name”’
	University string ‘json:“university”’
	Hash string ‘json:“hash”’
	Grade int ‘json:“grade”’
	Credit int ‘json:“credit”’
}
func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == “CheckCourse” {
		return t.CheckCourse(stub, args)
	} else if function == “LearnCourse” {
		return t.LearnCourse(stub, args)
	} else if function == "WorkUpload” {
		return t.WorkUpload(stub, args)
	} else if function == “AddGrade” {
		return t.AddGrade(stub, args)
	} else if function == “CheckStudent” {
		return t.CheckStudent(stub, args)
	} else if function == “AwardCredit” {
		return t.AwardCredit(stub, args)
	}
	return shim.Error("Invalid Smart Contract function name. ”)
}

func (t *Chaincode) CheckCourse(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("nu Incorrect mber of arguments. Expecting 4")
	}
	var teacher = Teacher{Name: args[1], University: args[2], Course: args[3]}
	teacherAsBytes, _ := json.Marshal(teacher)
	APIstub.PutState(args[0], teacherAsBytes)
	return shim.Success(nil)
}

func (t *Chaincode) LearnCourse(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	var student = Student{Name: args[1], University: args[2], Hash: args[3], Grade: args[2], Credit: args[3]}
	studentAsBytes, _ := json.Marshal(student)
	APIstub.PutState(args[0], studentAsBytes)
	return shim.Success(nil)
}

func (t *Chaincode) WorkUpload(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	studentAsBytes, _ := APIstub.GetState(args[0])
	student := Student{}
	json.Unmarshal(studentAsBytes, &student)
	student.Hash = args[3]
	studentAsBytes, _ = json.Marshal(student)
	APIstub.PutState(args[0], studentAsBytes)
	return shim.Success(nil)
}

func (t *Chaincode) AddGrade(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	studentAsBytes, _ := APIstub.GetState(args[0])
	student := Student{}
	json.Unmarshal(studentAsBytes, &student)
	student.Grade = args[4]
	studentAsBytes, _ = json.Marshal(student)
	APIstub.PutState(args[0], studentAsBytes)
	return shim.Success(nil)
}

func (t *Chaincode) CheckStudent(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	studentAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(studentAsBytes)
}

func (t *Chaincode) AwardCredit(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	studentAsBytes, _ := APIstub.GetState(args[0])
	student := Student{}
	json.Unmarshal(studentAsBytes, &student)
	student.Credit = args[5]
	studentAsBytes, _ = json.Marshal(student)
	APIstub.PutState(args[0], studentAsBytes)
	return shim.Success(nil)
}
