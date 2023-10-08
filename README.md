# Cognixus Assessment - Backend API
This project is to demostrate completion of assessment using tools below:
* Clean code practice
* Unit tests
* Docker
* Github authentication using access token

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Improvements](#improvements)

## General info
This project is structured based on clean architecture which also known as layered architecture that consists of:
1. Domain/Entities/Models
2. Use cases
3. Controllers/Adpaters 
4. Frameworks & Drivers 

The Github authentication is uses access token instead of JWT token as a simplicity to demostrate authentication process. 

There are 4 endpoints in this project:
* GET /api/v1/todo
```
Example response
{
    "message": "Get todo list successfully",
    "data": [
        {
            "ID": "d3a79c7f-6cd8-4b9b-8751-b77650d1ecae",
            "Name": "Test Todo",
            "Description": "Lorem Ipsum",
            "Status": true,
            "UserID": "77889911"
        }
    ]
}
```

* POST /api/v1/todo
```
Example request
{
    "name": "Test Todo",
    "description": "Lorem Ipsum"
}

Example response
{
    "message": "Task created successfully",
    "data": null
}
```

* PUT /api/v1/todo
Example
```
Example request
{
    "id": "d3a79c7f-6cd8-4b9b-8751-b77650d1ecae",
}

Example response
{
    "message": "Updated todo successfully",
    "data": null
}
```

4. DELETE /api/v1/todo
Example
```
Example request
{
    "id": "d3a79c7f-6cd8-4b9b-8751-b77650d1ecae",
}

Example response
{
    "message": "Delete todo successfully",
    "data": null
}
```

## Technologies
1. Docker
2. Github
3. Golang 1.21.2
4. PostgresSQL
5. Makefile

## Setup
To run the project locally, you need create a .env file and copy the content of .env.example to .env. Do forget to fill in the fields with relevant values.

There is a Makefile command to run the application.
* make run -> to run the application
* make stop -> to stop the application
* make clean -> tear down and clean the application

## Improvements
Few ideas for improvements as below:
1. CI/CD pipeline for deployment
2. Hosting
3. Oauth2.0 flow setup
4. Logging and monitoring

