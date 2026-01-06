package cache

import (
	"sync"
	"time"

	"github.com/shaxiaozz/sangfor-ad-exporter/model"
)

type TokenCache struct {
	mu        sync.RWMutex
	token     *model.SangforAdLoginResp
	expiresAt time.Time
}

func NewTokenCache() *TokenCache {
	return &TokenCache{}
}

func (c *TokenCache) IsValid() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.token == nil {
		return false
	}
	return time.Now().Before(c.expiresAt)
}

func (c *TokenCache) Get(
	loginFunc func() (*model.SangforAdLoginResp, error),
) (*model.SangforAdLoginResp, error) {

	// 快速路径
	if c.IsValid() {
		c.mu.RLock()
		defer c.mu.RUnlock()
		return c.token, nil
	}

	// 慢路径（加写锁）
	c.mu.Lock()
	defer c.mu.Unlock()

	// Double Check（防止并发重复登录）
	if c.token != nil && time.Now().Before(c.expiresAt) {
		return c.token, nil
	}

	token, err := loginFunc()
	if err != nil {
		return nil, err
	}

	// 关键：用 expired_timestamp 计算 TTL
	expireAt := time.Unix(int64(token.ExpiredTimestamp), 0)

	// 安全兜底：提前 30 秒过期
	expireAt = expireAt.Add(-30 * time.Second)

	c.token = token
	c.expiresAt = expireAt

	return token, nil
}
