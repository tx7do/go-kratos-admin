package logging

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/mileusna/useragent"

	"github.com/tx7do/go-utils/timeutil"
	"github.com/tx7do/go-utils/trans"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	op := options{}
	for _, o := range opts {
		o(&op)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			startTime := time.Now()

			reply, err = handler(ctx, req)

			// 统计耗时
			latency := time.Since(startTime).Seconds()

			var operationLogData *adminV1.AdminOperationLog
			var loginLogData *adminV1.AdminLoginLog

			if tr, ok := transport.FromServerContext(ctx); ok {
				var htr *http.Transport
				if htr, ok = tr.(*http.Transport); ok {
					loginLogData = fillLoginLog(htr)
					operationLogData = fillOperationLog(htr)
				}
			}

			// 获取错误码和是否成功
			statusCode, reason, success := getStatusCode(err)

			if operationLogData != nil {
				operationLogData.CostTime = timeutil.Float64ToDurationpb(latency)
				operationLogData.StatusCode = trans.Ptr(statusCode)
				operationLogData.Reason = trans.Ptr(reason)
				operationLogData.Success = trans.Ptr(success)
			}
			if loginLogData != nil {
				loginLogData.StatusCode = trans.Ptr(statusCode)
				loginLogData.Reason = trans.Ptr(reason)
				loginLogData.Success = trans.Ptr(success)
			}

			if op.writeOperationLogFunc != nil {
				_ = op.writeOperationLogFunc(ctx, operationLogData)
			}
			if op.writeLoginLogFunc != nil {
				_ = op.writeLoginLogFunc(ctx, loginLogData)
			}

			return
		}
	}
}

// fillLoginLog 填充登录日志
func fillLoginLog(htr *http.Transport) *adminV1.AdminLoginLog {
	if htr.Operation() != adminV1.OperationAuthenticationServiceLogin {
		return nil
	}

	loginLogData := &adminV1.AdminLoginLog{}

	clientIp := getClientRealIP(htr.Request())

	loginLogData.LoginIp = trans.Ptr(clientIp)
	loginLogData.LoginTime = timeutil.TimeToTimestamppb(trans.Ptr(time.Now()))

	loginLogData.Location = trans.Ptr(clientIpToLocation(clientIp))

	if username, err := BindLoginRequest(htr.Request()); err == nil {
		loginLogData.UserName = trans.Ptr(username)
	}

	// 获取客户端ID
	loginLogData.ClientId = trans.Ptr(getClientID(htr.Request(), nil))

	// 用户代理信息
	strUserAgent := htr.RequestHeader().Get(HeaderKeyUserAgent)
	ua := useragent.Parse(strUserAgent)

	var deviceName string
	if ua.Device != "" {
		deviceName = ua.Device
	} else {
		if ua.Desktop {
			deviceName = "PC"
		}
	}

	loginLogData.UserAgent = trans.Ptr(ua.String)
	loginLogData.BrowserVersion = trans.Ptr(ua.Version)
	loginLogData.BrowserName = trans.Ptr(ua.Name)
	loginLogData.OsName = trans.Ptr(ua.OS)
	loginLogData.OsVersion = trans.Ptr(ua.OSVersion)
	loginLogData.ClientName = trans.Ptr(deviceName)

	return loginLogData
}

// fillOperationLog 填充操作日志
func fillOperationLog(htr *http.Transport) *adminV1.AdminOperationLog {
	if htr.Operation() == adminV1.OperationAuthenticationServiceLogin {
		return nil
	}

	operationLogData := &adminV1.AdminOperationLog{}

	clientIp := getClientRealIP(htr.Request())
	referer, _ := url.QueryUnescape(htr.RequestHeader().Get(HeaderKeyReferer))
	requestUri, _ := url.QueryUnescape(htr.Request().RequestURI)
	bodyBytes, _ := io.ReadAll(htr.Request().Body)

	operationLogData.Method = trans.Ptr(htr.Request().Method)
	operationLogData.Operation = trans.Ptr(htr.Operation())
	operationLogData.Path = trans.Ptr(htr.PathTemplate())
	operationLogData.Referer = trans.Ptr(referer)
	operationLogData.ClientIp = trans.Ptr(clientIp)
	operationLogData.RequestId = trans.Ptr(getRequestId(htr.Request()))
	operationLogData.RequestUri = trans.Ptr(requestUri)
	operationLogData.RequestBody = trans.Ptr(string(bodyBytes))
	operationLogData.Location = trans.Ptr(clientIpToLocation(clientIp))

	ut := extractAuthToken(htr)
	if ut != nil {
		operationLogData.UserId = trans.Ptr(ut.UserId)
		operationLogData.UserName = trans.Ptr(ut.Username)
	}

	// 获取客户端ID
	operationLogData.ClientId = trans.Ptr(getClientID(htr.Request(), ut))

	// 用户代理信息
	strUserAgent := htr.RequestHeader().Get(HeaderKeyUserAgent)
	ua := useragent.Parse(strUserAgent)

	var deviceName string
	if ua.Device != "" {
		deviceName = ua.Device
	} else {
		if ua.Desktop {
			deviceName = "PC"
		}
	}

	operationLogData.UserAgent = trans.Ptr(ua.String)
	operationLogData.BrowserVersion = trans.Ptr(ua.Version)
	operationLogData.BrowserName = trans.Ptr(ua.Name)
	operationLogData.OsName = trans.Ptr(ua.OS)
	operationLogData.OsVersion = trans.Ptr(ua.OSVersion)
	operationLogData.ClientName = trans.Ptr(deviceName)

	return operationLogData
}
