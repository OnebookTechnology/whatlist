package server

import (
	"github.com/OnebookTechnology/whatlist/server/models"
	"sort"
	"sync"
	"fmt"
)

var UserMap sync.Map // map[string]*model.User
var BookList map[int][]*models.Book
var UserSuitMap sync.Map     //map[user_id][]*models.Book
var UserUnSuit30Map sync.Map //map[user_id][]*models.Book
var UserUnSuit10Map sync.Map //map[user_id][]*models.Book

type ListResult struct {
	List   []*models.Book
	Weight []int
}

func LoadAllBooks() error {
	var err error
	BookList, err = server.DB.FindAllBooks()
	if err != nil {
		return err
	}
	return nil
}

func doRecommend(user *models.User) {
	if !IsNeedUpdateRecommend(user) {
		return
	} else {
		mySuitList := new(ListResult)
		myUnSuitList := new(ListResult)

		var wg sync.WaitGroup

		//区分匹配喜好的books和不匹配的books
		for _, tagId := range CategoryArray {
			bookList, ok := BookList[tagId]
			if !ok {
				continue
			}
			if sliceIntContains(user.Hobby, tagId) {
				mySuitList.List = append(mySuitList.List, bookList...)
			} else {
				myUnSuitList.List = append(myUnSuitList.List, bookList...)
			}

		}

		wg.Add(1)
		go SuitRecommend(user, mySuitList, wg)
		wg.Add(1)
		go UnSuitRecommend(user, myUnSuitList, wg)
		wg.Wait()
		fmt.Println("wait ok")
	}

	user.NeedUpdateRecommend = false
	UserMap.Store(user.UserId, user)
	return
}

func SuitRecommend(user *models.User, mySuitList *ListResult, wg sync.WaitGroup) {
	for i := range mySuitList.List {
		mySuitList.Weight[i] = calculateWeightOfBook(user, mySuitList.List[i])
	}
	//降序排序
	sort.Sort(mySuitList)
	UserSuitMap.Store(user.UserId, mySuitList)
	fmt.Println("s done")
	wg.Done()
}

func UnSuitRecommend(user *models.User, myUnSuitList *ListResult, wg sync.WaitGroup) {
	//30% recommend and 10% recommend
	myUnSuit30List := new(ListResult)
	myUnSuit10List := new(ListResult)
	for i := range myUnSuitList.List {
		weight := calculateWeightOfBook(user, myUnSuitList.List[i])
		if weight >= 10 {
			myUnSuit30List.List = append(myUnSuit30List.List, myUnSuitList.List[i])
			myUnSuit30List.Weight = append(myUnSuit30List.Weight, weight)
		} else if weight >= 5 {
			myUnSuit10List.List = append(myUnSuit10List.List, myUnSuitList.List[i])
			myUnSuit10List.Weight = append(myUnSuit10List.Weight, weight)
		}
	}

	sort.Sort(myUnSuit30List)
	UserSuitMap.Store(user.UserId, myUnSuit30List)

	sort.Sort(myUnSuit10List)
	UserSuitMap.Store(user.UserId, myUnSuit10List)
	fmt.Println("us done")
	wg.Done()
}

func calculateWeightOfBook(user *models.User, book *models.Book) (weight int) {
	// 计算权值
	//Field1 年龄id范围
	if user.Field1 == AgeAny {
		weight += 1
	} else if sliceIntContains(book.Field1, user.Field1) {
		weight += 3
	}

	//Field2 性别
	if user.Field2 == SexAny {
		weight += 1
	} else if book.Field2 == user.Field2 {
		weight += 3
	}

	//Field3 婚姻状况id
	if user.Field3 == MarriageAny {
		weight += 1
	} else if sliceIntContains(book.Field3, user.Field3) {
		weight += 3
	}

	//Field4 教育程度
	if user.Field4 == AgeAny {
		weight += 1
	} else if sliceIntContains(book.Field4, user.Field4) {
		weight += 3
	}

	//Field5 最小收入
	if user.Field5 == IncomeAny {
		weight += 1
	} else if sliceIntContains(book.Field5, user.Field5) {
		weight += 3
	}

	//Field6 工作行业id
	if user.Field6 == WorkAny {
		weight += 1
	} else if sliceIntContains(book.Field6, user.Field6) {
		weight += 3
	}

	//Field7 身高体重比例
	if user.Field7 == WeightAny {
		weight += 1
	} else if book.Field7 == user.Field7 {
		weight += 3
	}
	return
}

func IsNeedUpdateRecommend(user *models.User) bool {
	if _, ok := UserSuitMap.Load(user); !ok {
		//未查到记录需要重新推荐
		return true
	} else {
		if user.NeedUpdateRecommend {
			//需要重新推荐
			return true
		} else {

			return false
		}
	}
}

//Data interface of recommend,
func (sl *ListResult) Len() int {
	return len(sl.List)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (sl *ListResult) Less(i, j int) bool {
	if sl.Weight[i] < sl.Weight[j] {
		return false
	}
	return true
}

// Swap swaps the elements with indexes i and j.
func (sl *ListResult) Swap(i, j int) {
	tempBook := sl.List[i]
	tempWeight := sl.Weight[i]
	sl.List[i] = sl.List[j]
	sl.Weight[i] = sl.Weight[j]
	sl.List[j] = tempBook
	sl.Weight[j] = tempWeight
}
