package mock

import (
	"cloud-api/domain/entity" // Import the entity package
	"math/rand"
	"time"
)

func GenerateModules() []entity.Module {
	modules := make([]entity.Module, 3)

	for i := 0; i < 3; i++ {
		module := entity.Module{
			ID:          rand.Intn(1000),                              // Random ID (change as needed)
			Name:        "Module " + string(rune(i+1)),                // Replace with actual names
			Description: "Description of Module " + string(rune(i+1)), // Replace with actual descriptions
			Difficulty:  rand.Intn(5),                                 // Random difficulty (change as needed)
			Tags:        []string{"Tag1", "Tag2"},                     // Replace with actual tags
			MadeBy:      "Author " + string(rune(i+1)),                // Replace with actual authors
			Private:     false,                                        // Random private flag (change as needed)
			Key:         "Key " + string(i+1),                         // Replace with actual keys
			CreatedAt:   time.Now().String(),                          // Current timestamp as string
			UpdatedAt:   time.Now().String(),                          // Current timestamp as string
			SoftDeleted: "",                                           // Soft deleted empty string
		}

		modules[i] = module
	}

	return modules
}

func GenerateCourses(modules []entity.Module) []entity.Course {

	courses := make([]entity.Course, len(modules))

	for i, module := range modules {
		course := entity.Course{
			ID:          rand.Intn(1000),                        // Random ID (change as needed)
			ModuleID:    module.ID,                              // Use the ModuleID from the provided Module
			Name:        "Course " + string(i+1),                // Replace with actual course names
			Description: "Description of Course " + string(i+1), // Replace with actual descriptions
			Difficulty:  rand.Intn(5),                           // Random difficulty (change as needed)
			Tags:        []string{"Tag1", "Tag2"},               // Replace with actual tags
			CreatedAt:   time.Now().String(),                    // Current timestamp as string
			UpdatedAt:   time.Now().String(),                    // Current timestamp as string
			SoftDeleted: "",                                     // Soft deleted empty string
		}

		courses[i] = course
	}

	return courses
}

// GenerateExercises generates a slice of Exercise instances with random data.
func GenerateExercises(courses []entity.Course, exerciseTypes []entity.ExerciseType) []entity.Exercise {

	exercises := make([]entity.Exercise, len(courses))

	for i, course := range courses {
		exerciseType := exerciseTypes[rand.Intn(len(exerciseTypes))] // Select a random exercise type
		exercise := entity.Exercise{
			ID:             rand.Intn(1000),           // Random ID (change as needed)
			CourseID:       course.ID,                 // Use the CourseID from the provided Course
			Name:           "Exercise " + string(i+1), // Replace with actual exercise names
			Question:       "Question " + string(i+1), // Replace with actual questions
			Answer:         "Answer " + string(i+1),   // Replace with actual answers
			QuestionTypeID: exerciseType.ID,           // Use the ID from the selected ExerciseType
			Tags:           []string{"Tag1", "Tag2"},  // Replace with actual tags
			Difficulty:     rand.Intn(5),              // Random difficulty (change as needed)
			MadeBy:         "Author " + string(i+1),   // Replace with actual authors
			CreatedAt:      time.Now().String(),       // Current timestamp as string
			UpdatedAt:      time.Now().String(),       // Current timestamp as string
			SoftDeleted:    "",                        // Soft deleted empty string
		}

		exercises[i] = exercise
	}

	return exercises
}

// GenerateExerciseTypes generates a slice of ExerciseType instances with random data.
func GenerateExerciseTypes() []entity.ExerciseType {
	exerciseTypes := make([]entity.ExerciseType, 3) // Generate three entity.ExerciseType instances

	for i := 0; i < 3; i++ {
		exerciseType := entity.ExerciseType{
			ID:          rand.Intn(1000),            // Random ID (change as needed)
			Settings:    GenerateExerciseSettings(), // You can generate ExerciseSettings as needed
			CreatedAt:   time.Now().String(),        // Current timestamp as string
			UpdatedAt:   time.Now().String(),        // Current timestamp as string
			SoftDeleted: "",                         // Soft deleted empty string
		}

		exerciseTypes[i] = exerciseType
	}

	return exerciseTypes
}

// GenerateExerciseSettings generates an ExerciseSettings instance with random data.
func GenerateExerciseSettings() entity.ExerciseSettings {
	exerciseSettings := entity.ExerciseSettings{
		Timer:         rand.Intn(2) == 1, // Random boolean value for Timer
		TimeInSeconds: rand.Intn(300),    // Random time in seconds (adjust as needed)
	}

	return exerciseSettings
}

