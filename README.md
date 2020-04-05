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

## Why?

I find [Golang](https://go.dev) intriguing and fun, and I want to learn more about it. I've also never created a command line utility before.

The COVID19 pandemic is bascically going to dominate much of our lives for at least the next while. I thought why not let something positive come out of it and learn something new.

This project is also part a [#100DaysOfCode](https://github.com/devillexio/100-days-of-code) challenge I'm doing.