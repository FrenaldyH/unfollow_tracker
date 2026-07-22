package models

type Relation struct {
	Title          string           `json:"title"`
	StringListData []StringListData `json:"string_list_data"`
}
