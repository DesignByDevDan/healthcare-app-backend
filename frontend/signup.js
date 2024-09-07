document.getElementById('signup-form').addEventListener('submit', async function(e) {
    e.preventDefault(); // Prevent the form from submitting the traditional way

    // Get form values
    const name = document.getElementById('signup-name').value;
    const email = document.getElementById('signup-email').value;
    const password = document.getElementById('signup-password').value;

    // Password validation
    const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,}$/;
    if (!passwordRegex.test(password)) {
        alert('Password must be at least 8 characters long, and include at least one uppercase letter, one lowercase letter, one number, and one special character.');
        return;
    }

    try {
        // Make POST request to signup endpoint
        const response = await fetch('http://localhost:8080/signup', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                name: name,
                email: email,
                password: password
            })
        });

        const result = await response.json();

        if (response.ok) {
            // Success alert message
            alert('Signup successful! You can now log in.');
        } else {
            // Error alert message
            alert(`Signup failed: ${result.error}`);
        }
    } catch (error) {
        console.error('Error during signup:', error);
        alert('An error occurred during signup. Please try again later.');
    }
});
