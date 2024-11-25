package types

type Poem struct{
	ID		uint64 `json:"id"`
	Title   string `json:"title"` 
	Pages   uint64 `json:"pages"` 
	Genre	string `json:"genre"`
}