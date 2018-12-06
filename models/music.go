package models

import (
	"fmt"
)

// Music -
type Music struct {
	Id      int64  `json:"id"`
	Name    string `xorm:"UNIQUE NOT NULL" json:"name"`
	Picture string `json:"picture"`
	Url     string `json:"url"`
	Source  string `json:"source"`
	Created int64  `xorm:"created"`
	Updated int64  `xorm:"updated"`
}

// CreateMusic -
func CreateMusic(name, picture, url, source string) (*Music, error) {
	_, has, err := ExistMusicByName(name)
	if err != nil {
		return nil, err
	}
	if has {
		return nil, fmt.Errorf("CreateMusic has the music already")
	}
	music := Music{Name: name, Picture: picture, Url: url, Source: source}
	err = AddMusic(&music)
	if err != nil {
		return nil, err
	}
	return &music, nil
}

// GetMusics -
func GetMusics(page, pageSize int) ([]*Music, error) {
	musics := make([]*Music, 0, pageSize)
	return musics, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&musics)
}

// AddMusic -
func AddMusic(m *Music) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(m); err != nil {
		return err
	}

	return sess.Commit()
}

// GetMusicByID -
func GetMusicByID(id int64) (*Music, error) {
	m := new(Music)
	has, err := x.Id(id).Get(m)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("music does not exist [id: %d]", id)
	}
	return m, nil
}

// ExistProjectByName -
func ExistMusicByName(name string) (*Music, bool, error) {
	if len(name) == 0 {
		return nil, false, fmt.Errorf("ExistMusicByName'param name is empty")
	}
	m := &Music{Name: name}
	has, err := x.Get(m)
	return m, has, err
}
