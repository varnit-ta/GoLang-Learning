- ```go run ./<file_name>``` to run go program

- ```go build ./<file_name>``` to build executable

- ***Data types*** in GoLang :

    1. int (like int8, int16, int32, int64)
    2. float (like float32, float64)
    3. Complex (like complex64, complex128)
    4. bool
    5. string
    6. character (byte)
    7. pointer (*T)


- ***Declaring values***

    1. using ```var``` keyword (```var variablename type = value```)
        - can be used inside and outside a function
        - value assignment can be done separately

    2. with the ```:=``` sign (```variablename := value```)
        - can be used only inside a fucntion
        - value assignment cannot be done separately

    3. multiple value declaration
        - example : ```var a, b, c, d int = 1, 3, 5, 7```  or  ```c, d := 7, "World!"```

    4. declaring constants
        - ```const CONSTNAME type = vale```
        - constants can also be declared in blocks
        - There are two types of constant
            - ```const A int = 1``` typed constant
            - ```const A = 1``` untyped constant


- ***Go output functions***

    1. ```Print()```
        - prints arguments in their default formats

    2. ```Println()```
        - similar to ```Print()``` rather than the fact that it justs add a newline at the end

    3. ```Printf()```
        - The ```Printf()``` function first formats its argument based on the given formatting verb and then prints them.
        - Formatting verbs:
            - ```%v``` is used to print the **value** of the argument
            - ```%T``` is used to print the **type** of the argument

                ```go
                var(
                    i string = "Hello"
                    j int = 15
                )

                fmt.Printf("i has value: %v and type: %T\n", i, i)
                fmt.Printf("j has value: %v and type: %T", j, j)
                ```
        - other verbs can be seen in https://www.w3schools.com/go/go_formatting_verbs.php

- ***Declare an array***

    1. with the ```var``` keyword
        - with length defined ```var array_name = [length]datatype{values}```
        - wher length is inferred ```var array_name = [...]datatype{values}```

    2. with the ```:=``` sign 
        - ```array_name := [length / ...]datatype{values}```

- ***Array initialization***

    1. not initialized
        - ```arr1 := [5]int{}```

    2. partially initialized
        - ```arr2 := [5]int{1,2}```

    3. fully initialized
        - ```arr3 := [5]int{1,2,3,4,5}```

    4. initialize only specific elements
        - ```arr1 := [5]int{1:10,2:40}```

***NOTE: You can find length using the ```len``` keyword, for example, ```len(arr1)```***

- ***Go Slices***
    
    1. Similar to arrays, but unlike arrays, the length of slice can grow and shrink as you see fit

    2. In Go, there are two functions that can be used to return the length and capacity of a slice:
        - ```len()``` function - returns the length of the slice (the number of elements in the slice)
        - ```cap()``` function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

    3. Create a Slice 
        - With []datatype{values}
            - ```myslice := []int{}```
        
        - from an array
            - ```myslice := myarray[start:end]```
        
        - with ```make()``` command
            - ```slice_name := make([]type, length, capacity)```

    4. Append elements to slice
        - ```slice_name = append(slice_name, element1, element2, ...)```

        - append one slice to another slice
            - ```slice3 = append(slice1, slice2...)```

- ***Go Loops***

    1. ```for``` loops 

        ```go
        for statement1; statement2; statement3 {
            // code to be executed for each iteration
        }
        ```

        example:

        ```go
        for i:=0; i < 5; i++ {
            fmt.Println(i)
        }
        ```

    2. ```range``` keyword

        ```go
        for index, value := array|slice|map {
            // code to be executed for each iteration
        }
        ```

        example:

        This example uses ```range``` to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value):
        ```go
        fruits := [3]string{"apple", "orange", "banana"}
        for idx, val := range fruits {
            fmt.Printf("%v\t%v\n", idx, val)
        }
        ```

        Here, we want to omit the indexes (idx stores the index, val stores the value):
        ```go
        fruits := [3]string{"apple", "orange", "banana"}
        for _, val := range fruits {
            fmt.Printf("%v\n", val)
        }
        ```

- ***Go Maps***

    ```go
    mapName := map[key_dataType]value_dataType{
        key1: value1,
        key2: value2,
    }
    ```

    example:

    ```go
    b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
    ```

    Create maps using ```make()``` function

    ```go
    var a = make(map[KeyType]ValueType)
    b := make(map[KeyType]ValueType)
    ```

    example:

    ```go
    b := make(map[string]int)
    b["Oslo"] = 1
    b["Bergen"] = 2
    b["Trondheim"] = 3
    b["Stavanger"] = 4
    ```

- ***Go Struct***

    To declare a structure in Go, use the type and struct keywords

    ```go
    type struct_name struct {
        member1 datatype;
        member2 datatype;
        member3 datatype;
        ...
    }
    ```

    example:

    ```go
    type Person struct {
        name string
        age int
        job string
        salary int
    }
    ```


- ***Go functions***

    GO functions can also take fuctions as parameters and cna also return multiple values

    ```go
    func functionsName(paramters ,parameters) (returnType){
        //code
    }
    ```

    example:

    ```go
    func circleArea(r float64) (float64){
        return math.Pi * r * r
    }
    ```

    ```go
    func cycleName(n []string, f func(string)) {
        for _, v := range n {
            f(v)
        }
    }
    ```

    GO named return values

    ```go
    func myFunction(x int, y int) (result int) {
        result = x + y
        return  
    }
    ```

    GO multiple return multiple values

    ```go
    func myFunction(x int, y string) (result int, txt1 string) {
        result = x + x
        txt1 = y + " World!"
        return
    }
    ```

- **GO Receiver Functions**

    You can declare methods with pointer receivers.

    This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

    For example, the Scale method here is defined on *Vertex.

    Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

    Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.

    With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

    ```go
    type Vertex struct {
	X, Y float64
    }

    func (v Vertex) Abs() float64 {
        return math.Sqrt(v.X*v.X + v.Y*v.Y)
    }

    func (v *Vertex) Scale(f float64) {
        v.X = v.X * f
        v.Y = v.Y * f
    }

    func main() {
        v := Vertex{3, 4}
        fmt.Println(v.X, v.Y)
        v.Scale(10)
        fmt.Println(v.X, v.Y)
        fmt.Println(v.Abs())
    }
    ```

    Output:

    ```go
    3 4
    30 40
    50
    ```

- **GO Pointers**

    The type *T is a pointer to a T value. Its zero value is nil.

    ```go
    var p *int
    ```

    The & operator generates a pointer to its operand.

    ```go
    i := 42
    p = &i
    ```

    The * operator denotes the pointer's underlying value.

    ```go
    fmt.Println(*p) // read i through the pointer p
    *p = 21         // set i through the pointer p
    ```