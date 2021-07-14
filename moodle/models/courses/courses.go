package courses

import (
	"fmt"
	"net/url"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Course struct {
	Fullname            string               `validate:"required"` //full name
	Shortname           string               `validate:"required"` //course short name
	CategoryId          int                  //category id - default 0
	IdNumber            string               //Optional - id number
	Summary             string               //Optional - summary
	SummaryFormat       int                  //Default to "1" - summary format (1 = HTML, 0 = MOODLE, 2 = PLAIN or 4 = MARKDOWN)
	Format              string               //Default to "topics" - course format: weeks, topics, social, site,..
	ShowGrades          int                  //Default to "1" - 1 if grades are shown, otherwise 0
	NewsItems           int                  //Default to "5" - number of recent items appearing on the course page
	StartDate           time.Time            //Optional - timestamp when the course start
	EndDate             time.Time            //Optional - timestamp when the course end
	NumSections         int                  //Optional - (deprecated, use courseformatoptions) number of weeks/topics
	MaxBytes            int                  //Default to "0" - largest size of file that can be uploaded into the course
	ShowReports         int                  //Default to "0" - are activity report shown (yes = 1, no =0)
	Visible             int                  //Optional - 1: available to student, 0:not available
	HiddenSections      int                  //Optional - (deprecated, use courseformatoptions) How the hidden sections in the course are displayed to students
	GroupMode           int                  //Default to "0" - no group, separate, visible
	GroupModeForce      int                  //Default to "0" - 1: yes, 0: no
	DefaultGroupingId   int                  //Default to "0" - default grouping id
	EnableCompletion    int                  //Optional - Enabled, control via completion and activity settings. Disabled, not shown in activity settings.
	CompletionNotify    int                  //Optional - 1: yes 0: no
	Lang                string               //Optional - forced course language
	ForceTheme          string               //Optional - name of the force theme
	CourseFormatOptions []CourseFormatOption //Optional - additional options for particular course format
	CustomFields        []CustomField        //Optional - custom fields for the course
}

type CourseFormatOption struct {
	Name  string `validate:"required"` //course format option name
	Value string `validate:"required"` //course format option value
}

type CustomField struct {
	Shortname string `validate:"required"` //The shortname of the custom field
	Value     string `validate:"required"` //The value of the custom field
}

func verifyCourseDataRequired(course Course) error {
	return validator.New().Struct(course)
}

func verifyCourseFormatOptionDataRequired(option CourseFormatOption) error {
	return validator.New().Struct(option)
}

func verifyCustomFieldDataRequired(field CustomField) error {
	return validator.New().Struct(field)
}

func ParseCourseToFormData(indice int, course Course) (string, error) {
	err := verifyCourseDataRequired(course)
	if err != nil {
		return "", err
	}
	//Required data
	data := fmt.Sprintf("&courses[%d][fullname]=%s&courses[%d][shortname]=%s&courses[%d][categoryid]=%d", indice, url.QueryEscape(course.Fullname), indice, url.QueryEscape(course.Shortname), indice, course.CategoryId)

	//Optional data
	if course.Summary == "" {
		data += fmt.Sprintf("&courses[%d][fullname]=%s", indice, url.QueryEscape(course.Fullname))
	}
	if course.SummaryFormat != 0 {
		data += fmt.Sprintf("&courses[%d][summaryformat]=%d", indice, course.SummaryFormat)
	}
	if course.Format == "" {
		data += fmt.Sprintf("&courses[%d][format]=%s", indice, url.QueryEscape(course.Format))
	}
	if course.ShowGrades != 0 {
		data += fmt.Sprintf("&courses[%d][showgrades]=%d", indice, course.ShowGrades)
	}
	if course.NewsItems != 0 {
		data += fmt.Sprintf("&courses[%d][newsitems]=%d", indice, course.NewsItems)
	}
	if course.StartDate.Unix() != -62135596800 {
		data += fmt.Sprintf("&courses[%d][startdate]=%d", indice, course.StartDate.Unix())
	}
	if course.EndDate.Unix() != -62135596800 {
		data += fmt.Sprintf("&courses[%d][enddate]=%d", indice, course.EndDate.Unix())
	}
	if course.NumSections != 0 {
		data += fmt.Sprintf("&courses[%d][numsections]=%d", indice, course.NumSections)
	}
	if course.MaxBytes != 0 {
		data += fmt.Sprintf("&courses[%d][maxbytes]=%d", indice, course.MaxBytes)
	}
	if course.ShowReports != 0 {
		data += fmt.Sprintf("&courses[%d][showreports]=%d", indice, course.ShowReports)
	}
	if course.Visible != 0 {
		data += fmt.Sprintf("&courses[%d][visible]=%d", indice, course.Visible)
	}
	if course.HiddenSections != 0 {
		data += fmt.Sprintf("&courses[%d][hiddensections]=%d", indice, course.HiddenSections)
	}
	if course.GroupMode != 0 {
		data += fmt.Sprintf("&courses[%d][groupmode]=%d", indice, course.GroupMode)
	}
	if course.GroupModeForce != 0 {
		data += fmt.Sprintf("&courses[%d][groupmodeforce]=%d", indice, course.GroupModeForce)
	}
	if course.DefaultGroupingId != 0 {
		data += fmt.Sprintf("&courses[%d][defaultgroupingid]=%d", indice, course.DefaultGroupingId)
	}
	if course.EnableCompletion != 0 {
		data += fmt.Sprintf("&courses[%d][enablecompletion]=%d", indice, course.EnableCompletion)
	}
	if course.CompletionNotify != 0 {
		data += fmt.Sprintf("&courses[%d][completionnotify]=%d", indice, course.CompletionNotify)
	}
	if course.Lang != "" {
		data += fmt.Sprintf("&courses[%d][lang]=%s", indice, url.QueryEscape(course.Lang))
	}
	if course.ForceTheme != "" {
		data += fmt.Sprintf("&courses[%d][forcetheme]=%s", indice, url.QueryEscape(course.ForceTheme))
	}
	if course.ForceTheme != "" {
		data += fmt.Sprintf("&courses[%d][forcetheme]=%s", indice, url.QueryEscape(course.ForceTheme))
	}

	for idx, op := range course.CourseFormatOptions {
		if err := verifyCourseFormatOptionDataRequired(op); err == nil {
			data += fmt.Sprintf("&courses[%d][courseformatoptions][%d][name]=%s&courses[%d][courseformatoptions][%d][value]=%s", indice, idx, url.QueryEscape(op.Name), indice, idx, url.QueryEscape(op.Value))
		}
	}

	for idx, cu := range course.CustomFields {
		if err := verifyCustomFieldDataRequired(cu); err == nil {
			data += fmt.Sprintf("&courses[%d][customfields][%d][name]=%s&courses[%d][customfields][%d][value]=%s", indice, idx, url.QueryEscape(cu.Shortname), indice, idx, url.QueryEscape(cu.Value))
		}
	}

	return data, nil
}
