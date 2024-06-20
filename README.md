# GO
# Slice & Struct 
``` go
package main

import "fmt"

// Define a struct type
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create a slice of Person struct
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    // Accessing elements of the slice
    for _, person := range people {
        fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
    }

    // Modifying an element in the slice
    people[0].Age = 31
    fmt.Printf("Updated age of Alice: %d\n", people[0].Age)
}

```
When you create a slice of a struct type, each element in the slice is an instance of that struct. You can access and modify the fields of each struct in the slice just like you would with any other struct instance.
