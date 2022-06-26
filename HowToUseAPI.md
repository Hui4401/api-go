### 用户注册

```postman
curl --location --request POST 'http://127.0.0.1:8080/user/register' \
--json 'username="test"' \
--json 'password="test1234"' \
--json 'password_confirm="test1234"'
```

### 用户登陆
```postman
curl --location --request POST 'http://127.0.0.1:8080/user/login' \
--json 'username="test"' \
--json 'password="test1234"'
```


### 用户信息

假如你的token如下

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2MzQyMTksImlhdCI6MTYzNjU0NzgxOSwiVXNlcklEIjoxfQ.dYuewho6A2EOC22t0_qyDdcwY7Ap80QgkMKkadlpnuE

```postman
curl --location --request GET 'http://127.0.0.1:8080/user/me?user_id=test' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2MzQyMTksImlhdCI6MTYzNjU0NzgxOSwiVXNlcklEIjoxfQ.dYuewho6A2EOC22t0_qyDdcwY7Ap80QgkMKkadlpnuE' \
```


### 用户退出
```postman
curl --location --request POST 'http://127.0.0.1:8080/user/logout' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2MzQyMTksImlhdCI6MTYzNjU0NzgxOSwiVXNlcklEIjoxfQ.dYuewho6A2EOC22t0_qyDdcwY7Ap80QgkMKkadlpnuE' \
```

