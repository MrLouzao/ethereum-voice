const express = require('express');
const router = express.Router();

const signerController = require('./controllers/signerController');
const userController = require('./controllers/userController');

router.get('/signedTransaction', signerController.signTransaction);
router.get('/users/signedTrasnaction',signerController.signForUser);
router.get('/test', signerController.test);
router.get('/users',userController.getUsers);
router.post('/users',userController.addUsers);
module.exports = router;
