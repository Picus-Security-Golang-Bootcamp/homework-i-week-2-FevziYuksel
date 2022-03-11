## Homework | Week 2

We have a list of books. This application has 2 tasks.
1. To show all the books in the application as output.
2. If the name of the book given as input is available, printing it on the screen; if not, printing a printout stating that the book does not exist.

### list command
```
go run main.go list
```
With the "list" command, we suppose to access whole list of the books in the application as output.

### search command 
```
go run main.go search <bookName>

With the "search" command , we suppose to access the name of the book that we give the application if the list contains it. 
If not, it suppose to show us the message "<bookName> is not found"
