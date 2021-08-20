package Services

import (
	models "url-shortener/Model"
)

func AddClick(url_id int, refer_type string) {
	var click models.Click
	var url models.Url
	models.Db.Where(models.Click{Url_id: url_id, Referer: refer_type}).Assign(models.Click{Count: 0}).FirstOrCreate(&click)
	models.Db.Where(models.Url{Url_id: uint(url_id)}).FirstOrCreate(&url)

	count := click.Count + 1
	total_count := url.TotalCount + 1
	// fmt.Println(click.Click_id)
	// fmt.Println(count)
	models.Db.Model(&models.Click{}).Where(&models.Click{Click_id: click.Click_id}).Update("count", count)
	models.Db.Model(&models.Url{}).Where(&models.Url{Url_id: url.Url_id}).Update("total_count", total_count)
}
