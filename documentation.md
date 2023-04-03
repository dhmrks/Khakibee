# Paradox Project API
Documentation for Paradox Project public API.  

## Content negotiation

### URI Schemes
  * https

### Consumes
  * application/json  

### Produces
  * application/json 

## Access control

* **Type**: apikey
* **Header**: Authorization


## Endpoints 

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/pub/rooms/{id}/bookings | [Available Bookings](#getBookings) | Returns the available bookings for a three week period. |
| POST | /api/pub/rooms/{id}/bookings | [Reserve Booking](#reserveBooking) | Reserves a booking slot |
| DELETE | /api/pub/rooms/{id}/bookings | [Cancel Reservation](#cancelReservation) | Cancels a booking slot reservation |


## Paths

### <span id="getBookings"></span> Get available booking slots

```
GET /api/pub/rooms/{id}/bookings
```

Returns the available booking slots for a three week period for a specific room id. The booking slots are organized per callendar date

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Description |
|------|--------|------|-----------|
| id | `path param` | int | Room ID |

#### All responses
| Code | Status | Description | Schema |
|------|--------|-------------|-----------|
| `200` | OK | Contains the available booking slots. | [Response JSON](#available-booking-response) |
| `400` | Bad Request | Request body is not valid |

#### <span id="available-booking-response"></span> Response JSON

```
structure
{
  "{date}": ["timeslot","timeslot","..."]
}

example
{
    "2021-11-14": [ "11:00", "15:00" ],
    "2021-11-16": [ "11:00" ],
    "2021-11-17": [ "11:00", "15:00","20:30" ]
}
```
  


### <span id="reserveBooking"></span> Reserve Booking slot

```
POST /api/pub/rooms/{id}/bookings
```

Reserve Booking slot for a specific room id

#### Parameters

| Name | Source | Type | Description |
|------|--------|------|---------|
| id | `path param` | int |  Room ID |
| Body | `body` | [ReserveBookingRequest](#reserve-booking-request)| Reservation details |

#### All responses
| Code | Status | Description | Schema |
|------|--------|-------------|-----------|
| `200` | OK | Booking slot booked successfully  |  |
| `400` | Bad Request | Request body is not valid | |
| `404` | Not Found | Booking slot does not exists | |
| `409` | Conflict | Booking slot is not available | |


### <span id="cancelReservation"></span> Cancel Reservation

```
DELETE /api/pub/rooms/{id}/bookings
```

Cancels booking slot reservation for a specific room id.

#### Produces
  * application/json

#### Parameters

| Name | Source | Type  | Description |
|------|--------|------|---------|
| id | `path param` | int | Room ID |
| Body | `body` | [CancelReservationRequest](#cancel-reservation-request)| Reservations date and time |

#### All responses
| Code | Status | Description | Schema |
|------|--------|-------------|:-----------:|-----------|
| `200` | OK | Booking slot canceled successfully | |
| `400` | Bad Request | Request body is not valid | |
| `404` | Not Found | Reservation does not exists | |
  



## Models

### <span id="reserve-booking-request"></span> ReserveBookingRequest

> ReserveBookingRequest Contains all the details for the specified booking slot


**Properties**

| Name | Type | Required | Description | Expected Value |
|------|------|:--------:|-------------|---------|
| dte | string | true | Booking date | YYYY-MM-DD | 
| tme | string | true | Booking time | HH:MM |
| nme | string | true | | `/[α-ωΑ-Ω-ωίϊΐόάέύϋΰήώa-zA-Z']{3,75}/` |
| phone | string | true | | Valid greek mobile number without prefix |
| email | string | true | |valid email format |
| plr_num | int | true | Number of players joining | 2-7 |
| diff | string | true | Game dificulty | "H","N" |
| plr_lvl | string | true | Team members level | "0-1","2-5","6-10","11-20","21-35","36-50","51-100","101-200","201+" | 
| learned_us | string | true | Player learned about us | "fb","go","ta","fr","mn","bo","co","si","lf" |
| lng | string | true | Speaking language | "el","en" |
| under_aged | boolean | true | Has the team any under aged members | |
| team_nme | string | false | Teams Name | `/[α-ωΑ-Ω-ωίϊΐόάέύϋΰήώa-zA-Z' 0-9!^%\.,@#&*-_+=:']{3,75}/` |
| notes | string | false | |`/[α-ωΑ-Ω-ωίϊΐόάέύϋΰήώa-zA-Z 0-9!^%\.,@#&*-_+=:']{0,150}/` |
| code | string | false | Promo code | `/[\w-]{0,10}/` |
| consent | bool | true | TnCs consent | true |
| played_room1 | bool | true for room id 2 | Team has already played room id 1 | |
| escaped_room1 | bool | true for room id 2 | Team has solved room id 1 | |

**Example**
```
{
  "dte": "2021-11-07",
  "tme":"23:59",
  "nme": "Kostas Panopoulos",
  "phone": "6971231234",
  "email": "kostas.pan@gmail.com",
  "plr_num": 3,
  "diff": "N",
  "plr_lvl": "2-5",
  "learned_us": "fb",
  "team_nme": "team",
  "lng": "el",
  "under_aged": true,
  "notes": "note",
  "code": "12312",
  "played_room1": false,
  "escaped_room1": false,
  "consent": true
}
```



### <span id="cancel-reservation-request"></span> CancelReservationRequest

> CancelReservationRequest Contains the date and time of the requested slot


**Properties**

| Name | Type | Required | Description | Expected value | Example |
|------|------|:--------:|-------------|---------|---------|
| dte | string | true | Booking date | YYYY-MM-DD | `"2021-11-18"` |
| tme | string | true | Booking time | HH:MM | `"16:00"` |


**Example**
```
{
  "dte": "2021-11-07",
  "tme":"23:59"
}
```


