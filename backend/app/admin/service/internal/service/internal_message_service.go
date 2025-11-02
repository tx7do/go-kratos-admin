package service

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"
	pagination "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
	internalMessageV1 "kratos-admin/api/gen/go/internal_message/service/v1"

	"kratos-admin/pkg/middleware/auth"
	"kratos-admin/pkg/utils/name_set"
)

type InternalMessageService struct {
	adminV1.InternalMessageServiceHTTPServer

	log *log.Helper

	internalMessageRepo          *data.InternalMessageRepo
	internalMessageCategoryRepo  *data.InternalMessageCategoryRepo
	internalMessageRecipientRepo *data.InternalMessageRecipientRepo
	userRepo                     *data.UserRepo
}

func NewInternalMessageService(
	logger log.Logger,
	internalMessageRepo *data.InternalMessageRepo,
	internalMessageCategoryRepo *data.InternalMessageCategoryRepo,
	internalMessageRecipientRepo *data.InternalMessageRecipientRepo,
	userRepo *data.UserRepo,
) *InternalMessageService {
	l := log.NewHelper(log.With(logger, "module", "internal-message/service/admin-service"))
	return &InternalMessageService{
		log:                          l,
		internalMessageRepo:          internalMessageRepo,
		internalMessageCategoryRepo:  internalMessageCategoryRepo,
		internalMessageRecipientRepo: internalMessageRecipientRepo,
		userRepo:                     userRepo,
	}
}

func (s *InternalMessageService) ListMessage(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListInternalMessageResponse, error) {
	resp, err := s.internalMessageRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	var categorySet = make(name_set.UserNameSetMap)

	for _, v := range resp.Items {
		if v.CategoryId != nil {
			categorySet[v.GetCategoryId()] = nil
		}
	}

	ids := make([]uint32, 0, len(categorySet))
	for id := range categorySet {
		ids = append(ids, id)
	}

	categories, err := s.internalMessageCategoryRepo.GetCategoriesByIds(ctx, ids)
	if err == nil {
		for _, c := range categories {
			categorySet[c.GetId()] = &name_set.UserNameSet{
				UserName: c.GetName(),
			}
		}

		for k, v := range categorySet {
			if v == nil {
				continue
			}

			for i := 0; i < len(resp.Items); i++ {
				if resp.Items[i].CategoryId != nil && resp.Items[i].GetCategoryId() == k {
					resp.Items[i].CategoryName = &v.UserName
				}
			}
		}
	}

	return resp, nil
}

func (s *InternalMessageService) GetMessage(ctx context.Context, req *internalMessageV1.GetInternalMessageRequest) (*internalMessageV1.InternalMessage, error) {
	resp, err := s.internalMessageRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.CategoryId != nil {
		category, err := s.internalMessageCategoryRepo.Get(ctx, &internalMessageV1.GetInternalMessageCategoryRequest{Id: resp.GetCategoryId()})
		if err == nil && category != nil {
			resp.CategoryName = category.Name
		} else {
			s.log.Warnf("Get internal message category failed: %v", err)
		}
	}

	return resp, nil
}

