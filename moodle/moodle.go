package moodle

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brunossouza/gomoodle-api/config"
	"github.com/brunossouza/gomoodle-api/models/categories"
	"github.com/brunossouza/gomoodle-api/models/courses"
	"github.com/brunossouza/gomoodle-api/models/enrollments"
	"github.com/brunossouza/gomoodle-api/models/response"
	"github.com/brunossouza/gomoodle-api/models/users"
)

func generateEndpoint(method string) string {
	cfg, err := config.NewMoodleApiConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return fmt.Sprintf("%s/webservice/rest/server.php?wstoken=%s&wsfunction=%s&moodlewsrestformat=json", cfg.URL, cfg.Token, method)
}

func sendDataToAPI(endpoint string, data string) (response.Responses, error) {
	resp, err := http.Post(endpoint, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalf("error: openlms.sendDataToAPI() error=%s", err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	result, err := response.ParseResponseToStruct(body)

	return result, err
}

func CreateUsers(usersSlice []users.User) (response.Responses, error) {
	endpoint := generateEndpoint("core_user_create_users")

	var data string

	for idx, user := range usersSlice {
		if userData, err := users.ParseUserToFormData(idx, user); err == nil {
			data = fmt.Sprintf("%s%s", data, userData)
		}
	}
	return sendDataToAPI(endpoint, data)
}

func CreateCategory(indice int, category categories.Category) (response.Responses, error) {
	endpoint := generateEndpoint("core_course_create_categories")

	if data, err := categories.ParsecategoryToFormData(indice, category); err == nil {
		fmt.Println(endpoint, data)
		return sendDataToAPI(endpoint, data)
	} else {
		return nil, err
	}
}

func CreateCourses(coursesSlice []courses.Course) (response.Responses, error) {
	endpoint := generateEndpoint("core_course_create_courses")

	var data string

	for idx, course := range coursesSlice {
		if courseData, err := courses.ParseCourseToFormData(idx, course); err == nil {
			data = fmt.Sprintf("%s%s", data, courseData)
		}
	}
	return sendDataToAPI(endpoint, data)
}

func SetEnrollment(enrolls []enrollments.Enrollments) (response.Responses, error) {
	endpoint := generateEndpoint("enrol_manual_enrol_users")

	var data string

	for idx, enroll := range enrolls {
		if enrollData, err := enrollments.ParseEnrollmentToFormData(idx, enroll); err == nil {
			data = fmt.Sprintf("%s%s", data, enrollData)
		} else {
			log.Fatalln(err.Error())
		}
	}
	fmt.Println("POST DATA:", data)
	return sendDataToAPI(endpoint, data)
}
