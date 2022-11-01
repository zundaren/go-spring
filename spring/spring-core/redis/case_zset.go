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
	"testing"

	"github.com/go-spring/spring-base/assert"
)

func (c *Cases) ZAdd() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 1, "uno")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 2, "two", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(2))

			r4, err := c.ZRangeWithScores(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []ZItem{{"one", 1}, {"uno", 1}, {"two", 2}, {"three", 3}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 uno",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two 3 three",
				"Response": "\"2\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1 WITHSCORES",
				"Response": "\"one\",\"1\",\"uno\",\"1\",\"two\",\"2\",\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZCard() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZCard(ctx, "myzset")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(2))
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZCARD myzset",
				"Response": "\"2\""
			}]
		}`,
	}
}

func (c *Cases) ZCount() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZCount(ctx, "myzset", "-inf", "+inf")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(3))

			r5, err := c.ZCount(ctx, "myzset", "(1", "3")
			assert.Nil(t, err)
			assert.Equal(t, r5, int64(2))
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZCOUNT myzset -inf +inf",
				"Response": "\"3\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZCOUNT myzset (1 3",
				"Response": "\"2\""
			}]
		}`,
	}
}

func (c *Cases) ZDiff() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "zset1", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "zset1", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "zset1", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZAdd(ctx, "zset2", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(1))

			r5, err := c.ZAdd(ctx, "zset2", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r5, int64(1))

			r6, err := c.ZDiff(ctx, "zset1", "zset2")
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"three"})

			r7, err := c.ZDiffWithScores(ctx, "zset1", "zset2")
			assert.Nil(t, err)
			assert.Equal(t, r7, []ZItem{{"three", 3}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD zset1 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset1 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset1 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZDIFF 2 zset1 zset2",
				"Response": "\"three\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZDIFF 2 zset1 zset2 WITHSCORES",
				"Response": "\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZIncrBy() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZIncrBy(ctx, "myzset", 2, "one")
			assert.Nil(t, err)
			assert.Equal(t, r3, float64(3))

			r4, err := c.ZRangeWithScores(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []ZItem{{"two", 2}, {"one", 3}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZINCRBY myzset 2 one",
				"Response": "\"3\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1 WITHSCORES",
				"Response": "\"two\",\"2\",\"one\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZInter() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "zset1", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "zset1", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "zset2", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZAdd(ctx, "zset2", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(1))

			r5, err := c.ZAdd(ctx, "zset2", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r5, int64(1))

			r6, err := c.ZInter(ctx, 2, "zset1", "zset2")
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"one", "two"})

			r7, err := c.ZInterWithScores(ctx, 2, "zset1", "zset2")
			assert.Nil(t, err)
			assert.Equal(t, r7, []ZItem{{"one", 2}, {"two", 4}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD zset1 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset1 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZINTER 2 zset1 zset2",
				"Response": "\"one\",\"two\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZINTER 2 zset1 zset2 WITHSCORES",
				"Response": "\"one\",\"2\",\"two\",\"4\""
			}]
		}`,
	}
}

func (c *Cases) ZLexCount() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 0, "a", 0, "b", 0, "c", 0, "d", 0, "e")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(5))

			r2, err := c.ZAdd(ctx, "myzset", 0, "f", 0, "g")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(2))

			r3, err := c.ZLexCount(ctx, "myzset", "-", "+")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(7))

			r4, err := c.ZLexCount(ctx, "myzset", "[b", "[f")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(5))
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 0 a 0 b 0 c 0 d 0 e",
				"Response": "\"5\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 0 f 0 g",
				"Response": "\"2\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZLEXCOUNT myzset - +",
				"Response": "\"7\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZLEXCOUNT myzset [b [f",
				"Response": "\"5\""
			}]
		}`,
	}
}

