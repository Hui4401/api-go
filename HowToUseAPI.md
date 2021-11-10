### 用户注册

```postman
curl --location --request POST 'http://127.0.0.1:8080/api/v1/user/register' \
--form 'username="test"' \
--form 'password="test1234"' \
--form 'password_confirm="test1234"'
```

### 用户登陆
```postman
curl --location --request POST 'http://127.0.0.1:8080/api/v1/user/login' \
--form 'username="test"' \
--form 'password="test1234"'
```



### 用户信息

加入你的token如下

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2MzQyMTksImlhdCI6MTYzNjU0NzgxOSwiVXNlcklEIjoxfQ.dYuewho6A2EOC22t0_qyDdcwY7Ap80QgkMKkadlpnuE

```postman
curl --location --request GET 'http://127.0.0.1:8080/api/v1/user/me?user_id=test' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2MzQyMTksImlhdCI6MTYzNjU0NzgxOSwiVXNlcklEIjoxfQ.dYuewho6A2EOC22t0_qyDdcwY7Ap80QgkMKkadlpnuE' \
--form 'username="test"' \
--form 'password="test1234"'
```



### 用户退出
```postman
curl --location --request POST 'http://127.0.0.1:8080/api/v1/user/logout' \
--header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY2MzQyMTksImlhdCI6MTYzNjU0NzgxOSwiVXNlcklEIjoxfQ.dYuewho6A2EOC22t0_qyDdcwY7Ap80QgkMKkadlpnuE' \
--form 'username="test"'
```

