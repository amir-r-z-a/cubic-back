package repos

type UserRepo struct{
	Repo *AppRepo
}

func NewUserRepo (appRep *AppRepo) *UserRepo {
	return &UserRepo{
		Repo: appRep,
	}
}

type UserRepoInterface struct{

}