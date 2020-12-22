package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"go-gf-blog/app/model/articles"
	"go-gf-blog/app/model/categories"
)

var Article = new(serviceArticle)

type serviceArticle struct {
}

// 根据文章ID查找
func (a *serviceArticle) Get(id int) (articleEntity *articles.Entity, err error) {
	articleEntity, err = articles.Model.FindOne("id", id)
	return
}

// 根据条件查找
func (a *serviceArticle) ConditionGetList(req *articles.ApiArticlesListReq) (totalCount int, pageList gdb.Result, err error) {
	model := g.DB().Table(articles.Table + " a")
	if req.CategoryId != 0 {
		model = model.Where("a.category_id", req.CategoryId)
	}
	if req.Keywords != "" {
		model = model.Where("a.title like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Status > -1 {
		model = model.Where("a.status", req.Status)
	}
	totalCount, err = model.Count()
	if err != nil {
		err = gerror.New("按条件查询所有文章失败")
		return
	}
	if req.PageSize == 0 {
		req.PageSize = articles.PageSize
	}

	pageList, err = model.
		InnerJoin(categories.Table+" c", "a.category_id = c.id").
		Fields("a.*,c.name category_name").
		Page(req.PageNum, req.PageSize).
		Order("a.created_at desc").
		All()

	if err != nil {
		err = gerror.New("按条件查询分页文章失败")
		return
	}

	// 返回
	return
}

// 添加文章
func (a *serviceArticle) Add(req *articles.ApiAddReq) (res sql.Result, err error) {
	articleEntity := &articles.Entity{}
	articleEntity.Status = req.Status
	articleEntity.CategoryId = req.CategoryId
	articleEntity.Content = req.Content
	articleEntity.CreatedAt = gtime.Now()
	articleEntity.UpdatedAt = gtime.Now()
	articleEntity.From = req.From
	articleEntity.Content = req.Content
	articleEntity.MdContent = req.MdContent
	articleEntity.Summary = req.Summary
	articleEntity.Tags = req.Tags
	articleEntity.Cover = req.Cover
	res, err = articles.Model.Insert(articleEntity)
	if err != nil {
		err = gerror.New("添加文章失败")
		return
	}
	return
}

// 修改文章
func (a *serviceArticle) Edit(id int, req *articles.ApiAddReq) (result sql.Result, err error) {
	articleEntity, err := a.Get(id)
	if err != nil {
		err = gerror.New("要修改的文章已不存在")
		return
	}
	articleEntity.Status = req.Status
	articleEntity.CategoryId = req.CategoryId
	articleEntity.Content = req.Content
	articleEntity.UpdatedAt = gtime.Now()
	articleEntity.From = req.From
	articleEntity.Content = req.Content
	articleEntity.MdContent = req.MdContent
	articleEntity.Summary = req.Summary
	articleEntity.Tags = req.Tags
	articleEntity.Cover = req.Cover
	result, err = articles.Model.Replace(articleEntity)
	if err != nil {
		err = gerror.New("修改文章失败")
		return
	}
	return
}

// 删除文章
func (a *serviceArticle) Delete(id int) (result sql.Result, err error) {
	result, err = articles.Model.Delete("id", id)
	if err != nil {
		err = gerror.New("删除文章失败，请联系管理员")
		return
	}
	return
}
