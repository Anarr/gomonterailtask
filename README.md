`init.sql` file contains mysql datase structre before start application import it mysql database    

To start application follow the commands below:

    > cd gomonterail 
    > go build -o server
    > ./server

After starting application the database seedings run automatically

The task does not contain any details about user authentication.
So in booking API's you just need past `user_id` in request header

`user_id` must be `integer` value.

Exists endpoints:
-
* (GET) localhost:8001/api/events -- retrieve exists events
* (GET) localhost:8001/api/events/:id/tickets -- retrive details about event tickets
* (POST) localhost:8001/api/bookings -- new booking request
* (PATCH) localhost:8001/api/bookings/3 -- confirm booking


Booking api's requqire `user_id` in header. It accepts integer values. `user_id` describe booking owner
Detailed Api Doc shared in postman: https://documenter.getpostman.com/view/1163851/UV5Xgwai