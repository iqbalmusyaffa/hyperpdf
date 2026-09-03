import { ref } from 'vue'

const THEME_STORAGE_KEY = 'hyperpdf_theme'

function getInitialTheme(): boolean {
  if (typeof window === 'undefined') return false
  const saved = localStorage.getItem(THEME_STORAGE_KEY)
  if (saved) {
    return saved === 'dark'
  }
  return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches
}

const isDark = ref<boolean>(getInitialTheme())

function applyTheme(dark: boolean) {
  isDark.value = dark
  if (typeof document !== 'undefined') {
    if (dark) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }
  try {
    localStorage.setItem(THEME_STORAGE_KEY, dark ? 'dark' : 'light')
  } catch (err) {
    console.error('Failed to save theme preference', err)
  }
}

// Initial application
if (typeof document !== 'undefined') {
  applyTheme(isDark.value)
}

export function useDarkMode() {
  function toggleDarkMode() {
    applyTheme(!isDark.value)
  }

  function setDarkMode(dark: boolean) {
    applyTheme(dark)
  }

  return {
    isDark,
    toggleDarkMode,
    setDarkMode,
  }
}
