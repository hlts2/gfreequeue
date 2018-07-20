# gfreequeue


## Requirement
Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/gfreequeue
```
## Example

### Basic Example

Enqueu is `Enqueu(value interface{})`, so you can enqueu any type of object

```go

q := New()

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

q.Dequeue() // 1
q.Dequeue() // 2
q.Dequeue() // 3

```

### Iterator Support

```go

q := New()

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

q.Iterator() // [1, 2, 3]

/*
for _, v := range q.Iterator() {
  fmt.Println(v)
}
*/
```


## Author
[hlts2](https://github.com/hlts2)

## LICENSE
gfreequeue released under MIT license, refer [LICENSE](https://github.com/hlts2/gfreequeue/blob/master/LICENSE) file.
