package tagdrasil

type Tag struct {
	ID int64
	Name string
}

type TagTree struct {
	Tag
	ChildTags []*Tag
}

type User struct {
	ID           int
	FirstName    string
	LastName     string      // optional
	UserName     string       // optional
	LanguageCode string  // optional
}
