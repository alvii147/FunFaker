<p align="center">
    <img alt="FunFaker Logo" src="img/FunFakerLogo.png" width=200 />
</p>

<p align="center">
    <strong><i>FunFaker</i></strong> is a web API for generating fake data with pop culture references.
</p>

<p align="center">

[![](https://img.shields.io/github/workflow/status/alvii147/FunFaker/Go%20GitHub%20CI?label=Tests&logo=github)](https://github.com/alvii147/FunFaker/actions) [![](https://goreportcard.com/badge/github.com/alvii147/FunFaker)](https://goreportcard.com/report/github.com/alvii147/FunFaker) [![Live Demo](https://img.shields.io/badge/Northflank-Live%20Demo-02133e)](https://funfaker--api--cgvttg4279tq.code.run/name)

</p>

# Try it out

Try out the demo API at `https://funfaker--api--cgvttg4279tq.code.run`:

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/name"
```

Response:

```json
{
    "first-name": "Frank",
    "last-name": "Castle",
    "trivia": "Francis Castle, also known as The Punisher, is an antihero appearing in American comic books published by Marvel Comics."
}
```

# Running FunFaker API locally using Go

:one: Install Go from the [official website](https://go.dev/). Installation instructions may vary depending on the OS.

:two: Build Go module.

```bash
make build
```

:three: Run Go server

```bash
make server
```

This should launch the FunFaker API server on `http://localhost:8080/`.

# Running FunFaker API locally using Docker

:one: Install Docker from the [official website](https://www.docker.com/). Installation instructions may vary depending on the OS.

:two: Build Docker image

```bash
make docker/build
```

:three: Run Docker container

```bash
make docker/run
```

This should launch the FunFaker API server on `http://localhost:8080/`.

# API Reference

## Name

Generate first and last names of a pop-culture associated person.

**URL:** `/name`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`sex` | Sex of person | Yes | `male`, `female`, `other`
`group` | Group that the person should belong to | Yes | `comics`, `movies`

**Response:**

```json
{
    "first-name": "first name of person",
    "last-name": "last name of person",
    "trivia": "brief description of person"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/name?sex=female&group=comics"
```

```json
{
    "first-name": "Kate",
    "last-name": "Bishop",
    "trivia": "Kate Bishop, also known as Hawkeye, is the third character and first female to take the Hawkeye name, appearing in American comic books published by Marvel Comics."
}
```

## Email

Generate email of a pop-culture associated person.

**URL:** `/email`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`sex` | Sex of person | Yes | `male`, `female`, `other`
`group` | Group that the person should belong to | Yes | `comics`, `movies`
`domain-name` | Domain name of email (e.g. `gmail`) | Yes | Any
`domain-suffix` | Domain name of email (e.g. `com`) | Yes | Any

**Response:**

```json
{
    "email": "email of person",
    "trivia": "brief description of person"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/email?sex=male&group=comics&domain-name=outlook&domain-suffix=org"
```

```json
{
    "email": "damian.wayne@outlook.org",
    "trivia": "Damian Wayne, also known as Damian al Ghul is a superhero and the son of Batman and Talia al Ghul, appearing in comic books published by DC Comics."
}
```

## Address

Generate a pop-culture associated address.

**URL:** `/address`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`group` | Group that the address should belong to | Yes | `cartoons`, `comics`, `tv-shows`
`valid-only` | Only generate valid addresses | Yes | `true`, `false`

**Response:**

```json
{
    "street-name": "street name of address",
    "city": "city of address",
    "state": "state/province of address",
    "country": "country of address",
    "postal-code": "postal code of address",
    "valid": "whether or not the address is valid (true/false)",
    "trivia": "brief description of address"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/address?group=tv-shows"
```

```json
{
    "street-name": "308 Negra Arroyo Lane",
    "city": "Albuquerque",
    "state": "New Mexico",
    "country": "United States",
    "postal-code": "87104",
    "valid": false,
    "trivia": "The White Residence, located at 308 Negra Arroyo Lane, was the home of the White family including Walter, his wife Skyler, their son Walt Jr. and their infant daughter, Holly."
}
```

# Testing

Run all tests:

```bash
make test
```

Run tests in verbose mode:

```bash
make test/verbose
```

# Validation

Validate all data:

```bash
make validate
```

Validate all data and apply fixes where possible if validation fails:

```bash
make validate/autofix
```
