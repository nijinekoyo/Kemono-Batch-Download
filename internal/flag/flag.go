/*
 * @Author: nijineko
 * @Date: 2026-08-26 16:24:09
 * @LastEditTime: 2026-08-26 16:28:34
 * @LastEditors: nijineko
 * @Description: 命令参数解析模块
 * @FilePath: \kemonoDownload\internal\flag\flag.go
 */
package flag

import "flag"

type FlagData struct {
	Service              string // 服务平台
	User                 string // 创作者用户ID
	Query                string // 搜索关键词
	SavePath             string // 文件保存根路径
	FileNameFilter       string // 下载时过滤文件名
	ExtensionFilter      string // 下载时过滤文件扩展名
	Host                 string // Kemono Host 地址
	FileServerHost       string // Kemono 文件服务器地址
	FileServerPathPrefix string // Kemono 文件服务器路径前缀
	OutputFileLinkOnly   bool   // 仅输出文件链接
}

// 全局命令参数数据
var GlobalData *FlagData

/**
 * @description: 获取全局命令参数
 * @return {*FlagData} 全局命令参数
 */
func Get() *FlagData {
	return GlobalData
}

/**
 * @description: 初始化参数
 * @return {error} 错误
 */
func Init() *FlagData {
	// 参数解析
	Service := flag.String("service", "", "Service platform")
	User := flag.String("user", "", "Creator User ID")
	Query := flag.String("query", "", "Search Keywords")
	SavePath := flag.String("save_path", "data/", "File save root path")
	FileNameFilter := flag.String("file_name_filter", "", "Filter file names when downloading")
	ExtensionFilter := flag.String("extension_filter", "", "Filter file extensions when downloading")
	Host := flag.String("host", "https://kemono.su", "Kemono Host address")
	FileServerHost := flag.String("file_server_host", "https://kemono.su", "Kemono file server address")
	FileServerPathPrefix := flag.String("file_server_path_prefix", "", "Kemono file server path prefix")
	OutputFileLinkOnly := flag.Bool("output_file_link_only", false, "Only output file links")
	flag.Parse()

	// 初始化全局命令参数数据
	GlobalData = &FlagData{}

	// 将参数写入变量
	GlobalData.Service = *Service
	GlobalData.User = *User
	GlobalData.Query = *Query
	GlobalData.SavePath = *SavePath
	GlobalData.FileNameFilter = *FileNameFilter
	GlobalData.ExtensionFilter = *ExtensionFilter
	GlobalData.Host = *Host
	GlobalData.FileServerHost = *FileServerHost
	GlobalData.FileServerPathPrefix = *FileServerPathPrefix
	GlobalData.OutputFileLinkOnly = *OutputFileLinkOnly

	return GlobalData
}
