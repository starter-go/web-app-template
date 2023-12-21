package vo

import "github.com/starter-go/web-app-template/app/web/dto"

// Example ... VO
type Example struct {
	Base

	Items []*dto.Example `json:"items"`
}
