package blockchain

import (
	"encoding/json"
	"fmt"

	"github.com/mihikadusad/HypeFL/blockchain"
)

// GetAVD query the chaincode to get the state of avdID
func (setup *FabricSetup) GetAVD(avdID string) (*model.AVD, error) {

	// Prepare arguments
	funcName := "getAVD"
	args := []string{avdID}

	// invoke the chaincode and return
	payload, err := setup.Query(funcName, args)
	if err != nil {
		fmt.Println("error reading state of AVD : " + err.Error())
		return nil, err
	}
	var avd model.AVD
	err = json.Unmarshal(payload, &avd)
	if err != nil {
		fmt.Println("Error unmarshalling JSON : " + err.Error())
		return nil, err
	}
	return &avd, nil
}
