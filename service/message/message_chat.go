package message

import (
	"Douyin/common"
	"Douyin/repository"
	"time"
)

func GetMessageList(userId uint, toUserId uint, preTime int64) ([]common.Message, error) {
	// 检查用户id是否存在
	err := check(userId, toUserId)
	if err != nil {
		return nil, err
	}
	//2. 查询聊天记录
	messageList, err := repository.GetMessageList(userId, toUserId, time.Unix(preTime/1000, preTime%1000*1e6))
	if err != nil {
		return nil, err
	}
	//3.组装成合适的格式并返回
	var retList []common.Message
	for i := 0; i < len(messageList); i++ {
		x := messageList[i]
		message := common.Message{
			Id:         x.ID,
			ToUserId:   x.ToUserId,
			FromUserId: x.UserId,
			Content:    x.Content,
			CreateTime: x.CreatedAt.UnixNano() / 1e6,
		}
		retList = append(retList, message)
	}
	return retList, nil
}
