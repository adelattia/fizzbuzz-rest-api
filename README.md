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
 - [x] Structure the app
 - [x] DockerFile for deployment, seperate build and runnable images
 - [x] Write fizzbuzz controller
 - [ ] Postman collection
 - [ ] Docs for deployment and user guide

Optional ?
 - [ ] add very simple UI
 - [x] add benchmarks
 - [ ] add swagger documentation


Conclusions :



I've tried 2 implementations, one straightforwad using the modulo (%) and one I thought it will be way more optimized since I thought modulo operation is expensive

It turned out that both have almost the same performance

=> to try the benchmark :

```
go test -bench=.
```


