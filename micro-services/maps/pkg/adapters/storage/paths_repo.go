package storage

import (
	"context"

	"gholi-fly-maps/internal/paths/domain"
	"gholi-fly-maps/pkg/adapters/storage/mapper"
	"gholi-fly-maps/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

type PathRepo struct {
	db *gorm.DB
}

// NewPathRepo creates a new PathRepo instance.
func NewPathRepo(db *gorm.DB) *PathRepo {
	return &PathRepo{db: db}
}

// GetAll retrieves all paths from the database.
func (r *PathRepo) GetAll(ctx context.Context) ([]domain.Path, error) {
	var dbPaths []types.Path
	if err := r.db.Find(&dbPaths).Error; err != nil {
		return nil, err
	}

	// Map database models to domain models
	var domainPaths []domain.Path
	for _, p := range dbPaths {
		domainPaths = append(domainPaths, *mapper.PathToDomain(&p))
	}

	return domainPaths, nil
}

// Create adds a new path to the database.
func (r *PathRepo) Create(ctx context.Context, path *domain.Path) error {
	dbPath := mapper.DomainToPath(path)
	return r.db.Create(dbPath).Error
}

// GetByID retrieves a path by its ID.
func (r *PathRepo) GetByID(ctx context.Context, id string) (*domain.Path, error) {
	var dbPath types.Path
	if err := r.db.First(&dbPath, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return mapper.PathToDomain(&dbPath), nil
}

// Update modifies an existing path in the database.
func (r *PathRepo) Update(ctx context.Context, path *domain.Path) error {
	dbPath := mapper.DomainToPath(path)
	return r.db.Save(dbPath).Error
}

// Delete removes a path by its ID.
func (r *PathRepo) Delete(ctx context.Context, id string) error {
	return r.db.Delete(&types.Path{}, "id = ?", id).Error
}
