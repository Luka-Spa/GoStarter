# Go Starter API
[![Go Version][go-version-shield]][go-version-url]
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
## Table of contents

1. How to run
2. Build with
3. Folder structure

## How to run
1. Run ``` go get ``` in the current directory to install the dependencies
2. Run ``` go run . ``` to run the project

## Techstack


| Technology       | Description                                |
| ---------------- | ------------------------------------------ |
| Go               | Go is used as the programming language     |
| Gin              | Gin is the web framework                   |
| MongoDB          | MongoDB is used as the default database    |
| Docker           | Docker is used to containerize the service |

## Folder Structure
```c++
├───controller      // Http Routing & handling
├───logic           // Business logic
├───model           // Data models
└───repository      // Data repositories
    └───mongo       // MongoDB repository implementation
```
[forks-shield]: https://img.shields.io/github/forks/Luka-Spa/GoStarter?style=for-the-badge
[forks-url]: https://github.com/Luka-Spa/GoStarter/network/members
[stars-shield]: https://img.shields.io/github/stars/Luka-Spa/GoStarter?style=for-the-badge
[stars-url]: https://github.com/Luka-Spa/GoStarter/stargazers
[issues-shield]: https://img.shields.io/github/issues/Luka-Spa/GoStarter?style=for-the-badge
[issues-url]: https://github.com/Luka-Spa/GoStarter/issues
[license-shield]: https://img.shields.io/github/license/Luka-Spa/GoStarter?style=for-the-badge
[license-url]: https://github.com/Luka-Spa/GoStarter/blob/main/LICENSE
[contributors-shield]: https://img.shields.io/github/contributors/Luka-Spa/GoStarter?color=blue&style=for-the-badge
[contributors-url]: https://github.com/Luka-Spa/GoStarter/graphs/contributors
[go-version-shield]: https://img.shields.io/badge/Go%20Version-1.20.2-success?style=for-the-badge
[go-version-url]: https://go.dev/dl/