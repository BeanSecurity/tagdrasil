package repository

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"log"
	models "tagdrasil/models"
)

// type Repository interface {
// 	GetTagByName(tagName string, userID int) (TagNode, error)
// 	GetTagByID(tagID int64, userID int) (TagNode, error)
// 	GetTagLine(lastTag TagNode, userID int) (models.TagLine, error)
// 	SaveTag(tag TagNode, userID int) (int, error)
// 	DeleteTag(tag TagNode, userID int) (int, error)
// 	GetUserByID(tagID int, userID int) (models.User, error)
// 	SaveUser(user models.User, userID int) (int, error)
// 	DeleteUser(user models.User, userID int) (int, error)

type TagPostgresRepository struct {
	db *sql.DB
}

func NewTagPostgresRepository(dsn string) *TagPostgresRepository {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &TagPostgresRepository{db}
	// return TagPostgresRepository{nil}
}

func (r *TagPostgresRepository) GetTagByName(tagName string, userID int64) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetTagByID(tagID int64, userID int64) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetTagLine(tagName string, userID int64) (models.TagNode, error) {
	return models.TagNode{1, "a", []models.TagNode{
			models.TagNode{ID: 3, Name: "c"}}},
		nil
}

func (r *TagPostgresRepository) GetSubtree(rootTagName string, userID int64) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetMetaTag(userID int64) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetRootTags(userID int64) ([]models.TagNode, error) {
	return []models.TagNode{
			models.TagNode{ID: 1, Name: "a"},
			models.TagNode{ID: 1, Name: "a"}},
		nil
}

func (r *TagPostgresRepository) SaveTag(tagName string, parentTagName string, userID int64) (int, error) {
	return 1, nil
}

func (r *TagPostgresRepository) DeleteTag(tagName string, userID int64) (int, error) {
	return 1, nil
}

func (r *TagPostgresRepository) GetUserByID(tagID int, userID int64) (models.User, error) {
	return models.User{ID: 1, FirstName: "Vasya"}, nil
}

func (r *TagPostgresRepository) SaveUser(user models.User, userID int64) (int64, error) {
	return 1, nil
}

// func (r *TagPostgresRepository) DeleteUser(user models.User, userID int64) (int64, error) {
// 	return 1, nil
// }
