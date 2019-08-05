/**
 * Responds to any HTTP request.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
const fetch = require('node-fetch');

exports.helloWorld = async (req, res) => {
  let message = req.query.message || req.body.message || 'Hello World!';
  
  console.log("req.query = ", req.body);
  console.log("PARA: ", req.body.queryResult.parameters.namelist);
  console.log("CANTIDAD: ", req.body.queryResult.parameters.number[0]);
  
  const recipient = req.body.queryResult.parameters.namelist;
  const amount = req.body.queryResult.parameters.number[0];
  const paymentRequestBody = `{"name": "${recipient}", "amount": ${amount}}`;
  
  const URL = 'http://104.198.177.32/voice-payment';
  var voiceResponse = await fetch(URL, 
     {method: "POST", 
      headers: {
      	'Content-Type': 'application/json'
      },
      body: paymentRequestBody
     });
  let voiceResponseJson = await voiceResponse.json();
  console.log(voiceResponseJson);
  
  res.status(200).send(voiceResponseJson);
};

