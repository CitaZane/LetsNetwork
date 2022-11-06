# LetsNetwork
## Project description
Facebook-like social network that will contain the following features:
- Followers
- Profile
- Posts
- Groups
- Events
- Notification
- Chats

## Run the project
To test the project you need to have [NodeJS](https://nodejs.org/en/) installed.
1. Start frontend server by going to **/frontend** directory and running those commands
-  `npm run serve`
-  `npm install`
2. Start backend server by going to **/backend** directory and run `go run server.go`


## Stack
Frontend
- Vue
- HTML & CSS

Backend
- Go
- SQLite3

## Role
Project has been created in a team of 4 people.
My role in project -> back-end developer


### What I learned

- Structuring big scale go projects
- Seperating database logic completely from handlers using interfaces
- Websocket basics
- Real time messaging and notifications
- Creating a larger database structure. Also many-to-many relationship in databases
- Using channels/ gorutines with Go

### To be improved (in back-end part)

- Error handling 
- Better designing the api (routes not structured well)
- Implement testing
