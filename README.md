# FizzBuzz Rest API Server

Expose a REST API endpoint that accepts five parameters : <br>
two strings (say, string1 and string2), <br>
and three integers (say, int1, int2 and limit), <br>
and returns a JSON

It must return a list of strings with numbers from 1 to limit, where:<br>
all multiples of int1 are replaced by string1,<br>
all multiples of int2 are replaced by string2,<br>
all multiples of int1 and int2 are replaced by string1string2<br>

Todo :
 - [ ] Write Tests (try TDD approach)
 - [x] Structure the app
 - [x] DockerFile for deployment, seperate build and runnable images
 - [x] Write fizzbuzz controller
 - [ ] Postman collection
 - [x] Docs for deployment and user guide

Optional ?
 - [ ] add very simple UI
 - [x] add benchmarks
 - [ ] add swagger documentation
 - [ ] add help menu
 - [ ] add example requests

## Installation

Vendor packages are included on the repo
You can use docker:
```
docker build -t fizzbuzz .
docker run -d -p 8080:8080 fizzbuzz
``` 

Or you can build and execute the server using go (tested with 1.13 version)
```
go build -mod=vendor -o main
./main
```


## Conclusions

I've tried 2 implementations, one straightforwad using the modulo (%) and one I thought it will be way more optimized since I thought modulo operation is expensive

It turned out that both have almost the same performance

=> to try the benchmark :

```
go test -bench=.
```


