# GO recap

## Data Types

### Basic

#### Declare Variables

```go
/* STATIC */

//declare_keyword name type
var x int; 

/* DYNAMIC */
x := 5;
```

#### Arrays

```go
var arr [5]int; //outputs 0 0 0 0 0
var arr [5]int = {1,2,3,4,5}; //outputs 1 2 3 4 5
var arr [...]int = {1,2,3,4,5}; //automatically infers the size 
```

#### Vectors

```go
var arr []int; //not assigning anything in the brackets make it a vector

arr = append(arr, 5) //uninitialized vector can only be appended values
arr = append(arr, 7)

//alternatively,

arr := []int{1,2,3,4}
```

#### Maps

```go
//declare_keyword name data_type[key_type]value_type
var users map[string]string;

users = make(map[string]string) //because maps aren't initialized when declared. 

users["kinjalk"] = "cat" //you don't "append" value, but "set" it.
```

In Go, when you declare a `map` variable without initializing it, it's assigned a
`nil` value. A `nil` map in Go means it doesn't point to any initialized hash table. In other words, it doesn't have any allocated memory to store key-value pairs.

When you try to add key-value pairs to a `nil` map, Go will throw a runtime panic because there's no underlying data structure to hold the data. That's why you need to initialize the map with `make` before you can add to it. The `make` function allocates the necessary memory and sets up the underlying data structure, allowing you to add key-value pairs to the map.

### Data Modelling

#### Structs

##### Struct

`int`, `string` etc. are all data types.

`var` is used to declare values of data types.

`type` is used to declare the data type.

`struct` is a custom data type.

_it's essentially a data type like `int` or `string` that can contain multiple varying data types_

(poor phrasing but i hope you get it)

- it's a data type that's why we specify it after _name_ 👇.
- it's a special data type that's why we declare it with type.
- `type` is used to define a new type.

```go
//declare_keyword name type
type Person struct {
    name string
    age int
}

func main() {
    var man Person = Person{"npc", 29}
    woman := Person{"npc", 29}
}
```

##### Method

```go
type Person struct {
    name string
    age  int
}

//Method
func (p *Person) gobackintime() { //weird syntax | "simplicity"
    p.age = 5
}

func main() {
    man := Person{"npc", 30}
    man.gobackintime()
    fmt.Println(man)
}
```

##### Constructor

these are just basic functions, not even associated to the structs like in every other language.

```go
func NewPerson(name string, age int) *Person {
    return &Person{name: name, age: age}
}
```

#### Interface

same as _java_.

```go
type Speaker interface {
    Speak() string
}

func (p *Person) Speak() string {
    return "Hello, my name is " + p.name
}
```

In this code, `Person` implements the `Speaker` interface because it has a `Speak` method.

## Data Structure

### Linked List

#### Step-by-step

This one is from my memory of what I can recall I studied in 1st year.
- linkedlist is a collection of nodes.
- node is a custom datatype or struct.
- it has 2 fields-- data & next.
- `data` stores the data it is meant to store.
- `next` stores pointer to another node since linkedlist is a collection of nodes.
- storing pointer to another node is what makes linkedlist connected or linked.

*(should be called "nodelist" instead; i find that to be descriptive)*

##### Node

```go
type Node struct {
	data int
	next *Node
}
```

##### Creating List

**First Element**
```go
var node1 Node = Node{5, new(Node)} //`new` returns a null pointer
```

**Second Element**
```go
var node2 Node = Node{5, new(Node)}
node1.next = &node2 //connecting back to first node
```

##### Printing Node

```go
func (node *Node) ShowNode() {
	fmt.Println(node.data)
}
```

```go
node1.ShowNode();
node2.ShowNode();
```

##### Printing List

before writing the function for it, let's first extend the current list (**optional**).

```go
var node3 Node = Node{15, new(Node)}
node2.next = &node3

var node4 Node = Node{20, new(Node)}
node3.next = &node4

var node5 Node = Node{25, new(Node)}
node4.next = &node5
```

now, the **function**.

```go
func (node *Node) ShowList() {
current := node
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
```

