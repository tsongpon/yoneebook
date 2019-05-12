package repository

import (
	"database/sql"

	"github.com/tsongpon/yoneebook/model"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) *MysqlRepository {
	repo := new(MysqlRepository)
	repo.db = db
	return repo
}

func (repo *MysqlRepository) GetStories() ([]*model.Story, error) {
	stories := []*model.Story{}
	sql := "SELECT id, title, content, author FROM story"
	result, err := repo.db.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		s := new(model.Story)
		err := result.Scan(&s.ID, &s.Title, &s.Content, &s.Author)
		if err != nil {
			panic(err.Error())
		}
		stories = append(stories, s)
	}

	return stories, nil
}

// GetStory return story by given ID
func (repo *MysqlRepository) GetStory(id string) (*model.Story, error) {
	sql := "SELECT id, title, content, author FROM story WHERE id = ?"
	var s model.Story
	err := repo.db.QueryRow(sql, id).Scan(&s.ID, &s.Title, &s.Content, &s.Author)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return &s, nil
}

func (repo *MysqlRepository) SaveStory(story *model.Story) (*model.Story, error) {
	sql := "INSERT INTO story (id, title, content, author) VALUES (?, ?, ?, ?)"
	stmt, err := repo.db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	stmt.Exec(story.ID, story.Title, story.Content, story.Author)

	return story, nil
}