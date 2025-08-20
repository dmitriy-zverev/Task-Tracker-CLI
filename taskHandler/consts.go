package taskHandler

const (
	ADD              = "add"
	UPDATE           = "update"
	DELETE           = "delete"
	MARK_IN_PROGRESS = "mark-in-progress"
	MARK_DONE        = "mark-done"
	LIST             = "list"
)

const (
	TODO = iota
	IN_PROGRESS
	DONE
)

const DATA_FILENAME = "tasks.todo"
const DATA_FILE_SEPARATOR = "^"
