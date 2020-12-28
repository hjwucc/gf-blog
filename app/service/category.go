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

var Category = new(serviceCategory)

type serviceCategory struct {
}

// 根据条件查询，若没有传参数则是查询所有分类
func (a *serviceCategory) ConditionQueryList(req *model.ApiQueryCategoriesReq) (res gdb.Result, err error) {
	model := g.DB().Table(dao.Categories.Table + " a")
	if req.Id == 0 && req.Status == 0 {
		res, err = model.FindAll()
		if err != nil {
			err = gerror.New("查询文章分类失败")
			return
		}
		return
	}
	if req.Status != 0 {
		model = model.Where("status", req.Status)
	}
	if req.Id != 0 {
		model = model.Where("id", req.Id)
	}

	res, err = model.All()
	if err != nil {
		err = gerror.New("查询文章分类失败")
		return
	}
	return
}

// 增加分类
func (a *serviceCategory) Add(req *model.ApiAddCategoryReq) (res sql.Result, err error) {
	categoryEntity := &model.Categories{}
	categoryEntity.Status = req.Status
	categoryEntity.UpdatedAt = gtime.Now()
	categoryEntity.CreatedAt = gtime.Now()
	categoryEntity.Name = req.Name
	categoryEntity.Sort = req.Sort
	res, err = dao.Categories.Insert(categoryEntity)
	if err != nil {
		err = gerror.New("增加文章分类失败")
		return
	}
	return
}

// 编辑分类
func (a *serviceCategory) Edit(id int, req *model.ApiAddCategoryReq) (res sql.Result, err error) {
	entity := &model.Categories{}
	entity.Sort = req.Sort
	entity.Name = req.Name
	entity.UpdatedAt = gtime.Now()
	entity.Status = req.Status
	res, err = dao.Categories.Update(gconv.Map(entity), "id", id)
	if res == nil || err != nil {
		err = gerror.New("编辑分类失败")
		return
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = gerror.New("分类ID不存在")
		return
	}
	return
}

// 删除分类
func (a *serviceCategory) Delete(id int) (res sql.Result, err error) {
	res, err = dao.Categories.Delete("id", id)
	if err != nil {
		err = gerror.New("删除分类失败")
		return
	}
	return
}
