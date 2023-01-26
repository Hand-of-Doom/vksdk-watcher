package watcher

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"github.com/SevereCloud/vksdk/v2/object"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MessageNewHandler = func(obj events.MessageNewObject, client *api.VK) error

type MessageNewMiddleware = func(obj events.MessageNewObject, client *api.VK) (bool, error)

type MessageReplyHandler = func(obj events.MessageReplyObject, client *api.VK) error

type MessageReplyMiddleware = func(obj events.MessageReplyObject, client *api.VK) (bool, error)

type MessageEditHandler = func(obj events.MessageEditObject, client *api.VK) error

type MessageEditMiddleware = func(obj events.MessageEditObject, client *api.VK) (bool, error)

type MessageAllowHandler = func(obj events.MessageAllowObject, client *api.VK) error

type MessageAllowMiddleware = func(obj events.MessageAllowObject, client *api.VK) (bool, error)

type MessageDenyHandler = func(obj events.MessageDenyObject, client *api.VK) error

type MessageDenyMiddleware = func(obj events.MessageDenyObject, client *api.VK) (bool, error)

type MessageTypingStateHandler = func(obj events.MessageTypingStateObject, client *api.VK) error

type MessageTypingStateMiddleware = func(obj events.MessageTypingStateObject, client *api.VK) (bool, error)

type PhotoNewHandler = func(obj events.PhotoNewObject, client *api.VK) error

type PhotoNewMiddleware = func(obj events.PhotoNewObject, client *api.VK) (bool, error)

type PhotoCommentNewHandler = func(obj events.PhotoCommentNewObject, client *api.VK) error

type PhotoCommentNewMiddleware = func(obj events.PhotoCommentNewObject, client *api.VK) (bool, error)

type PhotoCommentEditHandler = func(obj events.PhotoCommentEditObject, client *api.VK) error

type PhotoCommentEditMiddleware = func(obj events.PhotoCommentEditObject, client *api.VK) (bool, error)

type PhotoCommentRestoreHandler = func(obj events.PhotoCommentRestoreObject, client *api.VK) error

type PhotoCommentRestoreMiddleware = func(obj events.PhotoCommentRestoreObject, client *api.VK) (bool, error)

type PhotoCommentDeleteHandler = func(obj events.PhotoCommentDeleteObject, client *api.VK) error

type PhotoCommentDeleteMiddleware = func(obj events.PhotoCommentDeleteObject, client *api.VK) (bool, error)

type AudioNewHandler = func(obj events.AudioNewObject, client *api.VK) error

type AudioNewMiddleware = func(obj events.AudioNewObject, client *api.VK) (bool, error)

type VideoNewHandler = func(obj events.VideoNewObject, client *api.VK) error

type VideoNewMiddleware = func(obj events.VideoNewObject, client *api.VK) (bool, error)

type VideoCommentNewHandler = func(obj events.VideoCommentNewObject, client *api.VK) error

type VideoCommentNewMiddleware = func(obj events.VideoCommentNewObject, client *api.VK) (bool, error)

type VideoCommentEditHandler = func(obj events.VideoCommentEditObject, client *api.VK) error

type VideoCommentEditMiddleware = func(obj events.VideoCommentEditObject, client *api.VK) (bool, error)

type VideoCommentRestoreHandler = func(obj events.VideoCommentRestoreObject, client *api.VK) error

type VideoCommentRestoreMiddleware = func(obj events.VideoCommentRestoreObject, client *api.VK) (bool, error)

type VideoCommentDeleteHandler = func(obj events.VideoCommentDeleteObject, client *api.VK) error

type VideoCommentDeleteMiddleware = func(obj events.VideoCommentDeleteObject, client *api.VK) (bool, error)

type WallPostNewHandler = func(obj events.WallPostNewObject, client *api.VK) error

type WallPostNewMiddleware = func(obj events.WallPostNewObject, client *api.VK) (bool, error)

type WallRepostHandler = func(obj events.WallRepostObject, client *api.VK) error

type WallRepostMiddleware = func(obj events.WallRepostObject, client *api.VK) (bool, error)

type WallReplyNewHandler = func(obj events.WallReplyNewObject, client *api.VK) error

