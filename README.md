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
1. (GET) localhost:8001/api/events
2. (GET) localhost:8001/api/events/:id/tickets
3. (POST) localhost:8001/api/bookings
4. (PATCH) localhost:8001/api/bookings/3