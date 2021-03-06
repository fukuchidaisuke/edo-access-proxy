// Copyright 2015 realglobe, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package session

import (
	"time"
)

// TA 間連携プロトコルのセッション。

type Element struct {
	id string
	// 有効期限。
	exp time.Time
	// 転送先 TA の ID。
	toTa string
	// アカウントタグからアカウント情報へのマップ。
	acnts map[string]*Account
}

func New(id string, exp time.Time, toTa string, acnts map[string]*Account) *Element {
	return &Element{
		id,
		exp,
		toTa,
		acnts,
	}
}

// ID を返す。
func (this *Element) Id() string {
	return this.id
}

// 有効期限を返す。
func (this *Element) Expires() time.Time {
	return this.exp
}

// 転送先 TA の ID を返す。
func (this *Element) ToTa() string {
	return this.toTa
}

// アカウントタグからアカウント情報へのマップを返す。
func (this *Element) Accounts() map[string]*Account {
	return this.acnts
}
