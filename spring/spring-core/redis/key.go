/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package redis

import (
	"context"
)

const (
	CommandDel       = "DEL"
	CommandDump      = "DUMP"
	CommandExists    = "EXISTS"
	CommandExpire    = "EXPIRE"
	CommandExpireAt  = "EXPIREAT"
	CommandKeys      = "KEYS"
	CommandPersist   = "PERSIST"
	CommandPExpire   = "PEXPIRE"
	CommandPExpireAt = "PEXPIREAT"
	CommandPTTL      = "PTTL"
	CommandRandomKey = "RANDOMKEY"
	CommandRename    = "RENAME"
	CommandRenameNX  = "RENAMENX"
	CommandTouch     = "TOUCH"
	CommandTTL       = "TTL"
	CommandType      = "TYPE"
)

type KeyCommand interface {

	// Del https://redis.io/commands/del
	// Integer reply: The number of keys that were removed.
	Del(ctx context.Context, keys ...string) (int64, error)

	// Dump https://redis.io/commands/dump
	// Bulk string reply: the serialized value.
	// If key does not exist a nil bulk reply is returned.
	Dump(ctx context.Context, key string) (string, error)

	// Exists https://redis.io/commands/exists
	// Integer reply: 1 if the key exists, 0 if the key does not exist.
	Exists(ctx context.Context, key string) (bool, error)

	// Expire https://redis.io/commands/expire
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	Expire(ctx context.Context, key string, expire int64) (bool, error)

	// ExpireAt https://redis.io/commands/expireat
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	ExpireAt(ctx context.Context, key string, expireAt int64) (bool, error)

	// Keys https://redis.io/commands/keys
	// Array reply: list of keys matching pattern.
	Keys(ctx context.Context, pattern string) ([]string, error)

	// Persist https://redis.io/commands/persist
	// Integer reply: 1 if the timeout was removed, 0 if key does
	// not exist or does not have an associated timeout.
	Persist(ctx context.Context, key string) (bool, error)

	// PExpire https://redis.io/commands/pexpire
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	PExpire(ctx context.Context, key string, expire int64) (bool, error)

	// PExpireAt https://redis.io/commands/pexpireat
	// Integer reply: 1 if the timeout was set, 0 if the timeout was not set.
	PExpireAt(ctx context.Context, key string, expireAt int64) (bool, error)

	// PTTL https://redis.io/commands/pttl
	// Integer reply: TTL in milliseconds, -1 if the key exists but
	// has no associated expire, -2 if the key does not exist.
	PTTL(ctx context.Context, key string) (int64, error)

	// RandomKey https://redis.io/commands/randomkey
	// Bulk string reply: the random key, or nil when the database is empty.
	RandomKey(ctx context.Context) (string, error)

	// Rename https://redis.io/commands/rename
	// Simple string reply.
	Rename(ctx context.Context, key, newKey string) (string, error)

	// RenameNX https://redis.io/commands/renamenx
	// Integer reply: 1 if key was renamed to newKey, 0 if newKey already exists.
	RenameNX(ctx context.Context, key, newKey string) (bool, error)

	// Touch https://redis.io/commands/touch
	// Integer reply: The number of keys that were touched.
	Touch(ctx context.Context, keys ...string) (int64, error)

	// TTL https://redis.io/commands/ttl
	// Integer reply: TTL in seconds, -1 if the key exists but has no
	// associated expire, -2 if the key does not exist.
	TTL(ctx context.Context, key string) (int64, error)

	// Type https://redis.io/commands/type
	// Simple string reply: type of key, or none when key does not exist.
	Type(ctx context.Context, key string) (string, error)
}

func (c *BaseClient) Del(ctx context.Context, keys ...string) (int64, error) {
	args := []interface{}{CommandDel}
	for _, key := range keys {
		args = append(args, key)
	}
	return c.Int64(ctx, args...)
}

func (c *BaseClient) Dump(ctx context.Context, key string) (string, error) {
	args := []interface{}{CommandDump, key}
	return c.String(ctx, args...)
}

func (c *BaseClient) Exists(ctx context.Context, key string) (bool, error) {
	args := []interface{}{CommandExists, key}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) Expire(ctx context.Context, key string, expire int64) (bool, error) {
	args := []interface{}{CommandExpire, key, expire}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) ExpireAt(ctx context.Context, key string, expireAt int64) (bool, error) {
	args := []interface{}{CommandExpireAt, key, expireAt}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) Keys(ctx context.Context, pattern string) ([]string, error) {
	args := []interface{}{CommandKeys, pattern}
	return c.StringSlice(ctx, args...)
}

func (c *BaseClient) Persist(ctx context.Context, key string) (bool, error) {
	args := []interface{}{CommandPersist, key}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) PExpire(ctx context.Context, key string, expire int64) (bool, error) {
	args := []interface{}{CommandPExpire, key, expire}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) PExpireAt(ctx context.Context, key string, expireAt int64) (bool, error) {
	args := []interface{}{CommandPExpireAt, key, expireAt}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) PTTL(ctx context.Context, key string) (int64, error) {
	args := []interface{}{CommandPTTL, key}
	return c.Int64(ctx, args...)
}

func (c *BaseClient) RandomKey(ctx context.Context) (string, error) {
	args := []interface{}{CommandRandomKey}
	return c.String(ctx, args...)
}

func (c *BaseClient) Rename(ctx context.Context, key, newKey string) (string, error) {
	args := []interface{}{CommandRename, key, newKey}
	return c.String(ctx, args...)
}

func (c *BaseClient) RenameNX(ctx context.Context, key, newKey string) (bool, error) {
	args := []interface{}{CommandRenameNX, key, newKey}
	return c.Bool(ctx, args...)
}

func (c *BaseClient) Touch(ctx context.Context, keys ...string) (int64, error) {
	args := []interface{}{CommandTouch}
	for _, key := range keys {
		args = append(args, key)
	}
	return c.Int64(ctx, args...)
}

func (c *BaseClient) TTL(ctx context.Context, key string) (int64, error) {
	args := []interface{}{CommandTTL, key}
	return c.Int64(ctx, args...)
}

func (c *BaseClient) Type(ctx context.Context, key string) (string, error) {
	args := []interface{}{CommandType, key}
	return c.String(ctx, args...)
}
