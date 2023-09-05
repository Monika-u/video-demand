package services

import (
	models "design-system/model"

	"go.elastic.co/apm/v2"
)

// // UserService represents a user service.
// type UserService struct {
//     users  []models.User
//     lastID int
//     mu     sync.Mutex
// }

// // NewUserService creates a new user service.
// func NewUserService() *UserService {
//     return &UserService{}
// }

// RegisterUser registers a new user.
func RegisterUser(ctx, username string, password string) (int, error) {
	span, _ := apm.StartSpan(ctx, "register", "service")
	defer span.End()
	cData := utils.ParseContext(ctx)
	// Get all apps from DB
	// apps, err := s.appRepo.GetAll(ctx)
	userID, err := db.CreateUser(username, password)
	if err != nil {
		config.Logger.Error(ctx, "Error while user register", cData.Email, err.Error())
		// return resources.ServiceResult{
		// 	Code:    http.StatusInternalServerError,
		// 	IsError: true,
		// 	Error: resources.ServiceError{
		// 		ErrorCode: constants.ErrorGetByMail,
		// 		ErrorMsg:  s.errorCache.GetByID(ctx, constants.ErrorGetByMail),
		// 	},
		// }/
		return 0, err
	}
	return userID, nil
}

// AuthenticateUser authenticates a user.
func (us *UserService) AuthenticateUser(username, password string) (models.User, error) {
	us.mu.Lock()
	defer us.mu.Unlock()

	for _, u := range us.users {
		if u.Username == username && u.Password == password { // In a real app, compare hashed passwords
			return u, nil
		}
	}

	return models.User{}, ErrUserNotFound
}
