package storage

import (
	"context"
	"fmt"
	"gholi-fly-maps/internal/terminals/domain"
	"gholi-fly-maps/internal/terminals/port"
	"gholi-fly-maps/pkg/adapters/storage/mapper"
	"gholi-fly-maps/pkg/adapters/storage/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TerminalRepo struct {
	db *gorm.DB
}

// NewTerminalRepo creates a new TerminalRepo instance.
func NewTerminalRepo(db *gorm.DB) *TerminalRepo {
	return &TerminalRepo{db: db}
}

// GetAll retrieves all terminals from the database.
func (r *TerminalRepo) GetAll(ctx context.Context) ([]domain.Terminal, error) {
	var dbTerminals []types.Terminal
	if err := r.db.Find(&dbTerminals).Error; err != nil {
		return nil, err
	}

	var domainTerminals []domain.Terminal
	for _, t := range dbTerminals {
		domainTerminals = append(domainTerminals, *mapper.TerminalToDomain(&t))
	}

	return domainTerminals, nil
}

func (r *TerminalRepo) Create(ctx context.Context, terminal *domain.Terminal) error {
	dbTerminal := mapper.DomainToTerminal(terminal)

	if dbTerminal.Location == "" {
		return fmt.Errorf("location is empty in repository")
	}

	return r.db.WithContext(ctx).Create(dbTerminal).Error
}

// GetByID retrieves a terminal by its ID.
func (r *TerminalRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Terminal, error) {
	var dbTerminal types.Terminal
	if err := r.db.First(&dbTerminal, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return mapper.TerminalToDomain(&dbTerminal), nil
}

// Update modifies an existing terminal.
func (r *TerminalRepo) Update(ctx context.Context, terminal *domain.Terminal) error {
	dbTerminal := mapper.DomainToTerminal(terminal)
	return r.db.Save(dbTerminal).Error
}

// Delete removes a terminal by its ID.
func (r *TerminalRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.Delete(&types.Terminal{}, "id = ?", id).Error
}

func (r *TerminalRepo) Search(ctx context.Context, filter port.TerminalFilter) ([]domain.Terminal, error) {
    var dbTerminals []types.Terminal
    query := r.db.Model(&types.Terminal{})

    // Apply filters dynamically
    if filter.ID != "" {
        query = query.Where("id = ?", filter.ID)
    }
    if filter.Name != "" {
        query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
    }
    if filter.City != "" {
        query = query.Where("city ILIKE ?", "%"+filter.City+"%")
    }
    if filter.Type != "" {
        query = query.Where("type = ?", filter.Type)
    }

    if err := query.Find(&dbTerminals).Error; err != nil {
        return nil, err
    }

    var domainTerminals []domain.Terminal
    for _, t := range dbTerminals {
        domainTerminals = append(domainTerminals, *mapper.TerminalToDomain(&t))
    }

    return domainTerminals, nil
}