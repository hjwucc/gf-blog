package service

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
)

var Link = new(serviceLink)

type serviceLink struct {
}

// 根据条件分页查找
func (a *serviceLink) ConditionPageList(req *model.ApiLinkListReq) (totalCount int, pageList gdb.Result, err error) {
	M := g.DB().Table(dao.Link.Table + " a")
	if req.Sort != 0 {
		M = M.Where("a.sort", req.Sort)
	}
	totalCount, err = M.Count()

	if err != nil {
		err = gerror.New("按条件查询所有链接失败")
		return
	}

	if req.PageSize == 0 {
		req.PageSize = 20
	}

	pageList, err = M.
		Fields("a.*").
		Page(req.PageNum, req.PageSize).
		Order("a.created_at desc").
		All()

	if err != nil {
		err = gerror.New("按条件查询分页链接失败")
		return
	}

	return
}

// 添加链接
func (a *serviceLink) Add(req *model.ApiAddLinkReq) (res sql.Result, err error) {
	linkEntity := &model.Link{}
	linkEntity.Sort = req.Sort
	linkEntity.UpdatedAt = gtime.Now()
	linkEntity.CreatedAt = gtime.Now()
	linkEntity.IconUrl = req.IconUrl
	linkEntity.LinkName = req.LinkName
	linkEntity.LinkUrl = req.LinkUrl
	res, err = dao.Link.Insert(linkEntity)
	if err != nil {
		err = gerror.New("添加链接失败")
		return
	}

	return
}

// 修改链接
func (a *serviceLink) Edit(id int, req *model.ApiAddLinkReq) (result sql.Result, err error) {
	linkEntity := &model.Link{}
	linkEntity.LinkUrl = req.LinkUrl
	linkEntity.LinkName = req.LinkName
	linkEntity.IconUrl = req.IconUrl
	linkEntity.Sort = req.Sort
	linkEntity.UpdatedAt = gtime.Now()
	result, err = dao.Link.Update(gconv.Map(linkEntity), "id", id)
	if result == nil || err != nil {
		err = gerror.New("修改链接失败")
		return
	}

	if affected, _ := result.RowsAffected(); affected == 0 {
		err = gerror.New("链接ID不存在")
	}

	return
}

// 删除链接
func (a *serviceLink) Delete(id int) (result sql.Result, err error) {
	result, err = dao.Link.Delete("id", id)
	if err != nil {
		err = gerror.New("删除链接失败，请联系管理员")
		return
	}
	return
}
