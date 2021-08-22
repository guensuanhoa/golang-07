package common

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total"`
}

func (paging *Paging) Process() error {
	if paging.Page < 1 {
		paging.Page = 1
	}

	if paging.Limit <= 0 && paging.Limit > 100 {
		paging.Limit = 10
	}

	return nil
}
