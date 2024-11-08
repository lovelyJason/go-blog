package dao

import (
	"go-blog/models"
	"log"
)

func GetPostPage(page int, pageSize int) ([]models.Post, error) {
	rows, err := DB.Query("select * from goblog_post limit ?,?", (page-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		// 注意，这里字段的顺序严格和mysql列字段顺序一致
		err := rows.Scan(&post.Pid, &post.Title, &post.Slug, &post.Content, &post.CategoryId, &post.ViewCount, &post.Type, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println(err)

			return nil, err
		}
		posts = append(posts, post)
	}
	//log.Println(posts)
	return posts, nil
}
