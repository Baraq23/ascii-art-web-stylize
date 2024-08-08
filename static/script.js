const body = document.querySelector("body")
const container = document.querySelector(".container")
const themeToggler = document.querySelector(".theme-modes")

function toggleTheme() {
  body.classList.toggle('dark-mode');
  container.classList.toggle('dark-mode');
  themeToggler.classList.toggle('active');
}

themeToggler.addEventListener('click', toggleTheme);