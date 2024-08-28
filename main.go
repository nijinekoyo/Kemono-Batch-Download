/*
 * @Author: nijineko
 * @Date: 2024-08-26 19:54:53
 * @LastEditTime: 2024-08-28 09:13:32
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \kemonoDownload\main.go
 */
package main

import (
	"flag"
	"fmt"
	"kemonoDownload/internal/download"
	"kemonoDownload/internal/kemono"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	// 解析参数
	Service := flag.String("service", "", "服务")
	User := flag.String("user", "", "用户ID")
	Query := flag.String("query", "", "搜索关键词")
	SavePath := flag.String("save_path", "data/", "保存路径")
	FileNameFilter := flag.String("file_name_filter", "", "仅下载文件名包含的文件")
	ExtensionFilter := flag.String("extension_filter", "", "仅下载指定扩展名的文件")
	flag.Parse()

	if *Service == "" || *User == "" {
		fmt.Println("Usage: kemonoDownload -service <service> -user <user> [-query <query>] [-save_path <save_path>] [-file_name_filter <file_name_filter>] [-extension_filter <extension_filter>]")
		return
	}

	// 获取全部文章信息
	var CreatorPosts []kemono.CreatorPost
	for {
		Posts, err := kemono.GetCreatorPosts(*Service, *User, *Query, len(CreatorPosts))
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
			if *FileNameFilter != "" && !strings.Contains(Attachment.Name, *FileNameFilter) {
				// 文件名不包含过滤字符串则跳过
				continue
			}
			// 检查扩展名过滤
			if *ExtensionFilter != "" && (filepath.Ext(Attachment.Path) != *ExtensionFilter || filepath.Ext(Attachment.Name) != *ExtensionFilter) {
				// 扩展名不匹配则跳过
				continue
			}

			// 保存路径
			SavePath := path.Join(*SavePath, *Service, *User, CreatorPost.ID, Attachment.Name)
			// 检查文件是否存在
			if _, err := os.Stat(SavePath); err == nil {
				fmt.Println("File", Attachment.Name, "exists, skip")
				continue
			}

			// 下载文件
			Size, err := download.File(kemono.Host+Attachment.Path, SavePath)
			if err != nil {
				fmt.Println("Download", Attachment.Name, "failed:", err)
				continue
			}

			fmt.Println("Download", Attachment.Name, "success:", Size, "bytes")
		}
	}

	fmt.Println("Download success")
}