type WallReplyNewMiddleware = func(obj events.WallReplyNewObject, client *api.VK) (bool, error)

type WallReplyEditHandler = func(obj events.WallReplyEditObject, client *api.VK) error

type WallReplyEditMiddleware = func(obj events.WallReplyEditObject, client *api.VK) (bool, error)

type WallReplyRestoreHandler = func(obj events.WallReplyRestoreObject, client *api.VK) error

type WallReplyRestoreMiddleware = func(obj events.WallReplyRestoreObject, client *api.VK) (bool, error)

type WallReplyDeleteHandler = func(obj events.WallReplyDeleteObject, client *api.VK) error

type WallReplyDeleteMiddleware = func(obj events.WallReplyDeleteObject, client *api.VK) (bool, error)

type BoardPostNewHandler = func(obj events.BoardPostNewObject, client *api.VK) error

type BoardPostNewMiddleware = func(obj events.BoardPostNewObject, client *api.VK) (bool, error)

type BoardPostEditHandler = func(obj events.BoardPostEditObject, client *api.VK) error

type BoardPostEditMiddleware = func(obj events.BoardPostEditObject, client *api.VK) (bool, error)

type BoardPostRestoreHandler = func(obj events.BoardPostRestoreObject, client *api.VK) error

type BoardPostRestoreMiddleware = func(obj events.BoardPostRestoreObject, client *api.VK) (bool, error)

type BoardPostDeleteHandler = func(obj events.BoardPostDeleteObject, client *api.VK) error

type BoardPostDeleteMiddleware = func(obj events.BoardPostDeleteObject, client *api.VK) (bool, error)

type MarketCommentNewHandler = func(obj events.MarketCommentNewObject, client *api.VK) error

type MarketCommentNewMiddleware = func(obj events.MarketCommentNewObject, client *api.VK) (bool, error)

type MarketCommentEditHandler = func(obj events.MarketCommentEditObject, client *api.VK) error

type MarketCommentEditMiddleware = func(obj events.MarketCommentEditObject, client *api.VK) (bool, error)

type MarketCommentRestoreHandler = func(obj events.MarketCommentRestoreObject, client *api.VK) error

type MarketCommentRestoreMiddleware = func(obj events.MarketCommentRestoreObject, client *api.VK) (bool, error)

type MarketCommentDeleteHandler = func(obj events.MarketCommentDeleteObject, client *api.VK) error

type MarketCommentDeleteMiddleware = func(obj events.MarketCommentDeleteObject, client *api.VK) (bool, error)

type MarketOrderNewHandler = func(obj events.MarketOrderNewObject, client *api.VK) error

type MarketOrderNewMiddleware = func(obj events.MarketOrderNewObject, client *api.VK) (bool, error)

type MarketOrderEditHandler = func(obj events.MarketOrderEditObject, client *api.VK) error

type MarketOrderEditMiddleware = func(obj events.MarketOrderEditObject, client *api.VK) (bool, error)

type GroupLeaveHandler = func(obj events.GroupLeaveObject, client *api.VK) error

type GroupLeaveMiddleware = func(obj events.GroupLeaveObject, client *api.VK) (bool, error)

type GroupJoinHandler = func(obj events.GroupJoinObject, client *api.VK) error

type GroupJoinMiddleware = func(obj events.GroupJoinObject, client *api.VK) (bool, error)

type UserBlockHandler = func(obj events.UserBlockObject, client *api.VK) error

type UserBlockMiddleware = func(obj events.UserBlockObject, client *api.VK) (bool, error)

type UserUnblockHandler = func(obj events.UserUnblockObject, client *api.VK) error

type UserUnblockMiddleware = func(obj events.UserUnblockObject, client *api.VK) (bool, error)

type PollVoteNewHandler = func(obj events.PollVoteNewObject, client *api.VK) error

type PollVoteNewMiddleware = func(obj events.PollVoteNewObject, client *api.VK) (bool, error)

type GroupOfficersEditHandler = func(obj events.GroupOfficersEditObject, client *api.VK) error

type GroupOfficersEditMiddleware = func(obj events.GroupOfficersEditObject, client *api.VK) (bool, error)

