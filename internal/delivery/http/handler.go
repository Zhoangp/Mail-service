package http

import (
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
