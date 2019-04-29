package domain

type SiteConfig struct {
	Id int
	Sitename string
	Base_dir string
	Get_current_branch_command string
	Update_command string
	Generate_command string
	Test_release_command string
	Release_command string
	Cache_dir string
	Cache_exclude_dir string
	Cache_urls string
}