##### Adding to List

Creating a new element for the linked list can be quite cumbersome, but it doesn't really have to be.

1. Create a List
```go
func InitiateList() *Node {
	var headnode Node = Node{0, nil} //0? not neccessary; just my design
	return &headnode;
}
```

It is safe to return a reference (pointer) to a local variable in Go. When you create a local variable in Go, it's allocated on the stack. However, when you take the address of that variable to create a pointer, and that pointer escapes the local scope (for example, by being returned from the function), Go's compiler will automatically allocate that variable on the heap instead. This is known as escape analysis.

*this doesn't happen in Rust, that's why lifetimes are a thing*

2. Add to List

```go
func (headnode *Node) AddElement(data int) {

	//create a node
	new_element := &Node{data, nil}

	//if it's the first element
	if headnode.next == nil {
	
		headnode.next = new_element
	
	} else {

	//traverse to the last element
	
	current := headnode
	
	for current.next != nil {
		current = current.next
	}
	
	current.next = new_element
	
	}
}
```

##### Interface

```go
type LinkedList interface {
	AddElement(data int)
	ShowList()
}
```

done!

we can, if we want, add a little change to the function `InitializeList` to make the code more flexible and decoupled.

```go
func InitiateList() LinkedList {
	var headnode Node = Node{0, nil}
	return &headnode;
}
```

here, after the change, the function is capable of returning anything that implements `LinkedList`. Since we're returning `&headnode` which is `&Node` data type, that means our struct is implementing the interface correctly.

There is no explicity declaration for this,
`Node` implements `LinkedList` only because the their methods/functions match.

#### Full Code

```go
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList interface {
	AddElement(data int)
	ShowList()
}

func InitiateList() LinkedList {
	var headnode Node = Node{0, nil}
	return &headnode;
} 

func (headnode *Node) AddElement(data int) {
	new_element := &Node{data, nil}
	
	if headnode.next == nil {
	
		headnode.next = new_element

	} else {
	
		current := headnode
	
		for current.next != nil {
			current = current.next
		}
		
		current.next = new_element	
	}
}


func (node *Node) ShowList() {

	current := node
	current = current.next //skipping 1st element as it's always zero.

	for current.next != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	list := InitiateList()
	
	list.AddElement(5)
	list.AddElement(10)
	list.AddElement(15)
	list.AddElement(20)	
	list.AddElement(25)
	
	list.ShowList()
}
```
## Error Handling

### Default Errors
#### Returning Errors

Just like Rust, any function is capable of returning error in form `(data, error)`.

First import `errors` package.

```go
import (
	"errors"
)
```

now you can return error from any function. 

For instance, this function ⬇

```go
func divide(a, b float64) float64 {
	return a/b
}
```

can be written as:

```go
func divide(a, b float64) (float64, error) {
	return a / b, nil
}
```

we can then add a condition to return a value.

```go
func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("division by 0 is undefined")
	}

	return a / b, nil
}
```

*simple!*
#### Now Catching Errors

```go
func main() {

	answer, err := divide(15,0)

	if err != nil {
		fmt.Println(err)
		return //early return so the function ends here itself.
	}

	fmt.Println(answer)
}

```

### Custom Errors

1. Create a struct of what information you want your custom error to report.
```go
   type CustomError struct {
	message string
	code    int
}
```

2. Add interface of `Error` and specify the presentation of report
```go
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %d\n%s", e.code, e.message)
}
```

3. Return it from any function

```go
func divide(a, b float64) (float64, *CustomError) {

	if b == 0 {
		return 0, &CustomError{
			message: "why you dividing with 0?",
			code:    403,
		}
	}

	return a / b, nil
}
```

### Panic & Recovery

#### Panic

same as `Rust 🦀`, use it when you want to crash

```go
func main() {
    fmt.Println("start")
    panic("something bad happened")
    fmt.Println("end")
}
```

#### Recover

There are 2 components to recovery-- **catching panic** & **doing something with it**.
we'll have to talk about this in reverse order:

