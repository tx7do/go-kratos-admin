package logging

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mileusna/useragent"
	"github.com/tx7do/go-utils/geoip/qqwry"
	authnEngine "github.com/tx7do/kratos-authn/engine"

	"kratos-admin/app/admin/service/internal/data"

	adminV1 "kratos-admin/api/gen/go/admin/service/v1"
)

var ipClient *qqwry.Client = qqwry.NewClient()

// extractAuthToken 从JWT Token中提取用户信息
func extractAuthToken(authToken string, authenticator authnEngine.Authenticator) *data.UserTokenPayload {
	if len(authToken) == 0 {
		return nil
	}

	jwtToken := strings.TrimPrefix(authToken, "Bearer ")
	authnClaims, _ := authenticator.AuthenticateToken(jwtToken)
	if authnClaims == nil {
		return nil
	}

	ut, _ := data.NewUserTokenPayloadWithClaims(authnClaims)
	if ut == nil {
		return nil
	}

	return ut
}

// getClientRealIP 获取客户端真实IP
func getClientRealIP(request *http.Request) string {
	if request == nil {
		return ""
	}

	// 先检查 X-Forwarded-For 头
	xff := request.Header.Get(HeaderKeyXForwardedFor)
	if xff != "" {
		return xff
	}

	// 接着检查反向代理的 X-Real-IP 头
	xri := request.Header.Get(HeaderKeyXRealIP)
	if xri != "" {
		return xri
	}

	return request.RemoteAddr
}

// getRequestId 获取请求ID
func getRequestId(request *http.Request) string {
	if request == nil {
		return ""
	}

	// 先检查 X-Request-ID 头
	// 这是比较常见的用于标识请求的自定义头部字段。
	// 例如，在一个微服务架构的系统中，当一个请求从前端应用发送到后端的多个微服务时，
	// 每个微服务都可以在 X-Request-ID 字段中获取到相同的请求标识，从而方便追踪请求在各个服务节点中的处理情况。
	xri := request.Header.Get(HeaderKeyXRequestID)
	if xri != "" {
		return xri
	}

	// 接着检查 X-Correlation-ID 头
	// 它和 X-Request-ID 类似，用于关联一系列相关的请求或者事务。
	// 比如，在一个包含多个子请求的复杂业务流程中，X-Correlation-ID 可以用于跟踪整个业务流程中各个子请求之间的关系。
	xci := request.Header.Get(HeaderKeyXCorrelationID)
	if xci != "" {
		return xci
	}

	// 函数计算的请求ID
	xfcri := request.Header.Get(HeaderKeyXFcRequestID)
	if xfcri != "" {
		return xfcri
	}

	return ""
}

// getClientID 获取客户端ID
func getClientID(request *http.Request, userToken *data.UserTokenPayload) string {
	if request == nil {
		return ""
	}

	// 我们可以自定义一个Header叫做：X-Client-ID。
	xci := request.Header.Get(HeaderKeyXClientIP)
	if xci != "" {
		return xci
	}

	// 从JWT Token中获取ClientID也是可行的。
	if userToken != nil {
		return userToken.ClientId
	}

	return ""
}

// getStatusCode 状态码
func getStatusCode(err error) (int32, string, bool) {
	// 1. 信息响应 (100–199)
	// 2. 成功响应 (200–299)
	// 3. 重定向消息 (300–399)
	// 4. 客户端错误响应 (400–499)
	// 5. 服务端错误响应 (500–599)
	if se := errors.FromError(err); se != nil {
		return se.Code, se.Reason, se.Code < 400
	} else {
		return 200, "", true
	}
}

func PrintUserAgent(strUserAgent string) {
	ua := useragent.Parse(strUserAgent)

	fmt.Println("User-Agent", ua)
	fmt.Println()
	fmt.Println(ua.String)
	fmt.Println(strings.Repeat("=", len(ua.String)))
	fmt.Println("Name:", ua.Name, "v", ua.Version)
	fmt.Println("OS:", ua.OS, "v", ua.OSVersion)
	fmt.Println("Device:", ua.Device)
	if ua.Mobile {
		fmt.Println("(Mobile)")
	}
	if ua.Tablet {
		fmt.Println("(Tablet)")
	}
	if ua.Desktop {
		fmt.Println("(Desktop)")
	}
	if ua.Bot {
		fmt.Println("(Bot)")
	}
	if ua.URL != "" {
		fmt.Println(ua.URL)
	}
}

func BindLoginRequest(r *http.Request, v *adminV1.LoginRequest) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("BindQuery ReadAll", err)
		return err
	}
	defer r.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println("BindQuery ParseQuery", err)
		return err
	}

	v.Username = values.Get("username")

	return nil
}

// clientIpToLocation 获取客户端IP的地理位置
func clientIpToLocation(ip string) string {
	res, err := ipClient.Query(ip)
	if err != nil {
		return ""
	}
	return res.City
}
