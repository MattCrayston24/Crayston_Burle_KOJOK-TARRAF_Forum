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

//fonction gerant la photo de profil

function updateProfileIcon() {
    var profileIcon = document.getElementById('profile-icon');
    var usernameElement = document.getElementById('username');

    if(isUserLoggedIn()) {
        // Mettez à jour la source de l'image et le nom d'utilisateur
        profileIcon.src = "assets/img/profile_icon.png";
        usernameElement.innerText = getUsername();  // Supposons que getUsername() est une fonction qui renvoie le nom d'utilisateur actuel.
    } else {
        profileIcon.src = "assets/img/not_connected.png";
        usernameElement.innerText = "";
    }
}

// Vous pouvez appeler cette fonction chaque fois que l'état de connexion change, ou à chaque rechargement de la page.
updateProfileIcon();
