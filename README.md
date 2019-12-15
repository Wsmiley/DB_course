# Database course
---
## Introduce
Student grade management system design by golang

## Usage
* docker:Run the SQL Server 2017 container
* walk:Build GUI for app
* gorm:Database operations
  
## Tool
* VS code

## Database
* sql server 2017


## How to use
### docker 
```bash
docker pull mcr.microsoft.com/mssql/server:2017-latest
docker run -e "ACCEPT_EULA=Y"-e"SA_PASSWORD=<YourStrong@Passw0rd>" -p 1433:1433 --name sql1  -d mcr.microsoft.com/mssql/server:2017-latest
```
### walk
  
#### To install
```bash
go get github.com/lxn/walk
```
Then either compile the manifest using the rsrc tool, like this:
```bash
go get github.com/akavel/rsrc
rsrc -manifest app.manifest -o rsrc.syso
```
#### Build app
```bash
go build
```
To get rid of the cmd window, instead run
```
go build -ldflags="-H windowsgui"
```
#### Run app
```
DB_course.exe
```
---
### test
![demo](demo.png)

