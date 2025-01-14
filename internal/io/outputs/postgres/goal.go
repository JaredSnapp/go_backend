package postgres

import "github.com/google/uuid"

func (r Repo[T]) DeleteGoal(id uuid.UUID) error {
	result := r.Pg.db.Delete(new(T), id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
