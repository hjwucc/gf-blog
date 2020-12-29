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
func (a *serviceArticle) Get(id int) (articleEntity *model.Articles, err error) {
	articleEntity, err = dao.Articles.FindOne("id", id)
	return
}

// 根据条件分页查找
func (a *serviceArticle) ConditionPageList(req *model.ApiArticlesListReq) (totalCount int, pageList gdb.Result, err error) {
	M := g.DB().Table(dao.Articles.Table + " a")
	if req.CategoryId != 0 {
		M = M.Where("a.category_id", req.CategoryId)
	}
	if req.Keywords != "" {
		M = M.Where("a.title like ?", fmt.Sprintf("%%%s%%", req.Keywords))
	}
	if req.Status > -1 {
		M = M.Where("a.status", req.Status)
	}
	totalCount, err = M.Count()
	if err != nil {
		err = gerror.New("按条件查询所有文章失败")
		return
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	pageList, err = M.
		InnerJoin(dao.Articles.Table+" c", "a.category_id = c.id").
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
func (a *serviceArticle) Add(req *model.ApiAddArticleReq) (res sql.Result, err error) {
	articleEntity := &model.Articles{}
	articleEntity.Status = req.Status
	articleEntity.CategoryId = req.CategoryId
	articleEntity.Content = req.Content
	articleEntity.CreatedAt = gtime.Now()
	articleEntity.UpdatedAt = gtime.Now()
	articleEntity.From = req.From
	articleEntity.Title = req.Title
	articleEntity.MdContent = req.MdContent
	articleEntity.Summary = req.Summary
	articleEntity.Tags = req.Tags
	articleEntity.Cover = req.Cover
	articleEntity.Author = req.Author
	articleEntity.IsTop = req.IsTop
	res, err = dao.Articles.Insert(articleEntity)
	if err != nil {
		err = gerror.New("添加文章失败")
		return
	}
	return
}

// 修改文章
func (a *serviceArticle) Edit(id int, req *model.ApiAddArticleReq) (result sql.Result, err error) {
	articleEntity := &model.Articles{}
	articleEntity.Status = req.Status
	articleEntity.CategoryId = req.CategoryId
	articleEntity.Content = req.Content
	articleEntity.UpdatedAt = gtime.Now()
	articleEntity.From = req.From
	articleEntity.Title = req.Title
	articleEntity.MdContent = req.MdContent
	articleEntity.Summary = req.Summary
	articleEntity.Tags = req.Tags
	articleEntity.Cover = req.Cover
	articleEntity.Author = req.Author
	articleEntity.IsTop = req.IsTop
	result, err = dao.Articles.Update(gconv.Map(articleEntity), "id", id)
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
	result, err = dao.Articles.Delete("id", id)
	if err != nil {
		err = gerror.New("删除文章失败，请联系管理员")
		return
	}
	return
}
