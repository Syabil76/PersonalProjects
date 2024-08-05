# Databaseless CRUD API in GO
  A program in go that allows you to Create, Read, Update and Delete your own json files in GO!
## Overview
The program uses Postman- an API platform that can be hosted locally on port 8000. This program utilises postman by
having an interactable interface to update JSON files\
\
The program is relatively limited as there is no database as stated by name. Therefore new fields of information cannot be stored.
For example in the program the data is held by a *struct* but cannot be tampered to add something like an ID field.\
\
The program allows basica CRUD operations to mess around with your 'database' by using JSON. 
Therefore, if desired, the data can created or removed or manipulated.\
\
It's a very basic example using cars, but can realistically be done with things like movies or classrooms or whatever, really.

## Logic
  The Program follows a 4 step program:
  - Routes
  - Functions
  - Endpoints
  - Methods
  \
  The routes and functions show the workings of the program- very straightforward.\
  The endpoints show where the data is actually being updated and specified by ID, sort of like a primary key.\
  The methods show what operation is used by postman/API to actually update data.
  ![image](https://github.com/Syabil76/PersonalProjects/blob/main/Golang/CRUD_API_GO/images/Logic.png?raw=true)
  ![postman](https://voyager.postman.com/illustration/rest-client/rest-client-send-and-save-postman.png)
## Usage
Usage is relatively simple, run the program in your CLI (1) and make sure your port is set to 8000. After that go to postman and
remember to use their agent otherwise it wont be able to connect to your network, or just download the app (2).\
\
Create your files where you want to put all your functions (3) and set the connection to http://localhost:8000/whateveryouput (in my case, cars). 
make sure to set the correct method eg. GET for READ and POST for CREATE (4).\
\
Have fun! (5)\

1. ![cli](https://github.com/Syabil76/PersonalProjects/blob/main/Golang/CRUD_API_GO/images/CLI.png)
2. ![postman2](https://github.com/Syabil76/PersonalProjects/blob/main/Golang/CRUD_API_GO/images/postman.png)
3. ![getall](https://github.com/Syabil76/PersonalProjects/blob/main/Golang/CRUD_API_GO/images/getall.png)
4. ![getbyid](https://github.com/Syabil76/PersonalProjects/blob/main/Golang/CRUD_API_GO/images/getbyid.png)
5. ![good](https://www.pngall.com/wp-content/uploads/14/Thumbs-Up-Emoji-PNG-Image.png)
## Errors and known issues
  No comments- code hard to read if unfamiliar.\
  No go.mod and go.sum file.\
  Nil.
## Libraries used
  - "encoding/json" --> to update (generally) JSON files.
	- "fmt" --> formatting strings for printing
	- "log" --> for logging errors or successful operations
	- "math/rand" --> random math operation
	- "net/http" --> HTTP operations
	- "strconv" --> String conversion
	- "github.com/gorilla/mux" --> For API operations

## Documentation
  Gorrila MUX- https://pkg.go.dev/github.com/gorilla/mux
