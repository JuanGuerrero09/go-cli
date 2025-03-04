# To-Do Api

## Item

```go
type item struct {
Task        string
Done        bool
CreatedAt   time.Time
CompletedAt time.Time
}
```

add := flag.Bool("add", false, "Add task to the ToDo list")
del := flag.Int("del", 0, "Item to be deleted")
pen := flag.Bool("pen", false, "List all pending items")
list := flag.Bool("list", false, "List all tasks")
listTime := flag.Bool("listime", false, "List all tasks with time")
complete := flag.Int("complete", 0, "Item to be completed")
