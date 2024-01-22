# PasswordManager

A password manager built with Go, MongoDB, and Wails (using Vue) for a school project. (Scroll down for installation guide) <br /><br />
The Password Manager project is a robust and secure application with a carefully crafted backend in the Go language. The backend features a MongoDB database interaction, strong cryptographic functions, and well-defined application logic. Each password is encrypted using the user's initial login password, adding an extra layer of security to reinforce the confidentiality of sensitive user data.<br />
On the frontend, orchestrated with Vue.js and Vuetify, the application provides a seamless and intuitive user experience. It includes features such as login, secure password management, and a responsive dashboard. The harmonious integration of the backend and frontend makes Password Manager a user-friendly solution for secure password management.

## Application

First, we get to a login page, where we enter the password from which we get the key that encrypts every password in the manager.
![Login](/pictures/Login.png)

Then we get to the home page or the dashboard.
![Home](/pictures/Home.png)

We can add a new password with an auto-generated password or our own.
![Add](/pictures/Add.png)

And we can see all the passwords with the GetAll method or get only one by platform (Get method).
![GetAll](/pictures/GetAll.png)

At the end, we can also delete any password by platform.

## Installation

To run the project locally, follow these steps:

1. Download the project files.
2. Ensure you have Go and Wails installed. If not, refer to the [installation guide](https://wails.io/docs/gettingstarted/installation).
3. After installation, make sure your paths are set correctly.
4. Run the command `wails dev`.

