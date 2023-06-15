# go-ddd-sample

This is a DDD sample using golang.
This is a golang version of [kotlin-dddd-sample](https://github.com/htnk128/kotlin-ddd-sample).

- [echo](https://github.com/labstack/echo)
- [sqlboiler](https://github.com/volatiletech/sqlboiler)
- [Staticcheck](https://staticcheck.io/)

## Run Application

```bash
$ make up-all
$ make migrate-up app=account
$ make migrate-up app=address
```
