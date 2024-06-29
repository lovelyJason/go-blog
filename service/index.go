package service

import (
	"go-blog/config"
	"go-blog/dao"
	"go-blog/models"
	"html/template"
	"log"
)

func GetAllInfo(page int, pageSize int) (*models.HomeData, error) {
	//var categorys = []models.Category{
	//	{
	//		Cid:  1,
	//		Name: "go",
	//	},
	//}
	categories, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	//var posts = []models.PostMore{
	//	{
	//		Pid:          1,
	//		Title:        "go博客",
	//		Content:      "内容",
	//		UserName:     "jasonhuang",
	//		ViewCount:    123,
	//		CreateAt:     "2022-02-20",
	//		CategoryId:   1,
	//		CategoryName: "go",
	//		Type:         0,
	//	},
	//}
	posts, err := dao.GetPostPage(page, pageSize)
	log.Println("view index: posts", posts)
	var postMores []models.PostMore
	for _, post := range posts {
		postMore := models.PostMore{
			Pid:        post.Pid,
			Title:      post.Title,
			Slug:       post.Slug,
			Content:    template.HTML(post.Content), // 注意， 这里结构体定义的是tempalte.HTML格式
			CategoryId: post.CategoryId,
		}
		postMores = append(postMores, postMore)
	}
	var homeData = &models.HomeData{
		Viewer:    config.Cfg.Viewer,
		Categorys: categories,
		//Posts:     posts, // 注意：Cannot use 'posts' (type []models.Post) as the type []PostMore
		Posts:   postMores,
		Total:   1,
		Page:    1,
		Pages:   []int{1},
		PageEnd: true,
	}
	return homeData, nil

}
