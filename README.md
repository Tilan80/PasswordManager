# PasswordManager

A password manager built with Go, MongoDB, and Wails (using Vue) for a school project.

## Installation

To run the project locally, follow these steps:

1. Download the project files.
2. Ensure you have Go and Wails installed. If not, refer to the [installation guide](https://wails.io/docs/gettingstarted/installation).
3. After installation, make sure your paths are set correctly.
4. Run the command `wails dev`.

## Issue

I am currently facing an issue connecting the backend and frontend, specifically in using backend functions in Vue. The problem is highlighted in the `/frontend/src/components/Login.vue` file. I am attempting to call the `CheckKey` function when a button is pressed.

However, I encounter an error that states: "Cannot read properties of undefined" at the line where I'm calling the `CheckKey` function. The full error message is as follows:
Login.vue:18 Error calling CheckKey: TypeError: Cannot read properties of undefined (reading 'main')