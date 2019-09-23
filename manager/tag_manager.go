package manager

import (
	"errors"
	"log"
	"tagdrasil/models"
)

type Repository interface {
	GetTagByName(tagName string, userID int) (models.TagNode, error)
	GetTagByID(tagID int64, userID int) (models.TagNode, error)
	GetTagLine(tagName string, userID int) (models.TagNode, error)
	GetSubtree(rootTagName string, userID int) (models.TagNode, error)
	GetMetaTag(userID int) (models.TagNode, error)
	GetRootTags(userID int) ([]models.TagNode, error)
	SaveTag(tagName string, parentTagName string, userID int) error
	DeleteTag(tagName string, userID int) (int, error)
	GetUserByID(userID int) (models.User, error)
	SaveUser(user models.User) error
	DeleteUserByID(userID int) error
}

type TagManager struct {
	Repo Repository
}

func NewTagManager(repository Repository) (*TagManager, error) {
	return &TagManager{Repo: repository}, nil
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
	switch len(tags) {
	case 0:
		err := errors.New("empty tags")
		log.Println("%s", err)
		return models.TagNode{}, err
	case 1:
		tagLine, err := m.Repo.GetTagLine(tags[0].Name, user.ID)
		if err != nil {
			log.Println("%s", err)
			return models.TagNode{}, err
		}
		return tagLine, nil
		// default:
		// 	for _, tag := range tags {
		// 		tagLine, err := m.Repo.GetTagLine(tag.Name, user.ID)
		// 	}
	}

	return models.TagNode{}, nil
}

func (m *TagManager) GetTagBoardTree(user models.User) (models.TagNode, error) {
	metaTag, err := m.Repo.GetMetaTag(user.ID)
	if err != nil {
		log.Println("%s", err)
		return models.TagNode{}, err
	}

	tag, err := m.Repo.GetSubtree(metaTag.Name, user.ID)
	if err != nil {
		log.Println("%s", err)
		return models.TagNode{}, err
	}
	return tag, nil
}

func (m *TagManager) GetSubtree(tag models.TagNode, user models.User) (models.TagNode, error) {
	tag, err := m.Repo.GetSubtree(tag.Name, user.ID)
	if err != nil {
		log.Println("%s", err)
		return models.TagNode{}, err
	}
	return tag, nil
}

func (m *TagManager) AddRootTag(tagName string, user models.User) error {
	metaTag, err := m.Repo.GetMetaTag(user.ID)
	if err != nil {
		log.Println("%s", err)
		return err
	}
	err = m.Repo.SaveTag(tagName, metaTag.Name, user.ID)
	if err != nil {
		log.Println("%s", err)
		return err
	}
	return nil

}

func (m *TagManager) AddLeafTag(childTagName, parentTagName string, user models.User) error {
	err := m.Repo.SaveTag(childTagName, parentTagName, user.ID)
	if err != nil {
		log.Println("%s", err)
		return err
	}
	return nil
}

func (m *TagManager) DeleteTagByName(tag models.TagNode, user models.User) error {
	// you cant delete metatag
	return nil
}

// func (m *TagManager) CheckUser(user models.User) error {
// 	return nil
// }

func (m *TagManager) SaveUser(user models.User) error {
	err := m.Repo.SaveUser(user)
	if err != nil {
		log.Println(err)
		return err
	}
	// m.
	return nil
}

func (m *TagManager) GetUserByID(userID int) (models.User, error) {
	return models.User{}, nil
}

func (m *TagManager) DeleteUserByID(userID int) error {
	return nil
}
