package dto


type InterestPointDTO struct {
	Title string `json:"title"`
	Description string `json:"description"`
	ImagePath string `json:"image_path"`
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type TourDTO struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Difficulty string `json:"difficulty"`
	Tags []string `json:"tags"`
	Status string `json:"status"`
	Price float32 `json:"price"`
	AuthorId string `json:"author_id"`
	StartPoint InterestPointDTO `json:"start_point"`
	EndPoint InterestPointDTO `json:"end_point"`
}
