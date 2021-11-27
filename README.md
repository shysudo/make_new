# Golang make and new

In golang both make and new are premitives for allocating memory, but still there are some difference between them.

We can explore more on them below.

## defination of the 'new' function
    
        func new(T) *T

It can be seen that its function parameter is a type T and it return value is a pointer to memory address of its type *T.

We can also put, 

new(T) allocate "zeroed" storage for new item of type T and returns its address, a value of type *T. In golang we can say, it returns a pointer to a newly allocated zero value of the type.

Even simplest we can put, new allocates the memory but not initalize memory

//Refer the code in newpackage
## main.go
        package main
        func main() {
                var i *int
	        var m *[2]int

                i = new(int)
	        m = new([2]int)
                
                fmt.Println(i, "    ", *i)
	        fmt.Println(m, "    ", *m)
        }

## output        
        0xc0000120b0      0
        &[0 0]      [0 0]

In output we can see the initialized memory address of i and m respectively and *i and *m are the zero values of int and [2]int type

## main.go
        package main
        func main() {
               type Person struct {
		        Name string
		        DOB  time.Time
	        }

                var p *Person
	        p = new(Person)
                
                fmt.Println(p, "    ", *p)
        }
## output        
       &{ 0001-01-01 00:00:00 +0000 UTC}      { 0001-01-01 00:00:00 +0000 UTC}

We can see the output, the initialized memory address of struct and the zero values of the type Person.

## defination of the 'make' function
    
        func make(t Type, size ...IntegerType) Type

It can be seen that its parameters are type and length(variadic) and returns a type.

The make() function, is a special built-in function that is used to initialize slices, maps, and channels. make() can only be used to initialize slices, maps and channels and it does not return a pointer.

make is designed to create these three built-in generic types.

As slice, maps and channels are reference types, 
        
	mc := new(map[string]string) 

Where new return *map[string]string and that mean mc is initialized to nil.
	
	(*mc)["str"] = "someone" 
	fmt.Println(mc) // this line will panic, assignment to entry in nil map.

nil can not be assigned directly.

End note :
The reason for the distinction in memory allocation is that these three type represent, under the covers, references to the data structures that must be initialize before use. for slice, map and channel make initialize the internal data structure and prepare the value for use.