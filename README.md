<p align="center">
    <img alt="FunFaker Logo" src="img/FunFakerLogo.png" width=200 />
</p>

<p align="center">
    <strong><i>FunFaker</i></strong> is a web API for generating fake data with pop culture references.
</p>

<div align="center">

[![](https://img.shields.io/github/workflow/status/alvii147/FunFaker/Go%20GitHub%20CI?label=Tests&logo=github)](https://github.com/alvii147/FunFaker/actions) [![](https://goreportcard.com/badge/github.com/alvii147/FunFaker)](https://goreportcard.com/report/github.com/alvii147/FunFaker) [![Live Demo](https://img.shields.io/badge/Northflank-Live%20Demo-02133e)](https://funfaker--api--cgvttg4279tq.code.run/name)

</div>

# Try it out

Try out the demo API at `https://funfaker--api--cgvttg4279tq.code.run`:

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/person"
```

Response:

```json
{
    "first-name": "Michael",
    "last-name": "Scott",
    "email": "michael.scott@dunder-mifflin.com",
    "trivia": "Michael Gary Scott is the regional manager of the Scranton, Pennsylvania branch of paper company, Dunder Mifflin and the central character of the sitcom The Office."
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

## Company

Generate a pop-culture associated company name.

**URL:** `/company`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`group` | Group that the company should belong to | Yes | `cartoons`, `comics`, `movies`, `tv-shows`

**Response:**

```json
{
    "name": "name of company",
    "trivia": "brief description of company"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/company?group=movies"
```

```json
{
    "name":"Cyberdyne Systems",
    "trivia":"Cyberdyne Systems is the tech corporation responsible for the development of Skynet in the Terminator movies."
}
```

## Date

Generate a pop-culture associated date.

**URL:** `/date`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`after` | Only generate dates after this date | Yes | Any date in `YYYY-MM-DD` format
`before` | Only generate dates before this date | Yes | Any date in `YYYY-MM-DD` format
`group` | Group that the company should belong to | Yes | `games`, `movies`, `tv-shows`

**Response:**

```json
{
    "day": "day of date",
    "month": "month of date",
    "year": "year of date",
    "trivia": "brief description of the event that occured on the date"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/date?after=1995-09-23&group=movies"
```

```json
{
    "day": 29,
    "month": 29,
    "year": 1997,
    "trivia": "Skynet becomes self-aware in movie Terminator 2 - Judgement Day."
}
```

## Person

Generate a pop-culture associated person.

**URL:** `/person`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`sex` | Sex of person | Yes | `male`, `female`, `other`
`group` | Group that the person should belong to | Yes | `comics`, `movies`, `tv-shows`
`domain-name` | Domain name of person's email (e.g. `gmail`) | Yes | Any
`domain-suffix` | Domain name of person's email (e.g. `com`) | Yes | Any

**Response:**

```json
{
    "first-name": "first name of person",
    "last-name": "last name of person",
    "email": "email of person",
    "trivia": "brief description of person"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/person?sex=female&group=tv-shows&domain-name=outlook&domain-suffix=org"
```

```json
{
    "first-name": "Margaery",
    "last-name": "Tyrell",
    "email": "margaery.tyrell@outlook.org",
    "trivia": "Margaery Tyrell is the only daughter of Lord Mace Tyrell and his wife, Lady Alerie Hightower, appearing in TV-show Game of Thrones."
}
```

## Website

Generate a pop-culture associated website URL.

**URL:** `/website`

**URL Parameters:**

Parameter | Description | Optional | Valid Values
--- | --- | --- | ---
`group` | Group that the website should belong to | Yes | `tv-shows`

**Response:**

```json
{
    "url": "url address of website",
    "trivia": "brief description of website"
}
```

**Example:**

```bash
curl --request GET --url "https://funfaker--api--cgvttg4279tq.code.run/website?group=tv-shows"
```

```json
{
    "url": "http://www.wuphf.com",
    "trivia": "Wuphf.com is a fictional website and social tool developed by Ryan Howard, but allegedly stolen from his then girlfriend Kelly Kapoor, in TV-show The Office."
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
