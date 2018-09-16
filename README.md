# Growth Book API

# Run

```
./runner.sh
```

# Dependences
You can jump into this step if you want to know what dependences are used and install them yourself.

```
go get -v github.com/azbshiri/common/db
go get -v github.com/go-pg/pg
go get -v github.com/joho/godotenv
go get -v github.com/labstack/echo
```


# To-Do
- [x] Add Echo intergration
- [x] Use PostgreSQL as database
- [x] Add database migrations (using `dbmate`)
- [x] Use Repository pattern to access database layer
- [ ] Add authentication mechanism
- [ ] Dockerize application
- [ ] Add unit and integeration tests
