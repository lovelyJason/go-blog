package dao

import (
	"go-blog/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select * from goblog_category")
	if err != nil {
		log.Printf("query failed, err:%v\n", err)
		return nil, err
	}
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Printf("scan failed, err:%v\n", err)
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
