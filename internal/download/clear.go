/*
 * @Author: nijineko
 * @Date: 2024-08-28 09:01:21
 * @LastEditTime: 2024-08-28 09:12:55
 * @LastEditors: nijineko
 * @Description: 清理文件
 * @FilePath: \kemonoDownload\internal\download\clear.go
 */
package download

import (
	"os"
	"os/signal"
	"syscall"
)

/**
 * @description: 启动文件清理任务
 */
func StartClear() {
	// 捕获强制终止信号
	SigintChan := make(chan os.Signal, 1)
	signal.Notify(SigintChan, syscall.SIGINT, syscall.SIGTERM)

	// 等待信号
	<-SigintChan

	// 清理当前正在下载的文件
	if DownloadingFilePath != "" && DownloadingFile != nil {
		// 关闭文件
		DownloadingFile.Close()

		// 检查文件是否存在
		if _, err := os.Stat(DownloadingFilePath); err == nil {
			// 删除文件
			err := os.Remove(DownloadingFilePath)
			if err != nil {
				panic(err)
			}
		}
	}

	// 退出程序
	os.Exit(0)
}
