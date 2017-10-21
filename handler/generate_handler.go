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
	"math/rand"
)

func Generate(conf_path string) error{
	cnf := config.NewConfig(conf_path)

	sourceDir := cnf.GetString(utils.DIR_SOURCE)
	postsDir  := cnf.GetString(utils.DIR_POSTS)
	publicDir := cnf.GetString(utils.DIR_PUBLIC)
	themeDir := cnf.GetString(utils.DIR_THEME)

	// 1. 检查 source 文件夹是否存在
	if !utils.ExistDir(sourceDir) {
		return errors.New(fmt.Sprintf("source directory '%s' doesn't exist, cann't generate", sourceDir))
	}

	// 2. 删除 public 文件夹
	if utils.ExistDir(publicDir) {
		if err := utils.RemoveDir(publicDir); err != nil {
			return err
		}
	}

	// 3. 创建新的 public 文件夹
	if !utils.Mkdir(publicDir) {
		return errors.New(fmt.Sprintf("create directory %s failed", publicDir))
	}

	if utils.ExistDir(postsDir) {

		// 4. 创建 public/posts 文件夹
		if !utils.Mkdir(path.Join(publicDir, "posts")) {
			return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts")))
		}

		// 5. 遍历source/posts/ 目录中的所有的markdown文件， 转化为html文件
		files, err := utils.ListDir(postsDir, "md")
		if err != nil {
			return err
		}

		parse := render.New()

		// 生成文章的静态html
		for _, fname := range files {
			mdContent, err := utils.ReadFile(path.Join(postsDir, fname))
			if err != nil {
				return err
			}

			postsTpl, err := utils.ReadFile(path.Join(themeDir, "post.hbs"))
			if err != nil {
				return err
			}

			// 调用markdown－>html方法, 得到文章信息、文章内容
			fileInfo, htmlContent, err := parse.DoRender(mdContent, postsTpl, nil)
			if err != nil {
				return err
			}

			now := time.Now()
			year := strconv.Itoa(now.Year())
			month := strconv.Itoa(int(now.Month()))
			day := strconv.Itoa(now.Day())
			title := "Untitled"+ strconv.Itoa(rand.Int())

			// 根据文章信息创建文件夹
			for k, v := range fileInfo {

				// 确定日期文件夹目录
				if k == "date" {
					ds := strings.Split(v.(string),"/")

					if len(ds) == 3 {
						year = ds[0]
						month = ds[1]
						day = ds[2]
					}
				}

				// 确定文章文件夹目录
				if k == "title" {
					title = v.(string)
				}
			}

			// 检查年份文件夹是否存在
			if !utils.ExistDir(path.Join(publicDir, "posts", year)) {
				if !utils.Mkdir(path.Join(publicDir, "posts", year)) {
					return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts", year)))
				}
			}

			// 检查月份文件夹是否存在
			if !utils.ExistDir(path.Join(publicDir, "posts", year, month)) {
				if !utils.Mkdir(path.Join(publicDir, "posts", year, month)) {
					return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts", year, month)))
				}
			}

			// 检查日期文件夹是否存在
			if !utils.ExistDir(path.Join(publicDir, "posts", year, month, day)) {
				if !utils.Mkdir(path.Join(publicDir, "posts", year, month, day)) {
					return errors.New(fmt.Sprintf("create directory %s failed", path.Join(publicDir, "posts", year, month, day)))
				}
			}

			newTitle := strings.Replace(title, " ", "_", -1)
			if !utils.Mkdir(path.Join(publicDir, "posts", year, month, day, newTitle)) {
				return errors.New(fmt.Sprintf("create directory %s failed",
					path.Join(publicDir, "posts", year, month, day, newTitle)))
			}

			// 根据文章内容创建html文件
			if !utils.Mkfile(path.Join(publicDir,"posts", year, month, day, newTitle, "index.html"), htmlContent) {
				return errors.New(fmt.Sprintf("create file %s failed",
					path.Join(publicDir,"posts", year, month, day, newTitle, "index.html")))
			}

		}

		// 6. 生成首页的html
		if err := generateIndexHtml(cnf, path.Join(publicDir, "index.html")); err != nil {
			return err
		}

		// 7. 复制样式文件
		//if err := utils.CopyDir(path.Join(themeDir, "assets"), publicDir); err != nil {
		//	return err
		//}
	}
	return nil
}

func generateIndexHtml(cnf *config.Config, indexPath string) error {
	parse := render.New()

	themeDir := cnf.GetString(utils.DIR_THEME)
	postsDir := cnf.GetString(utils.DIR_POSTS)

	indexCtx := make(map[string]interface{})

	// 首页的基本信息
	indexCtx["title"] = cnf.GetString(utils.COMMON_TITLE)
	indexCtx["description"] = cnf.GetString(utils.COMMON_DESC)
	indexCtx["author"] = cnf.GetString(utils.COMMON_AUTHOR)

	// 首页的文章信息
	files, err := utils.ListDir(postsDir, "md")
	if err != nil {
		return errors.New(fmt.Sprintf("read directory %s failed", postsDir))
	}

	indexCtx["articles"] = make([]map[string]interface{}, len(files))

	for i, fname := range files {
		mdContent, err := utils.ReadFile(path.Join(postsDir, fname))
		if err != nil {
			return err
		}

		var dateSlice []string
		if meta, err := render.GetMeta(mdContent); err == nil {
			date := meta["date"]
			title := meta["title"]

			dateSlice = strings.Split(date, "/")

			indexCtx["articles"].([]map[string]interface{})[i] = make(map[string]interface{})
			indexCtx["articles"].([]map[string]interface{})[i]["date"] = date
			indexCtx["articles"].([]map[string]interface{})[i]["title"] = title

			articleURL := path.Join("posts", dateSlice[0], dateSlice[1], dateSlice[2], title, "index.html")
			indexCtx["articles"].([]map[string]interface{})[i]["url"] = "/"+articleURL
		} else {
			return err
		}
	}

	indexTpl, err := utils.ReadFile(path.Join(themeDir, "index.hbs"))
	if err != nil {
		return err
	}

	_, indexHtml, err := parse.DoRender(nil, indexTpl, indexCtx)
	if err != nil {
		return err
	}

	if !utils.Mkfile(indexPath, indexHtml) {
		return errors.New(fmt.Sprintf("create file %s failed", indexPath))
	}

	return nil
}
