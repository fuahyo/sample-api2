# sample-api2 [Readme First]

simple RestAPI application using Go-language from database sql

initialisasi (Additional if you don't have go mod/go sum in ur file):

1. go mod init sample-api2
2. go get -u github.com/gin-gonic/gin
3. go get -u github.com/jmoiron/sqlx
4. go get -u github.com/go-sql-driver/mysql

database: golang
username: root
password:

table: person
column: id (int), first_name (varchar), last_name (varchar)

project name: sample-api2

How to run: go run main.go
url: localhost:1234/persons

how to push on github for new repo:
echo "# sample-api2" >> README.md (then update ur readme first)

1. git init
2. git add README.md
3. git commit -m "first commit"
4. git branch -M main
5. git remote add origin https://github.com/fuahyo/sample-api2.git
6. git push -u origin main

how to push on github for existing repo:

1. git init
2. git add .
3. git commit -m "lalalala"
4. git remote add origin https://github.com/fuahyo/sample-api2.git
5. git branch -M main
6. git push -u origin main or git push -u --force origin main
