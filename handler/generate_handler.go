package handler

import (
	"github.com/gogank/papillon/configuration"
	"github.com/gogank/papillon/utils"
	"errors"
	"github.com/gogank/papillon/render"
	"path"
	"fmt"
	"strings"
	"time"
	"strconv"
)

func Generate() error{
	config := config.NewConfig("../configuration/config/config.toml")

	sourceDir := config.GetString(utils.DIR_SOURCE)
	postsDir  := config.GetString(utils.DIR_POSTS)
	publicDir := config.GetString(utils.DIR_PUBLIC)

	if !utils.ExistDir(sourceDir) {
		return errors.New("source directory doesn't exist, cann't generate")
	}

	//todo 1. create public dir
	if isSuccess := utils.Mkdir(publicDir); !isSuccess {
		return errors.New(fmt.Sprintf("create directory %s failed", publicDir))
	}

	if utils.ExistDir(postsDir) {

		// create public/posts dir
		if !utils.Mkdir(path.Join(publicDir, "posts")) {
			return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts")))
		}

		// 遍历source/posts/ 目录中的所有的markdown文件
		files, err := utils.ListDir(postsDir, "md")
		if err != nil {
			return err
		}

		parse := render.New()

		for _, fname := range files {
			content, err := utils.ReadFile(path.Join(postsDir, fname))
			if err != nil {
				return err
			}

			// 调用markdown－>html方法, 得到文章信息、文章内容
			fileInfo, htmlContent, err := parse.Single(content)
			if err != nil {
				return err
			}

			//todo 根据文章信息创建文件夹
			for k, v := range fileInfo {
				if k == "date" {
					ds := strings.Split(v,"/")

					var (
						year string
						month string
						day	string
					)

					// 如果日期未指定，则默认是当前日期
					if len(ds) != 3 {
						year = strconv.Itoa(time.Now().Year())
						month = strconv.Itoa(int(time.Now().Month()))
						day = strconv.Itoa(time.Now().Day())
					}

					year = ds[0]
					month = ds[1]
					day = ds[2]

					// 检查年份文件夹是否存在
					if !utils.ExistDir(path.Join(publicDir, "posts", year)) {
						if !utils.Mkdir(path.Join(publicDir, "posts", year)) {
							return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts", year)))
						}
					}

					// 检查月份文件夹是否存在
					if !utils.ExistDir(path.Join(publicDir, "posts", month)) {
						if !utils.Mkdir(path.Join(publicDir, "posts", month)) {
							return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts", month)))
						}
					}

					// 检查日期文件夹是否存在
					if !utils.ExistDir(path.Join(publicDir, "posts", day)) {
						if !utils.Mkdir(path.Join(publicDir, "posts", day)) {
							return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts", day)))
						}
					}
				}
				if k == "title" {

				}
			}

			//todo 根据文章内容创建html文件
			if !utils.Mkfile(path.Join(publicDir, "posts", fname), htmlContent) {
				return errors.New(fmt.Sprintf("create file %s failed", path.Join(publicDir, "posts", fname+".html")))
			}
		}


		//todo 4.
	}
	return nil
}
