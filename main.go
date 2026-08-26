/*
 * @Author: nijineko
 * @Date: 2024-08-26 19:54:53
 * @LastEditTime: 2026-08-26 16:29:31
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \kemonoDownload\main.go
 */
package main

import (
	"fmt"
	"kemonoDownload/internal/download"
	"kemonoDownload/internal/flag"
	"kemonoDownload/internal/kemono"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	// 初始化参数解析
	flag.Init()

	if flag.Get().Service == "" || flag.Get().User == "" {
		fmt.Println("Usage: kemonoDownload -service <service> -user <user> [-query <query>] [-save_path <save_path>] [-file_name_filter <file_name_filter>] [-extension_filter <extension_filter>]")
		return
	}

	// 获取全部文章信息
	var CreatorPosts []kemono.CreatorPost
	for {
		Posts, err := kemono.GetCreatorPosts(flag.Get().Service, flag.Get().User, flag.Get().Query, len(CreatorPosts))
		if err != nil {
			panic(err)
		}

		// 追加到文章列表
		CreatorPosts = append(CreatorPosts, Posts...)

		if len(Posts) < 50 {
			// 如果文章数量小于50则退出
			break
		}
	}

	// 启动文件清理任务，程序退出时清理当前正在下载的文件
	go download.StartClear()

	// 下载文章附件
	for Index, CreatorPost := range CreatorPosts {
		fmt.Println("Download", Index+1, "of", len(CreatorPosts), ":", CreatorPost.Title)

		// 下载所有附件
		for _, Attachment := range CreatorPost.Attachments {
			// 检查文件名过滤
			if flag.Get().FileNameFilter != "" && !strings.Contains(Attachment.Name, flag.Get().FileNameFilter) {
				// 文件名不包含过滤字符串则跳过
				continue
			}
			// 检查扩展名过滤
			if flag.Get().ExtensionFilter != "" && (filepath.Ext(Attachment.Path) != flag.Get().ExtensionFilter || filepath.Ext(Attachment.Name) != flag.Get().ExtensionFilter) {
				// 扩展名不匹配则跳过
				continue
			}

			// 保存路径
			SavePath := path.Join(flag.Get().SavePath, flag.Get().Service, flag.Get().User, CreatorPost.ID, Attachment.Name)
			// 检查文件是否存在
			if _, err := os.Stat(SavePath); err == nil {
				fmt.Println("File", Attachment.Name, "exists, skip")
				continue
			}

			// 下载文件
			Size, err := download.File(flag.Get().FileServerHost+flag.Get().FileServerPathPrefix+Attachment.Path, SavePath)
			if err != nil {
				fmt.Println("Download", Attachment.Name, "failed:", err)
				continue
			}

			fmt.Println("Download", Attachment.Name, "success:", Size, "bytes")
		}
	}

	fmt.Println("Download success")
}
