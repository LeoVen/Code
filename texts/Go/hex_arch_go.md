# Hexagonal Architecture in Go

[Source](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3)

## Core

* Everything is surrounding the core of the application
* Core shouldn't be aware of how the application is served or where the data is
* A box that holds all business logic

## Actors

* Actors are real world things that want to interact with the core
  * Drivers (or primary) actors: those who trigger the communication with Core
  * Driven (or secondary) actors: those who are expecting Core to trigger communication

![Diagram](https://miro.medium.com/max/1400/1*kEomMfgNPu1srEAH7-Z_LA.png)

## Ports

> Ports belong to Core

* Ports for driver actors
* Ports for driven actors

![Diagram](https://miro.medium.com/max/550/1*b_c6bnop4qRjbK4ypUcWAg.png)

## Adapters

> Transformation between the Core's Ports and Actors

* Adapter for a driver port
* Adapter for a driven port

![Diagram](https://miro.medium.com/max/1400/1*ERYx0IB1pN-5ZX98cKAoUw.png)

## Dependency Injection

Decide which adapters to connect at each port

![Diagram](https://miro.medium.com/max/1400/1*tXttGUY2PCCXW8CO6_Xg2w.png)

## Advantages and Disadvantages

* __`+`__ Separation of concerns
* __`+`__ Focus on the business logic
* __`+`__ Parallelization of work
* __`+`__ Tests in isolation
* __`+`__ Easily change infrastructure
* __`+`__ Self-guided process
* __`-`__ Too complex for small or short-term projects
* __`-`__ Performance overhead
