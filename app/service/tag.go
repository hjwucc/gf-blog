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

var Tag = new(serviceTag)

type serviceTag struct {
}

// 增加标签
func (a *serviceTag) Add(name string) (id int, articleCount int, err error) {
	tagEntity, err := dao.Tag.FindOne("name", name)
	if tagEntity != nil {
		return tagEntity.Id, tagEntity.ArticleCount, nil
	}
	err = g.DB().Transaction(func(tx *gdb.TX) error {
		res, err := tx.Insert(dao.Tag.Table, &model.Tag{Name: name, CreatedAt: gtime.Now(), UpdatedAt: gtime.Now()})
		if err != nil {
			return gerror.New("增加标签[" + name + "]失败")
		}
		insertId, _ := res.LastInsertId()
		id = int(insertId)
		return nil
	})
	return id, 0, err
}

// 删除标签
func (a *serviceTag) Delete(id int) error {
	// 如果这个标签下有文章就不能删除
	tagEntity, err := dao.Tag.FindOne("id", id)

	if tagEntity == nil {
		return gerror.New("不存在标签ID为[" + strconv.Itoa(id) + "]的标签")
	} else if err == nil && tagEntity.ArticleCount > 0 {
		return gerror.New("标签ID为[" + strconv.Itoa(id) + "]的标签下还存在文章")
	} else {
		err = g.DB().Transaction(func(tx *gdb.TX) error {
			if err := ArticleTag.DeleteByTagId(id); err != nil {
				return err
			}
			if _, err := dao.Tag.Delete("id", id); err != nil {
				return gerror.New("删除标签ID为[" + strconv.Itoa(id) + "]的标签失败")
			}

			return nil
		})
	}

	return err
}
