# random-roles

## Overview

<img alt="notification_demo" width="300px" src="https://user-images.githubusercontent.com/38056766/69641243-3437b080-10a3-11ea-9cde-6914089bcd95.png">

Notification of the clearning roles for tech lovers in Fukuoka :heart:

```
random-roles
.
├── go.mod
├── go.sum
├── main.go
└── payload
    ├── attachment.go
    └── text.go
```

## Usage

### Installation

```
❯ go get -u github.com/ashwanthkumar/slack-go-webhook
```

### Setting 

After `git clone` , create a file `.envrc` and describe the following.
You can find the Webhook's URL from `https://slack.com/apps > Manage > Custom Integrations > Incoming Webhooks`  

```
export URL=https://hooks.slack.com/services/your/webhook's-url
```

### Build

If you build with OS and Architecture environment variables according to your env you will run, you can generated a binary file.
To specify a file name, execute the commnad with `-o` option.
For example, if you want to generate with the name `hoge` : `go build -o hoge main.go` .

```
# Linux ( A binary file named "main" is output )
❯ GOOS=linux GOARCH=amd64 go build main.go

# macOS ( A binary file named "main" is output )
❯ GOOS=darwin GOARCH=amd64 go build main.go

# windowsOS ( A binary file named "main.exe" is output )
❯ GOOS=windows GOARCH=amd64 go build main.go
```

### Run

In addition to executing the binary file generated earlier like this `./main` , you can also execute the following command.

```
❯ go run main.go 
```
