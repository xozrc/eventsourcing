package eventsourcing

type System struct {
}

func (s *System) Init() error {
	return nil
}

func (s *System) Start() error {
	return nil
}

func (s *System) Stop() error {
	return nil
}

func (s *System) Destroy() error {
	return nil
}

func (s *System) ProcessCmd(cmd Command) error {
	return nil
}

func (s *System) ApplyEvent(e Event) error {
	return nil
}
