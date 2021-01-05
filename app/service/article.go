package service

import (
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
	"go-gf-blog/library/constants"
)

var Article = new(serviceArticle)

type serviceArticle struct {
}

// 根据文章ID查找
func (a *serviceArticle) Get(id int) (article model.ArticleRes, err error) {
	am := g.DB().Table(dao.Article.Table+" a").
		InnerJoin(dao.Category.Table+" c", "a.category_id = c.id").
		Where("a.id", id)

	if count, _ := am.Count(); count == 0 {
		err = gerror.New("查询的文章不存在，请联系管理员")
		return
	}

	err = am.
		Fields("a.*,c.name category_name").
		Scan(&article.Article)

	err = g.DB().Table(dao.ArticleTag.Table+" aTag").
		InnerJoin(dao.Tag.Table+" t", "aTag.tag_id = t.id").
		Fields("aTag.article_id,t.id as tag_id,t.name as tag_name").
		Scan(&article.Tags, "aTag.article_id", id)

	if err != nil {
		err = gerror.New("根据文章ID查找失败")
		return
	}

	return
}

// 根据条件分页查找
func (a *serviceArticle) ConditionPageList(req *model.ServiceArticlesListReq) (articles []model.ArticleListRes, err error) {
	am := g.DB().Table(dao.Article.Table+" a").
		InnerJoin(dao.Category.Table+" c", "a.category_id = c.id").
		Fields("DISTINCT a.id,a.title,a.summary,a.cover,a.author,a.click_count,a.is_top,a.created_at,a.updated_at,c.name category_name")
	if req.CategoryId != 0 {
		am = am.Where("a.category_id", req.CategoryId)
	}
	if req.Keywords != "" {
		am = am.Where("a.title like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Status > -1 {
		am = am.Where("a.status", req.Status)
	}

	if req.TagId != 0 {
		am.
			InnerJoin(dao.ArticleTag.Table+" aTag", "a.id = aTag.article_id").
			Where("aTag.tag_id", req.TagId)
	}

	if req.PageSize == 0 {
		req.PageSize = 8
	}

	err = am.Page(req.PageNum, req.PageSize).
		Order("a.is_top desc , a.created_at desc").
		ScanList(&articles, "Article")

	tm := g.DB().Table(dao.ArticleTag.Table+" aTag").
		InnerJoin(dao.Tag.Table+" t", "aTag.tag_id = t.id").
		Where("aTag.article_id", gdb.ListItemValues(articles, "Article", "Id"))

	if count, _ := tm.Count(); count == 0 {
		err = gerror.New("没有符合条件的文章")
		return
	}

	err = tm.
		Fields("aTag.article_id,t.id as tag_id,t.name as tag_name").
		ScanList(&articles, "Tags", "Article", "article_id:Id")

	if err != nil {
		err = gerror.New("按条件查询分页文章失败")
		return
	}
	return
}

// 发布（新增或修改）文章
func (a *serviceArticle) Publish(req *model.ServicePublishArticleReq, userId int) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		var articleId int
		var err error
		if req.Id != 0 {
			articleId = req.Id
			err = Article.Edit(articleId, req)
		} else {
			articleId, err = Article.Add(req, userId)
		}

		if err != nil {
			return gerror.New("发布文章失败")
		}

		// 设置文章标签
		tags := gstr.Split(req.Tags, ",")
		if tagsLen := len(tags); tagsLen > 0 {
			for _, v := range tags {
				tagId, articleCount, err := Tag.Add(v)
				if err != nil {
					glog.Error(err.Error())
				} else {
					err = ArticleTag.Add(tagId, articleId, articleCount)
					if err != nil {
						glog.Error(err.Error())
					}
				}
			}
		}
		return nil
	})
	return err
}

// 新增文章
func (a *serviceArticle) Add(req *model.ServicePublishArticleReq, userId int) (articleId int, err error) {
	err = g.DB().Transaction(func(tx *gdb.TX) error {
		articleEntity := &model.Article{}
		articleEntity.Status = req.Status
		articleEntity.CategoryId = req.CategoryId
		articleEntity.Content = req.Content
		articleEntity.CreatedAt = gtime.Now()
		articleEntity.UpdatedAt = gtime.Now()
		articleEntity.From = req.From
		articleEntity.Title = req.Title
		articleEntity.MdContent = req.MdContent
		articleEntity.Summary = req.Summary
		articleEntity.Cover = req.Cover
		articleEntity.Author = req.Author
		articleEntity.UserId = userId
		articleEntity.IsTop = req.IsTop
		res, err := tx.Insert(dao.Article.Table, articleEntity)
		if err != nil {
			return gerror.New("添加文章失败")
		}
		id, err := res.LastInsertId()
		articleId = int(id)
		return err
	})
	return
}

// 修改文章
func (a *serviceArticle) Edit(id int, req *model.ServicePublishArticleReq) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		articleEntity := &model.Article{}
		articleEntity.Status = req.Status
		articleEntity.CategoryId = req.CategoryId
		articleEntity.Content = req.Content
		articleEntity.UpdatedAt = gtime.Now()
		articleEntity.From = req.From
		articleEntity.Title = req.Title
		articleEntity.MdContent = req.MdContent
		articleEntity.Summary = req.Summary
		articleEntity.Cover = req.Cover
		articleEntity.Author = req.Author
		articleEntity.IsTop = req.IsTop
		result, err := tx.Update(dao.Article.Table, gconv.Map(articleEntity), "id", id)
		if result == nil || err != nil {
			return gerror.New("修改文章失败")
		}
		if affected, _ := result.RowsAffected(); affected == 0 {
			return gerror.New("文章ID不存在")
		}
		return nil
	})
	return err
}

// 删除文章
func (a *serviceArticle) Delete(id int) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		_, err := tx.Delete(dao.Article.Table, "id", id)

		if err != nil {
			return gerror.New("删除文章失败，请联系管理员")
		}

		tagIds, err := dao.ArticleTag.Array("tag_id", "article_id", id)

		_, err = tx.Update(dao.Tag.Table, g.Map{"article_count": gdb.Raw("article_count-1")}, "id", tagIds)

		if err != nil {
			return gerror.New("删除失败，因为文章对应的标签文章数量更新失败")
		}

		_, err = tx.Delete(dao.ArticleTag.Table, "article_id", id)

		if err != nil {
			return gerror.New("删除失败，因为删除文章标签关联表失败")
		}
		return nil
	})
	return err
}

// 修改文章属性（置顶/发布状态）
func (a *serviceArticle) UpdateAttributeById(serviceReq *model.ServiceUpdateArticleAttributeReq) error {
	one, _ := dao.Article.FindOne("id", serviceReq.ArticleId)
	if one == nil {
		return gerror.New("要修改属性的文章不存在")
	}
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		if constants.Top == serviceReq.Type {
			if one.IsTop == 0 {
				one.IsTop = 1
			} else {
				one.IsTop = 0
			}
			_, err := tx.Update(dao.Article.Table, g.Map{"is_top": one.IsTop}, "id", serviceReq.ArticleId)
			if err != nil {
				return gerror.New("修改文章置顶失败")
			}
			return nil
		}
		if constants.Status == serviceReq.Type {
			if one.Status == 0 {
				one.Status = 1
			} else {
				one.Status = 0
			}
			_, err := tx.Update(dao.Article.Table, g.Map{"status": one.Status}, "id", serviceReq.ArticleId)
			if err != nil {
				return gerror.New("修改文章发布状态失败")
			}
			return nil
		}
		return gerror.New("没有匹配要修改的文章属性")
	})

	return err
}
