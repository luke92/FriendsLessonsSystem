# FriendsLessonsSystem
Service that allows a user to know for what courses their friends have taken any lesson
## BackEnd (GO with GIN Framework)
### Requisites
- Install Go 1.17

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

### DEMO
https://friendslessonssystemapi.herokuapp.com/api/users

## FrontEnd (Angular)
### Requisites
- NodeJS
- Angular 13

### How to Run (FrontEnd/friends)
- Run `npm install`
- Run `ng serve --o` for development
- Run `ng serve` for testing or production
- Start in http://localhost:4200
- User: `admin`
- Password: `admin123`

### How to configure Endpoint API for development
- Change `target` in `proxy.conf.json` if you use `ng serve --o` (Comunicate localhost with localhost)
- Change `config.json` if you use `ng serve`

### DEMO
https://friendslessonssystem.netlify.app/
