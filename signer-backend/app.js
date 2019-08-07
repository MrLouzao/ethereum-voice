//express import
const express = require('express');
const bodyParser = require('body-parser');

//Database
const mongo = require('mongodb');
const monk = require('monk');
const db = monk('localhost:27017/testDb');

const router = require('./router');

const app = express();

app.use(function(req,res,next){
    req.db = db;
    next();
});
app.use(bodyParser.json());
app.use('/',router);
// Make our db accessible to our router

app.listen(process.env.PORT || 8080, () => {
    console.log("Serve started in port " + process.env.PORT || 8080);
    console.log(`${process.env.PRIVATEKEY}`);
});

module.exports = app;