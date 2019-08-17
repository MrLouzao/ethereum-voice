# VOICE-BACKEND

Voice Backend written in Golang that exposes the REST API to perform payments over Ethereum network.
This backend requires a signer backend running previously and an Infura endpoint.


## Environment variables

This backend should run with the following environment variables set:

* INFURA_URL: refers to the INFURA URL project, required to connect to a node.
* SIGNER_URL: refers to the container that signs your transactions.
* PORT: port exposed for this application (8000 by default).

## Deploy on Docker

This backend is intended to be deployed on cloud as a service. On development, we create the correspondant image with docker build and then we run the instance with the correspondant environment variables.

```bash
$ docker build -t mrlouzao/voice-backend . 
$ docker run -it --rm --name voice-ethereum -p 8000:8000 mrlouzao/voice-backend
```

To test the application, send a POST request to your docker:
```bash
$ curl -X POST -d '{"name": "Rablo", "amount": 0.23}' http://localhost:8000/voice-payment
```

You must receive a response like this:
```
{
  "fulfillmentText": "Rablo va a recibir 0.230000 ether en su wallet. Mira aquí la transacción: https://ropsten.etherscan.io/tx/0xa223f1cef82127af1ad545793f22fc0f3200be78de84dd82428898ea3adab602"
}
```

Once the application is running properly on your localhost, you must push the image to your registry (in our case Google Cloud Registry) and create a VM running this docker image. The container created requires the environment variables INFURA_URL and SIGNER_URL.


## Wallets for testing

**Important**: for this initial version (PoC) the wallet used is the Testing 1. In next versions, an env var will be created to pass the public address of the wallet as a parameter when creating the docker container instance.

We have the following wallets:

- Testing 1 wallet:
 - Public: 0x007BfF585Be4B690db8FBC9FDe4F936294Fa37De
 - Private: 109ACFDD4C5B16564EC4C3AC8EC5AC7EFC858B163841A95D15C85E1B877DB0BD


- Testing 2 wallet (Rablo):
 - Public: 0xd0CfE66448093dDA6cdC6525DB0C66BF1DD9c138
 - Private: C01CF93A477C39311DD2C12EF2E4DDC6E8315498C7B712849BA71A49EDAFFD8D


- Louzao wallet:
  - Password: ABCD1234efg5ZYX
  - Public: 0xde95aff743b29b72885a27d03795632f6e741fdf
  - Keystore: {"version":3,"id":"89fa7fd9-50fd-4f62-84df-bad68c83dc4f","address":"de95aff743b29b72885a27d03795632f6e741fdf","crypto":{"ciphertext":"4ac94f52514104037a5db645159998b1d64040099fe9889008a1544d360ecbe4","cipherparams":{"iv":"67bee63dcfd69cfcaa346c0b57741fd3"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"25446cc3b26a73a129e1c461a26d071ab3454fb8cde891dc09440c83d0212df9","n":131072,"r":8,"p":1},"mac":"02b690f23d58c6c17db2d6855cb40ec71c9669b945c06dc01be33257184c432d"}}


- Pablo wallet:
  - Password: ABCD1234efg5ZYX
  - Public: 0x4f41045c96c2e2eca0735e6a39002f4edce12fbe
  - Keystore: {"version":3,"id":"97df9897-043c-49ef-900b-e606d5a4f4c9","address":"4f41045c96c2e2eca0735e6a39002f4edce12fbe","crypto":{"ciphertext":"6b3fc38841fe3cfff2314ed1c400d88496ff3da5ea342364e91449e4d3ca6e2c","cipherparams":{"iv":"821bedd46a2b68f5c9af53093652ed3a"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"1681430953d1e7d494ea4c3054720bdad6d824e7c995186fb06af4c7aa36a0c7","n":131072,"r":8,"p":1},"mac":"287df5b59ce5e56a70b80b0b48357331f870baa6cc359216b45900187d81a20f"}}