# todolist
todo CLI application. Fun project in golang

`go build cmd/todo/`

Use the binary to use the CLI TODO list tool

The cli allows `add`, `delete`, `list`, `complete` flags.

### Add a task to TODO:

`./todo -task "A new task!"`

### List existing tasks:
`./todo -list`

### Mark an existing task as complete:
`./todo -complete <index of task>`

Index of task represents the number associated to task. If the task was added at the beginning, it will be considered to have index `1`.

### Delete an existing task:
`./todo -del <index of task>`

Index of task represents the number associated to task. If the task was added at the beginning, it will be considered to have index `1`.
