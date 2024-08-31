/*
 * @Author: nijineko
 * @Date: 2024-08-27 19:41:02
 * @LastEditTime: 2024-09-01 05:30:29
 * @LastEditors: nijineko
 * @Description: 下载文件封装
 * @FilePath: \kemonoDownload\internal\download\file.go
 */
package download

import (
	"io"
	"os"
	"path"

	"github.com/HyacinthusAcademy/yuzuhttp"
	"github.com/schollz/progressbar/v3"
)

var (
	DownloadingFilePath string
	DownloadingFile     *os.File
)

/**
 * @description: 下载文件
 * @param {string} URL 文件地址
 * @param {string} SavePath 文件保存路径
 * @return {int} 文件大小
 * @return {error} 错误
 */
func File(URL string, SavePath string) (int, error) {
	Response := yuzuhttp.Get(URL).Do()
	if Response.Error != nil {
		return 0, Response.Error
	}

	// 创建进度条
	Progressbar := progressbar.NewOptions(int(Response.ContentLength),
		progressbar.OptionEnableColorCodes(true), // 启用颜色
		progressbar.OptionShowBytes(true),        // 显示速度
		progressbar.OptionFullWidth(),            // 宽度设置为Full
		progressbar.OptionSetDescription("Download: "+truncatingString(path.Base(URL), 10)), // 设置描述
		progressbar.OptionClearOnFinish(), // 完成后清除进度条
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[light_blue]=[reset]", // 设置进度条的样式(中间)
			SaucerHead:    "[light_blue]>[reset]", // 设置进度条的样式(头部)
			SaucerPadding: " ",                    // 设置进度条的样式(空白部分)
			BarStart:      "[",                    // 设置进度条的开头
			BarEnd:        "]",                    // 设置进度条的结尾
		}))

	// 创建文件夹
	err := os.MkdirAll(path.Dir(SavePath), 0644)
	if err != nil {
		return 0, err
	}
	// 打开文件
	File, err := os.OpenFile(SavePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer File.Close()

	// 赋值当前下载文件信息
	DownloadingFilePath = SavePath
	DownloadingFile = File

	// 写入文件
	Size, err := io.Copy(io.MultiWriter(File, Progressbar), Response.Body)
	if err != nil {
		// 删除下载失败的文件
		err := File.Close()
		if err == nil {
			err = os.Remove(SavePath)
			if err != nil {
				return 0, err
			}
		}

		return 0, err
	}

	return int(Size), nil
}

/**
 * @description: 截取字符串
 * @param {string} Str 待截取字符串
 * @param {int} Size 截取长度
 * @return {string} 截取后字符串
 */
func truncatingString(Str string, Size int) string {
	var truncatingStr string
	// 截取摘要
	StrRune := []rune(Str)
	if len(StrRune) > Size {
		truncatingStr = string(StrRune[:Size]) + "..."
	} else {
		truncatingStr = string(StrRune)
	}

	return truncatingStr
}
