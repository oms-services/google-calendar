package calendar

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	key            = os.Getenv("GOOGLE_CALENDAR_KEY")
	calendarID     = os.Getenv("GOOGLE_CALENDAR_ID")
	eventID        = os.Getenv("GOOGLE_CALENDAR_EVENT_ID")
	eventStartDate = os.Getenv("GOOGLE_CALENDAR_EVENT_START_DATE")
	eventEndDate   = os.Getenv("GOOGLE_CALENDAR_EVENT_END_DATE")
)

var _ = Describe("Get Event List with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/eventList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EventList)
	handler.ServeHTTP(recorder, request)

	Describe("List Event", func() {
		Context("list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event List with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calendar := CalendarArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/eventList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EventList)
	handler.ServeHTTP(recorder, request)

	Describe("List Event", func() {
		Context("list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event List with valid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/eventList", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(EventList)
	handler.ServeHTTP(recorder, request)

	Describe("List Event", func() {
		Context("list", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Get Event by ID with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event by ID with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calendar := CalendarArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event by ID with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{CalendarID: "mockCalendarID", EventID: "mockEventID"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Event by ID with valid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{CalendarID: calendarID, EventID: eventID}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEventByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Event", func() {
		Context("event by id", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Get Calendar By ID with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalendarByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calendar", func() {
		Context("Calendar by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calendar By ID with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calendar := CalendarArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalendarByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calendar", func() {
		Context("Calendar by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calendar By ID with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{CalendarID: "mockCalendarID"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalendarByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calendar", func() {
		Context("Calendar by id", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calendar by ID with valid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{CalendarID: calendarID}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/getCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCalendarByID)
	handler.ServeHTTP(recorder, request)

	Describe("Get Calendar", func() {
		Context("Calendar by id", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Create Calendar with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalendar)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calendar", func() {
		Context("create calendar", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Get Calendar List with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calendar := CalendarArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalendar)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calendar", func() {
		Context("create calendar", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Calendar with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{EventID: "mockCalendar"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalendar)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calendar", func() {
		Context("create calendar", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Calendar with valid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{Summary: "Test Calendar", Description: "Test Description", Location: "Test Location"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createCalendar", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCalendar)
	handler.ServeHTTP(recorder, request)

	Describe("Create Calendar", func() {
		Context("create calendar", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Create Event with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Event with invalid key", func() {

	os.Setenv("KEY", "mockKey")
	calendar := CalendarArguments{}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Event with invalid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{EventID: "mockCalendar"}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Create Event with valid param", func() {

	os.Setenv("KEY", key)
	calendar := CalendarArguments{Summary: "Test Calendar", Description: "Test Description", CalendarID: calendarID, StartDate: eventStartDate, EndDate: eventEndDate}
	requestBody := new(bytes.Buffer)
	encodeErr := json.NewEncoder(requestBody).Encode(calendar)
	if encodeErr != nil {
		log.Fatal(encodeErr)
	}

	request, err := http.NewRequest("POST", "/createEvent", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEvent)
	handler.ServeHTTP(recorder, request)

	Describe("Create Event", func() {
		Context("create event", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
