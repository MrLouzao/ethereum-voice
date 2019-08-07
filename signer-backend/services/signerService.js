const EthereumTx = require('ethereumjs-tx').Transaction
const keythereum = require('keythereum');

/**
 * Sign the payload data
 * @param {object} transaction transaction to be signed
 * @return {string} tx hash
 */
exports.signTransaction = async function (transaction) {
    try {
        const txParams = {
            nonce: transaction.nonce,
            gasPrice: transaction.gasPrice,
            gasLimit: transaction.gasLimit,
            to: transaction.to,
            value: transaction.value,
            data: transaction.data,
          }
        //, { chain: 'mainnet', hardfork: 'petersburg' } could be useful
        const tx = new EthereumTx(txParams, { chain: 'ropsten', hardfork: 'petersburg' });
        var privateKey = Buffer.from(process.env.PRIVATEKEY.toString(), 'hex');
        tx.sign(privateKey);
        const signedTx = `0x${tx.serialize().toString('hex')}`;
        return signedTx;
    } catch (err) {
        console.log(err);
        throw err;
    }
}

exports.signForUser = function (transaction, user, db, password) {
    const tx = new EthereumTx(transaction, { chain: 'mainnet', hardfork: 'petersburg' });
    return new Promise((resolve,reject)=>{
        getFromDatabasePrivateKey(user, db, password)
        .then((privateKey)=>{
            const txParams = {
                nonce: transaction.nonce,
                gasPrice: transaction.gasPrice,
                gasLimit: transaction.gasLimit,
                to: transaction.to,
                value: transaction.value,
                data: transaction.data,
              }
            //, { chain: 'mainnet', hardfork: 'petersburg' } could be useful
            const tx = new EthereumTx(txParams);
            tx.sign(privateKey);
            let signedTx = `0x${tx.serialize().toString('hex')}`;
            resolve(signedTx);
        })
        .catch((e)=>{reject(e)});
    });
}

const getFromDatabasePrivateKey = (user, db, password) => {
    let users = db.get('users');
    return new Promise((resolve,reject)=>{
        users.find({ name: user })
        .then((targetUser)=>{
            let privateKey = keythereum.recover(password, targetUser[0].keystore);
            resolve(privateKey);
        })
        .catch((e)=>{reject(e)});
    });
}