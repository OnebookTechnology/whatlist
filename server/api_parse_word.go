package server

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"io"
	"strconv"
)

func ParseWord(c *gin.Context) {
	crossDomain(c)
	form, _ := c.MultipartForm()
	fileHeader := form.File["file"][0]
	file, err := fileHeader.Open()
	if err != nil {
		logger.Info("error: %s", err)
	}
	params := readWord(file, int(fileHeader.Size))
	// 构造book object
	b := generateBookJson(params)
	bJson, _ := jsoniter.MarshalToString(b)
	sendJsonResponse(c, 0, "%s", bJson)
}

func readWord(reader io.ReaderAt, size int) map[string]string {
	//var params map[string]string
	//params = make(map[string]string)
	////doc, err := document.Open("C:\\Users\\陈曦\\Desktop\\断舍离.docx")
	//doc, err := document.Read(reader, int64(size))
	//if err != nil {
	//	log.Fatalf("error opening document: %s", err)
	//}
	//table := doc.Tables()[0]
	//var key = ""
	//var value = ""
	//// 解析word中的表格
	//for _, row := range table.Rows() {
	//	for cell_index, cell := range row.Cells() {
	//		for _, paras := range cell.Paragraphs() {
	//			// 获取Key值
	//			if cell_index%2 == 0 {
	//				key = paras.Runs()[0].Text()
	//			} else {
	//				for _, run := range paras.Runs() {
	//					value = value + run.Text()
	//				}
	//			}
	//		}
	//		params[key] = value
	//		value = ""
	//	}
	//}
	return nil
}

func generateBookJson(params map[string]string) *models.Book {
	b := new(models.Book)
	for k, v := range params {
		switch k {
		case "书名":
			b.BookName = v
			break
		case "一句话简介":
			b.BookBriefIntro = v
			break
		case "作者":
			b.AuthorName = v
			break
		case "出版社":
			b.Press = v
			break
		case "出版时间":
			b.PublicationTime = v
			break
		case "原价":
			price, _ := strconv.ParseFloat(v, 10)
			b.BookPrice = price
			break
		case "版次":
			edition, _ := strconv.Atoi(v)
			b.Edition = uint8(edition)
			break
		case "印刷时间":
			b.PrintTime = v
			break
		case "开本":
			b.Format = v
			break
		case "纸张":
			b.Paper = v
			break
		case "包装":
			b.Pack = v
			break
		case "套装":
			if v == "否" {
				b.Suit = 0
			} else {
				b.Suit = 1
			}
			break
		case "ISBN":
			isbn, _ := strconv.ParseUint(v, 10, 64)
			b.ISBN = isbn
			break
		case "编辑推荐":
			b.EditorRecommend = v
			break
		case "内容简介":
			b.ContentIntro = v
			break
		case "作者简介":
			b.AuthorIntro = v
			break
		}
	}
	return b
}
