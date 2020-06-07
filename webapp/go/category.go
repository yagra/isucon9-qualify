package main

import (
	"errors"
	"github.com/jmoiron/sqlx"
)

var categories map[int]Category = map[int]Category{
	1:  {1, 0, "ソファー", nil},
	2:  {2, 1, "一人掛けソファー", "ソファー"},
	3:  {3, 1, "二人掛けソファー", "ソファー"},
	4:  {4, 1, "コーナーソファー", "ソファー"},
	5:  {5, 1, "二段ソファー", "ソファー"},
	6:  {6, 1, "ソファーベッド", "ソファー"},
	10: {10, 0, "家庭用チェア", nil},
	11: {11, 10, "スツール", "家庭用チェア"},
	12: {12, 10, "クッションスツール", "家庭用チェア"},
	13: {13, 10, "ダイニングチェア", "家庭用チェア"},
	14: {14, 10, "リビングチェア", "家庭用チェア"},
	15: {15, 10, "カウンターチェア", "家庭用チェア"},
	20: {20, 0, "キッズチェア", nil},
	21: {21, 20, "学習チェア", "キッズチェア"},
	22: {22, 20, "ベビーソファ", "キッズチェア"},
	23: {23, 20, "キッズハイチェア", "キッズチェア"},
	24: {24, 20, "テーブルチェア", "キッズチェア"},
	30: {30, 0, "オフィスチェア", nil},
	31: {31, 30, "デスクチェア", "オフィスチェア"},
	32: {32, 30, "ビジネスチェア", "オフィスチェア"},
	33: {33, 30, "回転チェア", "オフィスチェア"},
	34: {34, 30, "リクライニングチェア", "オフィスチェア"},
	35: {35, 30, "投擲用椅子", "オフィスチェア"},
	40: {40, 0, "折りたたみ椅子", nil},
	41: {41, 40, "パイプ椅子", "折りたたみ椅子"},
	42: {42, 40, "木製折りたたみ椅子", "折りたたみ椅子"},
	43: {43, 40, "キッチンチェア", "折りたたみ椅子"},
	44: {44, 40, "アウトドアチェア", "折りたたみ椅子"},
	45: {45, 40, "作業椅子", "折りたたみ椅子"},
	50: {50, 0, "ベンチ", nil},
	51: {51, 50, "一人掛けベンチ", "ベンチ"},
	52: {52, 50, "二人掛けベンチ", "ベンチ"},
	53: {53, 50, "アウトドア用ベンチ", "ベンチ"},
	54: {54, 50, "収納付きベンチ", "ベンチ"},
	55: {55, 50, "背もたれ付きベンチ", "ベンチ"},
	56: {56, 50, "ベンチマーク", "ベンチ"},
	60: {60, 0, "座椅子", nil},
	61: {61, 60, "和風座椅子", "座椅子"},
	62: {62, 60, "高座椅子", "座椅子"},
	63: {63, 60, "ゲーミング座椅子", "座椅子"},
	64: {64, 60, "ロッキングチェア", "座椅子"},
	65: {65, 60, "座布団", "座椅子"},
	66: {66, 60, "空気椅子", "座椅子"},
}

type Category struct {
	ID                 int    `json:"id" db:"id"`
	ParentID           int    `json:"parent_id" db:"parent_id"`
	CategoryName       string `json:"category_name" db:"category_name"`
	ParentCategoryName string `json:"parent_category_name,omitempty" db:"-"`
}

// エラーメッセージは適当です
// もしかしたらsqlxのエラーメッセージと同じものを返す必要あり?
func getCategoryByID(q sqlx.Queryer, categoryID int) (category Category, err error) {
	targetCategory, ok1 := categories[categoryID]
	if !ok1 {
		return Category{}, errors.New("not found")
	}
	return targetCategory, nil
}
