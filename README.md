# Task CLI

A simple command-line task manager built with Go.

This project was created as part of my Go learning journey, focusing on Go fundamentals, project structure, structs, methods, pointers, error handling, and separation of concerns.

## Features

* Add tasks
* List tasks
* Complete tasks
* Delete tasks
* Basic error handling

## Structure

```text
Task-CLI/
├── main.go
├── go.mod
└── internal/
    └── task/
        └── task.go
```

## Commands

```bash
go run . add "Learn Go"
go run . list
go run . done 1
go run . delete 1
```

## Concepts

* Go modules & packages
* Structs & methods
* Pointers
* Slices
* CLI arguments
* Error handling
* Separation of concerns

> Tasks are currently stored in memory and are lost when the program exits.
