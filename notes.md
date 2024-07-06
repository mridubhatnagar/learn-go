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
```
fmt.Scan(&userName)
```

## 2024-06-26

### Loops and Control Statement


## 2024-06-27

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


## Strings

1. Can be accessed using square bracktes
prints the ASCII value of the character
```
s := "Hello World" 
fmt.Println(s[0])
```
2. Use `%c` to print the character of string.

```
s := 'Hello World'
fmt.Printf("%c", s[0])
```
3. Printing a part of the string. 

```
// prints total 5 charaters. 0 to 4
fmt.Println(s[0:5])
fmt.Println(s[:5])
```
4. Print string given only the starting index.
Prints from W. 
```
fmt.Println(s[6:])
```
5. String concat works normally. 
This would print `Hello World Again`

```
s = s + " Again"
fmt.Println(s)
```

7. String Literals

```
fmt.Println("Hello\nWorld")
fmt.Println("Hello\tWorld")
fmt.Println("Hello\bWorld")
```

## String Library

1. Built-in method `ToUpper` for converting 
string from lower case to uppercase. 

```
s := "Hello World"
fmt.Println(strings.ToUpper(s))
```

2. Built-in method `ToLower` for converting
string from upper case to lower case.

```
fmt.Println(strings.ToLower(s))
```
3. For check if string contains the prefix.

```
fmt.Println(strings.HasPrefix(s, "Hello"))
```
4. For checking if string contains suffix

```
fmt.Println(strings.HasSuffix(s, "World))
```
5. Replace. 
1st argument - full string.
2nd argument - string you want to replace
3rd argument - new string you want to replace old string with
4th argument - If the string to be replaced appear more than once. Count of strings you want to replace.  

```
fmt.Println(strings.Replace(s, "Hello", "World", 2))
```
