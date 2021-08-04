package moodle

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/brunossouza/gomoodle-api/moodle/config"
	"github.com/brunossouza/gomoodle-api/moodle/models/categories"
	"github.com/brunossouza/gomoodle-api/moodle/models/courses"
	"github.com/brunossouza/gomoodle-api/moodle/models/enrollments"
	"github.com/brunossouza/gomoodle-api/moodle/models/response"
	"github.com/brunossouza/gomoodle-api/moodle/models/users"
)

//generateEndpoint Method responsible for concatenation of the base url, method, and access token. Returning the configured endpoint for interaction with the requested method.
func generateEndpoint(method string) string {
	cfg, err := config.NewMoodleApiConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return fmt.Sprintf("%s/webservice/rest/server.php?wstoken=%s&wsfunction=%s&moodlewsrestformat=json", cfg.URL, cfg.Token, method)
}

//sendDataToAPI Method used to send data to moodle api.
func sendDataToAPI(method string, data string) (response.Responses, error) {
	resp, err := http.Post(generateEndpoint(method), "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalf("error: openlms.sendDataToAPI() error=%s", err.Error())
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	result, err := response.ParseResponseToStruct(body)

	return result, err
}

func CreateUsers(usersSlice []users.User) (response.Responses, error) {
	var data string
	for idx, user := range usersSlice {
		if userData, err := users.ParseUserToCreateFormData(idx, user); err == nil {
			data = fmt.Sprintf("%s%s", data, userData)
		}
	}

	return sendDataToAPI("core_user_create_users", data)
}

func UpdateUsers(usersSlice []users.User) (response.Responses, error) {
	var data string
	for idx, user := range usersSlice {
		if userData, err := users.ParseUserToUpdateFormData(idx, user); err == nil {
			data = fmt.Sprintf("%s%s", data, userData)
		}
	}

	return sendDataToAPI("core_user_update_users", data)
}

func CreateCategory(indice int, category categories.Category) (response.Responses, error) {
	if data, err := categories.ParsecategoryToFormData(indice, category); err == nil {
		return sendDataToAPI("core_course_create_categories", data)
	} else {
		return nil, err
	}
}

func CreateCourses(coursesSlice []courses.Course) (response.Responses, error) {
	var data string
	for idx, course := range coursesSlice {
		if courseData, err := courses.ParseCourseToFormData(idx, course); err == nil {
			data = fmt.Sprintf("%s%s", data, courseData)
		}
	}

	return sendDataToAPI("core_course_create_courses", data)
}

func SetEnrollment(enrolls []enrollments.Enrollments) (response.Responses, error) {
	var data string
	for idx, enroll := range enrolls {
		if enrollData, err := enrollments.ParseEnrollmentToFormData(idx, enroll); err == nil {
			data = fmt.Sprintf("%s%s", data, enrollData)
		} else {
			log.Fatalln(err.Error())
		}
	}

	return sendDataToAPI("enrol_manual_enrol_users", data)
}
