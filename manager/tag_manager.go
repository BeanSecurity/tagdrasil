package tagdrasil

import (
	models "tagdrasil/models"
	repository "tagdrasil/repositories"
)

type TagManager interface {
	GetTagByName(tagName string) (models.Tag, error)
	GetTagLine(tag models.Tag) ([]models.Tag, error)
	GetAllTags() (models.TagTree, error)
	AddRootTag(tag models.Tag) error
	AddLeafTag(tag models.Tag) error
	DeleteTagByName(tagName string) error
	SaveUser(user models.User) error
}

type TagTelegramManager struct {
	TagRepo repository.TagRepository
}

func NewTagTelegramManager(repo repository.TagRepository) *TagTelegramManager {
	return &TagTelegramManager{TagRepo: repo}
}

func (m *TagTelegramManager) GetTagByName(tagName string) (models.Tag, error) {
	return models.Tag{ID: 1, Name: "a"}, nil
}

func (m *TagTelegramManager) GetTagLine(tag models.Tag) ([]models.Tag, error) {
	return []models.Tag{
			models.Tag{ID: 1, Name: "a"},
			models.Tag{ID: 1, Name: "a"}},
		nil
}

func (m *TagTelegramManager) GetAllTags() (models.TagTree, error) {
	return models.TagTree{}, nil
}

func (m *TagTelegramManager) AddRootTag(tag models.Tag) error {
	return nil
}

func (m *TagTelegramManager) AddLeafTag(tag models.Tag) error {
	return nil
}

func (m *TagTelegramManager) DeleteTagByName(tagName string) error {
	return nil
}

func (r *TagTelegramManager) SaveUser(user models.User) error {
	return nil
}
