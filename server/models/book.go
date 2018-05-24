package models

type Book struct {
	ISBN       uint64 `json:"isbn,omitempty"`
	BookName   string `json:"book_name,omitempty"`
	AuthorName string `json:"author_name,omitempty"`
	// 出版社
	Press           string `json:"press,omitempty"`
	PublicationTime string `json:"publication_time,omitempty"`
	// 印刷时间
	PrintTime string `json:"print_time,omitempty"`
	// 版次
	Edition   uint8   `json:"edition,omitempty"`
	BookPrice float64 `json:"price,omitempty"`
	// 开本
	Format string `json:"format,omitempty"`
	// 纸张
	Paper string `json:"paper,omitempty"`
	// 包装
	Pack string `json:"pack,omitempty"`
	// 是否套装
	Suit uint8 `json:"suit,omitempty"`
	// 目录
	TableOfContent string `json:"table_of_content,omitempty"`
	// 图书一句话简介
	BookBriefIntro string `json:"book_brief_intro,omitempty"`
	// 作者信息
	AuthorIntro string `json:"author_intro,omitempty"`
	// 图书内容简介
	ContentIntro string `json:"content_intro,omitempty"`
	// 编辑推荐
	EditorRecommend string `json:"editor_recommend,omitempty"`
	// 一级分类
	FirstClassification uint16 `json:"first_classification,omitempty"`
	// 二级分类
	SecondClassification uint16 `json:"second_classification,omitempty"`
	// 书的总分
	TotalScore uint32 `json:"total_score,omitempty"`
	// 书的评论次数
	CommentTimes uint32 `json:"comment_times,omitempty"`

	BookIcon   string  `json:"book_icon,omitempty"`
	BookPic    string  `json:"book_pic,omitempty"`
	BookDetail string  `json:"book_detail,omitempty"`
	Discount   float64 `json:"discount,omitempty,omitempty"`
	RealPrice  float64 `json:"real_price,omitempty,omitempty"`
	Category   int     `json:"category,omitempty"`
	Field1     []int   `json:"field_1,omitempty"` //年龄id范围
	Field2     int     `json:"field_2,omitempty"` //性别
	Field3     []int   `json:"field_3,omitempty"` //婚姻状况id
	Field4     []int   `json:"field_4,omitempty"` //教育程度
	Field5     []int   `json:"field_5,omitempty"` //收入id范围
	Field6     []int   `json:"field_6,omitempty"` //工作行业id
	Field7     float64 `json:"field_7,omitempty"` //身高体重比例

	IsInterested bool `json:"is_interested"` // 是否标记为喜欢，0不喜欢，1喜欢
}

type ExhibitionBook struct {
	ISBN           uint64  `json:"isbn"`
	ExhibitionID   uint8   `json:"exhibition_id"`
	BookName       string  `json:"book_name"`
	BookBriefIntro string  `json:"book_brief_intro"`
	BookDiscount   float64 `json:"book_discount"`
	Price          float64 `json:"price"`
	CommentTimes   uint32  `json:"comment_times"`
	TotalScore     uint32  `json:"total_score"`
	BookIcon       string  `json:"book_icon"`
	OffTime        string  `json:"off_time"`
	RealPrice      float64 `json:"real_price"`
}
