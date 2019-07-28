package abi

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"

	_ "math/big"

	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
)

type Compiler struct {
	mapContracts map[string]*compiler.Contract
	binData      string
}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (c *Compiler) SolidityCompile(filename string) error {
	var (
		err error
	)
	c.mapContracts, err = compiler.CompileSolidity("solc", filename)
	if err != nil {
		return err
	}

	return nil
}

func (c *Compiler) GetContractsNames() []string {
	var names = []string{}
	for name, _ := range c.mapContracts {
		part := strings.Split(name, ":")
		names = append(names, part[1])
	}
	return names
}

func (c *Compiler) GetContracts() map[string]*compiler.Contract {
	return c.mapContracts
}

func (c *Compiler) GetComipleData(contractName string) (string, string, error) {
	var contract *compiler.Contract
	for name, ct := range c.mapContracts {
		part := strings.Split(name, ":")
		if part[1] == contractName {
			contract = ct
		}
	}

	if contract == nil {
		return "", "", errors.New(contractName + " not exist ... ")
	}

	abiData, err := json.Marshal(contract.Info.AbiDefinition) // Flatten the compiler parse
	return contract.Code, string(abiData), err
}

// if method == "" , using the constructor method
func (c *Compiler) GetDeployData(contractName string, method string, params ...interface{}) (string, error) {
	binData, abiData, err := c.GetComipleData(contractName)
	if err != nil {
		return "", err
	}

	abiObj, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		return "", err
	}

	byt, err := abiObj.Pack(method, params...)
	if err != nil {
		return "", err
	}

	log.Print("abiBin : ", hex.EncodeToString(byt))
	return binData + hex.EncodeToString(byt), err
}
