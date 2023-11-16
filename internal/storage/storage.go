package storage

// Storage interface provides common methods which needs to be implemented by any storage implementation
type Storage interface {
	Produce([]byte) (int, error)
	Consume(int) ([]byte, error)
}