type GroupChangeSettingsHandler = func(obj events.GroupChangeSettingsObject, client *api.VK) error

type GroupChangeSettingsMiddleware = func(obj events.GroupChangeSettingsObject, client *api.VK) (bool, error)

type GroupChangePhotoHandler = func(obj events.GroupChangePhotoObject, client *api.VK) error

type GroupChangePhotoMiddleware = func(obj events.GroupChangePhotoObject, client *api.VK) (bool, error)

type VkpayTransactionHandler = func(obj events.VkpayTransactionObject, client *api.VK) error

type VkpayTransactionMiddleware = func(obj events.VkpayTransactionObject, client *api.VK) (bool, error)

type LeadFormsNewHandler = func(obj events.LeadFormsNewObject, client *api.VK) error

type LeadFormsNewMiddleware = func(obj events.LeadFormsNewObject, client *api.VK) (bool, error)

type AppPayloadHandler = func(obj events.AppPayloadObject, client *api.VK) error

type AppPayloadMiddleware = func(obj events.AppPayloadObject, client *api.VK) (bool, error)

type MessageReadHandler = func(obj events.MessageReadObject, client *api.VK) error

type MessageReadMiddleware = func(obj events.MessageReadObject, client *api.VK) (bool, error)

type LikeAddHandler = func(obj events.LikeAddObject, client *api.VK) error

type LikeAddMiddleware = func(obj events.LikeAddObject, client *api.VK) (bool, error)

type LikeRemoveHandler = func(obj events.LikeRemoveObject, client *api.VK) error

type LikeRemoveMiddleware = func(obj events.LikeRemoveObject, client *api.VK) (bool, error)

type DonutSubscriptionCreateHandler = func(obj events.DonutSubscriptionCreateObject, client *api.VK) error

type DonutSubscriptionCreateMiddleware = func(obj events.DonutSubscriptionCreateObject, client *api.VK) (bool, error)

type DonutSubscriptionProlongedHandler = func(obj events.DonutSubscriptionProlongedObject, client *api.VK) error

type DonutSubscriptionProlongedMiddleware = func(obj events.DonutSubscriptionProlongedObject, client *api.VK) (bool, error)

type DonutSubscriptionExpiredHandler = func(obj events.DonutSubscriptionExpiredObject, client *api.VK) error

type DonutSubscriptionExpiredMiddleware = func(obj events.DonutSubscriptionExpiredObject, client *api.VK) (bool, error)

type DonutSubscriptionCancelledHandler = func(obj events.DonutSubscriptionCancelledObject, client *api.VK) error

type DonutSubscriptionCancelledMiddleware = func(obj events.DonutSubscriptionCancelledObject, client *api.VK) (bool, error)

type DonutSubscriptionPriceChangedHandler = func(obj events.DonutSubscriptionPriceChangedObject, client *api.VK) error

type DonutSubscriptionPriceChangedMiddleware = func(obj events.DonutSubscriptionPriceChangedObject, client *api.VK) (bool, error)

type DonutMoneyWithdrawHandler = func(obj events.DonutMoneyWithdrawObject, client *api.VK) error

type DonutMoneyWithdrawMiddleware = func(obj events.DonutMoneyWithdrawObject, client *api.VK) (bool, error)

type DonutMoneyWithdrawErrorHandler = func(obj events.DonutMoneyWithdrawErrorObject, client *api.VK) error

type DonutMoneyWithdrawErrorMiddleware = func(obj events.DonutMoneyWithdrawErrorObject, client *api.VK) (bool, error)

type Watcher struct {
	client  *api.VK
	lp      *longpoll.LongPoll
	chanErr chan error
}

