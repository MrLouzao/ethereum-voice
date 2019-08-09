const keythereum = require('keythereum');

exports.getUsers = function (db) {
    let users = db.get('users');
    let result = new Promise((resolve, reject) => {
        users.find()
            .then((data) => { resolve(data) })
            .catch((e) => { reject(e) });

    });
    return result;
}

exports.getUser = function (db, _name) {
    let users = db.get('users');
    let result = new Promise((resolve, reject) => {
        users.find({ name: _name })
            .then((data) => { resolve(data) })
            .catch((e) => { reject(e) });

    });
    return result;
}

exports.addUser = function (db, _user) {
    let users = db.get('users');
    return new Promise((resolve, reject) => {
        let password = _user.password || "password";
        if (_user.name == "" || _user.name == undefined) { throw new Error("User not specified") }
        let dk = keythereum.create();
        // synchronous
        var keyObject = keythereum.dump(password, dk.privateKey, dk.salt, dk.iv);
        users.find({ name: _user.name }).then((data) => {
            console.log(data)
            if (data != []) { throw new Error('USer allready exists') }
            else { return 0}
        }).then(() => {
            users.insert({
                name: _user.name,
                keystore: keyObject
            })
                .then((data) => { resolve(data) })
                .catch((e) => { reject(e) })
        }).catch((e) => { reject(e) });
    })
}
