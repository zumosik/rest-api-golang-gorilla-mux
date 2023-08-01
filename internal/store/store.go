package store

type Store struct {
	config *Config
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	return nil
}

func (s *Store) Close() {

}
