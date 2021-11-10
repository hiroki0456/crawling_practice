package repository

import (
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"mf.crawling/config"
)

func TestInsert(t *testing.T) {
	sqlMock := config.InitDBMock()
	defer sqlMock.Db.Close()

	cases := map[string]struct {
		obj Repository
	}{
		"Users": {
			obj: &UserRepository{UserId: "test", CrawlingDate: time.Now()},
		},
		"Banks": {
			obj: &BankRepository{Name: "test銀行", UpdatedAt: time.Now(), UserId: 1},
		},
		"Details": {
			obj: &DetailRepository{BankName: "test銀行", TradingContent: "売り上げ", Amount: 100000, UpdatedAt: time.Now(), UserId: 1, BankId: 1},
		},
	}

	for name, v := range cases {
		t.Run(name, func(t *testing.T) {
			sqlMock.Mock.ExpectPrepare(fmt.Sprintf(`INSERT INTO %s`, name))
			sqlMock.Mock.ExpectExec(fmt.Sprintf(`INSERT INTO %s`, name)).WillReturnResult(sqlmock.NewResult(1, 1))

			id, err := Insert(sqlMock.Db, v.obj)
			fmt.Println(id)
			if err != nil {
				t.Errorf("failed to create: %s", err)
			}
			if err := sqlMock.Mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
