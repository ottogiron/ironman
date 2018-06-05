<p align="center">
 <img style="float: right;" src="ironman.png" alt="Ironman logo"/>
</p>

# Ironman 
[![Build Status](https://travis-ci.org/ironman-project/ironman.svg?branch=master)](https://travis-ci.org/ironman-project/ironman)
[![Build status](https://ci.appveyor.com/api/projects/status/yi1e02dy65nv96uy/branch/master?svg=true)](https://ci.appveyor.com/project/ottogiron/ironman/branch/master)
[![GoDoc](https://godoc.org/github.com/ironman-project/ironman?status.svg)](https://godoc.org/github.com/ironman-project/ironman)
[![Go Report Card](https://goreportcard.com/badge/github.com/ironman-project/ironman)](https://goreportcard.com/report/github.com/ironman-project/ironman)

**Ironman**  is CLI tool that provides a way to define and share project templates hosted in git repositories. Install and generate a new document based project using  **Ironman Templates** in seconds.

## Features 

 * Develop  Ironman based templates.
 * Manage local Ironman templates from remote sources (install, uninstall, upgrade) (git repositories)
 * Generate new projects based on ironman you or someone else created.

## Motivation

You are about to start a new project, with a common project structure, and few dozens of files.  Every time you have to change titles/subtitles, app identifiers, docs URL’s to get started. If you work with one or more teams, how do you enforce common standards?.  This is a repetitive task, how would you solve it?
 
An option is a git repo, everyone can clone their own copy and make the necessary to adapt it to their needs. From my experience that can lead to people not knowing the changes, they need to do to get a custom “correct” project. Also, outdated templates no one knows they even exist.
## Here comes the hero 

**Ironman** provides a declarative framework for you to define **Ironman Templates**. You can host **Ironman templates** in a git repository.

**Ironman Templates** can be managed (installed, upgraded, removed) using the **Ironman CLI**. With an installed template you can generate new projects many times as you want in a repetitive and reliable way. Share a template using a standard git URL.

## Install

Binary download is available for the following OS's

* Linux
* OSX
* Windows

Download your specific  binary from tar file from the [releases](https://github.com/ironman-project/ironman/releases) page

Unpack the binary and add it to your PATH.

### Verify 

```bash
$ ironman version
Ironman vv0.1.1-5d02c19 Build date: 20180411.034900
``` 

Run ```ironman help``` or ```ironman help <command>``` to get help about more commands.



## Docs

Get started with the [Quick Start guide](docs/quickstart.md) you can find all the additional documentation [here](docs)
