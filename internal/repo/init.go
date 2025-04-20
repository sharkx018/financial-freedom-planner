package repo

type ResourceRepo interface {
}

type ResourceRepository struct{}

func NewResource() *ResourceRepository {
	return &ResourceRepository{}
}
