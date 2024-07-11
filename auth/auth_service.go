package auth


type AuthService interface {
	GenerateToken() (string,error);
}

type authService struct {
	authRepository AuthRepository
}

func NewAuthService(repo AuthRepository) AuthService{
	return &authService {
		authRepository: repo,
	}
}


func (s *authService) GenerateToken() (string, error) {
	return s.authRepository.GenerateToken()
}