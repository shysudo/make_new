# Golang make and new

In golang both make and new are premitives for allocating memory, but still there are some difference between them

We can explore more on them below

## defination of the 'new' function
    
        func new(type) *type

It can be seen its function parameter is a type and it return value is a pointer to memory address of its type.

We can also put, 

new(type) allocate "zeroed" storage for new item of type and returns its address, a value of  *type. In golang we can say, it returns a pointer to a newly allocated zero value of the type.

Even simplest we can put, new allocates the memory but not initalize memory