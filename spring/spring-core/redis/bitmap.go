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
	CommandBitCount = "BITCOUNT"
	CommandBitOp    = "BITOP"
	CommandBitPos   = "BITPOS"
	CommandGetBit   = "GETBIT"
	CommandSetBit   = "SETBIT"
)

type BitmapCommand struct {
	c Redis
}

func NewBitmapCommand(c Redis) *BitmapCommand {
	return &BitmapCommand{c: c}
}

// BitCount https://redis.io/commands/bitcount
// Command: BITCOUNT key [start end]
// Integer reply: The number of bits set to 1.
func (c *BitmapCommand) BitCount(ctx context.Context, key string, args ...interface{}) (int64, error) {
	args = append([]interface{}{key}, args...)
	return c.c.Int64(ctx, CommandBitCount, args...)
}

// BitOpAnd https://redis.io/commands/bitop
// Command: BITOP AND destkey srckey1 srckey2 srckey3 ... srckeyN
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *BitmapCommand) BitOpAnd(ctx context.Context, destKey string, keys ...string) (int64, error) {
	args := []interface{}{"AND", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return c.c.Int64(ctx, CommandBitOp, args...)
}

// BitOpOr https://redis.io/commands/bitop
// Command: BITOP OR destkey srckey1 srckey2 srckey3 ... srckeyN
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *BitmapCommand) BitOpOr(ctx context.Context, destKey string, keys ...string) (int64, error) {
	args := []interface{}{"OR", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return c.c.Int64(ctx, CommandBitOp, args...)
}

// BitOpXor https://redis.io/commands/bitop
// Command: BITOP XOR destkey srckey1 srckey2 srckey3 ... srckeyN
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *BitmapCommand) BitOpXor(ctx context.Context, destKey string, keys ...string) (int64, error) {
	args := []interface{}{"XOR", destKey}
	for _, key := range keys {
		args = append(args, key)
	}
	return c.c.Int64(ctx, CommandBitOp, args...)
}

// BitOpNot https://redis.io/commands/bitop
// Command: BITOP NOT destkey srckey
// Integer reply: The size of the string stored in the destination key,
// that is equal to the size of the longest input string.
func (c *BitmapCommand) BitOpNot(ctx context.Context, destKey string, key string) (int64, error) {
	args := []interface{}{"NOT", destKey, key}
	return c.c.Int64(ctx, CommandBitOp, args...)
}

// BitPos https://redis.io/commands/bitpos
// Command: BITPOS key bit [start [end]]
// Integer reply: The command returns the position of the first bit
// set to 1 or 0 according to the request.
func (c *BitmapCommand) BitPos(ctx context.Context, key string, bit int64, args ...interface{}) (int64, error) {
	args = append([]interface{}{key, bit}, args...)
	return c.c.Int64(ctx, CommandBitPos, args...)
}

// GetBit https://redis.io/commands/getbit
// Command: GETBIT key offset
// Integer reply: the bit value stored at offset.
func (c *BitmapCommand) GetBit(ctx context.Context, key string, offset int64) (int64, error) {
	args := []interface{}{key, offset}
	return c.c.Int64(ctx, CommandGetBit, args...)
}

// SetBit https://redis.io/commands/setbit
// Command: SETBIT key offset value
// Integer reply: the original bit value stored at offset.
func (c *BitmapCommand) SetBit(ctx context.Context, key string, offset int64, value int) (int64, error) {
	args := []interface{}{key, offset, value}
	return c.c.Int64(ctx, CommandSetBit, args...)
}
