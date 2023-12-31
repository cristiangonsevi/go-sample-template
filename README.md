# How to Structure Your Golang Projects
This repository serves as an example of a possible structure for Go (Golang) projects. It's important to note that there isn't a single correct way to structure a Go project, and each developer or team has the freedom to adapt the structure according to their needs and preferences.

## Purpose
The purpose of this repository is to provide a basic guide on how to organize your Go project, separating responsibilities into different packages and layers. The proposed structure includes directories for cmd, internal, and files such as go.mod and go.sum.

## Project Structure

```
my-api-project/
|-- cmd/
|   |-- api/
|       |-- main.go
|
|-- internal/
|   |-- api/
|       |-- handler/
|           |-- handler.go
|       |-- middleware/
 model/
|           |-- user.go
|       |-- repository/
|           |-- user_repository.go
|       |-- router/
|           |-- router.go
|       |-- service/
|           |-- user_service.go
|-- go.mod
|-- go.sum
```
## How to Use
To use this structure as a starting point for your project, follow these steps:

- Clone this repository: git clone https://github.com/cristiangonsevi/go-sample-template.git.
- Modify packages and files according to your needs.
- Customize the go.mod file with your project information.
- Start developing your application!
Remember that this structure is only a suggestion, and you can adjust it based on your preferences and specific requirements.
