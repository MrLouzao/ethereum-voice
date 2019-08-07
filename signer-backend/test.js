const EthereumTx = require('ethereumjs-tx').Transaction
const privateKey = Buffer.from(
  '3a1076bf45ab87712ad64ccb3b10217737f7faacbf2872e88fdd9a537d8fe266',
  'hex',
)

const txParams = {
  nonce: '0x01',
  gasPrice: '0x00',
  gasLimit: '0x27100',
  to: '0xb7af14feb10df73b16c538753f8d34d2e7106c4f',
  value: '0x01',
  data: '0x00',
}


// The second parameter is not necessary if these values are used
const tx = new EthereumTx(txParams, { chain: 'mainnet', hardfork: 'petersburg' })
tx.sign(privateKey)
const serializedTx = `0x${tx.serialize().toString('hex')}`
console.log(serializedTx);