package route

import (
    "github.com/gorilla/mux"
    calendar "github.com/heaptracetechnology/google-calendar/calendar"
    "log"
    "net/http"
)

//Route struct
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

//Routes array
type Routes []Route

var routes = Routes{
    Route{
        "CreateCalendar",
        "POST",
        "/createCalendar",
        calendar.CreateCalendar,
    },
    Route{
        "GetCalendarByID",
        "POST",
        "/getCalendar",
        calendar.GetCalendarByID,
    },
    Route{
        "EventList",
        "POST",
        "/eventList",
        calendar.EventList,
    },
    Route{
        "CreateEvent",
        "POST",
        "/createEvent",
        calendar.CreateEvent,
    },
    Route{
        "GetEventByID",
        "POST",
        "/getEvent",
        calendar.GetEventByID,
    },
}

//NewRouter route
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
