env设置：

若SQL_DSN为空，则启用自带的的sqlite


## 注册

### request

POST http://localhost:8080/api/auth/users

>{
"username":"raphael",
"password":"1234acc"
}

>curl --location 'http://localhost:8080/api/auth/users' \
--header 'Content-Type: application/json' \
--data '{
"username":"raphael",
"password":"1234acc"
}'    

### response
>{
"message": "注册成功",
"success": true
}
## 登录
POST http://localhost:8080/api/auth/login

### request
>{
"username":"raphael",
"password":"1234acc"
}

>curl --location 'http://localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data '{
"username":"raphael",
"password":"1234acc"
}'

### response
>{
"data": {
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.WkuFfrFL9I-PqgF3Uigw_k2uunM0ap63PiMsK3S3Wnw",
"userId": 2
},
"message": "",
"success": true
}




## 翻译task

tips，翻译接口有鉴权，需要在header中添加token

token在login接口返回

POST http://localhost:8080/api/tasks

### request
>{
"username":"raphael",
"password":"1234acc"
}

>curl --location 'http://localhost:8080/api/tasks' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjJ9.WkuFfrFL9I-PqgF3Uigw_k2uunM0ap63PiMsK3S3Wnw' \
--form 'file=@"/input.json"' \
--form 'source_lang="en"' \
--form 'target_lang="zh"'

### response
>{
"name": "约翰·多伊",
"email": "johndoe@example.com",
"phone": "+1234567890",
"address": "123 Main St, Springfield, IL 翻译为：伊利诺伊州斯普林菲尔德市主街123号",
"website": "翻译全部：https://johndoe.com",
"description": "这是一段示例描述。",
"tags": [
"developer",
"designer",
"freelancer"
]
}