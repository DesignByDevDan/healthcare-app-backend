const backendUrl = "https://healthcare-app-backend.onrender.com";

document.getElementById("login-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    
    const email = document.getElementById("login-email").value;
    const password = document.getElementById("login-password").value;

    try {
        const response = await fetch(`${backendUrl}/login`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ email, password }),
        });

        const data = await response.json();

        if (response.ok) {
            alert("Login successful!");
            localStorage.setItem("token", data.token); // Store token locally
            window.location.href = "protected.html"; // Redirect to protected resource page
        } else {
            alert(`Error: ${data.error}`);
        }
    } catch (error) {
        alert("An error occurred during login. Please try again later.");
    }
});
