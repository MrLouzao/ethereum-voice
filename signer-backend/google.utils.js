const google = require('googleapis')
const data = require('./application.data')

const googleConfig = {
  clientId: data.client_email, // e.g. asdfghjkljhgfdsghjk.apps.googleusercontent.com
  clientSecret: data.private_key_id, // e.g. _ASDFA%DFASDFASDFASD#FAD-
  redirect: data.auth_uri // this must match your google api settings
};
exports.getGoogleConfig = function() {return googleConfig}

/**
 * Create the google auth object which gives us access to talk to google's apis.
 */
exports.createConnection = function() {
  return new google.auth.OAuth2(
    googleConfig.clientId,
    googleConfig.clientSecret,
    googleConfig.redirect
  );
}

/**
 * This scope tells google what information we want to request.
 */
const defaultScope = [
  'https://www.googleapis.com/auth/plus.me',
  'https://www.googleapis.com/auth/userinfo.email',
];

exports.getDEfaultScope = function() {return defaultScope}

/**
 * Get a url which will open the google sign-in page and request access to the scope provided (such as calendar events).
 */
exports.getConnectionUrl = function (auth) {
  return auth.generateAuthUrl({
    access_type: 'offline',
    prompt: 'consent', // access type and approval prompt will force a new refresh token to be made each time signs in
    scope: defaultScope
  });
}

/**
 * Create the google url to be sent to the client.
 */
exports.urlGoogle = function () {
  const auth = createConnection(); // this is from previous step
  const url = getConnectionUrl(auth);
  return url;
}