package agora

import (
	"ConferenceSpace/data/ws"
	logger2 "ConferenceSpace/logger"
	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
)

func GenerateRtcToken(param *ws.RtcToken) (string, error) {
	appID := "8dfd92ec1e78400095e75b7916ed88ac"
	appCertificate := "4a765d97aff44df39bea9d2f7f3511ee"
	// AccessToken2 过期的时间，单位为秒
	// 当 AccessToken2 过期但权限未过期时，用户仍在频道里并且可以发流，不会触发 SDK 回调
	// 但一旦用户和频道断开连接，用户将无法使用该 Token 加入同一频道。请确保 AccessToken2 的过期时间晚于权限过期时间
	tokenExpireTimeInSeconds := uint32(40)
	// 权限过期的时间，单位为秒
	// 权限过期30秒前会触发 token-privilege-will-expire 回调
	// 权限过期时会触发 token-privilege-did-expire 回调
	// 为作演示，在此将过期时间设为 40 秒。你可以看到客户端自动更新 Token 的过程
	privilegeExpireTimeInSeconds := uint32(40)
	var role rtctokenbuilder.Role
	switch param.Role {
	case 1:
		role = rtctokenbuilder.RolePublisher
	case 2:
		role = rtctokenbuilder.RoleSubscriber
	}

	result, err := rtctokenbuilder.BuildTokenWithUid(appID, appCertificate, param.ChannelName, param.UidRtcInt, role,
		tokenExpireTimeInSeconds, privilegeExpireTimeInSeconds)
	if err != nil {
		return "", err
	}
	logger2.InfoF("GenerateRtcToken Token with uid: %s\n uid is %d\n ChannelName is %s\n Role is %d\n", result, param.UidRtcInt, param.ChannelName, role)
	return result, nil
}
