## Overview

This command line utility makes use of the [Coronavirus COVID19 API](https://covid19api.com) (A free API for data on the Coronavirus) to query for the latest status based on the country slug (as returned by the `GET /countries` endpoint).

## Usage

`go run main.go <country-slug>`

## Example

`go run main.go south-africa`

```
go run main.go south-africa
confirmed -     1585
recovered -       95
deaths    -        9
```