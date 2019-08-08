# Ethereum Voice Payments

![Ethereum Voice](/docs/eth-voice-logo.jpg?raw=true "Ethereum Voice")

All files required to deploy an Ethereum voice payment agent (bot) on Dialogflow and integrate on Google Assistant. This bot will trigger Ethereum transactions against the Ethereum Ropsten testnet with Infura.

## Project content structure

* Ethereum-payment-agent.zip: dialogflow project that contains all required files to deploy the agent.
* ethereum-voice-caller-lambda: lambda function to deploy on gcloud as Node REST endpoint. The purpose of this lambda is to intercept all messages from assistant and cast message to one that the backend can understand.
* voice-backend: server that provides a REST API that handles Ethereum transactions. Written in GO lang.
* signer-backend: server that publishes an API to sign raw transactions on Ethereum.

## Architecture Overview

![Ethereum Voice Architecture Overview](/docs/architecture-overview-eth-voice.jpg?raw=true "Ethereum Voice Architecture Overview")


The proposed architecture is intended to be deployed on Gcloud infraestructure.


## Deploy on your cloud

To deploy this project on your cloud, follow the next steps:

1. Create correspondant images for voice-backend and signer-backend. Push them to GCR (Google Cloud Registry, use your own registry of your favourite provider) image repository.

2. Create on Infura your own project and copy the Ropsten URL. Don't use this application on mainnet!

3. Deploy a VM on GCloud with the voice-backend and signer-backend properly configured. Check the README files for each component to know the environment variables required. The scheme proposed is:
  - Run signer-backend with your private key on port 8080.
  - Run voice-backend pointing to Infura and signer-backend as localhost:8080
  - When running voice-backend, map the VM port 80 to 8000 of voice-backend instance
  - Get the voice-backend public IP.

4. Deploy Gcloud function to communicate with your backend. This lambda is intended to process all incoming messages from Google Assistant and forward in a format that the voice-backend is able to understand.
  - When deploying the lambda, use the URL from voice-backend on env variable BACKEND_IP.
  - Copy the https secure URL for this lambda.

5. On dialogflow, use the files on Ethereum-payment-agent.zip to deploy the project of voice recognition. Once deployed, on integrations point to the https URL from point 4. This is required to forward recognized messages from assistant to our voice-backend.

6. On Google Assistant test the application.

Thats all!


## For future versions

This is a Proof of Concept. On future versions the following features will be released:
 - Remove hardcoded wallets from voice-backend and add support new users with REST API
 - Integration of Google API to recognize who is sending the voice message and associate it to the user on voice-backend
 - Add support to contacts for each user
 - Add frontend to deploy an administration view in which the user will be able to perform:
   - Add a new contact
   - Set my wallet secret key as a keystore
   - Generate a recover seed for the wallet


A post on Medium will be released soon with more explanations about this project.


## Built With

* [Dialogflow](https://dialogflow.com) - Voice recognition tool
* [Google Cloud](https://cloud.google.com/) - Cloud platform to deploy all transactions

## Authors

* **Luis Miguel Louzao** -  [Github](https://github.com/MrLouzao)
* **Iñigo** - [Github](https://github.com/nefera606)

See also the list of [contributors](https://github.com/MrLouzao/ethereum-voice/contributors) who participated in this project.
