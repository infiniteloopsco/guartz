# Guartz
It's a simple job scheduler with HTTP interface and MySQL as Database.

## Setting

You can install this server with this command:
```
go get github.com/infiniteloopsco/guartz
```
Browse to the guartz directory and install the dependencies:

```
godep restore
```
Build the binary:
```
go build
```
This will create a binary called guartz, create the a .env file on the root of the project with this content:
```
MYSQL_DB=root:password@tcp(127.0.0.1:3306)/guartz?charset=utf8&parseTime=True
GIN_MODE=prod
PORT=8081
```
And set the enviroment variable:
```
export GUARTZ_MODE=dev
```
Now you can run the server:
```
./guartz
```
This will run the server where you will be able to schedule the command's execution in the operative system.

## API REST

You can manage the tasks on the server by

## Creating a task

You can create a task on the server with this rest call
```
POST http://localhost:3000/tasks
{
    "id": "task-id",
    "periodicity": "@every 1h",
    "command": "docker run demoApp"
}
```
With this request I can create a task that will run a docker container every hour. You can use any cron unix valid format or use these shortcuts:
```
Entry                  | Description                                | Equivalent To
-----                  | -----------                                | -------------
@yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
@monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
@weekly                | Run once a week, midnight on Sunday        | 0 0 0 * * 0
@daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
@hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
```
The server uses https://godoc.org/gopkg.in/robfig/cron.v2 library, there you can get more information.

## Stopping a task

You can stop a task by deleting it or updating it:
```
DELETE http://localhost:3000/tasks/task-id
```
```
POST http://localhost:3000/tasks
{
    "id": "task-id",
    "periodicity": "stop",
    "command": "docker run demoApp"
}
```

## Listing tasks

You can list all system tasks:
```
GET http://localhost:3000/tasks
```

## Get a task

You can show the info of a task using:
```
GET http://localhost:3000/tasks/task-id
```

Feel free to contribute and report bugs.
