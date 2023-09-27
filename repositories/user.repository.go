package repositories

type UserRepository struct {
	// dbContext
}

func NewUserRepository() *UserRepository {
	// Create connection to db here

	return &UserRepository{
		//dbContext: dbCon,
	}
}

func (u UserRepository) GetAll() {
}

func (u UserRepository) GetOne(id string) {

}

func (u UserRepository) Create() {

}

func (u UserRepository) UpdateOne() {

}

func (u UserRepository) DeleteOne(id string) {

}