func (c *Cases) ZMScore() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZMScore(ctx, "myzset", "one", "two", "nofield")
			assert.Nil(t, err)
			assert.Equal(t, r3, []float64{1, 2, 0})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZMSCORE myzset one two nofield",
				"Response": "\"1\",\"2\",NULL"
			}]
		}`,
	}
}

func (c *Cases) ZPopMax() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZPopMax(ctx, "myzset")
			assert.Nil(t, err)
			assert.Equal(t, r4, []ZItem{{"three", 3}})

			r5, err := c.ZPopMax(ctx, "nonexisting")
			assert.Equal(t, len(r5), 0)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZPOPMAX myzset",
				"Response": "\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZPopMaxN() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZPopMaxN(ctx, "myzset", 1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []ZItem{{"three", 3}})

			r5, err := c.ZPopMaxN(ctx, "nonexisting", 1)
			assert.Nil(t, err)
			assert.Nil(t, r5)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZPOPMAX myzset 1",
				"Response": "\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZPopMin() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZPopMin(ctx, "myzset")
			assert.Nil(t, err)
			assert.Equal(t, r4, []ZItem{{"one", 1}})

			r5, err := c.ZPopMin(ctx, "nonexisting")
			assert.Equal(t, len(r5), 0)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZPOPMIN myzset",
				"Response": "\"one\",\"1\""
			}]
		}`,
	}
}

func (c *Cases) ZPopMinN() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZPopMinN(ctx, "myzset", 1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []ZItem{{"one", 1}})

			r5, err := c.ZPopMinN(ctx, "nonexisting", 1)
			assert.Nil(t, err)
			assert.Nil(t, r5)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZPOPMIN myzset 1",
				"Response": "\"one\",\"1\""
			}]
		}`,
	}
}

func (c *Cases) ZRandMember() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "dadi", 1, "uno", 2, "due", 3, "tre", 4, "quattro", 5, "cinque", 6, "sei")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(6))

			r2, err := c.ZRandMember(ctx, "dadi")
			assert.Nil(t, err)
			assert.NotEqual(t, r2, "")

			r3, err := c.ZRandMember(ctx, "dadi")
			assert.Nil(t, err)
			assert.NotEqual(t, r3, "")

			r4, err := c.ZRandMemberWithScores(ctx, "dadi", -5)
			assert.Nil(t, err)
			assert.Equal(t, len(r4), 5)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD dadi 1 uno 2 due 3 tre 4 quattro 5 cinque 6 sei",
				"Response": "\"6\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANDMEMBER dadi",
				"Response": "\"sei\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANDMEMBER dadi",
				"Response": "\"sei\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANDMEMBER dadi -5 WITHSCORES",
				"Response": "\"uno\",\"1\",\"uno\",\"1\",\"cinque\",\"5\",\"sei\",\"6\",\"due\",\"2\""
			}]
		}`,
	}
}

func (c *Cases) ZRandMemberN() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "dadi", 1, "uno", 2, "due", 3, "tre", 4, "quattro", 5, "cinque", 6, "sei")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(6))

			r2, err := c.ZRandMemberN(ctx, "dadi", 1)
			assert.Nil(t, err)
			assert.SubInSlice(t, r2, []string{"uno", "due", "tre", "quattro", "cinque", "sei"})

			r3, err := c.ZRandMemberN(ctx, "dadi", 1)
			assert.Nil(t, err)
			assert.SubInSlice(t, r3, []string{"uno", "due", "tre", "quattro", "cinque", "sei"})

			r4, err := c.ZRandMemberWithScores(ctx, "dadi", -5)
			assert.Nil(t, err)
			assert.Equal(t, len(r4), 5)

			r5, err := c.ZRandMemberWithScores(ctx, "nonexisting", -5)
			assert.Nil(t, err)
			assert.Equal(t, len(r5), 0)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD dadi 1 uno 2 due 3 tre 4 quattro 5 cinque 6 sei",
				"Response": "\"6\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANDMEMBER dadi 1",
				"Response": "\"sei\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANDMEMBER dadi 1",
				"Response": "\"sei\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANDMEMBER dadi -5 WITHSCORES",
				"Response": "\"uno\",\"1\",\"uno\",\"1\",\"cinque\",\"5\",\"sei\",\"6\",\"due\",\"2\""
			}]
		}`,
	}
}

func (c *Cases) ZRange() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRange(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"one", "two", "three"})

			r5, err := c.ZRange(ctx, "myzset", 2, 3)
			assert.Nil(t, err)
			assert.Equal(t, r5, []string{"three"})

			r6, err := c.ZRange(ctx, "myzset", -2, -1)
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"two", "three"})

			r7, err := c.ZRangeWithScores(ctx, "myzset", 0, 1)
			assert.Nil(t, err)
			assert.Equal(t, r7, []ZItem{{"one", 1}, {"two", 2}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1",
				"Response": "\"one\",\"two\",\"three\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 2 3",
				"Response": "\"three\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset -2 -1",
				"Response": "\"two\",\"three\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 1 WITHSCORES",
				"Response": "\"one\",\"1\",\"two\",\"2\""
			}]
		}`,
	}
}

