# SIGNER-BACKEND

Signer Backend made with NodeJS and Express, that exposes the REST API to sign payments.
This backend requires the private key of your wallet. This software is a PoC, **so it's very unsafe to type your secret key here!**.


## Environment variables

This backend should run with the following environment variables set:

* PRIVATEKEY: refers to the INFURA URL project, required to connect to a node.
* PORT: port exposed for this application (8080 by default).

## Deploy on Docker

This backend is intended to be deployed on cloud as a service in conjunction with Voice-backend. On development, we create the correspondant image with docker build and then we run the instance with the correspondant environment variables.

```bash
$ docker build -t mrlouzao/signer-backend . 
$ docker run -it --rm --name signer-ethereum -e PRIVATEKEY=<your_key_here> -p 8080:8080 mrlouzao/signer-backend
```

To test the application, send a POST request to your docker:
```bash
http://localhost:8080/signedTransaction?transaction={"nonce":"0x1f","gasPrice":"0xb2d05e00","gasLimit":"0x27100","to":"0xCF46eC111492a5b00246eBBe260746A7E2E76558","value":"0x01","data":"0x00"}
```

You must receive a response like this:
```
{
    "status": 200,
    "data": "0xf8641f84b2d05e008302710094cf46ec111492a5b00246ebbe260746a7e2e7655801002aa0a81ddce221db36ab164ac9c98e17fb912cc590bf18ab12161e8425ed218e5334a05be20bb5a904c3cfcac363d0b86406cfab2550541328c5204547076e6b612e8e",
    "message": "Succesfully Signed Transaction"
}
```

The json shown below gives you the transaction signed to send to Infura on your voice-backend.

Once the application is running properly on your localhost, you must push the image to your registry (in our case Google Cloud Registry) and create a VM running this docker image. The container created requires the environment variable PRIVATEKEY.


## Wallets for testing

**Important**: for this initial version (PoC) the wallet used is the Testing 1. In next versions, an env var will be created to pass the public address of the wallet as a parameter when creating the docker container instance.

We have the following wallets:

- Testing 1 wallet:
 - Public: 0x007BfF585Be4B690db8FBC9FDe4F936294Fa37De
 - Private: 109ACFDD4C5B16564EC4C3AC8EC5AC7EFC858B163841A95D15C85E1B877DB0BD
