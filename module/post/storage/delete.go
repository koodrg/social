package restaurantstorage

import (
	"context"
	restaurantmodel "social/module/post/model"
)

func (s *sqlStore) Delete(
	context context.Context,
	id string,
) (bool, error) {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (s *sqlStore) SoftDelete(
	context context.Context,
	id string,
) (bool, error) {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return false, err
	}

	return true, nil
}
