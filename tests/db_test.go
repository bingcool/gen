package tests_test

import (
	"fmt"
	"github.com/bingcool/gen/tests/.gen/dal_1/model"
	"github.com/bingcool/gen/tests/.gen/dal_1/query"
	"testing"
)

func TestQuery1(t *testing.T) {
	tableName := query.Use(DB).TblUser.TableName()
	fmt.Println(tableName)
}

func TestQuery2(t *testing.T) {
	query.SetDefault(DB)
	res, _ := query.TblUser.WithContext(ctx).First()
	fmt.Println(res.Phone)
	tableName := query.TblUser.TableName()
	fmt.Println(tableName)
}

func TestQuery3(t *testing.T) {
	query.SetDefault(DB)
	res, _ := query.TblUser.WithContext(ctx).Limit(5).Find()
	fmt.Println(res)
}

func TestQuery4(t *testing.T) {
	query.SetDefault(DB)
	res, _ := query.TblUser.
		WithContext(ctx).
		Where(
			query.TblUser.UserID.Gt(1),
			query.TblUser.UserID.Lt(100),
		).
		Count()
	fmt.Println(res)
}

func TestQuery5(t *testing.T) {
	query.SetDefault(DB)
	res, _ := query.TblUser.
		WithContext(ctx).
		Where(
			query.TblUser.UserID.Gt(1),
			query.TblUser.UserID.Lt(100),
		).
		Join(query.TblOrder, query.TblOrder.UserID.EqCol(query.TblUser.UserID)).
		Count()
	fmt.Println(res)
}

func TestQuery6(t *testing.T) {
	query.SetDefault(DB)
	res, _ := query.TblUser.
		WithContext(ctx).
		Where(
			query.TblUser.UserID.Gt(1),
			query.TblUser.UserID.Lt(100),
		).
		Join(query.TblOrder, query.TblOrder.UserID.EqCol(query.TblUser.UserID)).
		Count()
	fmt.Println(res)
}

func TestUpdate1(t *testing.T) {
	query.SetDefault(DB)
	user, _ := query.TblUser.WithContext(ctx).First()
	fmt.Println("userId:", user.UserID, "phone:", user.Phone)

	newPhone := "12345678901"
	result := DB.Model(&model.TblUser{}).Where("user_id = ?", user.UserID).Updates(map[string]interface{}{
		"phone": newPhone,
	})
	affected := result.RowsAffected
	fmt.Println("DB.Updates()更新行数:", affected)

	user.Phone = "12345678902"
	result = DB.Save(user)

	affected = result.RowsAffected
	fmt.Println("DB.Save()更新行数:", affected)

	user, _ = query.TblUser.WithContext(ctx).Where(query.TblUser.UserID.Eq(user.UserID)).First()
	fmt.Println(user.Phone)
}

func TestUpdate2(t *testing.T) {
	query.SetDefault(DB)
	user, _ := query.TblUser.WithContext(ctx).First()
	fmt.Println("userId:", user.UserID, "phone:", user.Phone)

	newPhone := "12345678904"
	updates, err := query.TblUser.WithContext(ctx).Where(query.TblUser.UserID.Eq(user.UserID)).Updates(map[string]interface{}{
		"phone": newPhone,
	})

	if err != nil {
		return
	}

	fmt.Println("更新行数:", updates.RowsAffected)

	user, _ = query.TblUser.WithContext(ctx).Where(query.TblUser.UserID.Eq(user.UserID)).First()
	fmt.Println(user.Phone)
}
