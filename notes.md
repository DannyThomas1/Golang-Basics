## GO Notes

### Declaring Variables

- There are two ways to declare a variable

  ```
  var card string = "Ace of Spades"
  ```

  OR

  ```
  card := "Ace of Spades"
  ```

  - ':=' tells GO to infer the type based on the right sides value

### Arrays vs Slices

- **_Arrays:_** Fixed length list of things
- **_Slices:_** An array that can grow or shrink

  - Every element in a slice must be of same type
  - Slices are zero-indexed

    ```
    //Declaring a splice
    cards := []string{"Ace of Diamonds"}

    //Adding to a splice, append makes a new slice instead of applying the push to original splice
    cards = append(cards, "Six of Hearts")

    //Iterating through a splice
    for i,card := range cards {
      fmt.Println(card)
    }

    //Get a range of values
    cards[startIndexIncluding: upToNotIncluding]
    ```

  - When we create a slice, Go automatically creates an array and a structure that records the length of the slice, the capacity of the slice and a reference to the underlying array.

### Type Conversion

- To change the type of an exisiting type, we use the following syntax

```
  greeting := "Hi There"
  fmt.Println([]byte(greeting)) //[]TypeWeWant(TypeWeHave)
```

### Pointer Operations

- To get the memory address of the value the variable is pointing at, use:
  ```
  &variable
  ```
- To get the value the memory address is pointing at, use:

  ```
  *pointer
  ```

  - When there is a star with a type declaration (*person), it means we're working with a pointer to a person and known as a type description. When there is a star with a pointer (*pointerToPerson),it means we want to manipulate the value the pointer is referencing, this is known as an operator.
    ```
    func (pointerToPerson *person) updateName(){
      *pointerToPerson
    }
    ```

### Types

- Reference Types (a reference to another data structure in memory)
  - slices
  - maps
  - channels
  - pointers
  - functions
- Value Types (use pointers to change values)
  - int
  - float
  - string
  - bool
  - structs

### Maps

- All keys need to be of the same type, same with the values. Keys and values dont need to be the same type.
- A zero value for map is an empty map
- Accessing/making a key/value piar, bracket notation is needed, cannot use dot notation.
- Declaring maps:

  ```
    //Method 1
    colors := map[string]string{
      "red":   "#ff0000",
      "green": "#4bf934",
    }

    //Method 2
    var colors map[string]string

    //Method 3
    colors := make(map[string]string)

  ```

- Deleting a value:
  ```
    //Built in method, need to feed the map and the key
    delete(colors, "white")
  ```

### Differences between Map and Struct

- Map:
  - All keys must of the same type
  - Use to represent a collection of related properties
  - All values must be the same type
  - Dont need to knwo all the keys at compile time
  - Keys are indexed, we can iterate over them
  - Reference Type!
- Struct:
  - Values can be of different type
  - You need to know all the different fields at compile time
  - Keys dont support indexing
  - Use to represent a "thing" with a lot of different properties
  - Value Type!

### Interfaces

- When you want to reuse the same logic to apply to different types
- Interfaces are not generic types
- Interfaces are "implicit", we dont have to manually say that our custom type satisfies some interface
- Interfaces are a contract to help us manage types

### Routines

- Keyword 'go' before a function call is used to create new GO routines
- Used to run code using different threads
- A GO scheduler (By default GO tries to use one CPU core) is used to run one routine until it finishes or makes a blocking call, it then looks for other routines able to execute.
- If we use more than one CPU core, the scheduler runs one thread on each "logical" core
- **Concurreny** - We can have multiple threads executing code. If one thread blocks another one is picked up and worked on.
- **Parallelism** - Multiple threads executed at the exact same time. Requires multiple CPU's.
- There is one Main Routine that start when the program is executed, any other routines called with 'go' are called Child routines.

### Channels

- Channels are used for go routines to communicate between all the running routines.
- Channels are typed, information we pass into a channel must all be of the same type
- Can be used to tell the Main routine that Child routines are still running and check their statuses before Main quits

  ```
  //To make a channel
  c := make(chan string)

  //Send value '5' into this channel
  channel <- 5

  //Wait for a value to be sent into this channel. When we get one, assign the value to 'myNumber'
  myNumber <- channel

  //Wait for a value to be sent into the channel. When we get one, log it out immediately
  fmt.Println(<- channel)
  ```

- Receiving messages from a channel is considered a blocking code.
- Channel Gotchas
  - Never reference a variable inside of two different routines

### Function Literal

- Like Javascripts Anonymous Function

```
 func(){
  //code
 }()
```
