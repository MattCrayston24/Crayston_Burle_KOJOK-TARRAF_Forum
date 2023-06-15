// Ouvre la modale
function openModal() {
    document.getElementById("modal").style.display = "block";
}

// Ferme la modale
function closeModal() {
    document.getElementById("modal").style.display = "none";
}

// Bascule entre les formulaires Connexion et Inscription
function switchForm() {
    var loginForm = document.getElementById("login-form");
    var signupForm = document.getElementById("signup-form");

    if (loginForm.style.display === "none") {
        signupForm.style.display = "none";
        loginForm.style.display = "block";
    } else {
        loginForm.style.display = "none";
        signupForm.style.display = "block";
    }
}
