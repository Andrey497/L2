package pkg

type iOperation interface {
	Execute(Account)
	SetNext(nextOperation iOperation)
}
