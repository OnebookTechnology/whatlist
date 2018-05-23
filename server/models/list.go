package models

type List struct {
	ListID           uint64  `json:"list_id,omitempty"`             // 书单ID
	ListName         string  `json:"list_name,omitempty"`           // 书单名称
	ListAuthor       string  `json:"list_author,omitempty"`         // 书单作者
	ListCategoryID   uint64  `json:"list_category_id,omitempty"`    // 书单分类ID
	ListCategoryName string  `json:"list_category_name,omitempty"`  // 书单分类名称
	ListIntro        string  `json:"list_intro,omitempty"`          // 书单简介
	ListBriefIntro   string  `json:"list_brief_intro,omitempty"`              // 书单一句话简介
	ListImg          string  `json:"list_img,omitempty"`            // 书单图片
	ListCreateTime   string  `json:"list_create_time,omitempty"`    // 书单创建时间
	ListLastEditTime string  `json:"list_last_edit_time,omitempty"` // 书单最后编辑时间
	ListClickCount   uint64  `json:"list_click_count,omitempty"`    // 书单点击量
	ListBooks        []*Book `json:"list_books,omitempty"`          // 书单包含的图书
}
