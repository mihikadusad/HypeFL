# HypeFL
HypeFL is a novel machine learning framework built on Hyperledger Fabric that combines machine learning and blockchain to create a decentralized, privacy-preserving fully autonomous vehicle system. Through cooperative perception, each vehicle communicates its LiDAR sensor detections to other vehicles in the network to incorporate knowledge and localize surrounding vehicles in the system, combatting visual occlusions and corner cases. Our system utilizes federated learning to optimize data privacy by only sharing model parameters between vehicles, rather than raw data. The blockchain serves as an immutable, decentralized server that stores vehicles and model parameters as nodes, providing protection against single-point failures and eliminating the risk of malicious attacks.

## Prerequisites
* Install CARLA 0.9.12
