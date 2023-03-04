package blockchain

// UpdateAVD: Add new data to avd
func (setup *FabricSetup) UpdateAVD(avdID string, newID string, newData string) (string, error) {

	// Prepare arguments
	funcName := "updateAVD"
	args := []string{avdID, newID, newData}

	return setup.Invoke(funcName, args)
}
