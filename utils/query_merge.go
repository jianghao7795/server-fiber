package utils

import (
	"fmt"
	"strings"

	json "github.com/bytedance/sonic"

	"gorm.io/gorm"
)

func MergeQuery(db *gorm.DB, query interface{}, args ...string) (err error) {
	searchValue := strings.Join(args, ",")
	var b []byte
	b, err = json.Marshal(query)
	if err != nil {
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return
	}
	for k, v := range m {
		if strings.Contains(searchValue, k) {
			if fmt.Sprintf("%T", v) == "float64" {
				if v.(float64) != 0 {
					db = db.Where(fmt.Sprintf("%s = ?", k), v)
				}
			} else {
				if v != "" {
					if k == "method" {
						db = db.Where(fmt.Sprintf("%s = ?", k), v)
					} else {
						db = db.Where(fmt.Sprintf("%s like ?", k), "%"+v.(string)+"%")
					}

				}

			}

		}
	}
	return
}
