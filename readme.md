# Dynamic Field Validation

Obtains a paginated list of customers from an api and validates the given data, printing the id and invalid fields based on provided validation data.

Customers are dynamic json objects with an id. They can be validated based on the type and length of each field and fields can be marked as required.

Pagiantion data allows the program to make subsequent calls to the api to aggregate the full list of customers.

### Installation

Clone the git repository to the desired location.

Make sure you have the latest version of Golang installed.

Then simply run `go run main.go {api-url}` in the project.

To run the tests use the following commands to install `ginkgo` and `gomega`

```
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
```

### Example api response:
```
{
  "validations": [
    {
      "name": {
        "required": true,
        "type": "string",
        "length": {
          "min": 5
        }
      }
    },
    {
      "email": {
        "required": true
      }
    },
    {
      "age": {
        "type": "number",
        "required": false
      }
    },
    {
      "newsletter": {
        "required": true,
        "type": "boolean"
      }
    }
  ],
  "customers": [
    {
      "id": 6,
      "name": null,
      "email": "oscar@interview.com",
      "age": "29",
      "country": "El Salvador",
      "newsletter": true
    },
    {
      "id": 7,
      "name": "Tetsuro",
      "email": "tetsuro@interview.com",
      "age": 41,
      "country": "Japan",
      "newsletter": true
    },
    {
      "id": 8,
      "name": "Ricardo",
      "email": "ricardo@interview.com",
      "age": 26,
      "country": "Venezuela",
      "newsletter": false
    },
    {
      "id": 9,
      "name": "Adrian",
      "email": "adrian@interview.com",
      "age": 34,
      "country": "Ireland",
      "newsletter": true
    },
    {
      "id": 10,
      "name": "julien",
      "email": "julien@interview.com",
      "age": 25,
      "country": "Germany",
      "newsletter": true
    }
  ],
  "pagination": {
    "current_page": 2,
    "per_page": 5,
    "total": 16
  }
}
```