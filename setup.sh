# start mongo server
mongod --config "C:\develop\mongo\bin\mongod.cfg"

# auth
## mongo
use admin  
db.createUser({
  user: 'admin',  // 用户名
  pwd: 'admin',  // 密码
  roles:[{
    role: 'root',  // 角色
    db: 'admin'  // 数据库
  }]
})

## config
security:
  authorization: enabled

# run mongo client
mongo
use admin
db.auth('admin', 'admin')
mongo admin -u admin -p admin

# add user
use tit  // 跳转到需要添加用户的数据库
db.createUser({
  user: 'tit',  // 用户名
  pwd: 'tit',  // 密码
  roles:[{
    role: 'readWrite',  // 角色
    db: 'tit'  // 数据库名
  }]
})

# user command
show users  // 查看当前库下的用户

db.dropUser('testadmin')  // 删除用户

db.updateUser('admin', {pwd: '654321'})  // 修改用户密码

db.auth('admin', '654321')  // 密码认证

show dbs

show collections

db.dropDatabase()

db.foo.drop()

db.foo.remove({})

db.stats()

db.foo.status()

db.foo.insert({a:1, b:2})

db.foo.findOne()
db.foo.find().pretty()
db.foo.find().count()


