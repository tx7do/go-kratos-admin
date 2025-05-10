package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIPFromRemoteAddr(t *testing.T) {
	assert.Equal(t, "::1", getIPFromRemoteAddr("[::1]:7788"))
	assert.Equal(t, "127.0.0.1", getIPFromRemoteAddr("127.0.0.1:7788"))
	assert.Equal(t, "127.0.0.1", getIPFromRemoteAddr("127.0.0.1"))
	assert.Equal(t, "::1", getIPFromRemoteAddr("::1"))
	assert.Equal(t, "127.0.0.1", getIPFromRemoteAddr("127.0.0.1:12456"))
	assert.Equal(t, "192.0.2.1", getIPFromRemoteAddr("192.0.2.1:5566"))
	assert.Equal(t, "2001:db8::68", getIPFromRemoteAddr("2001:db8::68"))

	assert.Equal(t, "", getIPFromRemoteAddr("192.0.2"))
}
