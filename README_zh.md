# Kemono-Batch-Download
[Kemono](https://kemono.su/) 批量下载工具  
利用[Kemono](https://kemono.su/)提供的API，批量下载某一创作者的所有文章附件

## Use
1. 下载[release](https://github.com/nijinekoyo/Kemono-Batch-Download/releases)构建
2. 运行
``` shell
./kemonoDownload -service fanbox -user 6570768
```

## Command Parameters
|     Parameter     |      Description       |    Example     |
| :---------------: | :--------------------: | :------------: |
|     -service      |    创作者所属的平台    |     fanbox     |
|       -user       |     创作者的用户ID     |    6570768     |
|      -query       |       搜索关键词       | 始まりました！ |
|    -save_path     |     文件保存根路径     |    ./data/     |
| -file_name_filter | 下载时按文件名过滤文件 |      風俗      |
| -extension_filter | 下载时按扩展名过滤文件 |      .mp4      |

## Build
需要 `Go >= 1.22.2`
``` shell
> go mod tidy

> go build .
```

## License
本项目基于`Apache License 2.0`协议开源