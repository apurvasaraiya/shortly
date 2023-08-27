## Shortly

shortly is the url shortner written in golang.

### Project Structure

- Project mainly contains the controller, service and repository layers which helps making changes at each layer very easy without making changes in each others territories.
- Otherwise it has dto package for holding all the object definitions
- helper package for having all helper functions
- mocks package for generating mocks

### Commands for your help

```sh
make           # builds and starts the http server
make dockerise # builds docker image
make mocks     # generates test mocks for unit tests
make test      # runs all unit test for project
```

### Improvement areas (TODOs)

- have proper logging and use gin or gomux for drop in replacement
- have full test coverage for pending scenarios
- extract out common code and reuse it when needed
- for error scenario have proper response body
- improve documentation
- integrate swagger to generate openapi specs
- integration tests
- add redis/memcachd instead of in-memory database
