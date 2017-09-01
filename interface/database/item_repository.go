package database

import (
	"fmt"
	"errors"
	"log"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"mytodo/domain"
	"mytodo/interface/config"
)

type ItemRepository struct {
	dbSession *dbr.Session
}

func NewItemRepository(config config.Database) *ItemRepository {
	// DBファイルを開いてセッション作成
	// 失敗した場合は強制終了
	conn, err := dbr.Open(config.Engine, config.Source, nil)
	if err != nil {
		log.Fatalln(err)
	}
	sess := conn.NewSession(nil)

	// テーブルが存在しない場合は作成
	// 失敗した場合は強制終了
	createStmt := `CREATE TABLE IF NOT EXISTS items (
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT,
		description TEXT
	);`
	_, err = sess.Exec(createStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createStmt)
	}

	return &ItemRepository{
		dbSession: sess,
	}
}

// 名前と説明を指定してアイテムを追加
func (repo *ItemRepository) Add(item *domain.Item) (int, error) {
	// DBにアイテムを追加
	insertResult, err := repo.dbSession.InsertInto("items").
		Columns("title", "description").
		Values(item.Title, item.Description).
		Exec()
	if err != nil {
		log.Printf("%q", err)
		return -1, err
	}

	// 追加したアイテムのIDを返却
	id, _ := insertResult.LastInsertId()
	return int(id), nil
}

// 指定したIDのアイテムを取得
func (repo *ItemRepository) Get(id int) (*domain.Item, error) {
	// IDを指定してアイテムを取得
	var item domain.Item
	num, err := repo.dbSession.Select("*").From("items").Where("id = ?", id).Load(&item)
	if err != nil {
		log.Printf("%q\n", err)
	}

	// 存在しない場合はエラー
	if num == 0 {
		return nil, errors.New(fmt.Sprintf("item not found (id:%d)", id))
	}

	return &item, nil
}

// 全てのアイテムを取得
func (repo *ItemRepository) GetAll() ([]*domain.Item, error) {
	// 条件指定なしでアイテムを取得
	var items []*domain.Item
	repo.dbSession.Select("*").From("items").Load(&items)

	return items, nil
}

// 指定したIDのアイテムを更新
func (repo *ItemRepository) Update(id int, item *domain.Item) error {
	// IDを指定してアイテムを更新
	_, err := repo.dbSession.Update("items").
		Set("title", item.Title).
		Set("description", item.Description).
		Where("id = ?", id).
		Exec()
	if err != nil {
		log.Printf("%q\n", err)
		return err
	}

	return nil
}

// 指定したIDのアイテムを削除
func (repo *ItemRepository) Delete(id int) error {
	// IDを指定してアイテムを削除
	_, err := repo.dbSession.DeleteFrom("items").Where("id = ?", id).Exec()
	if err != nil {
		log.Printf("%q\n", err)
		return err
	}
	return nil
}
