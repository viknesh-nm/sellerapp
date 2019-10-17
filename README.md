Steps to run:

```go
glide install
```

```go
docker-compose up --build
```
`
#Note:
 Please watch the log where the mysql container is up, Since it takes a little time to install the service 
 `

Once the mysql is up, restore the dump data to the container

```
docker-compose exec -T mysql mysql -uuser -p"password" -D user_details < $GOPATH/src/github.com/viknesh-nm/sellerapp/resources/profile_access_auth.sql
```

```go
Port will be running on `:9090`
```