package app

import "github.com/mattermost/focalboard/server/model"

func (a *App) GetUserCategoryBlocks(userID, teamID string) ([]model.CategoryBlocks, error) {
	return a.store.GetUserCategoryBlocks(userID, teamID)
}

func (a *App) AddUpdateUserCategoryBlock(userID, categoryID, blockID string) error {
	return a.store.AddUpdateCategoryBlock(userID, categoryID, blockID)
}