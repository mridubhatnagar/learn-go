# Golang 101

## 2024-06-24
1. Everything is organized as Packages. 
2. All files must have only 1 entrypoint. main() is 
the entrypoint and gets executed on its own. 
3. To initialize a go project. Run the following command
```
go mod init <folder name>
```
This creates a go.mod file. Which defines go version and the module name. 
4. Main function doesn't accept any arguments. Neither
returns anything. 
5. Go package is a collection of source files. It gives
various functionalities. Containers of various functionalities that go gives you for ready to use. 


## 2024-06-25


### VARIABLES

1. If variable is constant then use `const` keyword before variable name. Else use `var`. 

```
const conferenceTickets = 50
```
As number of available tickets is fixed i.e 50.

2. To print formatted data. There is a function `printf` present in `fmt` package. `%v` is the default.
PS: For more details check printf documentation in `fmt`
standard library. 

### DATA TYPES

1. If we are assigning the value to a variable and declaring the variable at the same time. Then, Go implicitly assigns the data type to variable based on the value. 

2. If variable is only declared and no value is assigned. Then we need to explicitly mention the
data type of variable. 

```
var userName string
```

3. `%T` is used to print type of the variable

4. Syntatic Sugar in GO. Works only for variables
and not constants. 

Instead of writing

```
var conferenceName string = "Go Conference"
```

We can write it as

```
conferenceName := "Go Conference"
```

### Asking User for Input

1. Pointer is variable that points to memory address of
another variable. 
fmt.Scan(&userName)

### Loops and Control Statement


### Array and Slices

1. Array can contain only elements with same data type.
2. Array size and type has to be declared in the beginning.

Below is the syntax to initialize empty array
```
var bookings[50]string{}
```

Syntax to initialize array with list of elements.
```
var bookings = [50]string{"Mridu", "Sakshi", "Shubh"}
```

3. Dynamic arrays are slices. Slices are used when 
size is not fixed. 

```
var bookings[]string
```
4. To add an element to a slice. Use `append`.
First argument to append is slice.
Syntax.
```
bookings = append(bookings, firstName + " " + lastName) 
```
### Infinite Loop

No exit condition
```
for {

}
```

### Break and Continue

1. `break` breaks out of the loop.
2. `continue` moves to next iteration without executing remaining part of the body.