const service = require('../services/signerService');

exports.signTransaction = async function(req, res, next) {
    // Validate request parameters, queries using express-validator
    
    var tx = JSON.parse(req.query.transaction);
    service.signTransaction(tx)
    .then((signedTx)=>{res.status(200).json({ status: 200, data: signedTx, message: "Succesfully Signed Transaction" })})
    .catch((e)=>{res.status(400).json({ status: 400, message: e.message })})
};

exports.signForUser = async function(req, res, next) {

    var tx = JSON.parse(req.query.transaction);
    var user = req.query.user;
    var db = req.db;
    var password = req.query.password;
    service.signForUser(tx, user, db, password)
    .then((signedTx)=>{res.status(200).json({ status: 200, data: signedTx, message: "Succesfully Signed Transaction" })})
    .catch((e)=>{res.status(400).json({ status: 400, message: e.message })})
}

exports.test = async function(req,res, next) {

    var test = req.query;
    try {
        return res.status(200).json({status: 200, message: `Ok`, data: test});
    } catch (e) {
        return res.status(400).json({ status: 400, message: e.message });
    }
}