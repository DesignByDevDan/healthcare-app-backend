# Healthcare App - Frontend

This is the frontend for the Healthcare Fullstack App. It allows users to sign up, log in, and access a protected resource. The frontend is built using HTML, CSS, and vanilla JavaScript.

## Features

- **Signup**: Users can create an account by providing their name, email, and password.
- **Login**: Users can log in with their email and password.
- **Protected Resource**: After logging in, users can access a protected resource, which is served by the backend.

## Technologies Used

- HTML
- CSS
- JavaScript
- Hosted on [Netlify](https://www.netlify.com)

## Project Structure

```plaintext
frontend/
├── index.html        # Main page for signup, login, and protected resource access
├── styles.css        # Styles for the frontend
├── signup.js         # Handles the signup functionality
├── login.js          # Handles the login functionality
├── protected.js      # Handles fetching protected resources after login
