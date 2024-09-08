document.addEventListener("DOMContentLoaded", () => {
    fetchUserProfile();
    fetchAppointments();
    fetchPrescriptions();
    fetchMedicalHistory();

    // Fetch and display user profile
    function fetchUserProfile() {
        fetch('https://healthcare-app-backend.onrender.com/user/profile')
        .then(response => response.json())
        .then(data => {
            document.getElementById('profile-name').textContent = data.name;
            document.getElementById('profile-email').textContent = data.email;
        })
        .catch(error => console.error('Error fetching profile:', error));
    }

    // Fetch appointments, prescriptions, medical history (similar code)
});
