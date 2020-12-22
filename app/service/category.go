package service

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"go-gf-blog/app/model/categories"
)

var Category = new(serviceCategory)

type serviceCategory struct {
}

// 根据条件查询，若没有传参数则是查询所有分类
func (a *serviceCategory) ConditionQueryList(req *categories.ApiQueryReq) (res gdb.Result, err error) {
	model := g.DB().Table(categories.Table + " a")
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
func (a *serviceCategory) Add(req *categories.ApiAddReq) (res sql.Result, err error) {
	categoryEntity := &categories.Entity{}
	categoryEntity.Status = req.Status
	categoryEntity.UpdatedAt = gtime.Now()
	categoryEntity.CreatedAt = gtime.Now()
	categoryEntity.Name = req.Name
	categoryEntity.Sort = req.Sort
	res, err = categories.Model.Insert(categoryEntity)
	if err != nil {
		err = gerror.New("增加文章分类失败")
		return
	}
	return
}

// 编辑分类
func (a *serviceCategory) Edit(id int, req *categories.ApiAddReq) (res sql.Result, err error) {
	entity, err := categories.Model.FindOne("id", id)
	if err != nil {
		err = gerror.New("查询分类失败")
		return
	}
	entity.Sort = req.Sort
	entity.Name = req.Name
	entity.UpdatedAt = gtime.Now()
	entity.Status = req.Status
	res, err = categories.Model.Replace(entity)
	if err != nil {
		err = gerror.New("编辑分类失败")
		return
	}
	return
}

// 删除分类
func (a *serviceCategory) Delete(id int) (res sql.Result, err error) {
	res, err = categories.Model.Delete("id", id)
	if err != nil {
		err = gerror.New("删除分类失败")
		return
	}
	return
}