func (c *Cases) ZRangeByLex() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 0, "a", 0, "b", 0, "c", 0, "d", 0, "e", 0, "f", 0, "g")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(7))

			r2, err := c.ZRangeByLex(ctx, "myzset", "-", "[c")
			assert.Nil(t, err)
			assert.Equal(t, r2, []string{"a", "b", "c"})

			r3, err := c.ZRangeByLex(ctx, "myzset", "-", "(c")
			assert.Nil(t, err)
			assert.Equal(t, r3, []string{"a", "b"})

			r4, err := c.ZRangeByLex(ctx, "myzset", "[aaa", "(g")
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"b", "c", "d", "e", "f"})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 0 a 0 b 0 c 0 d 0 e 0 f 0 g",
				"Response": "\"7\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYLEX myzset - [c",
				"Response": "\"a\",\"b\",\"c\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYLEX myzset - (c",
				"Response": "\"a\",\"b\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYLEX myzset [aaa (g",
				"Response": "\"b\",\"c\",\"d\",\"e\",\"f\""
			}]
		}`,
	}
}

func (c *Cases) ZRangeByScore() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRangeByScore(ctx, "myzset", "-inf", "+inf")
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"one", "two", "three"})

			r5, err := c.ZRangeByScore(ctx, "myzset", "1", "2")
			assert.Nil(t, err)
			assert.Equal(t, r5, []string{"one", "two"})

			r6, err := c.ZRangeByScore(ctx, "myzset", "(1", "2")
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"two"})

			r7, err := c.ZRangeByScore(ctx, "myzset", "(1", "(2")
			assert.Nil(t, err)
			assert.Equal(t, len(r7), 0)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYSCORE myzset -inf +inf",
				"Response": "\"one\",\"two\",\"three\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYSCORE myzset 1 2",
				"Response": "\"one\",\"two\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYSCORE myzset (1 2",
				"Response": "\"two\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGEBYSCORE myzset (1 (2",
				"Response": ""
			}]
		}`,
	}
}

func (c *Cases) ZRank() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRank(ctx, "myzset", "three")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(2))

			_, err = c.ZRank(ctx, "myzset", "four")
			assert.True(t, IsErrNil(err))
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANK myzset three",
				"Response": "\"2\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANK myzset four",
				"Response": "NULL"
			}]
		}`,
	}
}

func (c *Cases) ZRem() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRem(ctx, "myzset", "two")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(1))

			r5, err := c.ZRangeWithScores(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r5, []ZItem{{"one", 1}, {"three", 3}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREM myzset two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1 WITHSCORES",
				"Response": "\"one\",\"1\",\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZRemRangeByLex() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 0, "aaaa", 0, "b", 0, "c", 0, "d", 0, "e")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(5))

			r2, err := c.ZAdd(ctx, "myzset", 0, "foo", 0, "zap", 0, "zip", 0, "ALPHA", 0, "alpha")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(5))

			r3, err := c.ZRange(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r3, []string{
				"ALPHA", "aaaa", "alpha", "b", "c", "d", "e", "foo", "zap", "zip",
			})

			r4, err := c.ZRemRangeByLex(ctx, "myzset", "[alpha", "[omega")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(6))

			r5, err := c.ZRange(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r5, []string{"ALPHA", "aaaa", "zap", "zip"})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 0 aaaa 0 b 0 c 0 d 0 e",
				"Response": "\"5\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 0 foo 0 zap 0 zip 0 ALPHA 0 alpha",
				"Response": "\"5\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1",
				"Response": "\"ALPHA\",\"aaaa\",\"alpha\",\"b\",\"c\",\"d\",\"e\",\"foo\",\"zap\",\"zip\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREMRANGEBYLEX myzset [alpha [omega",
				"Response": "\"6\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1",
				"Response": "\"ALPHA\",\"aaaa\",\"zap\",\"zip\""
			}]
		}`,
	}
}

