# gfreequeue


## Requirement
Go (>= 1.8)

## Installation

```shell
go get github.com/hlts2/gfreequeue
```
## Example

```go

q := New()

q.Enqueue(1)
q.Enqueue(2)
q.Enqueue(3)

q.Dequeue() // 1
q.Dequeue() // 2
q.Dequeue() // 3

```

## Author
[hlts2](https://github.com/hlts2)

## LICENSE
gfreequeue released under MIT license, refer [LICENSE](https://github.com/hlts2/gfreequeue/blob/master/LICENSE) file.
