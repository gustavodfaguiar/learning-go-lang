package task

import (
	"database/sql"
	"fmt"
)

// type Reader interface {
// 	GetAll() ([]*Task, error)
// 	Get(ID int64) (*Task, error)
// }

// type Writer interface {
// 	Store(t *Task) error
// 	Update(t *Task) error
// 	Remove(ID int64) error
// }

type UseCase interface {
	GetAll() ([]*Task, error)
	Get(ID int64) (*Task, error)
	Store(t *Task) error
	Update(t *Task) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Task, error) {
	var result []*Task

	rows, err := s.DB.Query("SELECT id, title, description, type FROM task")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Type)
		if err != nil {
			return nil, err
		}

		result = append(result, &t)
	}

	return result, nil
}

func (s *Service) Get(ID int64) (*Task, error) {
	var t Task

	stmt, err := s.DB.Prepare("SELECT id, title, description, type FROM task WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&t.ID, &t.Title, &t.Description, &t.Type)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *Service) Store(t *Task) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO task(id, title, description, type) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.ID, t.Title, t.Description, t.Type)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (s *Service) Update(t *Task) error {
	if t.ID == 0 {
		return fmt.Errorf("Invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE task SET title=?, description=?, type=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Title, t.Description, t.Type, t.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("Invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM task WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
