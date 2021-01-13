package userpaymentmethodstorage

import "github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"

func (s *sqlStore) ListUserPaymentMethod() ([]userpaymentmethodmodel.UserPaymentMethod, error) {
	db := s.db
	var userpaymentmethods []userpaymentmethodmodel.UserPaymentMethod

	if err := db.Find(&userpaymentmethods).Error; err != nil {
		return nil, err
	}

	return userpaymentmethods, nil
}
