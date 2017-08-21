package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"mytodo/domain"
)

const dbDriverName = "sqlite3"

type Sqlite3ItemRepository struct {
	dbDriverName string
	dbFileName   string
}

func NewSqlite3ItemRepository(dbFileName string) *Sqlite3ItemRepository {
	// DBファイル作成
	// 失敗した場合は強制終了
	db, err := sql.Open(dbDriverName, dbFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// テーブルが存在しない場合は作成
	// 失敗した場合は強制終了
	createStmt := `CREATE TABLE IF NOT EXISTS items (
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT,
		description TEXT
	);`
	_, err = db.Exec(createStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, createStmt)
	}

	return &Sqlite3ItemRepository{
		dbDriverName: dbDriverName,
		dbFileName:   dbFileName,
	}
}

// 名前と説明を指定してアイテムを追加
func (repo *Sqlite3ItemRepository) Add(item *domain.Item) (int, error) {
	// DBに接続
	db, err := sql.Open(repo.dbDriverName, repo.dbFileName)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer db.Close()

	// DBにアイテムを追加
	insertStmt := `INSERT INTO items (title, description) VALUES (?, ?);`
	insertResult, err := db.Exec(insertStmt, item.Title, item.Description)
	if err != nil {
		log.Printf("%q: %s\n", err, insertStmt)
		return -1, err
	}

	// 追加したアイテムのIDを返却
	id, _ := insertResult.LastInsertId()
	return int(id), nil
}

// 指定したIDのアイテムを取得
func (repo *Sqlite3ItemRepository) Get(id int) (*domain.Item, error) {
	// DBに接続
	db, err := sql.Open(repo.dbDriverName, repo.dbFileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	// IDを指定してアイテムを取得
	selectStmt := `SELECT * FROM items WHERE id = ?;`
	rows, err := db.Query(selectStmt, id)
	if err != nil {
		log.Printf("%q: %s\n", err, selectStmt)
		return nil, err
	}
	defer rows.Close()

	// アイテムが存在する場合は返却
	for rows.Next() {
		var id int
		var title string
		var description string
		err = rows.Scan(&id, &title, &description)
		if err != nil {
			log.Println(err)
		}
		item := domain.Item{
			ID:          id,
			Title:       title,
			Description: description,
		}
		return &item, nil
	}

	// 存在しない場合はエラー
	return nil, errors.New(fmt.Sprintf("item not found (id:%d)", id))
}

// 全てのアイテムを取得
func (repo *Sqlite3ItemRepository) GetAll() ([]*domain.Item, error) {
	// DBに接続
	db, err := sql.Open(repo.dbDriverName, repo.dbFileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer db.Close()

	// 条件指定なしでアイテムを取得
	selectStmt := `SELECT * FROM items;`
	rows, err := db.Query(selectStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, selectStmt)
		return nil, err
	}
	defer rows.Close()

	// アイテムを配列にして返却
	items := []*domain.Item{}
	for rows.Next() {
		var id int
		var title string
		var description string
		err = rows.Scan(&id, &title, &description)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, &domain.Item{
			ID:          id,
			Title:       title,
			Description: description,
		})
	}
	return items, nil
}

// 指定したIDのアイテムを更新
func (repo *Sqlite3ItemRepository) Update(id int, item *domain.Item) error {
	// DBに接続
	db, err := sql.Open(repo.dbDriverName, repo.dbFileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	// IDを指定してアイテムを更新
	updateStmt := `UPDATE items	SET title = ?, description = ? WHERE id = ?;`
	_, err = db.Exec(updateStmt, item.Title, item.Description, id)
	if err != nil {
		log.Printf("%q: %s\n", err, updateStmt)
		return err
	}
	return nil
}

// 指定したIDのアイテムを削除
func (repo *Sqlite3ItemRepository) Delete(id int) error {
	// DBに接続
	db, err := sql.Open(repo.dbDriverName, repo.dbFileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	// IDを指定してアイテムを削除
	deleteStmt := `DELETE FROM items WHERE id = ?;`
	_, err = db.Exec(deleteStmt, id)
	if err != nil {
		log.Printf("%q: %s\n", err, deleteStmt)
		return err
	}
	return nil
}