1. doing something with panic
```go
func RecoverPanic() {
    if r := recover(); r != nil { //confirming panic
        fmt.Println("Recovered from", r)
    }
}   
```

2. catch the panic

`defer` keyword before any function means that that function will be executed after all the other statements in the scope.

in this context, it means that `defer` will run `RecoverPanic` as soon as all statement in the scope of `main` function executes **BUT** before `main` returns anything.

```go
func main() {
	defer RecoverPanic()
	panic("gotta crash!")
}
```

thus when the `panic` happens, it ends the flow of program, causing `main` to abruptly return. It's at that moment `defer` catches the `panic` and does something with it.

## Concurrency

- **Goroutines** are functions or methods that run concurrently with other functions or methods. Goroutines are lightweight threads managed by the Go runtime. You can start a goroutine simply by adding the keyword `go` in front of a function call.
- **Channels** are a means of communication between goroutines. They allow you to pass data between goroutines in a safe and controlled manner. Channels ensure that operations (send/receive) are atomic, which means only one goroutine accesses the data at any given point of time.

### goroutines

#### Introduction

> Yes, goroutines in Go are similar to green threads.
> 
> However, Goroutines are more lightweight than green threads and are multiplexed onto a small number of OS threads, rather than each having a direct correspondence to an OS thread. This allows for efficient scheduling and avoids the context switch overhead of traditional threads. Also, goroutines inherently support communication via channels, which is not a typical feature of green threads.

\- GitHub Copilot

**Features:**
- independent
- shared memory
- non-blocking

#### How to Use

##### Basic

just add `go` before the function like `async`.

```go
func printNumbers() {
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func printLetters() {
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main() {
    go printNumbers()
    go printLetters()

    // Wait enough time to ensure both goroutines finish
    time.Sleep(3000 * time.Millisecond)
    fmt.Println("\nDone")
}
```


##### Proper Way

Earlier, we were making the main thread sleep so that other threads (*goroutines*) could finish their stuff. This way isn't reliable, hence this is the proper way.

We use `sync` package for *blocking* the main thread & using *mutex* to avoid data race conditions.

###### Enable Blocking

Blocking mechanism in `Go` is quite different, as the goroutines (green threads) are different, and quite simple actually!

- For blocking, we first create a `counter`.
- The role of that `counter` is, to be a tracking point.
- Meaning, the program ends when the `counter` is zero.
- When the `counter` is 0, it's a signal to the Go program that it can end now.
- The `counter` is provided by the `WaitGroup`; present in the `sync` package.
- We create the counter object by defining the type for a variable as `WaitGroup`.
- That `counter` has methods attatched to it for:
	- `adding to the counter`
	- `reducing from the counter`.

So the mechanism is,

- We add to the `counter`, the amount of gorountines we're going to run, at start of the program.
- As discussed before, `defer` keyword is used to run any function at the end of a scope.
- `defer` will run the function after everything in the scope is executed.
- When any goroutine starts, then it ends, it's supposed to reduce the counter.
- We achieve this by sending a reference to the counter object inside the goroutine function, then uses the reducing method.
- As soon as the counter reaches 0, the program might end, so to ensure the counter reaches 0 after everything is executed, we use `defer` keyword for the reducing method.

```go
func printNumbers(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}

func printLetters(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)

    go printNumbers(&wg)
    go printLetters(&wg)

    wg.Wait()
    fmt.Println("\nDone")
}
```

###### Mutex

We can use `Mutex` in 2 manners:-
1.  Directly.
2.  Group it with some value to create association.

*It doesn't affect the program either way, it's just upto the deveoper to prefer.*


**Directly:**

```go
var mutex sync.Mutex

func p1(arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*arr = append(*arr, 500)
	mutex.Unlock()
}
func p2(arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*arr = append(*arr, 650)
	mutex.Unlock()
}
func p3(arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*arr = append(*arr, 900)
	mutex.Unlock()
}
func p4(arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*arr = append(*arr, 100)
	mutex.Unlock()
}
func p5(arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*arr = append(*arr, 200)
	mutex.Unlock()
}
func p6(arr *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	*arr = append(*arr, 800)
	mutex.Unlock()
}

func main() {

	var wg sync.WaitGroup
	wg.Add(6)

	var donation []int

	go p1(&donation, &wg)
	go p2(&donation, &wg)
	go p3(&donation, &wg)
	go p4(&donation, &wg)
	go p5(&donation, &wg)
	go p6(&donation, &wg)

	wg.Wait()
	fmt.Println(donation)

}
```



