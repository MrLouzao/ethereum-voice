const service = require('../services/userService');

exports.getUsers = function (req, res, next) {
    // Validate request parameters, queries using express-validator
    var db = req.db;
    service.getUsers(db)
    .then((usersList) => { res.status(200).json({ status: 200, data: usersList, message: "Succesfully get users list" }) })
    .catch((e)=>{res.status(400).json({ status: 400, message: e.message })});
};

exports.addUsers = function (req, res, next) {
    // Validate request parameters, queries using express-validator
    var db = req.db;
    let user = req.body;
    console.log(user);
    service.addUser(db,user)
    .then((_id) => { res.status(200).json({ status: 200, data: _id, message: "Succesfully set user" }) })
    .catch((e)=>{res.status(400).json({ status: 400, message: e.message })});
};