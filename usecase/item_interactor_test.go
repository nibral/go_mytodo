package usecase

import (
	"os"
	"testing"
	"mytodo/interface/database"
	"mytodo/domain"
)

func TestItemInteractor_CRUD_Sqlite3(t *testing.T) {
	const dbFileName = "test.db"
	itemInteractor := &ItemInteractor{
		ItemRepository: database.NewSqlite3ItemRepository(dbFileName),
	}

	// 追加
	createItem := &domain.Item{
		Title:       "タイトル1",
		Description: "説明文1",
	}
	item, err := itemInteractor.AddItem(createItem)
	if err != nil {
		t.Error(err)
	}
	if item.Title != createItem.Title {
		t.Errorf("追加:タイトルが一致しない")
	}
	if item.Description != createItem.Description {
		t.Errorf("追加:説明文が一致しない")
	}

	// 更新
	updateItem := &domain.Item{
		Title:       "タイトル1更新",
		Description: "説明文1更新",
	}
	item, err = itemInteractor.UpdateItemById(item.ID, updateItem)
	if err != nil {
		t.Error(err)
	}
	if item.Title != updateItem.Title {
		t.Errorf("更新:タイトルが一致しない")
	}
	if item.Description != updateItem.Description {
		t.Errorf("更新:説明文が一致しない")
	}

	// もう1件追加して全件取得
	createItem2 := &domain.Item{
		Title:       "タイトル2",
		Description: "説明文2",
	}
	item2, err := itemInteractor.AddItem(createItem2)
	if err != nil {
		t.Error(err)
	}
	items, err := itemInteractor.Items()
	if err != nil {
		t.Error(err)
	}
	if len(items) != 2 {
		t.Errorf("全件取得:データ件数が一致しない")
	}

	// 削除
	itemInteractor.DeleteItemById(item.ID)
	itemInteractor.DeleteItemById(item2.ID)
	items, err = itemInteractor.Items()
	if err != nil {
		t.Error(err)
	}
	if len(items) != 0 {
		t.Errorf("削除:データ件数が一致しない")
	}

	// テスト用DBを削除
	_, err = os.Stat(dbFileName)
	if err == nil {
		os.Remove(dbFileName)
	}
}
