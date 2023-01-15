package watcher

import (
	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/longpoll-bot"
	"github.com/SevereCloud/vksdk/object"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type MessageNewHandler = func(obj object.MessageNewObject, client *api.VK) error

type MessageNewMiddleware = func(obj object.MessageNewObject, client *api.VK) (bool, error)

type MessageReplyHandler = func(obj object.MessageReplyObject, client *api.VK) error

type MessageReplyMiddleware = func(obj object.MessageReplyObject, client *api.VK) (bool, error)

type MessageEditHandler = func(obj object.MessageEditObject, client *api.VK) error

type MessageEditMiddleware = func(obj object.MessageEditObject, client *api.VK) (bool, error)

type MessageAllowHandler = func(obj object.MessageAllowObject, client *api.VK) error

type MessageAllowMiddleware = func(obj object.MessageAllowObject, client *api.VK) (bool, error)

type MessageDenyHandler = func(obj object.MessageDenyObject, client *api.VK) error

type MessageDenyMiddleware = func(obj object.MessageDenyObject, client *api.VK) (bool, error)

type MessageTypingStateHandler = func(obj object.MessageTypingStateObject, client *api.VK) error

type MessageTypingStateMiddleware = func(obj object.MessageTypingStateObject, client *api.VK) (bool, error)

type PhotoNewHandler = func(obj object.PhotoNewObject, client *api.VK) error

type PhotoNewMiddleware = func(obj object.PhotoNewObject, client *api.VK) (bool, error)

type PhotoCommentNewHandler = func(obj object.PhotoCommentNewObject, client *api.VK) error

type PhotoCommentNewMiddleware = func(obj object.PhotoCommentNewObject, client *api.VK) (bool, error)

type PhotoCommentEditHandler = func(obj object.PhotoCommentEditObject, client *api.VK) error

type PhotoCommentEditMiddleware = func(obj object.PhotoCommentEditObject, client *api.VK) (bool, error)

type PhotoCommentRestoreHandler = func(obj object.PhotoCommentRestoreObject, client *api.VK) error

type PhotoCommentRestoreMiddleware = func(obj object.PhotoCommentRestoreObject, client *api.VK) (bool, error)

type PhotoCommentDeleteHandler = func(obj object.PhotoCommentDeleteObject, client *api.VK) error

type PhotoCommentDeleteMiddleware = func(obj object.PhotoCommentDeleteObject, client *api.VK) (bool, error)

type AudioNewHandler = func(obj object.AudioNewObject, client *api.VK) error

type AudioNewMiddleware = func(obj object.AudioNewObject, client *api.VK) (bool, error)

type VideoNewHandler = func(obj object.VideoNewObject, client *api.VK) error

type VideoNewMiddleware = func(obj object.VideoNewObject, client *api.VK) (bool, error)

type VideoCommentNewHandler = func(obj object.VideoCommentNewObject, client *api.VK) error

type VideoCommentNewMiddleware = func(obj object.VideoCommentNewObject, client *api.VK) (bool, error)

type VideoCommentEditHandler = func(obj object.VideoCommentEditObject, client *api.VK) error

type VideoCommentEditMiddleware = func(obj object.VideoCommentEditObject, client *api.VK) (bool, error)

type VideoCommentRestoreHandler = func(obj object.VideoCommentRestoreObject, client *api.VK) error

type VideoCommentRestoreMiddleware = func(obj object.VideoCommentRestoreObject, client *api.VK) (bool, error)

type VideoCommentDeleteHandler = func(obj object.VideoCommentDeleteObject, client *api.VK) error

type VideoCommentDeleteMiddleware = func(obj object.VideoCommentDeleteObject, client *api.VK) (bool, error)

type WallPostNewHandler = func(obj object.WallPostNewObject, client *api.VK) error

type WallPostNewMiddleware = func(obj object.WallPostNewObject, client *api.VK) (bool, error)

type WallRepostHandler = func(obj object.WallRepostObject, client *api.VK) error

type WallRepostMiddleware = func(obj object.WallRepostObject, client *api.VK) (bool, error)

type WallReplyNewHandler = func(obj object.WallReplyNewObject, client *api.VK) error

type WallReplyNewMiddleware = func(obj object.WallReplyNewObject, client *api.VK) (bool, error)

type WallReplyEditHandler = func(obj object.WallReplyEditObject, client *api.VK) error

type WallReplyEditMiddleware = func(obj object.WallReplyEditObject, client *api.VK) (bool, error)

type WallReplyRestoreHandler = func(obj object.WallReplyRestoreObject, client *api.VK) error

type WallReplyRestoreMiddleware = func(obj object.WallReplyRestoreObject, client *api.VK) (bool, error)

type WallReplyDeleteHandler = func(obj object.WallReplyDeleteObject, client *api.VK) error

type WallReplyDeleteMiddleware = func(obj object.WallReplyDeleteObject, client *api.VK) (bool, error)

type BoardPostNewHandler = func(obj object.BoardPostNewObject, client *api.VK) error

type BoardPostNewMiddleware = func(obj object.BoardPostNewObject, client *api.VK) (bool, error)

type BoardPostEditHandler = func(obj object.BoardPostEditObject, client *api.VK) error

type BoardPostEditMiddleware = func(obj object.BoardPostEditObject, client *api.VK) (bool, error)

type BoardPostRestoreHandler = func(obj object.BoardPostRestoreObject, client *api.VK) error

type BoardPostRestoreMiddleware = func(obj object.BoardPostRestoreObject, client *api.VK) (bool, error)

type BoardPostDeleteHandler = func(obj object.BoardPostDeleteObject, client *api.VK) error

type BoardPostDeleteMiddleware = func(obj object.BoardPostDeleteObject, client *api.VK) (bool, error)

type MarketCommentNewHandler = func(obj object.MarketCommentNewObject, client *api.VK) error

type MarketCommentNewMiddleware = func(obj object.MarketCommentNewObject, client *api.VK) (bool, error)

type MarketCommentEditHandler = func(obj object.MarketCommentEditObject, client *api.VK) error

type MarketCommentEditMiddleware = func(obj object.MarketCommentEditObject, client *api.VK) (bool, error)

type MarketCommentRestoreHandler = func(obj object.MarketCommentRestoreObject, client *api.VK) error

type MarketCommentRestoreMiddleware = func(obj object.MarketCommentRestoreObject, client *api.VK) (bool, error)

type MarketCommentDeleteHandler = func(obj object.MarketCommentDeleteObject, client *api.VK) error

type MarketCommentDeleteMiddleware = func(obj object.MarketCommentDeleteObject, client *api.VK) (bool, error)

type GroupLeaveHandler = func(obj object.GroupLeaveObject, client *api.VK) error

type GroupLeaveMiddleware = func(obj object.GroupLeaveObject, client *api.VK) (bool, error)

type GroupJoinHandler = func(obj object.GroupJoinObject, client *api.VK) error

type GroupJoinMiddleware = func(obj object.GroupJoinObject, client *api.VK) (bool, error)

type UserBlockHandler = func(obj object.UserBlockObject, client *api.VK) error

type UserBlockMiddleware = func(obj object.UserBlockObject, client *api.VK) (bool, error)

type UserUnblockHandler = func(obj object.UserUnblockObject, client *api.VK) error

type UserUnblockMiddleware = func(obj object.UserUnblockObject, client *api.VK) (bool, error)

type PollVoteNewHandler = func(obj object.PollVoteNewObject, client *api.VK) error

type PollVoteNewMiddleware = func(obj object.PollVoteNewObject, client *api.VK) (bool, error)

type GroupOfficersEditHandler = func(obj object.GroupOfficersEditObject, client *api.VK) error

type GroupOfficersEditMiddleware = func(obj object.GroupOfficersEditObject, client *api.VK) (bool, error)

type GroupChangeSettingsHandler = func(obj object.GroupChangeSettingsObject, client *api.VK) error

type GroupChangeSettingsMiddleware = func(obj object.GroupChangeSettingsObject, client *api.VK) (bool, error)

type GroupChangePhotoHandler = func(obj object.GroupChangePhotoObject, client *api.VK) error

type GroupChangePhotoMiddleware = func(obj object.GroupChangePhotoObject, client *api.VK) (bool, error)

type VkpayTransactionHandler = func(obj object.VkpayTransactionObject, client *api.VK) error

type VkpayTransactionMiddleware = func(obj object.VkpayTransactionObject, client *api.VK) (bool, error)

type LeadFormsNewHandler = func(obj object.LeadFormsNewObject, client *api.VK) error

type LeadFormsNewMiddleware = func(obj object.LeadFormsNewObject, client *api.VK) (bool, error)

type AppPayloadHandler = func(obj object.AppPayloadObject, client *api.VK) error

type AppPayloadMiddleware = func(obj object.AppPayloadObject, client *api.VK) (bool, error)

type MessageReadHandler = func(obj object.MessageReadObject, client *api.VK) error

type MessageReadMiddleware = func(obj object.MessageReadObject, client *api.VK) (bool, error)

type LikeAddHandler = func(obj object.LikeAddObject, client *api.VK) error

type LikeAddMiddleware = func(obj object.LikeAddObject, client *api.VK) (bool, error)

type LikeRemoveHandler = func(obj object.LikeRemoveObject, client *api.VK) error

type LikeRemoveMiddleware = func(obj object.LikeRemoveObject, client *api.VK) (bool, error)

type Watcher struct {
	client  *api.VK
	lp      *longpoll.Longpoll
	chanErr chan error
}

func (w *Watcher) WatchMessageNew(handler MessageNewHandler, middlewares ...MessageNewMiddleware) *Watcher {
	w.lp.MessageNew(func(obj object.MessageNewObject, i int) {
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
	w.lp.MessageReply(func(obj object.MessageReplyObject, i int) {
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
	w.lp.MessageEdit(func(obj object.MessageEditObject, i int) {
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
	w.lp.MessageAllow(func(obj object.MessageAllowObject, i int) {
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
	w.lp.MessageDeny(func(obj object.MessageDenyObject, i int) {
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
	w.lp.MessageTypingState(func(obj object.MessageTypingStateObject, i int) {
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
	w.lp.PhotoNew(func(obj object.PhotoNewObject, i int) {
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
	w.lp.PhotoCommentNew(func(obj object.PhotoCommentNewObject, i int) {
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
	w.lp.PhotoCommentEdit(func(obj object.PhotoCommentEditObject, i int) {
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
	w.lp.PhotoCommentRestore(func(obj object.PhotoCommentRestoreObject, i int) {
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
	w.lp.PhotoCommentDelete(func(obj object.PhotoCommentDeleteObject, i int) {
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
	w.lp.AudioNew(func(obj object.AudioNewObject, i int) {
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
	w.lp.VideoNew(func(obj object.VideoNewObject, i int) {
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
	w.lp.VideoCommentNew(func(obj object.VideoCommentNewObject, i int) {
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
	w.lp.VideoCommentEdit(func(obj object.VideoCommentEditObject, i int) {
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
	w.lp.VideoCommentRestore(func(obj object.VideoCommentRestoreObject, i int) {
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
	w.lp.VideoCommentDelete(func(obj object.VideoCommentDeleteObject, i int) {
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
	w.lp.WallPostNew(func(obj object.WallPostNewObject, i int) {
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
	w.lp.WallRepost(func(obj object.WallRepostObject, i int) {
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
	w.lp.WallReplyNew(func(obj object.WallReplyNewObject, i int) {
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
	w.lp.WallReplyEdit(func(obj object.WallReplyEditObject, i int) {
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
	w.lp.WallReplyRestore(func(obj object.WallReplyRestoreObject, i int) {
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
	w.lp.WallReplyDelete(func(obj object.WallReplyDeleteObject, i int) {
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
	w.lp.BoardPostNew(func(obj object.BoardPostNewObject, i int) {
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
	w.lp.BoardPostEdit(func(obj object.BoardPostEditObject, i int) {
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
	w.lp.BoardPostRestore(func(obj object.BoardPostRestoreObject, i int) {
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
	w.lp.BoardPostDelete(func(obj object.BoardPostDeleteObject, i int) {
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
	w.lp.MarketCommentNew(func(obj object.MarketCommentNewObject, i int) {
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
	w.lp.MarketCommentEdit(func(obj object.MarketCommentEditObject, i int) {
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
	w.lp.MarketCommentRestore(func(obj object.MarketCommentRestoreObject, i int) {
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
	w.lp.MarketCommentDelete(func(obj object.MarketCommentDeleteObject, i int) {
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
	w.lp.GroupLeave(func(obj object.GroupLeaveObject, i int) {
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
	w.lp.GroupJoin(func(obj object.GroupJoinObject, i int) {
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
	w.lp.UserBlock(func(obj object.UserBlockObject, i int) {
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
	w.lp.UserUnblock(func(obj object.UserUnblockObject, i int) {
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
	w.lp.PollVoteNew(func(obj object.PollVoteNewObject, i int) {
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
	w.lp.GroupOfficersEdit(func(obj object.GroupOfficersEditObject, i int) {
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
	w.lp.GroupChangeSettings(func(obj object.GroupChangeSettingsObject, i int) {
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
	w.lp.GroupChangePhoto(func(obj object.GroupChangePhotoObject, i int) {
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
	w.lp.VkpayTransaction(func(obj object.VkpayTransactionObject, i int) {
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
	w.lp.LeadFormsNew(func(obj object.LeadFormsNewObject, i int) {
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
	w.lp.AppPayload(func(obj object.AppPayloadObject, i int) {
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
	w.lp.MessageRead(func(obj object.MessageReadObject, i int) {
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
	w.lp.LikeAdd(func(obj object.LikeAddObject, i int) {
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
	w.lp.LikeRemove(func(obj object.LikeRemoveObject, i int) {
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

	w.lp.MessageNew(func(obj object.MessageNewObject, i int) {
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

	lp, err := longpoll.NewLongpollCommunity(client)
	if err != nil {
		chanErr <- err
		return nil, nil
	}
	lp.FullResponse(func(response object.LongpollBotResponse) {
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
