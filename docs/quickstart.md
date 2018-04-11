# Quickstart Guide

This guide covers how to get started using Ironman.

## Install Ironman

Download a binary release of Ironman from the [releases](https://github.com/ironman-project/ironman/releases) page for your OS and put it under a PATH directory.

## Install an example template

In order to install an existing template, use the ```ironman install <template-id>``` command.

We will install an template for creating the base project structure for a  [Go HTTP](https://golang.org/pkg/net/http/) app.


### Install the Go HTTP Server example template

```bash
$ ironman install https://github.com/ironman-project/simple-gohttp-template.git
```

### Generate a new project based on the template

```bash
$ ironman generate simple-gohttp /path/to/app --set projectName="Some project name",projectDescription="Some project description"
```

Your new project should be generated under the /path/to/generate/myhttp-app. Open the directory with your favorite editor and check the results.

You can check the template definition here https://github.com/ironman-project/simple-gohttp-template.

### Run the Go Application

This step (and only this step) requires go to be installed (https://golang.org/), since the generated application is written in go. You could've created your own templates for different project types, languages, since Ironman is not limited to any specific language. 

```bash
$ go run /path/to/generate/myhttp-app/main.go #will run the http server
```

go to http://localhost:8080, and you should see the the message "Hello World"

