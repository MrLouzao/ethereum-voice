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

// TODO: add all required documentation here


## Built With

* [Dialogflow](https://dialogflow.com) - Voice recognition tool
* [Google Cloud](https://cloud.google.com/) - Cloud platform to deploy all transactions

## Authors

* **Luis Miguel Louzao** -  [Github](https://github.com/MrLouzao)
* **Iñigo** - [Github](https://github.com/nefera606)

See also the list of [contributors](https://github.com/MrLouzao/ethereum-voice/contributors) who participated in this project.
