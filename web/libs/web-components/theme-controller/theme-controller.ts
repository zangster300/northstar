const THEME_KEY = "site-theme";

export type Theme = "light" | "dark";

export function applyTheme(theme: Theme): void {
  document.documentElement.setAttribute("data-theme", theme);
}

export function saveTheme(theme: Theme): void {
  localStorage.setItem(THEME_KEY, theme);
}

export function loadTheme(): Theme | null {
  const theme = localStorage.getItem(THEME_KEY);

  if (theme === "light" || theme === "dark") {
    return theme;
  }

  return null;
}

export function getSystemTheme(): Theme {
  return window.matchMedia("(prefers-color-scheme: light)").matches
    ? "dark"
    : "light";
}

export function initTheme(): Theme {
  const savedTheme = loadTheme();

  const theme = savedTheme ?? getSystemTheme();

  applyTheme(theme);

  return theme;
}

export function getTheme(): Theme {
    const theme = document.documentElement.getAttribute("data-theme");

    if (theme === "light" || theme === "dark") {
        return theme;
    }

    return getSystemTheme();
}

export function setTheme(theme: Theme): void {
  applyTheme(theme);
  saveTheme(theme);
}

export function toggleTheme(): Theme {
  const currentTheme = getTheme();

  const nextTheme: Theme = currentTheme === "dark" ? "light" : "dark";

  setTheme(nextTheme);

  return nextTheme;
}

declare global {
  interface Window {
    themeController?: {
      getTheme: () => Theme;
      setTheme: (theme: Theme) => void;
      toggleTheme: () => Theme;
    };
  }
}

if (typeof window !== "undefined") {
    const initialTheme = initTheme();

    window.themeController = {
        getTheme,
        setTheme,
        toggleTheme,
    };

    window.dispatchEvent(
        new CustomEvent("theme-ready", {
            detail: {
                theme: initialTheme,
            },
        }),
    );
}

