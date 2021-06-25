package enrollments

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

// 1,"",manager,"",1,manager
// 2,"",coursecreator,"",2,coursecreator
// 3,"",editingteacher,"",3,editingteacher
// 4,"",teacher,"",4,teacher
// 5,"",student,"",5,student
// 6,"",guest,"",6,guest
// 7,"",user,"",7,user
// 8,"",frontpage,"",8,frontpage

type Enrollments struct {
	RoleID    int       `validate:"required"` //Role to assign to the user
	UserID    int       `validate:"required"` //The user that is going to be enrolled
	CourseID  int       `validate:"required"` //The course to enrol the user role in
	TimeStart time.Time //Optional - Timestamp when the enrolment start
	TimeEnd   time.Time //Optional - Timestamp when the enrolment end
	Suspend   int       //Optional - set to 1 to suspend the enrolment
}

func verifyEnrollmentsDataRequired(enroll Enrollments) error {
	return validator.New().Struct(enroll)
}

func ParseEnrollmentToFormData(indice int, enroll Enrollments) (string, error) {
	err := verifyEnrollmentsDataRequired(enroll)
	if err != nil {
		return "", err
	}

	//Required data
	data := fmt.Sprintf("&enrolments[%d][roleid]=%d&enrolments[%d][userid]=%d&enrolments[%d][courseid]=%d", indice, enroll.RoleID, indice, enroll.UserID, indice, enroll.CourseID)

	//Optional data
	if enroll.TimeStart.Unix() != -62135596800 {
		data += fmt.Sprintf("&enrolments[%d][timestart]=%d", indice, enroll.TimeStart.Unix())
	}
	if enroll.TimeEnd.Unix() != -62135596800 {
		data += fmt.Sprintf("&enrolments[%d][timeend]=%d", indice, enroll.TimeEnd.Unix())
	}
	if enroll.Suspend != 0 {
		data += fmt.Sprintf("&enrolments[%d][suspend]=%d", indice, enroll.Suspend)
	}

	return data, nil

}