**Associative approach:**

```go
type SafeSlice struct {
    mu sync.Mutex
    slice []int
}

func (ss *SafeSlice) Append(value int) {
    ss.mu.Lock()
    defer ss.mu.Unlock()
    ss.slice = append(ss.slice, value)
}

func p1(ss *SafeSlice, wg *sync.WaitGroup) {
    defer wg.Done()
    ss.Append(500)
}

// Similarly modify p2, p3, etc.

func main() {
    var wg sync.WaitGroup
    wg.Add(6)

    safeSlice := &SafeSlice{}

    go p1(safeSlice, &wg)
    // Start p2, p3, etc. as goroutines

    wg.Wait()
    fmt.Println(safeSlice.slice)
}
```

> **associative approach** is also correct but I see no use of this; as of now.
> 
> To use mutex:
> 1. I just have to call `mutex.Lock()`  
> 2. Write on the data
> 3. then call `mutex.Unlock()` 
> 
> It gets the job done. Go is supposed to be simple, why complicate it further? Hence **direct approach** is better for most cases.
> (but this approach of wrapping mutex utilizing lines of code into its own function seems cool.)

\- my cents

### channels

#### Introduction

Channels is a way to communicating between the goroutines.

If our green threads can communicate with each other then:
1. We don't need to manually block main thread using `WatiGroup`.
2. We don't need `mutex` (*as they will not attempt to write on the same data together since they can communicate*).
#### How to use

##### Basic

The `<-` symbol is used for receiving/sending values from a channel.

To use channels, we first create a variable (type: channel), and also specify the type of variable it will hold-- `make(chan string)`.

- `chan <- value`: This means we are sending a `value` into the `chan` channel.
- `value := <- chan`: This means we are receiving a value from the `chan` channel.

```go
func goroutine(message chan string) {
	message <- "ping"
}

func main() {
	messages := make(chan string)

	go goroutine(messages)

	msg := <-messages

	fmt.Println(msg)
}
```

##### Enable blocking via channels

```go
func goroutine(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    
    done <- true // sends a value to notify that this function is done.
}

func main() {
    done := make(chan bool, 1)

    go goroutine(done)
    
    <-done // blocks `main` until we receive a notification.
}
```

##### Channels over `mutex`

```go
func goroutine(items chan int) {
	for item := range items { //This line continuously receives values from the items channel.
		fmt.Println("Working on item", item)
		time.Sleep(time.Second)
		fmt.Println("Finished item", item)
	}
}

func main() {
	items := make(chan int)

	go goroutine(items) //here we're just sending a channel

	/* a channel is platform for communication (passing values for writing/reading) */
	
	i := 0
	for i <= 5 {
		i++
		items <- i //here we're passing values over the platform (channel)
	}

	close(items) //closing the platform; no more values can be passed!
}
```

#### Additionally,

In Go, channels can be either unbuffered (no capacity) or buffered (with a capacity).

- An **unbuffered channel** has no capacity and therefore, sends and receives must be ready at the same time. If a goroutine tries to send to an unbuffered channel and there's no goroutine ready to receive from that channel, the sender goroutine blocks until a receiver is ready. Similarly, if a goroutine tries to receive from an unbuffered channel and there's no goroutine ready to send to that channel, the receiver goroutine blocks until a sender is ready.
- A **buffered channel** has a queue of elements. The capacity of the channel is the size of this queue. A send operation on a buffered channel inserts an element at the back of the queue, and a receive operation removes an element from the front. If the queue is full (i.e., its size is equal to the channel's capacity), the sender goroutine blocks until there's space in the queue. If the queue is empty, the receiver goroutine blocks until there's an element in the queue.
