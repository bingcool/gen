// Code generated by github.com/bingcool/gen. DO NOT EDIT.
// Code generated by github.com/bingcool/gen. DO NOT EDIT.
// Code generated by github.com/bingcool/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/bingcool/gen"
	"github.com/bingcool/gen/field"
	"github.com/bingcool/gen/tests/.gen/dal_4/model"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.CreditCard{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.CreditCard{}) fail: %s", err)
	}
}

func Test_creditCardQuery(t *testing.T) {
	creditCard := newCreditCard(_gen_test_db)
	creditCard = *creditCard.As(creditCard.TableName())
	_do := creditCard.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(creditCard.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <credit_cards> fail:", err)
		return
	}

	_, ok := creditCard.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from creditCard success")
	}

	err = _do.Create(&model.CreditCard{})
	if err != nil {
		t.Error("create item in table <credit_cards> fail:", err)
	}

	err = _do.Save(&model.CreditCard{})
	if err != nil {
		t.Error("create item in table <credit_cards> fail:", err)
	}

	err = _do.CreateInBatches([]*model.CreditCard{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <credit_cards> fail:", err)
	}

	_, err = _do.Select(creditCard.ALL).Take()
	if err != nil {
		t.Error("Take() on table <credit_cards> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <credit_cards> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <credit_cards> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <credit_cards> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.CreditCard{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <credit_cards> fail:", err)
	}

	_, err = _do.Select(creditCard.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <credit_cards> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <credit_cards> fail:", err)
	}

	_, err = _do.Select(creditCard.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <credit_cards> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <credit_cards> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <credit_cards> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <credit_cards> fail:", err)
	}

	_, err = _do.ScanByPage(&model.CreditCard{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <credit_cards> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <credit_cards> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <credit_cards> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <credit_cards> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <credit_cards> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <credit_cards> fail:", err)
	}
}
