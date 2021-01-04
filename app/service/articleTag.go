package service

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
	"strconv"
)

var ArticleTag = new(serviceArticleTag)

type serviceArticleTag struct {
}

// 根据标签ID删除
func (a *serviceArticleTag) DeleteByTagId(id int) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		if _, err := dao.ArticleTag.Delete("tag_id", id); err != nil {
			return gerror.New("根据标签ID删除 article_tag 表失败")
		}
		return nil
	})

	return err
}

// 添加
func (a *serviceArticleTag) Add(tagId int, articleId int, articleCount int) error {
	one, _ := dao.ArticleTag.FindOne(g.Map{"tag_id": tagId, "article_id": articleId})
	if one != nil {
		return nil
	}
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		if _, err := tx.Insert(dao.ArticleTag.Table, &model.ArticleTag{TagId: tagId, ArticleId: articleId, CreatedAt: gtime.Now(), UpdatedAt: gtime.Now()}); err != nil {
			return gerror.New("添加 article_tag 表 失败")
		}

		if _, err := tx.Update(dao.Tag.Table, g.Map{"article_count": gdb.Raw("article_count+1")}, "id", tagId); err != nil {
			return gerror.New("更改标签ID为[" + strconv.Itoa(tagId) + "]的标签下文章数量出错")
		}
		return nil
	})

	return err
}
