import { ref, computed } from 'vue'

export interface User {
  id: string
  name: string
  email: string
  plan: 'free' | 'pro'
  avatar?: string
  joinedAt: string
}

const USER_STORAGE_KEY = 'hyperpdf_auth_user'
const USAGE_STORAGE_KEY = 'hyperpdf_daily_usage'
const MAX_FREE_DAILY_LIMIT = 1000

const user = ref<User | null>(loadInitialUser())
const dailyUsage = ref<number>(loadInitialDailyUsage())

function getTodayKey(): string {
  const now = new Date()
  return `${now.getFullYear()}-${now.getMonth() + 1}-${now.getDate()}`
}

function loadInitialDailyUsage(): number {
  try {
    const raw = localStorage.getItem(USAGE_STORAGE_KEY)
    if (raw) {
      const data = JSON.parse(raw)
      if (data.date === getTodayKey()) {
        return data.count || 0
      }
    }
  } catch (err) {
    console.error('Failed to load daily usage', err)
  }
  return 2 // Default mock starting with 2 used (998 left)
}

function saveDailyUsage(count: number) {
  dailyUsage.value = count
  try {
    localStorage.setItem(
      USAGE_STORAGE_KEY,
      JSON.stringify({
        date: getTodayKey(),
        count: count,
      })
    )
  } catch (err) {
    console.error('Failed to save daily usage', err)
  }
}

function loadInitialUser(): User | null {
  try {
    const saved = localStorage.getItem(USER_STORAGE_KEY)
    if (saved) {
      return JSON.parse(saved)
    }
  } catch (err) {
    console.error('Failed to load user from localStorage', err)
  }
  return null
}

function saveUser(u: User | null) {
  user.value = u
  if (u) {
    localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(u))
  } else {
    localStorage.removeItem(USER_STORAGE_KEY)
  }
}

export function useAuth() {
  const isLoggedIn = computed(() => !!user.value)
  const isPro = computed(() => user.value?.plan === 'pro')
  const currentPlan = computed(() => user.value?.plan || 'free')

  const maxDailyLimit = MAX_FREE_DAILY_LIMIT
  const remainingQuota = computed(() => {
    if (isPro.value) return 999999
    return Math.max(0, MAX_FREE_DAILY_LIMIT - dailyUsage.value)
  })

  function recordCompressionUsage() {
    if (!isPro.value) {
      saveDailyUsage(dailyUsage.value + 1)
    }
  }

  function login(email: string, _password: string): Promise<User> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (!email || !email.includes('@')) {
          reject(new Error('Please enter a valid email address'))
          return
        }

        const loggedUser: User = {
          id: 'usr_' + Math.random().toString(36).substring(2, 9),
          name: email.split('@')[0],
          email: email,
          plan: 'free',
          joinedAt: new Date().toISOString(),
        }

        saveUser(loggedUser)
        resolve(loggedUser)
      }, 500)
    })
  }

  function register(name: string, email: string, _password: string): Promise<User> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (!name || name.trim().length < 2) {
          reject(new Error('Please enter a valid name (at least 2 characters)'))
          return
        }
        if (!email || !email.includes('@')) {
          reject(new Error('Please enter a valid email address'))
          return
        }

        const newUser: User = {
          id: 'usr_' + Math.random().toString(36).substring(2, 9),
          name: name.trim(),
          email: email.trim(),
          plan: 'free',
          joinedAt: new Date().toISOString(),
        }

        saveUser(newUser)
        resolve(newUser)
      }, 500)
    })
  }

  function logout() {
    saveUser(null)
  }

  function upgradeToPro(): Promise<void> {
    return new Promise((resolve) => {
      setTimeout(() => {
        if (user.value) {
          user.value.plan = 'pro'
          saveUser({ ...user.value, plan: 'pro' })
        } else {
          const proGuest: User = {
            id: 'usr_' + Math.random().toString(36).substring(2, 9),
            name: 'Pro Member',
            email: 'pro.user@hyperpdf.io',
            plan: 'pro',
            joinedAt: new Date().toISOString(),
          }
          saveUser(proGuest)
        }
        resolve()
      }, 600)
    })
  }

  return {
    user,
    isLoggedIn,
    isPro,
    currentPlan,
    dailyUsage,
    maxDailyLimit,
    remainingQuota,
    recordCompressionUsage,
    login,
    register,
    logout,
    upgradeToPro,
  }
}