func (c *Cases) ZRemRangeByRank() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRemRangeByRank(ctx, "myzset", 0, 1)
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(2))

			r5, err := c.ZRangeWithScores(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r5, []ZItem{{"three", 3}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREMRANGEBYRANK myzset 0 1",
				"Response": "\"2\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1 WITHSCORES",
				"Response": "\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZRemRangeByScore() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRemRangeByScore(ctx, "myzset", "-inf", "(2")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(1))

			r5, err := c.ZRangeWithScores(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r5, []ZItem{{"two", 2}, {"three", 3}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREMRANGEBYSCORE myzset -inf (2",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE myzset 0 -1 WITHSCORES",
				"Response": "\"two\",\"2\",\"three\",\"3\""
			}]
		}`,
	}
}

func (c *Cases) ZRevRange() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRevRange(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"three", "two", "one"})

			r5, err := c.ZRevRange(ctx, "myzset", 2, 3)
			assert.Nil(t, err)
			assert.Equal(t, r5, []string{"one"})

			r6, err := c.ZRevRange(ctx, "myzset", -2, -1)
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"two", "one"})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGE myzset 0 -1",
				"Response": "\"three\",\"two\",\"one\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGE myzset 2 3",
				"Response": "\"one\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGE myzset -2 -1",
				"Response": "\"two\",\"one\""
			}]
		}`,
	}
}

func (c *Cases) ZRevRangeWithScores() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRevRangeWithScores(ctx, "myzset", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"three", "3", "two", "2", "one", "1"})

			r5, err := c.ZRevRangeWithScores(ctx, "myzset", 2, 3)
			assert.Nil(t, err)
			assert.Equal(t, r5, []string{"one", "1"})

			r6, err := c.ZRevRangeWithScores(ctx, "myzset", -2, -1)
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"two", "2", "one", "1"})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGE myzset 0 -1",
				"Response": "\"three\",\"two\",\"one\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGE myzset 2 3",
				"Response": "\"one\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGE myzset -2 -1",
				"Response": "\"two\",\"one\""
			}]
		}`,
	}
}

func (c *Cases) ZRevRangeByLex() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 0, "a", 0, "b", 0, "c", 0, "d", 0, "e", 0, "f", 0, "g")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(7))

			r2, err := c.ZRevRangeByLex(ctx, "myzset", "[c", "-")
			assert.Nil(t, err)
			assert.Equal(t, r2, []string{"c", "b", "a"})

			r3, err := c.ZRevRangeByLex(ctx, "myzset", "(c", "-")
			assert.Nil(t, err)
			assert.Equal(t, r3, []string{"b", "a"})

			r4, err := c.ZRevRangeByLex(ctx, "myzset", "(g", "[aaa")
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"f", "e", "d", "c", "b"})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 0 a 0 b 0 c 0 d 0 e 0 f 0 g",
				"Response": "\"7\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYLEX myzset [c -",
				"Response": "\"c\",\"b\",\"a\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYLEX myzset (c -",
				"Response": "\"b\",\"a\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYLEX myzset (g [aaa",
				"Response": "\"f\",\"e\",\"d\",\"c\",\"b\""
			}]
		}`,
	}
}

