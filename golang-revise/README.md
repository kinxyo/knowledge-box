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
