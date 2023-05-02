package http

import (
	"context"
	"fmt"
	"github.com/Zhoangp/Mail-service/internal/model"
	"github.com/Zhoangp/Mail-service/pb"
)

func (hdl mailHandler) SendTokenVerifyAccount(context context.Context, request *pb.SendTokenVerifyAccountRequest) (*pb.SendTokenVerifyAccountResponse, error) {
	err := hdl.uc.SendEmail(&model.Email{
		DestMail: request.Mail.DestMail,
		Subject:  request.Mail.Subject,
	}, &model.SendTokenContent{
		Name: request.Name,
		Url:  request.Url + request.Token,
	})
	if err != nil {
		fmt.Println(err)
		return &pb.SendTokenVerifyAccountResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.SendTokenVerifyAccountResponse{}, nil

}
