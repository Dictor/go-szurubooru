package model

import (
	//"github.com/jinzhu/gorm"
	"time"
)

//https://github.com/rr-/szurubooru/blob/master/doc/API.md
//Implement version in only "micro" component

type MicroUser struct {
	Version           int
	Password          string // not included in resource definition
	Rank              string
	LastLoginTime     time.Time
	CreationTime      time.Time
	AvartarStyle      string
	CommentCount      int
	UploadedPostCount int
	LikedPostCount    int
	DislikedPostCount int
	FavoritePostCount int
}

type User struct {
	Name       string `gorm:"PRIMARY_KEY;UNIQUE"`
	AvartarUrl string
	User       MicroUser
}

type UserToken struct {
	Version        int
	User           MicroUser
	Token          string
	Note           string
	Enabled        bool
	ExpirationTime time.Time
	CreationTime   time.Time
	LastEditTime   time.Time
	LastUsageTime  time.Time
}

type TagCategory struct {
	Version int
	Name    string `gorm:"PRIMARY_KEY;UNIQUE"`
	Color   string
	Usages  int
	Default bool
}

type MicroTag struct {
	Version      int
	CreationTime time.Time
	LastEditTime time.Time
	Description  string
}

type Tag struct {
	Name     string
	Category string
	Usages   int
	Tag      MicroTag
}

type MicroPost struct {
	Version            int
	Id                 string
	CreationTime       time.Time
	LastEditTime       time.Time
	Safety             string
	Source             string
	Type               string
	Checksum           string
	CanvasWidth        int
	CanvasHeight       int
	ContentUrl         string
	Flags              []string
	Tags               []MicroTag
	Relations          []MicroPost
	Notes              []Note
	User               MicroUser
	Score              int
	OwnScore           int
	OwnFavorite        int
	TagCount           int
	FavoriteCount      int
	CommentCount       int
	NoteCount          int
	FeatureCount       int
	RelationCount      int
	LastFeatureTime    time.Time
	FavoriteBy         []MicroUser
	HasCustomThumbnail bool
	MimeType           string
	Comment            []Comment
}

type Post struct {
	ThumbnailUrl string
}

type Note struct {
	Polygon [][]int
	Text    string
}

type Comment struct {
	Version      int
	Id           string
	PostId       string
	User         MicroUser
	Text         string
	CreationTime time.Time
	LastEditTime time.Time
	Score        int
	OwnScore     int
}

type Snapshot struct {
	Operation string
	Type      string
	Id        string
	User      MicroUser
	Data      map[string]interface{}
	Time      time.Time
}

type UnpagedSearchResult struct {
	Results []interface{}
}

type PagedSearchResult struct {
	Query   string
	Offset  string
	Limit   int
	Total   int
	Results []interface{}
}

type ImageSearchResult struct {
	ExactPost    Post
	SimilarPosts []ImageSearchResultSimilar
}

type ImageSearchResultSimilar struct {
	Distance float32
	Post     Post
}
