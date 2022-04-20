package repositories

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/vale-app/vale-common/models"
)

type TaxonomiesService struct {
	client *resty.Client
}

func NewTaxonomiesService(client *resty.Client) *TaxonomiesService {
	return &TaxonomiesService{client: client}
}

func (s *TaxonomiesService) GetCategoryByID(categoryID uint) (*models.Category, error) {
	res, err := s.client.R().SetResult(&models.Category{}).Get(fmt.Sprintf("/taxonomies/categories/%d", categoryID))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("%s", res.Status())
	}

	c := res.Result().(*models.Category)

	return c, nil
}

func (s *TaxonomiesService) GetSubCategoryByID(subCategoryID uint) (*models.SubCategory, error) {
	res, err := s.client.R().SetResult(&models.SubCategory{}).Get(fmt.Sprintf("/taxonomies/subcategories/%d", subCategoryID))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, fmt.Errorf("%s", res.Status())
	}

	sc := res.Result().(*models.SubCategory)

	return sc, nil
}
