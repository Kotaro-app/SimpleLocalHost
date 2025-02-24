# SimpleLocalHost
This app is written in golang, is lightweight, and easy to set up on local host.

# Description
This program is a simple local HTTP server written in Go. It starts on a specified port number and serves the specified file (HTML, text, CSV, JSON, XML). It is useful for checking web pages under development and for easily sharing local files.

# Features
Start by specifying the port number and file name
Can serve HTML, text, CSV, JSON, XML files
Cross-platform support (Windows, macOS, Linux)
Simple command-line interface

# How to Use
Download the program and place the executable file (server.exe, server) in a directory in your PATH.
Run the server.exe or server command on the command line.
Enter the port number and file name.
Access http://localhost:<port number>/ in your browser.

# How to Build
You can build the program with the following command.

Bash

go build main.go
To cross-compile, use the following command.

# Windows
Bash<br>
GOOS=windows<br> GOARCH=amd64<br> go build -o server.exe

# macOS
Bash
<br>
GOOS=darwin<br> GOARCH=amd64<br> go build -o main_amd64<br>
GOOS=darwin<br> GOARCH=arm64<br> go build -o main_arm64　　
# Universal Binary
lipo -create -output main main_amd64 main_arm64

# Linux
Bash
GOOS=linux GOARCH=amd64 go build -o main

# Contribution
Please report bug reports and feature suggestions in GitHub Issues or Pull Requests.

# License
This program is released under the MIT license.

Short English Description:
A simple local HTTP server written in Go. Serves specified files (HTML, text, CSV, JSON, XML) on a given port. Cross-platform.
