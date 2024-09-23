package restaurantstorage

import (
	"context"

	restaurantmodel "social/module/post/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result *restaurantmodel.Restaurant

	db := s.db.Where("status in (1)")

	if f := filter; f != nil {
		if len(f.OwnerId) > 0 {
			db = db.Where("owner_id=?", f.OwnerId)
		}
	}

	if err := db.First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
