document.getElementById('login-form').addEventListener('submit', async function(e) {
    e.preventDefault(); // Prevent the form from submitting the traditional way

    // Get form values
    const email = document.getElementById('login-email').value;
    const password = document.getElementById('login-password').value;

    try {
        // Make POST request to login endpoint
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                email: email,
                password: password
            })
        });

        const result = await response.json();

        if (response.ok) {
            // Store the JWT token in localStorage (or cookies if preferred)
            localStorage.setItem('token', result.token);
            // Success alert message
            alert('Login successful!');
        } else {
            // Error alert message
            alert(`Login failed: ${result.error}`);
        }
    } catch (error) {
        console.error('Error during login:', error);
        alert('An error occurred during login. Please try again later.');
    }
});
