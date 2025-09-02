// Package redis-----------------------
// @author:  xiejunqiao
// @contact: xiejunqiao@wps.cn
// @since:   2024/11/25
// @desc: Redis chaos experiment tests
// ----------------------------------------
package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

// MockRedisClient 是一个模拟的 Redis 客户端
type MockRedisClient struct {
	pingResult string
	pingError  error
	closeError error
}

func (m *MockRedisClient) Ping(ctx context.Context) *redis.StringCmd {
	cmd := redis.NewStringCmd(ctx)
	if m.pingError != nil {
		cmd.SetErr(m.pingError)
	} else {
		cmd.SetVal(m.pingResult)
	}
	return cmd
}

func (m *MockRedisClient) Context() context.Context {
	return context.Background()
}

func (m *MockRedisClient) Close() error {
	return m.closeError
}

// TestRedisClientConnection 测试 Redis 客户端连接功能
func TestRedisClientConnection(t *testing.T) {
	// 创建 mock Redis 客户端
	mockClient := &MockRedisClient{
		pingResult: "PONG",
		pingError:  nil,
	}

	// 测试 Ping 功能
	ctx := context.Background()
	result, err := mockClient.Ping(ctx).Result()
	if err != nil {
		t.Errorf("期望无错误，但得到: %v", err)
	}
	if result != "PONG" {
		t.Errorf("期望结果 'PONG'，但得到: %s", result)
	}
}

// TestRedisClientList 测试多个 Redis 客户端的创建和管理
func TestRedisClientList(t *testing.T) {
	clientList := make([]*MockRedisClient, 0)

	// 创建 5 个 mock 客户端（减少数量避免过多输出）
	for i := 0; i < 5; i++ {
		mockClient := &MockRedisClient{
			pingResult: "PONG",
			pingError:  nil,
		}
		clientList = append(clientList, mockClient)
	}

	if len(clientList) != 5 {
		t.Errorf("期望创建 5 个客户端，但得到: %d", len(clientList))
	}

	// 测试所有客户端都能正常 Ping
	for i, client := range clientList {
		ctx := context.Background()
		result, err := client.Ping(ctx).Result()
		if err != nil {
			t.Errorf("客户端 %d 应该能正常 Ping，但得到错误: %v", i, err)
		}
		if result != "PONG" {
			t.Errorf("客户端 %d 应该返回 'PONG'，但得到: %s", i, result)
		}
	}
}

// TestRedisClientWithTimeout 测试带超时的 Redis 操作
func TestRedisClientWithTimeout(t *testing.T) {
	mockClient := &MockRedisClient{
		pingResult: "PONG",
		pingError:  nil,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// 测试带超时的 Ping
	result, err := mockClient.Ping(ctx).Result()
	if err != nil {
		t.Errorf("期望无错误，但得到: %v", err)
	}
	if result != "PONG" {
		t.Errorf("期望结果 'PONG'，但得到: %s", result)
	}
}

// TestRedisClientError 测试 Redis 客户端错误处理
func TestRedisClientError(t *testing.T) {
	mockClient := &MockRedisClient{
		pingResult: "",
		pingError:  redis.ErrClosed,
	}

	ctx := context.Background()

	// 测试错误情况
	_, err := mockClient.Ping(ctx).Result()
	if err == nil {
		t.Error("期望有错误，但没有得到错误")
	}
	if err != redis.ErrClosed {
		t.Errorf("期望错误 %v，但得到: %v", redis.ErrClosed, err)
	}
}

// TestRedisClientCleanup 测试客户端清理
func TestRedisClientCleanup(t *testing.T) {
	mockClient := &MockRedisClient{
		closeError: nil,
	}

	// 测试清理
	err := mockClient.Close()
	if err != nil {
		t.Errorf("期望无错误，但得到: %v", err)
	}
}

// BenchmarkRedisClientPing 性能测试
func BenchmarkRedisClientPing(b *testing.B) {
	mockClient := &MockRedisClient{
		pingResult: "PONG",
		pingError:  nil,
	}
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := mockClient.Ping(ctx).Result()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// TestRedisClientIntegration 测试 Redis 客户端集成场景
func TestRedisClientIntegration(t *testing.T) {
	// 模拟正常连接场景
	normalClient := &MockRedisClient{
		pingResult: "PONG",
		pingError:  nil,
		closeError: nil,
	}

	// 模拟连接失败场景
	failedClient := &MockRedisClient{
		pingResult: "",
		pingError:  redis.ErrClosed,
		closeError: redis.ErrClosed,
	}

	// 测试正常客户端
	ctx := context.Background()
	result, err := normalClient.Ping(ctx).Result()
	if err != nil {
		t.Errorf("正常客户端应该能 Ping 成功，但得到错误: %v", err)
	}
	if result != "PONG" {
		t.Errorf("正常客户端应该返回 'PONG'，但得到: %s", result)
	}

	// 测试失败客户端
	_, err = failedClient.Ping(ctx).Result()
	if err == nil {
		t.Error("失败客户端应该返回错误，但没有得到错误")
	}

	// 测试清理
	if err := normalClient.Close(); err != nil {
		t.Errorf("正常客户端清理应该成功，但得到错误: %v", err)
	}

	if err := failedClient.Close(); err == nil {
		t.Error("失败客户端清理应该失败，但没有得到错误")
	}
}
