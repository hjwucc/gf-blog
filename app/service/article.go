package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
)

var Article = new(serviceArticle)

type serviceArticle struct {
}

// 根据文章ID查找
func (a *serviceArticle) Get(id int) (article model.ArticleRes, err error) {
	am := g.DB().Table(dao.Article.Table + " a").
		InnerJoin(dao.Category.Table+" c", "a.category_id = c.id").
		Where("a.id",id)

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
func (a *serviceArticle) ConditionPageList(req *model.ApiArticlesListReq) (articles []model.ArticleListRes, err error) {
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

// 添加文章
func (a *serviceArticle) Add(req *model.ApiAddArticleReq, userId int) (res sql.Result, err error) {
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
	res, err = dao.Article.Insert(articleEntity)
	if err != nil {
		err = gerror.New("添加文章失败")
		return
	}
	return
}

// 修改文章
func (a *serviceArticle) Edit(id int, req *model.ApiAddArticleReq) (result sql.Result, err error) {
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
	result, err = dao.Article.Update(gconv.Map(articleEntity), "id", id)
	if result == nil || err != nil {
		err = gerror.New("修改文章失败")
		return
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		err = gerror.New("文章ID不存在")
		return
	}
	return
}

// 删除文章
func (a *serviceArticle) Delete(id int) (result sql.Result, err error) {
	result, err = dao.Article.Delete("id", id)
	if err != nil {
		err = gerror.New("删除文章失败，请联系管理员")
		return
	}
	return
}
