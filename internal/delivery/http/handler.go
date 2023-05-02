package http

import (
	"context"
	"fmt"
	"github.com/Zhoangp/Mail-service/internal/model"
	"github.com/Zhoangp/Mail-service/pb"
	"github.com/Zhoangp/Mail-service/pkg/common"
)

type MailUsecase interface {
	SendEmail(email *model.Email, content *model.SendTokenContent) error
}
type mailHandler struct {
	uc MailUsecase
	pb.UnimplementedMailServiceServer
}

func NewMailHandler(uc MailUsecase) *mailHandler {
	return &mailHandler{uc: uc}
}
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
func HandleError(err error) *pb.ErrorResponse {
	if errors, ok := err.(*common.AppError); ok {
		return &pb.ErrorResponse{
			Code:    int64(errors.StatusCode),
			Message: errors.Message,
		}
	}
	appErr := common.ErrInternal(err.(error))
	return &pb.ErrorResponse{
		Code:    int64(appErr.StatusCode),
		Message: appErr.Message,
	}
}
