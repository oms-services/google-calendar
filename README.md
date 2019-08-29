# _Google Calender_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/google-calendar.svg?branch=master)](https://travis-ci.com/omg-services/google-calendar)
[![codecov](https://codecov.io/gh/omg-services/google-calendar/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/google-calendar)


An OMG service for google calender, it allows users to create and edit events. Reminders can be enabled for events, with options available for type and time. Event locations can also be added, and other users can be invited to events.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Create Calender
```coffee
google-calender createCalender summary:'calender summary' description:'calender description' location:'geographic location'
{"conferenceProperties": {"object"},"description": "calender description","etag": "etag","id": "calender Id","kind": "calendar#calendar","location": "geographic location","summary": "calender summary", "timeZone": "UTC"}
```
##### Create Event
```coffee
google-calender createEvent calenderId:'calender Id' summary:'Event summary' description:'Event description' location:'location' attendeesList:'["abc@example.com","xyz@example.com"]'startDate:'2019-08-02T12:00:00.000Z' endDate:'2019-08-02T16:00:00.000Z'
{"attendees": ["attendees list"],"created": "2019-07-30T10:45:20.000Z","creator": {"creator details"},"description": "Event description","end": {"dateTime": "2019-08-02T16:00:00Z"},"etag": "etag","htmlLink": "htmlLink","iCalUID": "iCalUID","id": "event id","kind": "calendar#event","location": "Pune","organizer": {"organizer details"},"start": {"dateTime": "2019-08-02T12:00:00Z"},"status": "confirmed","summary": "Event summary","updated": "2019-07-30T10:45:20.628Z"}
```
##### Get Calender By ID
```coffee
google-calender getCalender calenderId:'calender Id'
{"accessRole": "owner","backgroundColor": "#cca6ac","colorId": "21","conferenceProperties": {"object"},"description": "calender description","etag": "etag","id": "calender Id","kind": "calendar#calendar","location": "geographic location","summary": "calender summary", "timeZone": "UTC"}
```
##### Event List
```coffee
google-calender eventList
{"etag": "etag",[{"accessRole": "owner","backgroundColor": "#cca6ac","colorId": "21","conferenceProperties": {"object"},"description": "calender description","etag": "etag","id": "calender Id","kind": "calendar#calendar","location": "geographic location","summary": "calender summary", "timeZone": "UTC"}] ,"kind": "calendar#calendarList", "nextPageToken": "nextPageToken"}
```
##### Get Event By ID
```coffee
google-calender getEvent calenderId:'calender Id' eventId:'event Id'
{"attendees": ["attendees list"],"created": "2019-07-30T10:45:20.000Z","creator": {"creator details"},"description": "Event description","end": {"dateTime": "2019-08-02T16:00:00Z"},"etag": "etag","htmlLink": "htmlLink","iCalUID": "iCalUID","id": "event id","kind": "calendar#event","location": "Pune","organizer": {"organizer details"},"start": {"dateTime": "2019-08-02T12:00:00Z"},"status": "confirmed","summary": "Event summary","updated": "2019-07-30T10:45:20.628Z"}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Create Calender
```shell
$ omg run createCalender -a summary=<SUMMARY> -a description=<DESCRIPTION> -a location=<LOCATION>  -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
##### Create Event
```shell
$ omg run createEvent -a calenderId=<CALENDER_ID> -a summary=<SUMMARY> -a description=<DESCRIPTION> -a location=<LOCATION> -a attendeesList=<LIST_ATTENDEES_EMAIL_ADDRESS> -a startDate=<START_DATE> -a endDate=<END_DATE> -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
**Note**: The Start and End date should be string in this "2006-01-02T15:04:05.000Z" format.
##### Get Event By ID
```shell
$ omg run getCalender -a calenderId=<CALENDER_ID> -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
##### Event List
```shell
$ omg run eventList -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
##### Get Event By ID
```shell
$ omg run getEvent -a calenderId=<CALENDER_ID> -a eventId=<EVENT_ID> -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```


**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/google-calender/blob/master/LICENSE).
