## SimpleApp

### Quickstart

Clone repo

    git clone https://github.com/jstalex/simpleapp

Change directory

    cd simpleapp

Run app in docker

    make run

Stop app

    make stop


### Requests

* Add user

```
curl -X POST -H "Content-Type: application/json" -d '{
    "name": "Alexander",
    "surname": "Parfenov",
    "patronymic": "Sergeevich",
    "email": "jstalex@gmail.com"
}' http://localhost:8080/user

```

* Get user

```
curl -X GET http://localhost:8080/user?email=jstalex@gmail.com
```