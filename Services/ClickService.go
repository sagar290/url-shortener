package Services

import (
	models "url-shortener/Model"
)

func AddClick(url_id int, refer_type string) {
	var click models.Click
	models.Db.Where(models.Click{Url_id: url_id, Referer: refer_type}).Assign(models.Click{Count: 0}).FirstOrCreate(&click)

	count := click.Count + 1
	// fmt.Println(click.Click_id)
	// fmt.Println(count)
	models.Db.Model(&models.Click{}).Where(&models.Click{Click_id: click.Click_id}).Update("count", count)
}
