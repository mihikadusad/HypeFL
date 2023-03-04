package blockchain

// CreateAVD
func (setup *FabricSetup) CreateAVD(licensePlate string, vehicleModel string,
	vinNumber string, userName string) (string, error) {

	// Prepare arguments
	funcName := "createAVD"
	args := []string{licensePlate, vehicleModel, vinNumber, userName}

	return setup.Invoke(funcName, args)
