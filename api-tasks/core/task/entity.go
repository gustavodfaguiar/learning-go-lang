package task

/*

sqlite3 data/task.db

CREATE TABLE task (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   title text NOT NULL,
   description text NOT NULL,
   type integer not null
);
*/

type Task struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Type        TaskType `json:"type"`
}

type TaskType int64

const (
	TypePersonal = 1
	TypeWork     = 2
	TypeEnglish  = 3
)

// Desta forma eu posso transformar uma struct em um método e chamar assim por exemplo
// var type TypePersonal
// fmt.Println(type.String())

func (t TaskType) String() string {
	switch t {
	case TypePersonal:
		return "Personal"
	case TypeWork:
		return "Work"
	case TypeEnglish:
		return "English"
	}
	return "Unknown"
}
