package categories

import (
	"fmt"
	"net/url"

	"gopkg.in/go-playground/validator.v9"
)

type Category struct {
	Name              string `validate:"required"` //new category name
	Parent            int    //Default to "0" - the parent category id inside which the new category will be created - set to 0 for a root category
	Idnumber          string //Optional - the new category idnumber
	Description       string //Optional - the new category description
	DescriptionFormat int    //Default to "1" //description format (1 = HTML, 0 = MOODLE, 2 = PLAIN or 4 = MARKDOWN)
	Theme             string //Optional - the new category theme. This option must be enabled on moodle
}

func verifyCategoryDataRequired(category Category) error {
	return validator.New().Struct(category)
}

func ParsecategoryToFormData(indice int, category Category) (string, error) {
	err := verifyCategoryDataRequired(category)
	if err != nil {
		return "", err
	}
	//Required data
	data := fmt.Sprintf("&categories[%d][name]=%s&categories[%d][parent]=%d", indice, url.QueryEscape(category.Name), indice, category.Parent)

	//Optional data
	if category.Idnumber == "" {
		data += fmt.Sprintf("&categories[%d][idnumber]=%s", indice, url.QueryEscape(category.Idnumber))
	}
	if category.Description == "" {
		data += fmt.Sprintf("&categories[%d][description]=%s", indice, url.QueryEscape(category.Description))
	}
	if category.DescriptionFormat != 0 {
		data += fmt.Sprintf("&categories[%d][descriptionformat]=%d", indice, category.DescriptionFormat)
	}
	if category.Theme == "" {
		data += fmt.Sprintf("&categories[%d][theme]=%s", indice, url.QueryEscape(category.Theme))
	}

	return data, nil
}
