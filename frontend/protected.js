protected.js
// Get the stored JWT token from localStorage
const token = localStorage.getItem("token");

// Select the element to display messages
const protectedMessage = document.getElementById("protectedMessage");

if (!token) {
    // If there's no token, redirect to the login page
    window.location.href = "/frontend/login.html";
} else {
    // Make a request to the protected route with the JWT token
    fetch('http://localhost:8080/api/protected-resource', {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}` // Include the token in the Authorization header
        }
    })
    .then(response => response.json())
    .then(data => {
        if (data.message) {
            // If we successfully access the protected resource, display the message
            protectedMessage.textContent = data.message;
            protectedMessage.style.color = "green";
        } else {
            // If something goes wrong, display an error and redirect to login
            protectedMessage.textContent = "Unauthorized! Redirecting to login...";
            protectedMessage.style.color = "red";
            setTimeout(() => {
                window.location.href = "/frontend/login.html";
            }, 2000);
        }
    })
    .catch(error => {
        // In case of an error, redirect to login
        console.error("Error accessing protected resource:", error);
        protectedMessage.textContent = "Error accessing protected resource! Redirecting to login...";
        setTimeout(() => {
            window.location.href = "/frontend/login.html";
        }, 2000);
    });
}
