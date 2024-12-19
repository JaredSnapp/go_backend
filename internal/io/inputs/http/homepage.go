package HTTPHandler

import (
	"context"

	"github.com/swaggest/usecase"
)

type homePageInput struct {
}

type homePageOuput struct {
	title string
}

func (h Handler) homePage() usecase.Interactor {
	f := func(ctx context.Context, input homePageInput, output *homePageOuput) error {

		output.title = "Welcome TO HELL"

		return nil
	}

	interactor := usecase.NewInteractor(f)
	interactor.SetDescription("Homepage")
	interactor.SetTitle("Homepage")
	return interactor
}
