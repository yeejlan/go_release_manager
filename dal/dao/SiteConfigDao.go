package dao

import (
	"github.com/yeejlan/maru"
	"release_manager/dal"
	"release_manager/domain"
)

var (
	SiteConfig = &siteConfigDao{}
)

type siteConfigDao struct{}

func (this *siteConfigDao) Add(siteConfig *domain.SiteConfig) (int, error) {
	sql := "insert into siteconfig "+
		"(`sitename`,`base_dir`,`get_current_branch_command`,`update_command`,"+
		"generate_command,`test_release_command`,`release_command`,`cache_dir`,"+
		"`cache_exclude_dir`,`cache_urls`) values "+
		"(:sitename,:base_dir,:get_current_branch_command,:update_command,"+
		":generate_command,:test_release_command,:release_command,:cache_dir,"+
		":cache_exclude_dir,:cache_urls)"

	return dal.DB.Insert(sql, siteConfig)
}

func (this *siteConfigDao) List(offset int, pageSize int) (result []domain.SiteConfig, err error) {

	p := map[string]interface{}{
		"offset": offset,
		"pageSize": pageSize,
	}
	sql := "select * from siteconfig limit :offset , :pageSize"
	result = []domain.SiteConfig{}
	err = dal.DB.Select(&result, sql, p)
	return
}

func (this *siteConfigDao) GetById(id int) (result *domain.SiteConfig, err error) {
	if id < 1 {
		return nil, maru.NewError("invalid id")
	}
	p := map[string]interface{}{
		"id": id,
	}
	sql := "select * from siteconfig where id = :id"
	result = &domain.SiteConfig{}
	err = dal.DB.SelectOne(result, sql, p)
	return
}

func (this *siteConfigDao) Update(siteConfig *domain.SiteConfig) (int, error) {

	sql := "update siteconfig set "+
		"`sitename` = :sitename, "+
		"`base_dir` = :base_dir, "+
		"`get_current_branch_command` = :get_current_branch_command, "+
		"`update_command` = :update_command, "+
		"`generate_command` = :generate_command, "+
		"`test_release_command` = :test_release_command, "+
		"`release_command` = :release_command,  "+
		"`cache_dir` = :cache_dir,  "+
		"`cache_exclude_dir` = :cache_exclude_dir,  "+
		"`cache_urls` = :cache_urls "+
		" WHERE id = :id"

	return dal.DB.Update(sql, siteConfig)
}

func (this *siteConfigDao) Delete(id int) (int, error) {
	if id < 1 {
		return 0, maru.NewError("invalid id")
	}
	p := map[string]interface{}{
		"id": id,
	}
	sql := "delete from siteconfig where id = :id"
	return dal.DB.Update(sql, p)
}