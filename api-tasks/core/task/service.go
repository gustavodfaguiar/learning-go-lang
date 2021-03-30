package task

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
	Remove(ID int64) errorß
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetAll() ([]*Task, error) {
	return nil, nil
}

func (s *Service) Get(ID int64) ([]*Task, error) {
	return nil, nil
}

func (s *Service) Store(t *task) error {
	return nil, nil
}

func (s *Service) Update(t *task) error {
	return nil, nil
}

func (s *Service) Remove(ID int64) error {
	return nil, nil
}
