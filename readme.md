# cato example project

## Quick Start
```shell
 go run internal/api/main.go 
```
and visit http://localhost:8080/api/bms/v1/search?category=aaa
will get response
```json
{"Code":"0","Message":"ok","Body":[{"Id":0,"Name":"aaa","Cover":"this is cover"}]}
```
