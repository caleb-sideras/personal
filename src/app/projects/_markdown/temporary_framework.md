## What is the Temporary Framework?

Temporary is a [Go](https://go.dev/) framework for building full-stack web applications using [Templ](https://templ.guide/) and [HTMX](https://htmx.org/). It tightly couples these technologies together and introduces other primitives to help you build web-applications easier and faster. Under the hood Temporary abstracts routing, building, serving and more.


## Main Features

Some of the main Temporary features include:

- __Routing:__ A file-system based router built on top of Templ Components that supports indexing, nested routing, slugs and more.
- __Rendering:__ Dynamic and Static Rendering of Server Components. 
- __Data Fetching:__ Simple data fetching within Server Components with lazy loading using HTMX.
- __HTMX & Templ:__ Tightly coupled integration with HTMX & Templ extending both their functionality.
- __Go:__ The best language on Earth.


## Getting Started

Visit <a href="https://temporary-framework.org">https://temporary-framework.org</a> to get started with Temporary.

## Documentation

Visit [https://temporary-framework.org/docs](https://temporary-framework.org/docs) to view the full documentation.

## Getting Started

These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- Bun: [Download and Install Bun](https://bun.sh/)
- Go Programming Language: [Download and Install Go](https://golang.org/doc/install)
- Templ: [Download and Install Templ](https://templ.guide/quick-start/installation)

## Installing

1. Clone the repository to your local machine.

   ```bash
   git clone --recursive https://github.com/caleb-sideras/temporary-framework-website.git
   ```

2. Navigate to the project's directory.

   ```bash
   cd temporary-framework-website
   ```

3. Install the project dependencies.

   ```bash
   bun install
   ```

## Building the Application

Run the Temporary build process through Bun:

  ```bash
  bun run build
  ```

## Running the Application

Run the Temporary start process through Bun:

   ```bash
   bun run start
   ```

## License

This project is licensed under the MIT License.
