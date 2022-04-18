# BACKEND FriendsLessonsSystem

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

## How to Deploy in Heroku
- https://devcenter.heroku.com/articles/getting-started-with-go

### Using Windows
- Install Heroku CLI
- Run `go env -w GOOS=linux`
- Run `go build -o bin/friendslessonssystemapi -v`
- Run `go env -w GOOS=windows` (necessary for run `go run .` again in Windows)

### Configure Buildpacks
- If you have the folder of project in a subfolder like me, you need put first https://github.com/timanovsky/subdir-heroku-buildpack 
- After put heroku/go

### Configure Vars
-  If you have the folder of project in a subfolder like me, Configure PROJECT_PATH=BackEnd/web-service-gin
- PORT=80

### Deploying
- Run `go build -o bin/friendslessonssystemapi -v`
- Run `git push heroku main`
- Run `heroku logs --tail`

## DEMO
- http://friendslessonssystemapi.herokuapp.com/api/users

