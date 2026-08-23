// Minimal bootstrap code to apply the user's saved theme before the page renders, preventing a flash of the wrong theme.
const THEME_KEY = "site-theme";

const saved = localStorage.getItem(THEME_KEY);

const theme =
  saved === "light" || saved === "dark"
    ? saved
    : window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light";

document.documentElement.setAttribute("data-theme", theme);
