API Reference

API is running on an EC2 instance, accessible using http://52.2.37.231:8080/.

## Methods

### Auth
* **Register**
  * Example Request
  ```bash
  curl -XPOST http://52.2.37.231:8080/register \
   -d username="toto" \
   -d password="foobar"
  ```
  * Example Response
  ```javascript 
    {
      "id":"tpoXasd092",
      "status":"Created"
    }
  ```
* **Login** Need to pass in access_token in all subsequent requests. 
  * Example Request
  ```bash
  curl -XPOST http://52.2.37.231:8080/login \
   -d username="toto" \
   -d password="foobar"
  ```
  * Example Response
  ```javascript 
    {
      "access_token":"230piaasdl9Lnasp98yasd"
    }
  ```
* **Logout**
  * Example Request
  ```bash
  curl -XPOST http://52.2.37.231:8080/logout \
   -H "Authorization: Bearer {access_token}" \
   -d username="toto"
  ```
  * Example Response
  ```javascript 
    {
      "status":"OK"
    }
  ```

### Configurations

* **Create a configuration**
  * Example Request
  ```bash
  curl -XPOST http://52.2.37.231:8080/configurations \
   -H "Authorization: Bearer {access_token}" \
   -d name="host1" \
   -d hostname="host-a.com" \
   -d port=1241 \
   -d username="toto"
  ```
  * Example Response
  ```javascript 
    {
      "id":"o8923p0dkWedoK",
      "status":"Created"
    }
  ```
* **Retrive a configuration**
  * Example Request
  ```bash
  curl -XGET http://52.2.37.231:8080/configurations/o8923p0dkWedoK \
 -H "Authorization: Bearer {access_token}"
  ```
  * Example Response
  ```javascript
  {
    "id":"o8923p0dkWedoK",
    "name":"host1",
    "hostname":"host-a.com",
    "port":1241,
    "username":"toto"
  }
  ```

* **Update a configuration**
  * Example Request
 ```bash
 curl -XPUT http://52.2.37.231:8080/configurations/o8923p0dkWedoK \
  -H "Authorization: Bearer {access_token}" \
  -d name="otherhost" \
  -d port=1244
 ```
   * Example Response
   ```javascript
   {
     "id":"o8923p0dkWedoK",
     "status":"OK"
   }
   ```

* **Delete a configuration**
  * Example Request
  ```bash
   curl -XDELETE http://52.2.37.231:8080/configurations/o8923p0dkWedoK \
   -H "Authorization: Bearer {access_token}"
  ```
  * Example Response
  ```javascript
  {
    "id":"o8923p0dkWedoK",
    "status":"OK"
  }
  ```
* **List configurations**

 * Example Request
 ```bash
 curl -XGET 'http://52.2.37.231:8080/configurations?size=2&sort=username:desc' \
 -H "Authorization: Bearer {access_token}"
 ```
 * Example Response
 ```javascript
 {
   "configurations":[
     { 
       "id":"o8923p0dkWedoK",
       "name":"otherhost",
       "hostname":"host-a.com",
       "port":1244,
       "username":"toto"
     },
     {
       "id": "0san09sadljqa",
       "name":"host",
       "hostname":"host-b.com",
       "port":3384,
       "username":"admin"
     }
   ]
 }
 ```
 * Pagination and Sorting
   - **size** - The number of records to return. Ex: size=100
    - **from** - The number of records to skip. Ex: size=1000&from=5000
    - **sort** - The field name(s) to use for sorting. Can sort by multiple fields(order matters). Ex: sort=fieldname or sort=fieldname:desc&sort=otherfieldname:asc

## Errors
### HTTP Status Codes
Code          | Status      
------------- | ------------- 
200           | OK 
201           | Created 
400           | Bad Request
401           | Unauthorized
402           | Request Failed
404           | Not Found
409           | Conflict
500-504       | Server Errors

