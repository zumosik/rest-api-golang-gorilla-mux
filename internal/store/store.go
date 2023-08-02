package store

type Store interface {
	User() UserRepository
	Link() LinkRepository
}
