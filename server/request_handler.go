package main

import (
	"context"
	"database/sql"
	"time"

	"mf.crawling/pb"
	"mf.crawling/repository"
)

type serviceHandler struct {
	db *sql.DB
}

func initService() *serviceHandler {
	return &serviceHandler{db: repository.Init()}
}

func (s *serviceHandler) CrawlingHandler(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	userId, err := repository.Insert(s.db, &repository.UserRepository{UserId: req.UserId, CrawlingDate: time.Now()})
	if err != nil {
		return
	}

	bankId, err := repository.Insert(s.db, &repository.BankRepository{Name: "みずほ銀行", UpdatedAt: time.Now(), UserId: userId})
	if err != nil {
		return
	}

	_, err = repository.Insert(s.db, &repository.DetailRepository{BankName: "みずほ銀行", TradingContent: "売り上げ", Amount: 200000, UpdatedAt: time.Now(), UserId: userId, BankId: bankId})
	if err != nil {
		return
	}
	defer s.db.Close()
	res = &pb.UserResponse{IsSuccess: true}
	return
}
