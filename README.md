# cache-service

### Description
The following project has 2 RPC services 
- Set         (on port:8080)
- Get         (on port:8080)
- SetUser     (on port:8000)
- GetUserByID (on port:8000)

 
 ##### The project has 4 folders  
 - proto -> the proto file and the generated code are present here
 - client -> the client function is present here for testing purpose
 - server -> files to initiate two gRPC server on different port (raw.go and user.go)
 - main.go -> main.go file to start both server and start testing function


## How to use 

### Downloading the repo in your local system 

 ```
    git clone https://github.com/ankitdmon/cache-service.git
 ```
  
  enter the folder
  
 ```
    cd cache-service
 ```
### Starting the server 
There is a main.go file present which starts the 2 RPC services  
 
 ```
    go run main.go
 ```
 

### ScreenShot

![image](https://user-images.githubusercontent.com/76701875/190658766-e89b0a65-8391-49dd-8449-7dc845752404.png)



