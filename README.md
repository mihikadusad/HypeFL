# HypeFL
HypeFL is a novel machine learning framework built on Hyperledger Fabric that combines machine learning and blockchain to create a decentralized, privacy-preserving fully autonomous vehicle system. Through cooperative perception, each vehicle communicates its LiDAR sensor detections to other vehicles in the network to incorporate knowledge and localize surrounding vehicles in the system, combatting visual occlusions and corner cases. Our system utilizes federated learning to optimize data privacy by only sharing model parameters between vehicles, rather than raw data. The blockchain serves as an immutable, decentralized server that stores vehicles and model parameters as nodes, providing protection against single-point failures and eliminating the risk of malicious attacks.

## Prerequisites
* Windows 10 PC with at least 165 GB of free space
* 6 GB GPU (for Unreal Engine)
* Install [CARLA 0.9.12](https://carla.readthedocs.io/en/latest/start_quickstart/#carla-installation)
* Install [Python 3.8](https://www.python.org/downloads/release/python-380/)
* Install [Node.js](https://nodejs.org/en/)
* Install [Go IDE](https://www.jetbrains.com/go/promo/?source=google&medium=cpc&campaign=10160687272&term=go%20compiler%20download&content=631311299925&gclid=Cj0KCQjwn9CgBhDjARIsAD15h0B48e-3zgtB2nwGr8Qn19N5LmIV8Bqj04xgzf1p2wQouJG3eX5qkWUaAqF8EALw_wcB)

### Install Docker
``` 
    chmod +x docker.sh
    sudo ./docker.sh
    usermod -a -G docker ${USER}
```

## Training
After cloning this Github into your directory, begin by training the federated learning object detection models.
```
cd Federated Learning
python3 fl_train.py
```

## Testing
```
    python3 fl_test.py
```

## Running HypeFL
```
    go run main.go
```
