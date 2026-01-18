package app

type App struct {
	repo Repository
}

func New(repo Repository) *App {
	return &App{
		repo: repo,
	}
}
