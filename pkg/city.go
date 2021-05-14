package test_api

type City struct {
	NameRu string `json:"nameEn" db:"name_en"`
	NameEn string `json:"nameRu" db:"name_ru"`
}
