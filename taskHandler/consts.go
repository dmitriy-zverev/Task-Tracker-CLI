package taskHandler

const (
	ADD = iota
	UPDATE
	DELETE
	MARK_IN_PROGRESS
	MARK_DONE
	LIST_ALL
	LIST_DONE
	LIST_TODO
	LIST_IN_PROGRESS
)

const (
	TODO = iota
	IN_PROGRESS
	DONE
)

const DATA_FILENAME = "tasks.todo"
const DATA_FILE_SEPARATOR = "^"
