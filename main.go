package main

import (
	"fmt"

	"github.com/brunossouza/gomoodle-api/moodle"
	"github.com/brunossouza/gomoodle-api/moodle/models/categories"
	"github.com/brunossouza/gomoodle-api/moodle/models/courses"
	"github.com/brunossouza/gomoodle-api/moodle/models/enrollments"
	"github.com/brunossouza/gomoodle-api/moodle/models/users"
)

func main() {

	var usersSlice []users.User

	user := users.User{
		Username:  "201610651",
		Password:  "Teste@123",
		Firstname: "User",
		Lastname:  "Moodle Teste Api",
		Email:     "ra201610650@universidadedevassouras.edu.br",
	}

	usersSlice = append(usersSlice, user)

	// user = users.User{
	// 	Username:  "201610651",
	// 	Password:  "Teste@123",
	// 	Firstname: "User",
	// 	Lastname:  "Moodle Teste Api",
	// 	Email:     "ra201610651@universidadedevassouras.edu.br",
	// }

	// usersSlice = append(usersSlice, user)

	usersResponse, _ := moodle.CreateUsers(usersSlice)
	for _, u := range usersResponse {
		fmt.Println(u)
	}

	category := categories.Category{
		Name:   "Segurança",
		Parent: 0,
	}

	categoriesResponse, _ := moodle.CreateCategory(0, category)

	for _, c := range categoriesResponse {
		fmt.Println(c)
	}

	course := []courses.Course{
		{
			Fullname:   "Hacking - 01",
			Shortname:  "hacking01",
			CategoryId: int(categoriesResponse[0].ID),
		},
	}

	coursesResponse, _ := moodle.CreateCourses(course)
	for _, c := range coursesResponse {
		fmt.Println(c)
	}

	enrolls := []enrollments.Enrollments{
		{
			RoleID:   5,
			UserID:   int(usersResponse[0].ID),
			CourseID: int(coursesResponse[0].ID),
		},
	}

	e, _ := moodle.SetEnrollment(enrolls)
	fmt.Println(e)
	for _, v := range e {
		fmt.Println(v)
	}

}
