/*
 * Copyright (c) 2019. Qing Cheng Technology Co., Ltd.
 */

package alipay

import (
	"strings"
)

func aliEscape(data string) string {
	data = strings.ReplaceAll(data, `!`, `%21`)
	data = strings.ReplaceAll(data, `*`, `%2A`)
	data = strings.ReplaceAll(data, `'`, `%27`)
	data = strings.ReplaceAll(data, `(`, `%28`)
	data = strings.ReplaceAll(data, `)`, `%29`)
	data = strings.ReplaceAll(data, `;`, `%3B`)
	data = strings.ReplaceAll(data, `:`, `%3A`)
	data = strings.ReplaceAll(data, `@`, `%40`)
	data = strings.ReplaceAll(data, `&`, `%26`)
	data = strings.ReplaceAll(data, `=`, `%3D`)
	data = strings.ReplaceAll(data, `+`, `%2B`)
	data = strings.ReplaceAll(data, `$`, `%24`)
	data = strings.ReplaceAll(data, `,`, `%2C`)
	data = strings.ReplaceAll(data, `/`, `%2F`)
	data = strings.ReplaceAll(data, `?`, `%3F`)
	data = strings.ReplaceAll(data, `%`, `%25`)
	data = strings.ReplaceAll(data, `#`, `%23`)
	data = strings.ReplaceAll(data, `[`, `%5B`)
	data = strings.ReplaceAll(data, `]`, `%5D`)
	return data
}
