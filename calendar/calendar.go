package calendar

import (
	"context"
	"encoding/base64"
	"encoding/json"
	result "github.com/heaptracetechnology/google-calendar/result"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	"net/http"
	"os"
	"time"
)

//CalendarArguments struct
type CalendarArguments struct {
	Summary       string   `json:"summary,omitempty"`
	Description   string   `json:"description,omitempty"`
	Location      string   `json:"location,omitempty"`
	EventID       string   `json:"eventId,omitempty"`
	AttendeesList []string `json:"attendeesList,omitempty"`
	StartDate     string   `json:"startDate,omitempty"`
	EndDate       string   `json:"endDate,omitempty"`
	CalendarID    string   `json:"calendarId,omitempty"`
}

const layout = "2006-01-02T15:04:05.000Z"

//CreateCalendar Google-Calendar
func CreateCalendar(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param CalendarArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, calendar.CalendarScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}
	client := conf.Client(context.TODO())

	service, serviceErr := calendar.New(client)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	calendarProperties := calendar.Calendar{
		Summary:     param.Summary,
		Description: param.Description,
		Location:    param.Location,
	}

	newCalendar := service.Calendars.Insert(&calendarProperties)

	newCal, newCalErr := newCalendar.Do()
	if newCalErr != nil {
		result.WriteErrorResponseString(responseWriter, newCalErr.Error())
		return
	}

	bytes, _ := json.Marshal(newCal)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//GetCalendarByID Google-Calendar
func GetCalendarByID(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param CalendarArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, calendar.CalendarScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}
	client := conf.Client(context.TODO())

	service, serviceErr := calendar.New(client)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	calendar, calendarErr := service.CalendarList.Get(param.CalendarID).Do()
	if calendarErr != nil {
		result.WriteErrorResponseString(responseWriter, calendarErr.Error())
		return
	}

	bytes, _ := json.Marshal(calendar)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//EventList Google-Calendar
func EventList(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param CalendarArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, calendar.CalendarScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}
	client := conf.Client(context.TODO())

	service, serviceErr := calendar.New(client)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	eventList, eventListErr := service.CalendarList.List().Do()
	if eventListErr != nil {
		result.WriteErrorResponseString(responseWriter, eventListErr.Error())
		return
	}

	bytes, _ := json.Marshal(eventList)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//CreateEvent Google-Calendar
func CreateEvent(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param CalendarArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, calendar.CalendarScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}
	client := conf.Client(context.TODO())

	service, serviceErr := calendar.New(client)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	attendeesList := param.AttendeesList

	var attendees []*calendar.EventAttendee

	for _, v := range attendeesList {
		tempAttendee := new(calendar.EventAttendee)
		tempAttendee.Email = v
		attendees = append(attendees, tempAttendee)
	}

	startDate, startDateErr := time.Parse(layout, param.StartDate)
	endDate, endDateErr := time.Parse(layout, param.EndDate)

	if startDateErr != nil {
		result.WriteErrorResponseString(responseWriter, startDateErr.Error())
		return
	} else if endDateErr != nil {
		result.WriteErrorResponseString(responseWriter, endDateErr.Error())
		return
	}

	eventDetails := calendar.Event{
		Summary:     param.Summary,
		Attendees:   attendees,
		Description: param.Description,
		Location:    param.Location,
		Start: &calendar.EventDateTime{
			DateTime: startDate.Format(time.RFC3339)},
		End: &calendar.EventDateTime{
			DateTime: endDate.Format(time.RFC3339)},
	}

	createdEvent, eventErr := service.Events.Insert(param.CalendarID, &eventDetails).Do()
	if eventErr != nil {
		result.WriteErrorResponseString(responseWriter, eventErr.Error())
		return
	}

	bytes, _ := json.Marshal(createdEvent)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//GetEventByID Google-Calendar
func GetEventByID(responseWriter http.ResponseWriter, request *http.Request) {

	var key = os.Getenv("KEY")

	decodedJSON, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param CalendarArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	conf, confErr := google.JWTConfigFromJSON(decodedJSON, calendar.CalendarScope)
	if confErr != nil {
		result.WriteErrorResponseString(responseWriter, confErr.Error())
		return
	}
	client := conf.Client(context.TODO())

	service, serviceErr := calendar.New(client)
	if serviceErr != nil {
		result.WriteErrorResponseString(responseWriter, serviceErr.Error())
		return
	}

	event, eventErr := service.Events.Get(param.CalendarID, param.EventID).Do()
	if eventErr != nil {
		result.WriteErrorResponseString(responseWriter, eventErr.Error())
		return
	}

	bytes, _ := json.Marshal(event)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