func (c *Cases) ZRevRangeByScore() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRevRangeByScore(ctx, "myzset", "+inf", "-inf")
			assert.Nil(t, err)
			assert.Equal(t, r4, []string{"three", "two", "one"})

			r5, err := c.ZRevRangeByScore(ctx, "myzset", "2", "1")
			assert.Nil(t, err)
			assert.Equal(t, r5, []string{"two", "one"})

			r6, err := c.ZRevRangeByScore(ctx, "myzset", "2", "(1")
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"two"})

			r7, err := c.ZRevRangeByScore(ctx, "myzset", "(2", "(1")
			assert.Nil(t, err)
			assert.Equal(t, len(r7), 0)
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYSCORE myzset +inf -inf",
				"Response": "\"three\",\"two\",\"one\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYSCORE myzset 2 1",
				"Response": "\"two\",\"one\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYSCORE myzset 2 (1",
				"Response": "\"two\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANGEBYSCORE myzset (2 (1",
				"Response": ""
			}]
		}`,
	}
}

func (c *Cases) ZRevRank() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "myzset", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "myzset", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZRevRank(ctx, "myzset", "one")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(2))

			_, err = c.ZRevRank(ctx, "myzset", "four")
			assert.True(t, IsErrNil(err))
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD myzset 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANK myzset one",
				"Response": "\"2\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZREVRANK myzset four",
				"Response": "NULL"
			}]
		}`,
	}
}
func (c *Cases) ZScore() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "myzset", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZScore(ctx, "myzset", "one")
			assert.Nil(t, err)
			assert.Equal(t, r2, float64(1))
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD myzset 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZSCORE myzset one",
				"Response": "\"1\""
			}]
		}`,
	}
}

func (c *Cases) ZUnion() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "zset1", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "zset1", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "zset2", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZAdd(ctx, "zset2", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(1))

			r5, err := c.ZAdd(ctx, "zset2", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r5, int64(1))

			r6, err := c.ZUnion(ctx, 2, "zset1", "zset2")
			assert.Nil(t, err)
			assert.Equal(t, r6, []string{"one", "three", "two"})

			r7, err := c.ZUnionWithScores(ctx, 2, "zset1", "zset2")
			assert.Nil(t, err)
			assert.Equal(t, r7, []ZItem{{"one", 2}, {"three", 3}, {"two", 4}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD zset1 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset1 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZUNION 2 zset1 zset2",
				"Response": "\"one\",\"three\",\"two\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZUNION 2 zset1 zset2 WITHSCORES",
				"Response": "\"one\",\"2\",\"three\",\"3\",\"two\",\"4\""
			}]
		}`,
	}
}

func (c *Cases) ZUnionStore() *Case {
	return &Case{
		Func: func(t *testing.T, ctx context.Context, c *Client) {

			r1, err := c.ZAdd(ctx, "zset1", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r1, int64(1))

			r2, err := c.ZAdd(ctx, "zset1", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r2, int64(1))

			r3, err := c.ZAdd(ctx, "zset2", 1, "one")
			assert.Nil(t, err)
			assert.Equal(t, r3, int64(1))

			r4, err := c.ZAdd(ctx, "zset2", 2, "two")
			assert.Nil(t, err)
			assert.Equal(t, r4, int64(1))

			r5, err := c.ZAdd(ctx, "zset2", 3, "three")
			assert.Nil(t, err)
			assert.Equal(t, r5, int64(1))

			r6, err := c.ZUnionStore(ctx, "out", 2, "zset1", "zset2", "WEIGHTS", 2, 3)
			assert.Nil(t, err)
			assert.Equal(t, r6, int64(3))

			r7, err := c.ZRangeWithScores(ctx, "out", 0, -1)
			assert.Nil(t, err)
			assert.Equal(t, r7, []ZItem{{"one", 5}, {"three", 9}, {"two", 10}})
		},
		Data: `
		{
			"Session": "df3b64266ebe4e63a464e135000a07cd",
			"Actions": [{
				"Protocol": "REDIS",
				"Request": "ZADD zset1 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset1 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 1 one",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 2 two",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZADD zset2 3 three",
				"Response": "\"1\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZUNIONSTORE out 2 zset1 zset2 WEIGHTS 2 3",
				"Response": "\"3\""
			}, {
				"Protocol": "REDIS",
				"Request": "ZRANGE out 0 -1 WITHSCORES",
				"Response": "\"one\",\"5\",\"three\",\"9\",\"two\",\"10\""
			}]
		}`,
	}
}
