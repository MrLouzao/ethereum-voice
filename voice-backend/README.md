# VOICE-BACKEND

This is the voice backend to expose the REST API.


## Environment variables

This backend should run with the following environment variables set:

* INFURA_URL: refers to the INFURA URL project, required to connect to a node.
* SIGNER_URL: refers to the container that signs your transactions.

## Deploy on Docker

This backend is intended to be deployed on docker. First we create the image with docker build and then we run the instance with the correspondant environment variables.

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


## Wallets for testing

- Louzao wallet:
  - Password: ABCD1234efg5ZYX
  - Public: 0xde95aff743b29b72885a27d03795632f6e741fdf
  - Keystore: {"version":3,"id":"89fa7fd9-50fd-4f62-84df-bad68c83dc4f","address":"de95aff743b29b72885a27d03795632f6e741fdf","crypto":{"ciphertext":"4ac94f52514104037a5db645159998b1d64040099fe9889008a1544d360ecbe4","cipherparams":{"iv":"67bee63dcfd69cfcaa346c0b57741fd3"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"25446cc3b26a73a129e1c461a26d071ab3454fb8cde891dc09440c83d0212df9","n":131072,"r":8,"p":1},"mac":"02b690f23d58c6c17db2d6855cb40ec71c9669b945c06dc01be33257184c432d"}}


- Pablo wallet:
  - Password: ABCD1234efg5ZYX
  - Public: 0x4f41045c96c2e2eca0735e6a39002f4edce12fbe
  - Keystore: {"version":3,"id":"97df9897-043c-49ef-900b-e606d5a4f4c9","address":"4f41045c96c2e2eca0735e6a39002f4edce12fbe","crypto":{"ciphertext":"6b3fc38841fe3cfff2314ed1c400d88496ff3da5ea342364e91449e4d3ca6e2c","cipherparams":{"iv":"821bedd46a2b68f5c9af53093652ed3a"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"1681430953d1e7d494ea4c3054720bdad6d824e7c995186fb06af4c7aa36a0c7","n":131072,"r":8,"p":1},"mac":"287df5b59ce5e56a70b80b0b48357331f870baa6cc359216b45900187d81a20f"}}