const backendUrl = "https://healthcare-app-backend.onrender.com";

document.getElementById("signup-form").addEventListener("submit", async (e) => {
    e.preventDefault();
    
    const name = document.getElementById("name").value;
    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    try {
        const response = await fetch(`${backendUrl}/signup`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ name, email, password }),
        });

        const data = await response.json();

        if (response.ok) {
            alert("Signup successful!");
        } else {
            alert(`Error: ${data.error}`);
        }
    } catch (error) {
        alert("An error occurred during signup. Please try again later.");
    }
});
