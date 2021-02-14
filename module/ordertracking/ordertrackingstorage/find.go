package ordertrackingstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
)

func (s *sqlStore) FindOrderTracking(
	ctx context.Context, 
	conditions map[string]interface{}, 
	moreInfo ...string) (*ordertrackingmodel.OrderTracking, error) {

	db := s.db.Table(ordertrackingmodel.OrderTracking{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var ordertracking ordertrackingmodel.OrderTracking

	if err := db.Where(conditions).First(&ordertracking).Error; err != nil {
		return nil, err
	}

	return &ordertracking, nil
}