func (s *InternalMessageService) CreateMessage(ctx context.Context, req *internalMessageV1.CreateInternalMessageRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.CreatedBy = trans.Ptr(operator.UserId)

	if _, err = s.internalMessageRepo.Create(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *InternalMessageService) UpdateMessage(ctx context.Context, req *internalMessageV1.UpdateInternalMessageRequest) (*emptypb.Empty, error) {
	if req.Data == nil {
		return nil, adminV1.ErrorBadRequest("invalid parameter")
	}

	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	req.Data.UpdatedBy = trans.Ptr(operator.UserId)

	if err = s.internalMessageRepo.Update(ctx, req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *InternalMessageService) DeleteMessage(ctx context.Context, req *internalMessageV1.DeleteInternalMessageRequest) (*emptypb.Empty, error) {
	if err := s.internalMessageRepo.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

// ListUserInbox 获取用户的收件箱列表 (通知类)
func (s *InternalMessageService) ListUserInbox(ctx context.Context, req *pagination.PagingRequest) (*internalMessageV1.ListUserInboxResponse, error) {
	resp, err := s.internalMessageRecipientRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, d := range resp.Items {
		if d.MessageId == nil {
			continue
		}

		msg, err := s.internalMessageRepo.Get(ctx, &internalMessageV1.GetInternalMessageRequest{
			Id: d.GetMessageId(),
		})
		if err != nil {
			s.log.Errorf("list user inbox failed, get message failed: %s", err)
			continue
		}

		d.Title = msg.Title
		d.Content = msg.Content
	}

	return resp, nil
}

func (s *InternalMessageService) DeleteNotificationFromInbox(ctx context.Context, req *internalMessageV1.DeleteNotificationFromInboxRequest) (*emptypb.Empty, error) {
	var err error
	err = s.internalMessageRecipientRepo.DeleteNotificationFromInbox(ctx, req)
	return &emptypb.Empty{}, err
}

// MarkNotificationAsRead 将通知标记为已读
func (s *InternalMessageService) MarkNotificationAsRead(ctx context.Context, req *internalMessageV1.MarkNotificationAsReadRequest) (*emptypb.Empty, error) {
	var err error
	err = s.internalMessageRecipientRepo.MarkNotificationAsRead(ctx, req)
	return &emptypb.Empty{}, err
}

// MarkNotificationsStatus 标记特定用户的某些或所有通知的状态
func (s *InternalMessageService) MarkNotificationsStatus(ctx context.Context, req *internalMessageV1.MarkNotificationsStatusRequest) (*emptypb.Empty, error) {
	var err error
	err = s.internalMessageRecipientRepo.MarkNotificationsStatus(ctx, req)
	return &emptypb.Empty{}, err
}

// RevokeMessage 撤销某条消息
func (s *InternalMessageService) RevokeMessage(ctx context.Context, req *internalMessageV1.RevokeMessageRequest) (*emptypb.Empty, error) {
	var err error
	if err = s.internalMessageRepo.Delete(ctx, req.GetMessageId()); err != nil {
		s.log.Errorf("delete internal message failed: [%d]", req.GetMessageId())
	}

	if err = s.internalMessageRecipientRepo.RevokeMessage(ctx, req); err != nil {
		s.log.Errorf("delete internal message inbox failed: [%d][%d]", req.GetMessageId(), req.GetUserId())
	}

	return &emptypb.Empty{}, err
}

// SendMessage 发送消息
func (s *InternalMessageService) SendMessage(ctx context.Context, req *internalMessageV1.SendMessageRequest) (*internalMessageV1.SendMessageResponse, error) {
	// 获取操作人信息
	operator, err := auth.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	var msg *internalMessageV1.InternalMessage
	if msg, err = s.internalMessageRepo.Create(ctx, &internalMessageV1.CreateInternalMessageRequest{
		Data: &internalMessageV1.InternalMessage{
			Title:      req.Title,
			Content:    trans.Ptr(req.GetContent()),
			Status:     trans.Ptr(internalMessageV1.InternalMessage_PUBLISHED),
			Type:       trans.Ptr(req.GetType()),
			CategoryId: req.CategoryId,
			CreatedBy:  trans.Ptr(operator.GetUserId()),
			CreatedAt:  timeutil.TimeToTimestamppb(&now),
		},
	}); err != nil {
		s.log.Errorf("create internal message failed: %s", err)
		return nil, err
	}

	if req.GetTargetAll() {
		users, err := s.userRepo.List(ctx, &pagination.PagingRequest{NoPaging: trans.Ptr(true)})
		if err != nil {
			s.log.Errorf("send message failed, list users failed, %s", err)
		} else {
			for _, user := range users.Items {
				if err = s.internalMessageRecipientRepo.Create(ctx, &internalMessageV1.InternalMessageRecipient{
					MessageId:       msg.Id,
					RecipientUserId: user.Id,
					Status:          trans.Ptr(internalMessageV1.InternalMessageRecipient_SENT),
					CreatedBy:       trans.Ptr(operator.GetUserId()),
					CreatedAt:       timeutil.TimeToTimestamppb(&now),
				}); err != nil {
					s.log.Errorf("send message failed, send to user failed, %s", err)
				}
			}
		}
	} else {
		if req.RecipientUserId != nil {
			if err = s.internalMessageRecipientRepo.Create(ctx, &internalMessageV1.InternalMessageRecipient{
				MessageId:       msg.Id,
				RecipientUserId: req.RecipientUserId,
				Status:          trans.Ptr(internalMessageV1.InternalMessageRecipient_SENT),
				CreatedBy:       trans.Ptr(operator.GetUserId()),
				CreatedAt:       timeutil.TimeToTimestamppb(&now),
			}); err != nil {
				s.log.Errorf("send message failed, send to user failed, %s", err)
			}
		} else {
			if len(req.TargetUserIds) != 0 {
				for _, uid := range req.TargetUserIds {
					if err = s.internalMessageRecipientRepo.Create(ctx, &internalMessageV1.InternalMessageRecipient{
						MessageId:       msg.Id,
						RecipientUserId: trans.Ptr(uid),
						Status:          trans.Ptr(internalMessageV1.InternalMessageRecipient_SENT),
						CreatedBy:       trans.Ptr(operator.GetUserId()),
						CreatedAt:       timeutil.TimeToTimestamppb(&now),
					}); err != nil {
						s.log.Errorf("send message failed, send to user failed, %s", err)
					}
				}
			}
		}
	}

	return &internalMessageV1.SendMessageResponse{
		MessageId: msg.GetId(),
	}, nil
}
