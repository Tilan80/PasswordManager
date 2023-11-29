# PasswordManager
 Making a password manager with Go, MongoDB and Wails (using Vue) for a school project.

## Installation
 To get the project started locally, download the files, then make sure you have go and wails installed, if not then follow the installation guide:
 https://wails.io/docs/gettingstarted/installation

 When everything is installed and paths are set correctlly, run the command "wails dev".

## Error
 Currently the problem I have is with connecting backend and frontedn or rather how to use the backend functions in vue.
 If you open /frontend/src/components/Login.vue you will see, that I am trying to call a function CheckKey when a button is pressed.
 But wwhen trying it out there is an error that says it cannot read properties of undefiend "CheckKey"
