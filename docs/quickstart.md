# Quickstart Guide

This guide covers how to get started using Ironman.

## Install Ironman

Download a binary release of Ironman from the [releases](https://github.com/ironman-project/ironman/releases) page for your OS and put it under a PATH directory.

## Install an example template

In order to install an existing template, use the ```ironman install <template-id>``` command.

We will install an template for creating the base project structure for a go [Cobra](https://github.com/spf13/cobra) CLI app. [Cobra](https://github.com/spf13/cobra) has already a CLI for doing this, but it is a good example of a general porpuse use of ironman.

### Install the Cobra CLI library example template

```bash
$ ironman install https://github.com/ironman-project/cobra-cli-template.git
```

### Generate a new project based on the template

```bash
$ ironman generate cobra-cli-template /path/to/generate/mycli --set cliName=mycli,projectName="My CLI",projectDescription="This is an example generated cobra cli project"
```

