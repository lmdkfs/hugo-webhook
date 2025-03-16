package service

import (
	"context"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/viper"
)

type HugoWebHookService interface {
	UpdateWebSite() error
}

type hugoWebHookService struct {
	RepoURL string
	Branch  string
	HugoDIR string
	*Service
}

func NewHugoWebHookService(config *viper.Viper, service *Service) HugoWebHookService {

	return &hugoWebHookService{
		Service: service,
		RepoURL: config.GetString("hugo.repo_url"),
		Branch:  config.GetString("hugo.hugo_site_branch"),
		HugoDIR: config.GetString("hugo.hugo_site_dir"),
	}
}

func (s *hugoWebHookService) UpdateWebSite() error {
	fmt.Printf("----->\n")
	fmt.Printf("hugoWebHookService UpdateWebSite\n")
	return nil
}

func (s *hugoWebHookService) rebuildHugoSite(ctx context.Context) error {
	_, err := os.Stat(s.HugoDIR)
	if os.IsNotExist(err) {
		_, err := git.PlainCloneContext(ctx, s.HugoDIR, false, &git.CloneOptions{
			URL:           s.RepoURL,
			ReferenceName: plumbing.NewBranchReferenceName(s.Branch),
			SingleBranch:  true,
			Depth:         1,
			Progress:      os.Stdout,
		})
		if err != nil {
			return fmt.Errorf("git clone 失败: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("检查目录是否存在失败: %w", err)
	} else {
		fmt.Printf("目录存在，执行 git pull\n")
		repo, err := git.PlainOpen(s.HugoDIR)
		if err != nil {
			return fmt.Errorf("打开 git 仓库失败: %w", err)
		}
		w, err := repo.Worktree()
		if err != nil {
			return fmt.Errorf("获取工作区失败: %w", err)
		}
		err = w.PullContext(ctx, &git.PullOptions{
			RemoteName:    "origin",
			ReferenceName: plumbing.NewBranchReferenceName(s.Branch),
			SingleBranch:  true,
			Progress:      os.Stdout,
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return fmt.Errorf("git pull 失败: %w", err)
		}
		if err == git.NoErrAlreadyUpToDate {
			fmt.Println("代码已经是最新")
		}
	}
	return nil
}
