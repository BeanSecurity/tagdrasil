package manager

import (
	"errors"
	"log"
	"tagdrasil/models"
)

type Repository interface {
	GetTagByName(tagName string, userID int64) (models.TagNode, error)
	GetTagByID(tagID int64, userID int64) (models.TagNode, error)
	GetTagLine(tagName string, userID int64) (models.TagNode, error)
	GetSubtree(rootTagName string, userID int64) (models.TagNode, error)
	GetMetaTag(userID int64) (models.TagNode, error)
	GetRootTags(userID int64) ([]models.TagNode, error)
	SaveTag(tagName string, parentTagName string, userID int64) (int, error)
	DeleteTag(tagName string, userID int64) (int, error)
	GetUserByID(tagID int, userID int64) (models.User, error)
	SaveUser(user models.User, userID int64) (int64, error)
}

type TagManager struct {
	Repo Repository
}

func NewTagManager(repository Repository) *TagManager {
	return &TagManager{Repo: repository}
}

// func (m *TagManager) GetTagByName(tagName string) (models.TagNode, error) {
// 	m.Repo.GetTagByName(tagName)
// 	return models.TagNode{ID: 1, Name: "a"}, nil
// }

// func (m *TagManager) GetTagLine(tag models.TagNode, user models.User) (models.TagLine, error) {
// 	// return models.tagline{
// 	// 		models.tagnode{id: 1, name: "a"},
// 	// 		models.tagnode{id: 1, name: "a"}},
// 	// 	nil
// 	tagline, err := m.repo.gettagline(tag.name, user.id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tagline, nil
// 	// return m.repo.gettagline(tag.name, user.id)
// 	// return m.repo.gettagline(tag.name, user.id)
// }

func (m *TagManager) GetTagHeader(tags []models.TagNode, user models.User) (models.TagNode, error) {
	if len(tags) == 0 {
		err := errors.New("empty tags")
		log.Fatalf("%s", err)
		return models.TagNode{}, err
	}
	if len(tags) == 1 {
		tagLine, err := m.Repo.GetTagLine(tags[0].Name, user.ID)
		if err != nil {
			return models.TagNode{}, err
		}
		return tagLine, nil
	}

	return models.TagNode{}, nil
}

func (m *TagManager) GetTagBoardTree(user models.User) (models.TagNode, error) {
	return models.TagNode{}, nil
}

func (m *TagManager) GetSubtreeForTag(tag models.TagNode, user models.User) (models.TagNode, error) {
	// return models.TagNode{}, nil
	return m.Repo.GetSubtree(tag.Name, user.ID)
}

func (m *TagManager) AddRootTag(tag models.TagNode, user models.User) error {
	metaTag, err := m.Repo.GetMetaTag(user.ID)
	if err != nil {
		return err
	}
	_, err = m.Repo.SaveTag(tag.Name, metaTag.Name, user.ID)
	return err
	// return nil

}

func (m *TagManager) AddLeafTag(tag models.TagNode, user models.User) error {
	return nil
}

func (m *TagManager) DeleteTagByName(tag models.TagNode, user models.User) error {
	return nil
}

func (m *TagManager) CheckUser(user models.User) error {
	return nil
}
