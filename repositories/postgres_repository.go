package tagdrasil

import (
	models "tagdrasil/models"
)

type TagRepository interface {
	GetTagByName(tagName string) (models.Tag, error)
	GetTagById(tagId int64) (models.Tag, error)
	GetTagLine(tag models.Tag) ([]models.Tag, error)
	SaveTag(tag models.Tag) error
	DeleteTag(tag models.Tag) error
	GetUserById(tagId int) (models.User, error)
	SaveUser(user models.User) error
	DeleteUser(user models.User) error
}

type TagPostgresRepository struct {
}

func NewTagPostgresRepository() TagRepository {
	return &TagPostgresRepository{}
}

func (r *TagPostgresRepository) GetTagByName(tagName string) (models.Tag, error) {
	return models.Tag{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetTagById(tagId int64) (models.Tag, error) {
	return models.Tag{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetTagLine(tag models.Tag) ([]models.Tag, error) {
	return []models.Tag{
			models.Tag{ID: 1, Name: "a"},
			models.Tag{ID: 1, Name: "a"}},
		nil
}

func (r *TagPostgresRepository) SaveTag(tag models.Tag) error {
	return nil
}

func (r *TagPostgresRepository) DeleteTag(tag models.Tag) error {
	return nil
}

func (r *TagPostgresRepository) GetUserById(tagId int) (models.User, error) {
	return models.User{ID: 1, FirstName: "Vasya"}, nil
}

func (r *TagPostgresRepository) SaveUser(user models.User) error {
	return nil
}

func (r *TagPostgresRepository) DeleteUser(user models.User) error {
	return nil
}
