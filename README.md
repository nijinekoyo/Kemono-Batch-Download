# Kemono-Batch-Download
[Kemono](https://kemono.su/) Batch Download Tool  
Use the API provided by [Kemono](https://kemono.su/) to batch download all article attachments of a certain creator

[中文说明](/README_zh.md)

## Use
1. Download [release](https://github.com/nijinekoyo/Kemono-Batch-Download/releases)
2. Run
``` shell
./kemonoDownload -service fanbox -user 6570768
```

## Command Parameters
|     Parameter     |                    Description                    |    Example     |
| :---------------: | :-----------------------------------------------: | :------------: |
|     -service      | The service platform to which the creator belongs |     fanbox     |
|       -user       |                  Creator User ID                  |    6570768     |
|      -query       |                  Search Keywords                  | 始まりました！ |
|    -save_path     |                File save root path                |    ./data/     |
| -file_name_filter |        Filter file names when downloading         |      風俗      |
| -extension_filter |      Filter file extensions when downloading      |      .mp4      |

## Build
Need have `Go >= 1.22.2`
``` shell
> go mod tidy

> go build .
```

## License
This project is open source using the `Apache License 2.0`