// GenerateRoles generates roles.
func GenerateRoles() []entity.Role {
	roles := []entity.Role{
		{
			ID:          1,                   // Set the ID as needed
			Name:        "admin",             // Replace with actual role names
			Permissions: []string{"admin"},   // Replace with actual permissions
			CreatedAt:   time.Now().String(), // Current timestamp as string
			UpdatedAt:   time.Now().String(), // Current timestamp as string
			SoftDeleted: "",                  // Soft deleted empty string
		},
		{
			ID:          2,                   // Set the ID as needed
			Name:        "teacher",           // Replace with actual role names
			Permissions: []string{"teacher"}, // Replace with actual permissions
			CreatedAt:   time.Now().String(), // Current timestamp as string
			UpdatedAt:   time.Now().String(), // Current timestamp as string
			SoftDeleted: "",                  // Soft deleted empty string
		},
		{
			ID:          3,                   // Set the ID as needed
			Name:        "student",           // Replace with actual role names
			Permissions: []string{"student"}, // Replace with actual permissions
			CreatedAt:   time.Now().String(), // Current timestamp as string
			UpdatedAt:   time.Now().String(), // Current timestamp as string
			SoftDeleted: "",                  // Soft deleted empty string
		},
	}

	return roles
}

// GenerateUsers generates three User instances with specific roles.
func GenerateUsers() []entity.User {
	// Generate roles
	roles := GenerateRoles()

	users := make([]entity.User, 3)

	for i, roleName := range []string{"admin", "teacher", "student"} {
		var roleID int

		// Find the role ID based on the role name
		for _, role := range roles {
			if role.Name == roleName {
				roleID = role.ID
				break
			}
		}

		user := entity.User{
			ID:              rand.Intn(1000),           // Random ID (change as needed)
			Name:            roleName,                  // Set the user name based on the role
			Email:           roleName + "@example.com", // Replace with actual email logic
			Password:        "password",                // Replace with actual password logic
			SchoolID:        "school123",               // Replace with actual school ID logic
			WhitelistModule: []int{1, 2, 3},            // Replace with actual module IDs or generate as needed
			RoleID:          string(rune(roleID)),      // Set the role ID based on the role
			Rating:          "rating123",               // Replace with actual rating or generate as needed
			Settings:        roleName + "_settings",    // Replace with actual settings or generate as needed
			CreatedAt:       time.Now().String(),       // Current timestamp as string
			UpdatedAt:       time.Now().String(),       // Current timestamp as string
			SoftDeleted:     "",                        // Soft deleted empty string
		}

		users[i] = user
	}

	return users
}

// GenerateSchools generates school instances with random data.
func GenerateSchools() []entity.School {
	schools := make([]entity.School, 3) // Generate three School instances

	for i := 0; i < 3; i++ {
		school := entity.School{
			ID:          rand.Intn(1000),              // Random ID (change as needed)
			Name:        "School Name " + string(i+1), // Replace with actual school names
			Location:    "Location " + string(i+1),    // Replace with actual locations
			CreatedAt:   time.Now().String(),          // Current timestamp as string
			UpdatedAt:   time.Now().String(),          // Current timestamp as string
			SoftDeleted: "",                           // Soft deleted empty string
		}

		schools[i] = school
	}

	return schools
}

func GenerateExerciseResults(user entity.User, exercises []entity.Exercise) []entity.ExerciseResult {
	results := make([]entity.ExerciseResult, 3)

	for i := 0; i < 3; i++ {
		exercise := exercises[i]
		result := entity.ExerciseResult{
			ID:          rand.Intn(1000),        // Random ID (change as needed)
			ExerciseID:  exercise.ID,            // Use the ExerciseID from the provided Exercise
			UserID:      user.ID,                // Use the UserID from the provided User
			Input:       "Input " + string(i+1), // Replace with actual input data
			Result:      rand.Intn(2) == 1,      // Random boolean result (change as needed)
			CreatedAt:   time.Now().String(),    // Current timestamp as string
			UpdatedAt:   time.Now().String(),    // Current timestamp as string
			SoftDeleted: "",                     // Soft deleted empty string
		}

		results[i] = result
	}

	return results
}

func generateData() {
	modules := GenerateModules()

	courses := GenerateCourses(modules)

	exerciseTypes := GenerateExerciseTypes()

	exercises := GenerateExercises(courses, exerciseTypes)

	users := GenerateUsers()

	schools := GenerateSchools()

	exerciseResults := GenerateExerciseResults(users[0], exercises)

	println(schools, exerciseResults)
}
