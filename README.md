# FizzBuzz Rest API Server

Expose a REST API endpoint that accepts five parameters : 
    two strings (say, string1 and string2), 
    and three integers (say, int1, int2 and limit), 
    and returns a JSON

It must return a list of strings with numbers from 1 to limit, where:
all multiples of int1 are replaced by string1,
all multiples of int2 are replaced by string2,
all multiples of int1 and int2 are replaced by string1string2

Todo :
 - [ ] Write Tests (try TDD approach)
 - [ ] Structure the app
 - [x] DockerFile for deployment, seperate build and runnable images
 - [ ] Write fizzbuzz controller
 - [ ] Postman collection
 - [ ] Docs for deployment and user guide

Optional ?
 - [ ] add very simple UI


