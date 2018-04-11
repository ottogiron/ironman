<p align="center">
  <img style="float: right;" src="ironman.png" alt="Ironman logo"/>
</p>

# Ironman 
[![Build Status](https://travis-ci.org/ironman-project/ironman.svg?branch=master)](https://travis-ci.org/ironman-project/ironman)
[![Build status](https://ci.appveyor.com/api/projects/status/yi1e02dy65nv96uy/branch/master?svg=true)](https://ci.appveyor.com/project/ottogiron/ironman/branch/master)
[![GoDoc](https://godoc.org/github.com/ironman-project/ironman?status.svg)](https://godoc.org/github.com/ironman-project/ironman)
[![Go Report Card](https://goreportcard.com/badge/github.com/ironman-project/ironman)](https://goreportcard.com/report/github.com/ironman-project/ironman)

**Ironman** is a CLI tool that provides an easy way to create and share project templates hosted in git repositories for any regular document based project. Install and generate a new document based project using  **Ironman Templates** in seconds.

## Features 

  * Develop  Ironman based templates.
  * Manage local Ironman templates from remote sources (install, uninstall, upgrade) (git repositories)
  * Generate new projects based on ironman you or someone else created.

## Motivation

Let’s say you are about to start a new cool project, you need a common project directories structure, a few dozens of files,  every time you have to change README titles/subtitles, specific app identifiers, documentation URL’s and so on to get started with a new app.  What if you work with one or many teams and you want to provide an easy way to share common project templates and enforce common development good practices and standards.  How would you solve this repetitive task ? 

A way could be keeping your project "templates" in a git repo so everyone can clone their own copy and make the necessary changes to adapt those changes to their needs. From my own experience that could lead to people not knowing what changes they need to do in order to get a custom “correct” project,  also rapidly outdated templates no one knows they even exists.

## Here comes the hero 

**Ironman** provides an easy to use declarative framework for you to define **Ironman Templates**. **Ironman templates** can then be hosted in a git repository.

**Ironman Templates** can be  administered (installed, upgraded, removed) using the **Ironman CLI**. Once a template is installed you can easily generate new projects based on that template as many times as you want in a repetitive and reliable way, then share them with anyone with an standard git URL. 

## Install

Binary download is available for the following OS's

* Linux
* OSX
* Windows

Download the your specific  binary from tar file from the [releases](https://github.com/ironman-project/ironman/releases) page

Unpack the binary and add it to your PATH.

### Verify 

```bash
$ ironman version
Ironman vv0.1.1-5d02c19 Build date: 20180411.034900
``` 

Run ```ironman help``` or ```ironman help <command>``` to get help about additional commands.


## Docs

Get started with the [Quick Start guide](docs/quickstart.md) you can find all the additional documentation [here](docs)
