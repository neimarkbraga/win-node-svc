## Windows Node Service
win-node-svc turns node projects into a windows service

### Prerequisites
- Windows OS
- Go Lang
- Node JS

### Getting Started
Clone project: `git clone https://github.com/neimarkbraga/win-node-svc.git`    
Install dependencies with `go get`    
Modify config on `app/config.go`
- Name -  name of service. Space is not allowed here
- DisplayName - display name of service. Spaces are allowed here
- Description - description of service
- WorkingDirectory - root directory of node app/project
- EntryFile - entry file path (relative to WorkingDirectory)
- LogDirectory - directory where console outputs will be logged

Build exe file, run `go build .`. `win-node-svc.exe` will be created in the root directory, this will be used to install windows service... filename can be changed to anything. 

### Usage
Open a command prompt as administrator and navigate to the folder where the built exe file is.    
The following are the commands that can be used:
- `win-node-svc.exe debug` - this will run the app in the terminal. you may want to run this before installing th e service to make sure that your app is running as expected.
- `win-node-svc.exe install` - this will install the exe file as a windows service; after running, your service should appear in windows service list.
- `win-node-svc.exe remove` - uninstalls/removes service.
- `win-node-svc.exe start` - starts the service. alternatively you can do this through services window.
- `win-node-svc.exe stop` - stops the service. alternatively you can do this through services window.
