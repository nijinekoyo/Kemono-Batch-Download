/*
 * @Author: nijineko
 * @Date: 2024-08-26 19:57:25
 * @LastEditTime: 2024-08-27 19:35:43
 * @LastEditors: nijineko
 * @Description: Kemono API封装
 * @FilePath: \kemonoDownload\internal\kemono\api.go
 */
package kemono

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Host    = "https://kemono.su" // Kemono Host地址
	APIPath = Host + "/api/v1"    // Kemono API地址
)

// 文章文件类型
type file struct {
	Name string `json:"name"` // 文件名
	Path string `json:"path"` // 文件路径
}

// 创作者文章列表类型
type CreatorPost struct {
	ID          string         `json:"id"`          // 文章ID
	User        string         `json:"user"`        // 用户ID
	Service     string         `json:"service"`     // 服务
	Title       string         `json:"title"`       // 标题
	Content     string         `json:"content"`     // 内容
	Embed       map[string]any `json:"embed"`       // 嵌入内容
	Shared_file bool           `json:"shared_file"` // 是否是共享文件
	Added       string         `json:"added"`       // 添加时间
	Published   string         `json:"published"`   // 发布时间
	Edited      string         `json:"edited"`      // 编辑时间
	File        file           `json:"file"`        // 文件信息
	Attachments []file         `json:"attachments"` // 附件
}

/**
 * @description: 组成GET请求参数
 * @param {map[string]string} Parameters 参数
 * @return {string} 参数拼接
 */
func getParameterSplice(Parameters map[string]string) string {
	if len(Parameters) == 0 {
		return ""
	}

	var URL string
	for Key, Parameter := range Parameters {
		if Parameter == "" {
			continue
		}

		if URL == "" {
			URL += "?"
		} else {
			URL += "&"
		}

		URL += fmt.Sprintf("%s=%s", Key, Parameter)
	}

	return URL
}

/**
 * @description: 获取创作者文章列表
 * @param {string} Service 服务
 * @param {string} User 用户ID
 * @param {string} Query 查询
 * @param {int} Offset 偏移
 * @return {[]creatorPost} 文章列表
 * @return {error} 错误
 */
func GetCreatorPosts(Service string, User string, Query string, Offset int) ([]CreatorPost, error) {
	URL := APIPath + fmt.Sprintf("/%s/user/%s", Service, User)

	// 偏移量转换为String
	OffsetString := fmt.Sprintf("%d", Offset)

	// 添加GET参数
	GetParameters := map[string]string{
		"q": Query,
		"o": OffsetString,
	}
	URL += getParameterSplice(GetParameters)

	// 发起请求
	Response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	// 读取返回
	var CreatorPosts []CreatorPost
	err = json.NewDecoder(Response.Body).Decode(&CreatorPosts)
	if err != nil {
		return nil, err
	}

	return CreatorPosts, nil
}
