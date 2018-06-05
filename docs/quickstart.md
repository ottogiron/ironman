# Quickstart Guide

This guide covers how to get started using Ironman.



## Install Ironman

Download a binary release of Ironman from the [releases](https://github.com/ironman-project/ironman/releases) page for your OS and put it under a PATH directory.



## Usage

You can get help by running ```ironman help``` and ```ironman help <command>```



## Install an example template

To install an existing template, use the ```ironman install <template-id>``` command.

We will install a template to create the base project structure for a  [Go HTTP](https://golang.org/pkg/net/http/) app.



### Install the Go HTTP Server example template

```bash
$ ironman install https://github.com/ironman-project/simple-gohttp-template.git
```

Now you can list the available templates.

```bash
$ ironman list
Installed templates
+---------------+-------------------------+--------------------------------+
|      ID       |          NAME           |          DESCRIPTION           |
+---------------+-------------------------+--------------------------------+
| simple-gohttp | Simple Go HTTP Template | A simple HTTP Go library       |
| | | template                       |
+---------------+-------------------------+--------------------------------+
```



### Generate a new project based on the template

A template might have many generators. The default "app" generator will run if you don't specify one. The app generator of this template generates the base project.

```bash
$ ironman generate simple-gohttp /path/to/app --set projectName="Some project name",projectDescription="Some project description"
```

The generated project will be under the /path/to/app. Open the directory with your favorite editor and check the results.

You can check the template definition here https://github.com/ironman-project/simple-gohttp-template.



#### Inline values and values files

You could have also passed the necessary values using a file.

```bash
$ ironman generate simple-gohttp /path/to/app -f /path/to/values.yaml
```

where ```values.yaml``` would look something like:

```yaml
projectName: Some Project Name
projectDescription: Some project Description
```

or multiple files:

```bash
$ ironman generate simple-gohttp /path/to/app -f /path/to/values.yaml,/path/to/values2.yaml
```

the rightmost file takes precedence. 

You can also mix it with inline values, which have the higher precedence: 

```bash
$ ironman generate simple-gohttp /path/to/app -f /path/to/values.yaml --set projectName="Higher Precedence Project Name"
```



#### Run the App

Go is necessary for this step (https://golang.org/), given the generated application is a Go app. 

```bash
cd /path/to/app
go build 
$ ./app #will run the http server
```

go to http://localhost:8080, and you should see the message "Hello World"



### Generate a new endpoint file using the template "endpoint" generator

Ironman also supports file generators. This template implements an "endpoint" generator

```bash
 cd /path/to/app
 ironman generate simple-gohttp:endpoint myendpoint.go --set endpoint="/myendpoint"
```

Stop and rebuild your app:

```bash
cd /path/to/app
go build 
$ ./app #will run the http server
```

You should be able to reach your new endpoint http://localhost:8080/myendpoint
