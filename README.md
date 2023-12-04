# PasswordManager

A password manager built with Go, MongoDB, and Wails (using Vue) for a school project.

## Installation

To run the project locally, follow these steps:

1. Download the project files.
2. Ensure you have Go and Wails installed. If not, refer to the [installation guide](https://wails.io/docs/gettingstarted/installation).
3. After installation, make sure your paths are set correctly.
4. Run the command `wails dev`.

## Issue
Most of the important and "should change" code is in the files:
  Frontend:
  1. `/frontend/src/components/Login.vue`
  2. `/frontend/src/App.vue`
  3. `/frontend/src/router/index.js`

  Backend:
  1. `main.go`
  2. `app.go`
  3. `dbmanage.go`

I currently have an issue, where I started using vue router, but the router-view does not render correctlly or rather app.vue does not show anything.
