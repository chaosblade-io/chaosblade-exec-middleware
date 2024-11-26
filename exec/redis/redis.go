/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package redis

import (
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

type RedisCommandSpec struct {
	spec.BaseExpModelCommandSpec
}

func (*RedisCommandSpec) Name() string {
	return "redis"
}

func (*RedisCommandSpec) ShortDesc() string {
	return "Redis experiment"
}

func (*RedisCommandSpec) LongDesc() string {
	return "Redis experiment"
}

func NewRedisCommandSpec() spec.ExpModelCommandSpec {
	return &RedisCommandSpec{
		spec.BaseExpModelCommandSpec{
			ExpActions: []spec.ExpActionCommandSpec{
				NewCacheExpireActionSpec(),
				NewCacheLimitActionSpec(),
				NewCacheHotKeyActionSpec(),
				NewClientsLimitActionSpec(),
			},
			ExpFlags: []spec.ExpFlagSpec{},
		},
	}
}
