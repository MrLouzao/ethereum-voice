

exports.getUsers = function(db){
    let users = db.get('users');
    let result = new Promise((resolve,reject)=>{
        users.find()
        .then((data)=>{resolve(data)})
        .catch((e)=>{reject(e)});

    });
    return result;
  }

exports.addUser = function(db, _user) {
    let users = db.get('users');
    return new Promise((resolve,reject)=>{
    users.insert(_user)
    .then((data)=>{resolve(data)})
    .catch((e)=>{reject(e)});
    })
}
