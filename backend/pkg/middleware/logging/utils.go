package logging

import (
	"fmt"
	"io"
	"net"
	"strings"

	"encoding/json"
	"net/url"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/mileusna/useragent"
	"github.com/tx7do/go-utils/geoip/geolite"
	"github.com/tx7do/go-utils/jwtutil"

	authenticationV1 "kratos-admin/api/gen/go/authentication/service/v1"

	"kratos-admin/pkg/jwt"
)

var ipClient, _ = geolite.NewClient()

// extractAuthToken 从JWT Token中提取用户信息
func extractAuthToken(htr *http.Transport) *authenticationV1.UserTokenPayload {
	authToken := htr.RequestHeader().Get(HeaderKeyAuthorization)
	if len(authToken) == 0 {
		return nil
	}

	jwtToken := strings.TrimPrefix(authToken, "Bearer ")

	claims, err := jwtutil.ParseJWTPayload(jwtToken)
	if err != nil {
		log.Errorf("extractAuthToken ParseJWTPayload failed: %v", err)
		return nil
	}

	ut, err := jwt.NewUserTokenPayloadWithJwtMapClaims(claims)
	if err != nil {
		log.Errorf("extractAuthToken NewUserTokenPayloadWithJwtMapClaims failed: %v", err)
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
	// 由于它可以记录整个代理链中的IP地址，因此适用于多级代理的情况。
	// 当请求经过多个代理服務器时，X-Forwarded-For字段可以完整地记录原始请求的客户端IP地址和所有代理服務器的IP地址。
	// 需要注意：
	// 最外层Nginx配置为：proxy_set_header X-Forwarded-For $remote_addr; 如此做可以覆写掉ip。以防止ip伪造。
	// 里层Nginx配置为：proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
	xff := request.Header.Get(HeaderKeyXForwardedFor)
	if xff != "" {
		// X-Forwarded-For字段的值是一个逗号分隔的IP地址列表，
		// 一般来说，第一个IP地址是原始请求的客户端IP地址（当然，它可以被伪造）。
		ips := strings.Split(xff, ",")

		for _, ip := range ips {
			// 去除空格
			ip = strings.TrimSpace(ip)
			// 检查是否是合法的IP地址
			if net.ParseIP(ip) != nil {
				return ip
			}
		}
	}

	// 接着检查反向代理的 X-Real-IP 头
	// 通常只在反向代理服務器中使用，并且只记录原始请求的客户端IP地址。
	// 它不适用于多级代理的情况，因为每经过一个代理服務器，X-Real-IP字段的值都会被覆盖为最新的客户端IP地址。
	xri := request.Header.Get(HeaderKeyXRealIP)
	if xri != "" {
		if net.ParseIP(xri) != nil {
			return xri
		}
	}

	return getIPFromRemoteAddr(request.RemoteAddr)
}

func getIPFromRemoteAddr(hostAddress string) string {
	// Check if the host address contains a port
	if strings.Contains(hostAddress, ":") {
		// Attempt to split the host address into host and port
		host, _, err := net.SplitHostPort(strings.TrimSpace(hostAddress))
		if err == nil {
			// Validate the host as an IP address
			if net.ParseIP(host) != nil {
				return host
			}
		}
	}
	// Validate the host address as an IP address
	if net.ParseIP(hostAddress) != nil {
		return hostAddress
	}
	return ""
}

// getRequestId 获取请求ID
func getRequestId(request *http.Request) string {
	if request == nil {
		return ""
	}

	// 先检查 X-Request-ID 头
	// 这是比较常见的用于标识请求的自定义头部字段。
	// 例如，在一个微服務架构的系统中，当一个请求从前端应用发送到后端的多个微服務时，
	// 每个微服務都可以在 X-Request-ID 字段中获取到相同的请求标识，从而方便追踪请求在各个服務节点中的处理情况。
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
func getClientID(request *http.Request, userToken *authenticationV1.UserTokenPayload) string {
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
		return userToken.GetClientId()
	}

	return ""
}

// getStatusCode 状态码
func getStatusCode(err error) (int32, string, bool) {
	// 1. 信息响应 (100–199)
	// 2. 成功响应 (200–299)
	// 3. 重定向消息 (300–399)
	// 4. 客户端错误响应 (400–499)
	// 5. 服務端错误响应 (500–599)
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

func BindLoginRequest(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("BindQuery ReadAll", err)
		return "", err
	}
	defer r.Body.Close()

	var loginRequest authenticationV1.LoginRequest
	if err = json.Unmarshal(body, &loginRequest); err == nil {
		//fmt.Println("BindLoginRequest Unmarshal JSON failed", err)
		return loginRequest.GetUsername(), nil
	}

	if values, err := url.ParseQuery(string(body)); err == nil {
		//fmt.Println("BindLoginRequest Unmarshal Query", err)
		return values.Get("username"), nil
	}

	return "", err
}

// clientIpToLocation 获取客户端IP的地理位置
func clientIpToLocation(ip string) string {
	res, err := ipClient.Query(ip)
	if err != nil {
		return ""
	}
	return res.City
}
