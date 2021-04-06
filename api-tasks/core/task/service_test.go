package task_test

import (
	"database/sql"
	"testing"

	"github.com/gustavodfaguiar/learning-go-lang/core/task"
	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	taskStore := &task.Task{
		ID:          1,
		Title:       "Estudar golang",
		Description: "Aprender golang",
		Type:        task.TypePersonal,
	}

	db, err := sql.Open("sqlite3", "../../data/task_test.db")
	defer db.Close()
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}

	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando dados %s", err.Error())
	}

	service := task.NewService(db)
	err = service.Store(taskStore)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados %s", err.Error())
	}

	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados %s", err.Error())
	}

	if saved.ID == 1 {
		t.Fatalf("Dados inválidos. Esperado %d, recebido %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM task")

	tx.Commit()
	return err
}