func (w *Watcher) WatchMessageNew(handler MessageNewHandler, middlewares ...MessageNewMiddleware) *Watcher {
	w.lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMessageReply(handler MessageReplyHandler, middlewares ...MessageReplyMiddleware) *Watcher {
	w.lp.MessageReply(func(ctx context.Context, obj events.MessageReplyObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMessageEdit(handler MessageEditHandler, middlewares ...MessageEditMiddleware) *Watcher {
	w.lp.MessageEdit(func(ctx context.Context, obj events.MessageEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMessageAllow(handler MessageAllowHandler, middlewares ...MessageAllowMiddleware) *Watcher {
	w.lp.MessageAllow(func(ctx context.Context, obj events.MessageAllowObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMessageDeny(handler MessageDenyHandler, middlewares ...MessageDenyMiddleware) *Watcher {
	w.lp.MessageDeny(func(ctx context.Context, obj events.MessageDenyObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMessageTypingState(handler MessageTypingStateHandler, middlewares ...MessageTypingStateMiddleware) *Watcher {
	w.lp.MessageTypingState(func(ctx context.Context, obj events.MessageTypingStateObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchPhotoNew(handler PhotoNewHandler, middlewares ...PhotoNewMiddleware) *Watcher {
	w.lp.PhotoNew(func(ctx context.Context, obj events.PhotoNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchPhotoCommentNew(handler PhotoCommentNewHandler, middlewares ...PhotoCommentNewMiddleware) *Watcher {
	w.lp.PhotoCommentNew(func(ctx context.Context, obj events.PhotoCommentNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchPhotoCommentEdit(handler PhotoCommentEditHandler, middlewares ...PhotoCommentEditMiddleware) *Watcher {
	w.lp.PhotoCommentEdit(func(ctx context.Context, obj events.PhotoCommentEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchPhotoCommentRestore(handler PhotoCommentRestoreHandler, middlewares ...PhotoCommentRestoreMiddleware) *Watcher {
	w.lp.PhotoCommentRestore(func(ctx context.Context, obj events.PhotoCommentRestoreObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchPhotoCommentDelete(handler PhotoCommentDeleteHandler, middlewares ...PhotoCommentDeleteMiddleware) *Watcher {
	w.lp.PhotoCommentDelete(func(ctx context.Context, obj events.PhotoCommentDeleteObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchAudioNew(handler AudioNewHandler, middlewares ...AudioNewMiddleware) *Watcher {
	w.lp.AudioNew(func(ctx context.Context, obj events.AudioNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchVideoNew(handler VideoNewHandler, middlewares ...VideoNewMiddleware) *Watcher {
	w.lp.VideoNew(func(ctx context.Context, obj events.VideoNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchVideoCommentNew(handler VideoCommentNewHandler, middlewares ...VideoCommentNewMiddleware) *Watcher {
	w.lp.VideoCommentNew(func(ctx context.Context, obj events.VideoCommentNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchVideoCommentEdit(handler VideoCommentEditHandler, middlewares ...VideoCommentEditMiddleware) *Watcher {
	w.lp.VideoCommentEdit(func(ctx context.Context, obj events.VideoCommentEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchVideoCommentRestore(handler VideoCommentRestoreHandler, middlewares ...VideoCommentRestoreMiddleware) *Watcher {
	w.lp.VideoCommentRestore(func(ctx context.Context, obj events.VideoCommentRestoreObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchVideoCommentDelete(handler VideoCommentDeleteHandler, middlewares ...VideoCommentDeleteMiddleware) *Watcher {
	w.lp.VideoCommentDelete(func(ctx context.Context, obj events.VideoCommentDeleteObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchWallPostNew(handler WallPostNewHandler, middlewares ...WallPostNewMiddleware) *Watcher {
	w.lp.WallPostNew(func(ctx context.Context, obj events.WallPostNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchWallRepost(handler WallRepostHandler, middlewares ...WallRepostMiddleware) *Watcher {
	w.lp.WallRepost(func(ctx context.Context, obj events.WallRepostObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchWallReplyNew(handler WallReplyNewHandler, middlewares ...WallReplyNewMiddleware) *Watcher {
	w.lp.WallReplyNew(func(ctx context.Context, obj events.WallReplyNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchWallReplyEdit(handler WallReplyEditHandler, middlewares ...WallReplyEditMiddleware) *Watcher {
	w.lp.WallReplyEdit(func(ctx context.Context, obj events.WallReplyEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchWallReplyRestore(handler WallReplyRestoreHandler, middlewares ...WallReplyRestoreMiddleware) *Watcher {
	w.lp.WallReplyRestore(func(ctx context.Context, obj events.WallReplyRestoreObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchWallReplyDelete(handler WallReplyDeleteHandler, middlewares ...WallReplyDeleteMiddleware) *Watcher {
	w.lp.WallReplyDelete(func(ctx context.Context, obj events.WallReplyDeleteObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchBoardPostNew(handler BoardPostNewHandler, middlewares ...BoardPostNewMiddleware) *Watcher {
	w.lp.BoardPostNew(func(ctx context.Context, obj events.BoardPostNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchBoardPostEdit(handler BoardPostEditHandler, middlewares ...BoardPostEditMiddleware) *Watcher {
	w.lp.BoardPostEdit(func(ctx context.Context, obj events.BoardPostEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchBoardPostRestore(handler BoardPostRestoreHandler, middlewares ...BoardPostRestoreMiddleware) *Watcher {
	w.lp.BoardPostRestore(func(ctx context.Context, obj events.BoardPostRestoreObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchBoardPostDelete(handler BoardPostDeleteHandler, middlewares ...BoardPostDeleteMiddleware) *Watcher {
	w.lp.BoardPostDelete(func(ctx context.Context, obj events.BoardPostDeleteObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMarketCommentNew(handler MarketCommentNewHandler, middlewares ...MarketCommentNewMiddleware) *Watcher {
	w.lp.MarketCommentNew(func(ctx context.Context, obj events.MarketCommentNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMarketCommentEdit(handler MarketCommentEditHandler, middlewares ...MarketCommentEditMiddleware) *Watcher {
	w.lp.MarketCommentEdit(func(ctx context.Context, obj events.MarketCommentEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMarketCommentRestore(handler MarketCommentRestoreHandler, middlewares ...MarketCommentRestoreMiddleware) *Watcher {
	w.lp.MarketCommentRestore(func(ctx context.Context, obj events.MarketCommentRestoreObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMarketCommentDelete(handler MarketCommentDeleteHandler, middlewares ...MarketCommentDeleteMiddleware) *Watcher {
	w.lp.MarketCommentDelete(func(ctx context.Context, obj events.MarketCommentDeleteObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMarketOrderNew(handler MarketOrderNewHandler, middlewares ...MarketOrderNewMiddleware) *Watcher {
	w.lp.MarketOrderNew(func(ctx context.Context, obj events.MarketOrderNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMarketOrderEdit(handler MarketOrderEditHandler, middlewares ...MarketOrderEditMiddleware) *Watcher {
	w.lp.MarketOrderEdit(func(ctx context.Context, obj events.MarketOrderEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchGroupLeave(handler GroupLeaveHandler, middlewares ...GroupLeaveMiddleware) *Watcher {
	w.lp.GroupLeave(func(ctx context.Context, obj events.GroupLeaveObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchGroupJoin(handler GroupJoinHandler, middlewares ...GroupJoinMiddleware) *Watcher {
	w.lp.GroupJoin(func(ctx context.Context, obj events.GroupJoinObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchUserBlock(handler UserBlockHandler, middlewares ...UserBlockMiddleware) *Watcher {
	w.lp.UserBlock(func(ctx context.Context, obj events.UserBlockObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchUserUnblock(handler UserUnblockHandler, middlewares ...UserUnblockMiddleware) *Watcher {
	w.lp.UserUnblock(func(ctx context.Context, obj events.UserUnblockObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchPollVoteNew(handler PollVoteNewHandler, middlewares ...PollVoteNewMiddleware) *Watcher {
	w.lp.PollVoteNew(func(ctx context.Context, obj events.PollVoteNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchGroupOfficersEdit(handler GroupOfficersEditHandler, middlewares ...GroupOfficersEditMiddleware) *Watcher {
	w.lp.GroupOfficersEdit(func(ctx context.Context, obj events.GroupOfficersEditObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchGroupChangeSettings(handler GroupChangeSettingsHandler, middlewares ...GroupChangeSettingsMiddleware) *Watcher {
	w.lp.GroupChangeSettings(func(ctx context.Context, obj events.GroupChangeSettingsObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchGroupChangePhoto(handler GroupChangePhotoHandler, middlewares ...GroupChangePhotoMiddleware) *Watcher {
	w.lp.GroupChangePhoto(func(ctx context.Context, obj events.GroupChangePhotoObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchVkpayTransaction(handler VkpayTransactionHandler, middlewares ...VkpayTransactionMiddleware) *Watcher {
	w.lp.VkpayTransaction(func(ctx context.Context, obj events.VkpayTransactionObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchLeadFormsNew(handler LeadFormsNewHandler, middlewares ...LeadFormsNewMiddleware) *Watcher {
	w.lp.LeadFormsNew(func(ctx context.Context, obj events.LeadFormsNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchAppPayload(handler AppPayloadHandler, middlewares ...AppPayloadMiddleware) *Watcher {
	w.lp.AppPayload(func(ctx context.Context, obj events.AppPayloadObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchMessageRead(handler MessageReadHandler, middlewares ...MessageReadMiddleware) *Watcher {
	w.lp.MessageRead(func(ctx context.Context, obj events.MessageReadObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchLikeAdd(handler LikeAddHandler, middlewares ...LikeAddMiddleware) *Watcher {
	w.lp.LikeAdd(func(ctx context.Context, obj events.LikeAddObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchLikeRemove(handler LikeRemoveHandler, middlewares ...LikeRemoveMiddleware) *Watcher {
	w.lp.LikeRemove(func(ctx context.Context, obj events.LikeRemoveObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutSubscriptionCreate(handler DonutSubscriptionCreateHandler, middlewares ...DonutSubscriptionCreateMiddleware) *Watcher {
	w.lp.DonutSubscriptionCreate(func(ctx context.Context, obj events.DonutSubscriptionCreateObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutSubscriptionProlonged(handler DonutSubscriptionProlongedHandler, middlewares ...DonutSubscriptionProlongedMiddleware) *Watcher {
	w.lp.DonutSubscriptionProlonged(func(ctx context.Context, obj events.DonutSubscriptionProlongedObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutSubscriptionExpired(handler DonutSubscriptionExpiredHandler, middlewares ...DonutSubscriptionExpiredMiddleware) *Watcher {
	w.lp.DonutSubscriptionExpired(func(ctx context.Context, obj events.DonutSubscriptionExpiredObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutSubscriptionCancelled(handler DonutSubscriptionCancelledHandler, middlewares ...DonutSubscriptionCancelledMiddleware) *Watcher {
	w.lp.DonutSubscriptionCancelled(func(ctx context.Context, obj events.DonutSubscriptionCancelledObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutSubscriptionPriceChanged(handler DonutSubscriptionPriceChangedHandler, middlewares ...DonutSubscriptionPriceChangedMiddleware) *Watcher {
	w.lp.DonutSubscriptionPriceChanged(func(ctx context.Context, obj events.DonutSubscriptionPriceChangedObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutMoneyWithdraw(handler DonutMoneyWithdrawHandler, middlewares ...DonutMoneyWithdrawMiddleware) *Watcher {
	w.lp.DonutMoneyWithdraw(func(ctx context.Context, obj events.DonutMoneyWithdrawObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

func (w *Watcher) WatchDonutMoneyWithdrawError(handler DonutMoneyWithdrawErrorHandler, middlewares ...DonutMoneyWithdrawErrorMiddleware) *Watcher {
	w.lp.DonutMoneyWithdrawError(func(ctx context.Context, obj events.DonutMoneyWithdrawErrorObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		err := handler(obj, w.client)
		if err != nil {
			w.chanErr <- err
		}
	})

	return w
}

type MessageParams struct {
	regexp  *regexp.Regexp
	matches []string
}

func (p *MessageParams) Exists(name string) bool {
	return len(p.matches) != 0
}

func (p *MessageParams) GetString(name string) string {
	pos := p.regexp.SubexpIndex(name)
	return p.matches[pos]
}

func (p *MessageParams) GetInt64(name string) (int64, error) {
	param := p.GetString(name)
	return strconv.ParseInt(param, 0, 64)
}

func (p *MessageParams) GetInt32(name string) (int32, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseInt(param, 0, 32)
	return int32(paramCasted), err
}

func (p *MessageParams) GetInt16(name string) (int16, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseInt(param, 0, 16)
	return int16(paramCasted), err
}

func (p *MessageParams) GetInt8(name string) (int8, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseInt(param, 0, 8)
	return int8(paramCasted), err
}

func (p *MessageParams) GetInt(name string) (int, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseInt(param, 0, 32)
	return int(paramCasted), err
}

func (p *MessageParams) GetUint64(name string) (uint64, error) {
	param := p.GetString(name)
	return strconv.ParseUint(param, 0, 64)
}

func (p *MessageParams) GetUint32(name string) (uint32, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseUint(param, 0, 32)
	return uint32(paramCasted), err
}

func (p *MessageParams) GetUint16(name string) (uint16, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseUint(param, 0, 16)
	return uint16(paramCasted), err
}

func (p *MessageParams) GetUint8(name string) (uint8, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseUint(param, 0, 8)
	return uint8(paramCasted), err
}

func (p *MessageParams) GetUint(name string) (uint, error) {
	param := p.GetString(name)

	paramCasted, err := strconv.ParseUint(param, 0, 32)
	return uint(paramCasted), err
}

func (p *MessageParams) GetStrings(name string, sep string) []string {
	param := p.GetString(name)
	return strings.Split(param, sep)
}

func (p *MessageParams) GetInts64(name string, sep string) ([]int64, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]int64, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, paramCasted)
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetInts32(name string, sep string) ([]int32, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]int32, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseInt(param, 0, 32)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, int32(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetInts16(name string, sep string) ([]int16, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]int16, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseInt(param, 0, 16)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, int16(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetInts8(name string, sep string) ([]int8, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]int8, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseInt(param, 0, 8)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, int8(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetInts(name string, sep string) ([]int, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]int, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseInt(param, 0, 32)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, int(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetUints64(name string, sep string) ([]uint64, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]uint64, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseUint(param, 0, 64)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, paramCasted)
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetUints32(name string, sep string) ([]uint32, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]uint32, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseUint(param, 0, 32)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, uint32(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetUints16(name string, sep string) ([]uint16, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]uint16, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseUint(param, 0, 16)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, uint16(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetUints8(name string, sep string) ([]uint8, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]uint8, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseUint(param, 0, 8)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, uint8(paramCasted))
	}

	return paramsCasted, errs
}

func (p *MessageParams) GetUints(name string, sep string) ([]uint, []error) {
	params := p.GetStrings(name, sep)
	var errs []error

	paramsCasted := make([]uint, 0, len(params))
	for _, param := range params {
		paramCasted, err := strconv.ParseUint(param, 0, 32)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		paramsCasted = append(paramsCasted, uint(paramCasted))
	}

	return paramsCasted, errs
}

func newMessageParams(regexp *regexp.Regexp, message string) *MessageParams {
	p := new(MessageParams)
	p.regexp = regexp
	p.matches = regexp.FindStringSubmatch(message)

	return p
}

type Answer struct {
	params api.Params
	client *api.VK
}

// Message (Required if 'attachments' is not set.) Text of the message.
func (a *Answer) Message(v ...string) *Answer {
	a.params["message"] = strings.Join(v, " ")
	return a
}

// Lat geographical latitude of a check-in, in degrees (from -90 to 90).
func (a *Answer) Lat(v float64) *Answer {
	a.params["lat"] = v
	return a
}

// Long geographical longitude of a check-in, in degrees (from -180 to 180).
func (a *Answer) Long(v float64) *Answer {
	a.params["long"] = v
	return a
}

// Attachment (Required if 'message' is not set.) List of objects attached to
// the message, separated by commas, in the following format:
// "<\owner_id>_<\media_id>", ” — Type of media attachment: 'photo' — photo,
// 'video' — video, 'audio' — audio, 'doc' — document, 'wall' — wall post,
// '<\owner_id>' — ID of the media attachment owner. '<\media_id>' — media
// attachment ID. Example: "photo100172_166443618".
func (a *Answer) Attachment(v interface{}) *Answer {
	a.params["attachment"] = v
	return a
}

// ReplyTo parameter.
func (a *Answer) ReplyTo(v int) *Answer {
	a.params["reply_to"] = v
	return a
}

// ForwardMessages ID of forwarded messages, separated with a comma. Listed
// messages of the sender will be shown in the message body at the
// recipient's. Example: "123,431,544".
func (a *Answer) ForwardMessages(v []int) *Answer {
	a.params["forward_messages"] = v
	return a
}

// Forward parameter.
func (a *Answer) Forward(v string) *Answer {
	a.params["forward"] = v
	return a
}

// StickerID sticker id.
func (a *Answer) StickerID(v int) *Answer {
	a.params["sticker_id"] = v
	return a
}

// Keyboard parameter.
// https://vk.com/dev/bots_docs_3
func (a *Answer) Keyboard(v interface{}) *Answer {
	a.params["keyboard"] = v
	return a
}

// Template parameter.
// https://vk.com/dev/bot_docs_templates
func (a *Answer) Template(v interface{}) *Answer {
	a.params["template"] = v
	return a
}

// Payload parameter.
func (a *Answer) Payload(v string) *Answer {
	a.params["payload"] = v
	return a
}

// DontParseLinks parameter.
func (a *Answer) DontParseLinks(v bool) *Answer {
	a.params["dont_parse_links"] = v
	return a
}

// DisableMentions parameter.
func (a *Answer) DisableMentions(v bool) *Answer {
	a.params["disable_mentions"] = v
	return a
}

// Intent parameter.
//
// https://vk.com/dev/bots_docs_4
func (a *Answer) Intent(v string) *Answer {
	a.params["intent"] = v
	return a
}

// SubscribeID parameter.
//
// TODO: write subscribe_id documentation.
func (a *Answer) SubscribeID(v int) *Answer {
	a.params["subscribe_id"] = v
	return a
}

func (a *Answer) Send() (response int, err error) {
	a.params["random_id"] = time.Now().UnixNano()
	response, err = a.client.MessagesSend(a.params)

	a.params = api.Params{}

	return
}

func newAnswer(peerID int, client *api.VK) *Answer {
	a := new(Answer)
	a.params = api.Params{
		"peer_id": peerID,
	}
	a.client = client

	return a
}

type MessageContext struct {
	Obj        object.MessagesMessage
	ClientInfo object.ClientInfo
	Params     *MessageParams
	Answer     *Answer
}

func newMessageContext(obj object.MessagesMessage, clientInfo object.ClientInfo, params *MessageParams, answer *Answer) *MessageContext {
	ctx := new(MessageContext)

	ctx.Obj = obj
	ctx.ClientInfo = clientInfo
	ctx.Params = params
	ctx.Answer = answer

	return ctx
}

type MessageNewHandlerExtended = func(ctx *MessageContext, client *api.VK) error

func (w *Watcher) HandleMessage(pattern string, handler MessageNewHandlerExtended, middlewares ...MessageNewMiddleware) *Watcher {
	regexpMessageText := regexp.MustCompile(pattern)

	w.lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		for _, middleware := range middlewares {
			ok, err := middleware(obj, w.client)
			if err != nil {
				w.chanErr <- err
				return
			}
			if !ok {
				return
			}
		}

		if regexpMessageText.MatchString(obj.Message.Text) {
			params := newMessageParams(regexpMessageText, obj.Message.Text)
			answer := newAnswer(obj.Message.PeerID, w.client)

			ctx := newMessageContext(obj.Message, obj.ClientInfo, params, answer)

			err := handler(ctx, w.client)
			if err != nil {
				w.chanErr <- err
			}
		}
	})

	return w
}

type StartWatcherFunc func()

func NewWatcher(client *api.VK, chanErr chan error) (*Watcher, StartWatcherFunc) {
	w := new(Watcher)
	w.chanErr = chanErr
	w.client = client

	lp, err := longpoll.NewLongPollCommunity(client)
	if err != nil {
		chanErr <- err
		return nil, nil
	}
	lp.FullResponse(func(response longpoll.Response) {
		log.Println(response)
	})

	w.lp = lp

	startWatcher := func() {
		if err = lp.Run(); err != nil {
			chanErr <- err
		}
	}

	return w, startWatcher
}
