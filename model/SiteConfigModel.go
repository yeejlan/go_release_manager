package model

import(
	"release_manager/domain"
	"release_manager/dal/dao"
)

var (
	SiteConfig = &siteConfigModel{}
)

type siteConfigModel struct{}

func (this *siteConfigModel) Add(siteConfig *domain.SiteConfig) (int, error) {
	return dao.SiteConfig.Add(siteConfig)
}

func (this *siteConfigModel) List(offset int, pageSize int) (result *[]domain.SiteConfig, err error) {
	return dao.SiteConfig.List(offset, pageSize)
}

func (this *siteConfigModel) GetById(id int) (result *domain.SiteConfig, err error) {
	return dao.SiteConfig.GetById(id)
}

func (this *siteConfigModel) Update(siteConfig *domain.SiteConfig) (int, error) {
	return dao.SiteConfig.Update(siteConfig)
}

func (this *siteConfigModel) Delete(id int) (int, error) {
	return dao.SiteConfig.Delete(id)
}