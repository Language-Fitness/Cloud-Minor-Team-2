package auth

import (
	"ExerciseMicroservice/graph/model" // Update this with the correct package name
	"ExerciseMicroservice/internal/repository"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const InvalidTokenMessage = "invalid token"
const NotFoundMessage = "exercise not found"
const InvalidPermissionsMessage = "invalid permissions for this action"

type IExercisePolicy interface {
	CreateExercise(bearerToken string) (string, error)
	UpdateExercise(bearerToken string, id string) (*model.Exercise, error)
	DeleteExercise(bearerToken string, id string) (bool, *model.Exercise, error)
	GetExercise(bearerToken string, id string) (*model.Exercise, error)
	ListExercises(bearerToken string) (bool, error)
	HasPermissions(bearerToken string, role string) bool
}

type ExercisePolicy struct {
	Token              IToken
	ExerciseRepository repository.IExerciseRepository
}

func NewExercisePolicy(collection *mongo.Collection) IExercisePolicy {
	token := NewToken()

	return &ExercisePolicy{
		Token:              token,
		ExerciseRepository: repository.NewExerciseRepository(collection),
	}
}

func (p *ExercisePolicy) CreateExercise(bearerToken string) (string, error) {
	uuid, roles, err := p.getSubAndRoles(bearerToken)
	if err != nil {
		return "", err
	}

	if !p.hasRole(roles, "create_exercise") {
		return "", errors.New("invalid permissions for this action")
	}

	return uuid, nil
}

func (p *ExercisePolicy) UpdateExercise(bearerToken string, id string) (*model.Exercise, error) {
	uuid, roles, err := p.getSubAndRoles(bearerToken)
	if err != nil {
		return nil, err
	}

	exercise, err2 := p.ExerciseRepository.GetExerciseByID(id)
	if err2 != nil {
		return nil, errors.New(NotFoundMessage)
	}

	if p.hasRole(roles, "update_exercise") && exercise.MadeBy == uuid {
		return exercise, nil
	}

	if p.hasRole(roles, "update_exercise_all") {
		return exercise, nil
	}

	return nil, errors.New(InvalidPermissionsMessage)
}

func (p *ExercisePolicy) DeleteExercise(bearerToken string, id string) (bool, *model.Exercise, error) {
	uuid, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return false, nil, err2
	}

	exercise, err := p.ExerciseRepository.GetExerciseByID(id)
	if err != nil {
		return false, nil, errors.New(NotFoundMessage)
	}

	if p.hasRole(roles, "delete_exercise_all") {
		return true, exercise, nil
	}

	if p.hasRole(roles, "delete_exercise") && exercise.MadeBy == uuid {
		return false, exercise, nil
	}

	return false, nil, errors.New(InvalidPermissionsMessage)
}

func (p *ExercisePolicy) GetExercise(bearerToken string, id string) (*model.Exercise, error) {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return nil, err2
	}

	exercise, err := p.ExerciseRepository.GetExerciseByID(id)
	if err != nil {
		return nil, errors.New(NotFoundMessage)
	}

	if p.hasRole(roles, "get_exercise") {
		return exercise, nil
	}

	return nil, errors.New(InvalidPermissionsMessage)
}

func (p *ExercisePolicy) ListExercises(bearerToken string) (bool, error) {
	_, roles, err2 := p.getSubAndRoles(bearerToken)
	if err2 != nil {
		return false, err2
	}

	if p.hasRole(roles, "get_exercises_all") {
		return true, nil
	}

	if !p.hasRole(roles, "get_exercises") {
		return false, errors.New(InvalidPermissionsMessage)
	}

	return false, nil
}

func (p *ExercisePolicy) HasPermissions(bearerToken string, role string) bool {
	_, roles, _ := p.getSubAndRoles(bearerToken)

	return p.hasRole(roles, role)
}

func (p *ExercisePolicy) getSubAndRoles(bearerToken string) (string, []interface{}, error) {
	token, err := p.Token.IntrospectToken(bearerToken)
	if err != nil || token == false {
		return "", nil, errors.New("invalid token introspect")
	}

	decodeToken, err := p.Token.DecodeToken(bearerToken)
	if err != nil {
		return "", nil, err
	}

	sub, _ := decodeToken["sub"].(string)

	resourceAccess, ok := decodeToken["resource_access"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New(InvalidTokenMessage)
	}

	userManagementClient, ok := resourceAccess["user-management-client"].(map[string]interface{})
	if !ok {
		return "", nil, errors.New(InvalidTokenMessage)
	}

	roles, ok := userManagementClient["roles"].([]interface{})
	if !ok {
		return "", nil, errors.New(InvalidTokenMessage)
	}
	return sub, roles, nil
}

func (p *ExercisePolicy) hasRole(roles []interface{}, targetRole string) bool {
	for _, role := range roles {
		if role == targetRole {
			return true
		}
	}
	return false
}
