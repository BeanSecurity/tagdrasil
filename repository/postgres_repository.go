package repository

import (
	"database/sql"
	"errors"
	_ "github.com/bmizerany/pq"
	"log"
	"tagdrasil/models"
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

// func NewTagPostgresRepository(dsn string) *TagPostgresRepository {
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &TagPostgresRepository{db}
// 	// return TagPostgresRepository{nil}
// }

func NewTagPostgresRepository(db *sql.DB) (*TagPostgresRepository, error) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
		return &TagPostgresRepository{}, err
	}

	q := `
CREATE TABLE IF NOT EXISTS telegram_user (
	user_id SERIAL PRIMARY KEY,
	username TEXT NOT NULL,
	full_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tag (
	tag_id SERIAL PRIMARY KEY,
	name TEXT NOT NULL UNIQUE ,
	user_id INTEGER NOT NULL REFERENCES telegram_user ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tag_relations (
	parent_tag_id INTEGER REFERENCES tag ON DELETE CASCADE,
	child_tag_id INTEGER REFERENCES tag ON DELETE CASCADE,
	UNIQUE(parent_tag_id,child_tag_id)
);
`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatal(err)
		return &TagPostgresRepository{}, err
	}
	return &TagPostgresRepository{db}, nil
}

func (r *TagPostgresRepository) GetTagByName(tagName string, userID int) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetTagByID(tagID int64, userID int) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetTagLine(tagName string, userID int) (models.TagNode, error) {
	stmt, err := r.db.Prepare(`
-- recursive query for tag line without string
WITH RECURSIVE r AS (
  SELECT tag.tag_id, tag.name, (SELECT parent_tag_id FROM tag_relations WHERE child_tag_id=tag.tag_id) as parent_id, 0 as level
    FROM tag
   WHERE name = $1 AND user_id=$2-- line leaf

   UNION

  SELECT tag.tag_id, tag.name, (SELECT parent_tag_id FROM tag_relations WHERE child_tag_id=tag.tag_id) AS parent_id, r.level+1 as level
       FROM tag
              JOIN r
                  ON tag.tag_id = r.parent_id
   )
SELECT tag_id,name FROM r;
`)
	if err != nil {
		log.Fatal(err)
		return models.TagNode{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(tagName, userID)
	if err != nil {
		log.Fatal(err)
		return models.TagNode{}, err
	}
	defer rows.Close()

	var id int64
	var name string

	var lastTag models.TagNode
	rows.Next()
	rows.Scan(&id, &name)
	lastTag = models.TagNode{id, name, nil}

	for rows.Next() {
		rows.Scan(&id, &name)
		lastTag = models.TagNode{id, name, []models.TagNode{lastTag}}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return models.TagNode{}, err
	}
	if lastTag.Name == "" && lastTag.ID == 0 && lastTag.ChildTags == nil {
		return models.TagNode{}, errors.New("no such tag with name: " + tagName)
	}

	return lastTag, nil
}

func (r *TagPostgresRepository) GetSubtree(rootTagName string, userID int) (models.TagNode, error) {
	return models.TagNode{ID: 1, Name: "a"}, nil
}

func (r *TagPostgresRepository) GetMetaTag(userID int) (models.TagNode, error) {
	stmt, err := r.db.Prepare(`
-- get metatag
SELECT tag.tag_id, tag.name
    FROM tag
 WHERE tag_id NOT IN (SELECT child_tag_id
                        FROM tag_relations
                               JOIN tag ON child_tag_id=tag_id
                       WHERE tag.user_id=$1)
   AND user_id=$1;
`)
	if err != nil {
		log.Fatal(err)
		return models.TagNode{}, err
	}
	defer stmt.Close()

	var tagID int64
	var tagName string
	err = stmt.QueryRow(userID).Scan(&tagID, &tagName)
	if err != nil {
		log.Fatal(err)
		return models.TagNode{}, err
	}
	return models.TagNode{ID: tagID, Name: tagName}, nil
}

func (r *TagPostgresRepository) GetRootTags(userID int) ([]models.TagNode, error) {
	return []models.TagNode{
			models.TagNode{ID: 1, Name: "a"},
			models.TagNode{ID: 1, Name: "a"}},
		nil
}

func (r *TagPostgresRepository) SaveTag(tagName, parentTagName string, userID int) error {
	stmt, err := r.db.Prepare(`
WITH
	child_id AS (INSERT INTO tag (name, user_id) VALUES ($1, $2) RETURNING tag_id),
	parent_id AS (SELECT tag_id FROM tag WHERE name = $3 AND user_id=$2)
    INSERT
    INTO
	  tag_relations (parent_tag_id, child_tag_id)
    VALUES ((SELECT * FROM parent_id), (SELECT * FROM child_id));
`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(tagName, userID, parentTagName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	n, err := res.RowsAffected()
	if n == 0 {
		return errors.New("no tag added")
	}
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (r *TagPostgresRepository) DeleteTag(tagName string, userID int) (int, error) {
	return 1, nil
}

func (r *TagPostgresRepository) GetUserByID(userID int) (models.User, error) {

	// row, err := r.db.QueryRow(insertQ, "notes", 0, "tagdrasil")

	return models.User{ID: 1, FirstName: "Vasya"}, nil
}

func (r *TagPostgresRepository) SaveUser(user models.User) error {
	stmt, err := r.db.Prepare(`
INSERT INTO telegram_user (user_id,username, full_name)
VALUES ($1, $2, $3) ON CONFLICT DO NOTHING;
`)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.ID, user.UserName, user.FirstName+user.LastName)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// if its not a new user, add metatag
	n, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if n == 0 {
		_, err = r.db.Exec(
			"INSERT INTO tag (name, user_id) VALUES ($1, $2);",
			"tagdrasil",
			user.ID,
		)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

func (r *TagPostgresRepository) DeleteUserByID(userID int) error {
	return nil
}

// func (r *TagPostgresRepository) DeleteUser(user models.User, userID int) (int, error) {
// 	return 1, nil
// }
