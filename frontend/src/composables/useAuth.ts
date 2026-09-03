import { ref, computed } from 'vue'

export interface User {
  id: string
  name: string
  username: string
  email: string
  plan: 'free' | 'pro'
  avatar?: string
  authProvider?: 'password' | 'google'
  joinedAt: string
}

const USER_STORAGE_KEY = 'hyperpdf_auth_user'
const USAGE_STORAGE_KEY = 'hyperpdf_daily_usage'
const REMEMBER_STORAGE_KEY = 'hyperpdf_remember_me'
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
    const savedLocal = localStorage.getItem(USER_STORAGE_KEY)
    if (savedLocal) {
      return JSON.parse(savedLocal)
    }
    const savedSession = sessionStorage.getItem(USER_STORAGE_KEY)
    if (savedSession) {
      return JSON.parse(savedSession)
    }
  } catch (err) {
    console.error('Failed to load user from storage', err)
  }
  return null
}

function saveUser(u: User | null, remember: boolean = true) {
  user.value = u
  if (u) {
    if (remember) {
      localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(u))
      localStorage.setItem(REMEMBER_STORAGE_KEY, 'true')
      sessionStorage.removeItem(USER_STORAGE_KEY)
    } else {
      sessionStorage.setItem(USER_STORAGE_KEY, JSON.stringify(u))
      localStorage.removeItem(USER_STORAGE_KEY)
      localStorage.removeItem(REMEMBER_STORAGE_KEY)
    }
  } else {
    localStorage.removeItem(USER_STORAGE_KEY)
    localStorage.removeItem(REMEMBER_STORAGE_KEY)
    sessionStorage.removeItem(USER_STORAGE_KEY)
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

  function login(identifier: string, _password: string, remember: boolean = true): Promise<User> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        const cleanIdentifier = identifier.trim()
        if (!cleanIdentifier || cleanIdentifier.length < 3) {
          reject(new Error('Masukkan email atau username yang valid (minimal 3 karakter)'))
          return
        }

        const isEmail = cleanIdentifier.includes('@')
        const uname = isEmail ? cleanIdentifier.split('@')[0] : cleanIdentifier
        const mail = isEmail ? cleanIdentifier : `${cleanIdentifier}@hyperpdf.io`

        const loggedUser: User = {
          id: 'usr_' + Math.random().toString(36).substring(2, 9),
          name: uname.charAt(0).toUpperCase() + uname.slice(1),
          username: uname.toLowerCase().replace(/[^a-z0-9_]/g, ''),
          email: mail,
          plan: 'free',
          authProvider: 'password',
          joinedAt: new Date().toISOString(),
        }

        saveUser(loggedUser, remember)
        resolve(loggedUser)
      }, 500)
    })
  }

  function register(name: string, username: string, email: string, _password: string): Promise<User> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (!name || name.trim().length < 2) {
          reject(new Error('Masukkan nama lengkap yang valid (minimal 2 karakter)'))
          return
        }
        if (!username || username.trim().length < 3) {
          reject(new Error('Username minimal 3 karakter'))
          return
        }
        if (!email || !email.includes('@')) {
          reject(new Error('Masukkan alamat email yang valid'))
          return
        }

        const newUser: User = {
          id: 'usr_' + Math.random().toString(36).substring(2, 9),
          name: name.trim(),
          username: username.trim().toLowerCase().replace(/[^a-z0-9_]/g, ''),
          email: email.trim(),
          plan: 'free',
          authProvider: 'password',
          joinedAt: new Date().toISOString(),
        }

        saveUser(newUser, true)
        resolve(newUser)
      }, 500)
    })
  }

  function loginWithGoogle(): Promise<User> {
    return new Promise((resolve) => {
      setTimeout(() => {
        const googleUser: User = {
          id: 'usr_g_' + Math.random().toString(36).substring(2, 9),
          name: 'Alex Johnson (Google)',
          username: 'alex_google',
          email: 'alex.google@gmail.com',
          plan: 'free',
          authProvider: 'google',
          avatar: 'https://lh3.googleusercontent.com/a/default-user',
          joinedAt: new Date().toISOString(),
        }

        saveUser(googleUser, true)
        resolve(googleUser)
      }, 600)
    })
  }

  function forgotPassword(identifier: string): Promise<string> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        const cleanId = identifier.trim()
        if (!cleanId || cleanId.length < 3) {
          reject(new Error('Masukkan email atau username yang terdaftar'))
          return
        }
        resolve(cleanId.includes('@') ? cleanId : `${cleanId}@email.com`)
      }, 500)
    })
  }

  function confirmResetPassword(newPassword: string): Promise<void> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (!newPassword || newPassword.length < 6) {
          reject(new Error('Kata sandi baru minimal 6 karakter'))
          return
        }
        resolve()
      }, 600)
    })
  }

  function updateProfile(name: string, username: string): Promise<User> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (!name || name.trim().length < 2) {
          reject(new Error('Nama lengkap minimal 2 karakter'))
          return
        }
        if (!username || username.trim().length < 3) {
          reject(new Error('Username minimal 3 karakter'))
          return
        }

        if (user.value) {
          const updatedUser: User = {
            ...user.value,
            name: name.trim(),
            username: username.trim().toLowerCase().replace(/[^a-z0-9_]/g, ''),
          }
          saveUser(updatedUser, true)
          resolve(updatedUser)
        } else {
          reject(new Error('Tidak ada sesi pengguna aktif'))
        }
      }, 400)
    })
  }

  function changePassword(oldPassword: string, newPassword: string): Promise<void> {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        if (!oldPassword) {
          reject(new Error('Masukkan kata sandi lama Anda'))
          return
        }
        if (!newPassword || newPassword.length < 6) {
          reject(new Error('Kata sandi baru minimal 6 karakter'))
          return
        }
        if (oldPassword === newPassword) {
          reject(new Error('Kata sandi baru tidak boleh sama dengan kata sandi lama'))
          return
        }
        resolve()
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
          const proUser: User = { ...user.value, plan: 'pro' }
          saveUser(proUser, true)
        } else {
          const proGuest: User = {
            id: 'usr_' + Math.random().toString(36).substring(2, 9),
            name: 'Pro Member',
            username: 'pro_member',
            email: 'pro.user@hyperpdf.io',
            plan: 'pro',
            joinedAt: new Date().toISOString(),
          }
          saveUser(proGuest, true)
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
    loginWithGoogle,
    forgotPassword,
    confirmResetPassword,
    updateProfile,
    changePassword,
    logout,
    upgradeToPro,
  }
}
