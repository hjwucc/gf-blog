package service

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
)

var Category = new(serviceCategory)

type serviceCategory struct {
}

// 初始化时查询分类列表并放入本地缓存
func init() {
	var categories []model.QueryCategoriesRes
	err := g.DB().Table(dao.Category.Table).Where("parent_id = 0").Where("status = 0").ScanList(&categories, "TopCategory")
	err = g.DB().Table(dao.Category.Table).Where("parent_id != 0").Where("status = 0").ScanList(&categories, "LowCategories", "TopCategory", "parent_id:Id")
	if err != nil {
		glog.Error("初始化分类列表失败")
		return
	}
	// 不过期
	gcache.Set("categoryList", categories, 0)
}

// 根据条件查询，若没有传参数则是查询所有分类
func (a *serviceCategory) ConditionQueryList(req *model.ApiQueryCategoriesReq) (categories []model.QueryCategoriesRes, err error) {
	if req.Id == 0 {
		// 由于查找所有分类比较耗时，所以放到本地缓存里
		v, _ := gcache.Get("categoryList")
		categories = v.([]model.QueryCategoriesRes)
		return
	}
	err = g.DB().Table(dao.Category.Table).Where("parent_id = 0").Where("status", req.Status).Where("id", req.Id).ScanList(&categories, "TopCategory")
	err = g.DB().Table(dao.Category.Table).Where("status", req.Status).Where("parent_id", req.Id).ScanList(&categories, "LowCategories", "TopCategory", "parent_id:Id")
	if err != nil {
		err = gerror.New("查询文章分类失败")
		return
	}
	return
}

// 刷新分类列表
func (a *serviceCategory) Fresh() error {
	var categories []model.QueryCategoriesRes
	err := g.DB().Table(dao.Category.Table).Where("parent_id = 0").Where("status = 0").ScanList(&categories, "TopCategory")
	err = g.DB().Table(dao.Category.Table).Where("parent_id != 0").Where("status = 0").ScanList(&categories, "LowCategories", "TopCategory", "parent_id:Id")
	if err != nil {
		err = gerror.New("刷新分类列表失败")
		return err
	}
	// 不过期
	if b, _ := gcache.Contains("categoryList"); b {
		_, _ = gcache.Remove("categoryList")
	}
	gcache.Set("categoryList", categories, 0)
	return nil
}

// 增加分类
func (a *serviceCategory) Add(req *model.ApiAddCategoryReq) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		categoryEntity := &model.Category{}
		categoryEntity.Status = req.Status
		categoryEntity.UpdatedAt = gtime.Now()
		categoryEntity.CreatedAt = gtime.Now()
		categoryEntity.Name = req.Name
		categoryEntity.Sort = req.Sort
		categoryEntity.ParentId = req.ParentId
		categoryEntity.Cover = req.Cover
		categoryEntity.Description = req.Description
		_, err := tx.Insert(dao.Category.Table, categoryEntity)
		if err != nil {
			return gerror.New("增加文章分类失败")
		}
		return nil
	})
	return err
}

// 编辑分类
func (a *serviceCategory) Edit(id int, req *model.ApiAddCategoryReq) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		entity := &model.Category{}
		entity.Sort = req.Sort
		entity.Name = req.Name
		entity.UpdatedAt = gtime.Now()
		entity.Status = req.Status
		entity.ParentId = req.ParentId
		res, err := tx.Update(dao.Category.Table, gconv.Map(entity), "id", id)
		if res == nil || err != nil {
			return gerror.New("编辑分类失败")
		}

		if affected, _ := res.RowsAffected(); affected == 0 {
			return gerror.New("分类ID不存在")
		}
		return nil
	})
	return err
}

// 删除分类
func (a *serviceCategory) Delete(id int) error {
	err := g.DB().Transaction(func(tx *gdb.TX) error {
		_, err := tx.Delete(dao.Category.Table, "id", id)
		if err != nil {
			return gerror.New("删除分类失败")
		}
		return nil
	})
	return err
}
