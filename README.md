# FriendsLessonsSystem
Service that allows a user to know for what courses their friends have taken any lesson
## BackEnd (GO with GIN Framework)
### Requisites
- Install Go 1.18

### How to Run (BackEnd/web-service-gin)
- Run `go get .`
- Run `go run .`
- Make a GET REQUEST in http://localhost:8080/api/users in POSTMAN or curl

### Endpoints
- Get All Users `/api/users`
- Get User by Id `/api/users/{id}`
- Get All Friendships `/api/friendships`
- Get All Friends by User Id `/api/users/{id}/friends`
- Get All Lessons by User Id `/api/users/{id}/lessons`
## FrontEnd (Angular)
### Requisites
- NodeJS
- Angular 13

### How to Run (FrontEnd/friends)
- Run `npm install`
- Run `ng serve --o`
- Start in http://localhost:4200
- User: admin
- Password: admin123

### How to configure Endpoint API
- Change `target` in `proxy.conf.json`