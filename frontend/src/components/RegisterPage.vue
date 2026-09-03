<template>
  <div class="min-h-screen w-full flex flex-col justify-between bg-slate-50 dark:bg-slate-950 text-slate-800 dark:text-slate-100 transition-colors">
    <!-- Clean Minimal Auth Header -->
    <header class="w-full max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 h-20 flex items-center justify-between">
      <!-- Brand Logo -->
      <div
        @click="$emit('navigate', 'home')"
        class="flex items-center space-x-2.5 cursor-pointer group select-none"
      >
        <div class="w-9 h-9 rounded-xl bg-gradient-to-tr from-brand-600 via-brand-500 to-rose-400 text-white flex items-center justify-center shadow-md shadow-brand-500/25 group-hover:scale-105 transition-transform duration-200">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 animate-pulse" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
          </svg>
        </div>
        <span class="text-xl sm:text-2xl font-black tracking-tight text-slate-900 dark:text-white">
          Hyper<span class="text-transparent bg-clip-text bg-gradient-to-r from-brand-500 to-rose-500">PDF</span>
        </span>
      </div>

      <!-- Quick Back & Login Link -->
      <div class="flex items-center space-x-3 sm:space-x-4 text-xs sm:text-sm">
        <span class="hidden sm:inline text-slate-500 dark:text-slate-400">Sudah punya akun?</span>
        <button
          type="button"
          @click="$emit('navigate', 'login')"
          class="font-bold text-brand-600 dark:text-brand-400 hover:underline"
        >
          Masuk
        </button>
        <span class="text-slate-300 dark:text-slate-700">|</span>
        <button
          type="button"
          @click="$emit('navigate', 'home')"
          class="inline-flex items-center space-x-1 font-semibold text-slate-600 dark:text-slate-300 hover:text-slate-900 dark:hover:text-white transition-colors"
        >
          <span>Beranda</span>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      </div>
    </header>

    <!-- Main Content: Centered Register Card -->
    <main class="flex-grow flex items-center justify-center px-4 py-6 sm:py-10">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 border border-slate-200/90 dark:border-slate-800 rounded-3xl shadow-xl overflow-hidden animate-scale-up">
        <!-- Card Header -->
        <div class="bg-gradient-to-r from-brand-500 via-brand-600 to-rose-500 p-6 sm:p-7 text-white text-center relative">
          <div class="w-12 h-12 mx-auto rounded-2xl bg-white/20 backdrop-blur-md flex items-center justify-center mb-2.5 shadow-inner">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
            </svg>
          </div>

          <h1 class="text-2xl sm:text-3xl font-black tracking-tight">
            Buat Akun Baru
          </h1>
          <p class="text-xs sm:text-sm text-white/85 mt-1 max-w-xs mx-auto">
            Daftar gratis dan nikmati pemrosesan PDF super cepat.
          </p>
        </div>

        <!-- Form Body -->
        <div class="p-6 sm:p-8 space-y-4">
          <!-- Error Alert -->
          <div
            v-if="errorMessage"
            class="p-3.5 rounded-2xl bg-red-50 dark:bg-red-900/30 border border-red-200 dark:border-red-800 text-red-700 dark:text-red-300 text-xs font-medium flex items-center space-x-2.5 animate-shake"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 shrink-0 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span class="flex-1">{{ errorMessage }}</span>
          </div>

          <!-- Google Sign Up Button -->
          <button
            type="button"
            @click="handleGoogleSignUp"
            :disabled="isGoogleLoading || isLoading"
            class="w-full py-3 px-4 rounded-xl border border-slate-300 dark:border-slate-700 hover:border-slate-400 bg-white dark:bg-slate-800 hover:bg-slate-50 text-slate-700 dark:text-slate-200 font-bold text-xs sm:text-sm flex items-center justify-center space-x-3 transition-all transform active:scale-98 shadow-xs"
          >
            <svg v-if="isGoogleLoading" class="animate-spin h-4 w-4 text-brand-600" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path>
            </svg>
            <svg v-else class="w-4 h-4 shrink-0" viewBox="0 0 24 24">
              <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
              <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
              <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.06H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.94l2.85-2.22.81-.63z"/>
              <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.06l3.66 2.84c.87-2.6 3.3-4.52 6.16-4.52z"/>
            </svg>
            <span>{{ isGoogleLoading ? 'Menghubungkan Google...' : 'Daftar dengan Google' }}</span>
          </button>

          <!-- Divider -->
          <div class="relative flex items-center justify-center my-2">
            <div class="border-t border-slate-200 dark:border-slate-800 w-full"></div>
            <span class="bg-white dark:bg-slate-900 px-3 text-[11px] font-semibold text-slate-400 uppercase tracking-wider">atau dengan formulir</span>
          </div>

          <form @submit.prevent="handleSubmit" class="space-y-3">
            <!-- Full Name Field -->
            <div class="space-y-1">
              <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                Nama Lengkap
              </label>
              <div class="relative">
                <input
                  type="text"
                  v-model="name"
                  required
                  placeholder="cth. Alex Johnson"
                  class="w-full pl-10 pr-4 py-2.5 rounded-xl border border-slate-300 dark:border-slate-700 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                </div>
              </div>
            </div>

            <!-- Username Field -->
            <div class="space-y-1">
              <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                Username
              </label>
              <div class="relative">
                <input
                  type="text"
                  v-model="username"
                  required
                  placeholder="cth. alexjohnson"
                  class="w-full pl-10 pr-4 py-2.5 rounded-xl border border-slate-300 dark:border-slate-700 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400 text-xs font-bold">
                  @
                </div>
              </div>
            </div>

            <!-- Email Field -->
            <div class="space-y-1">
              <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                Email Address
              </label>
              <div class="relative">
                <input
                  type="email"
                  v-model="email"
                  required
                  placeholder="nama@email.com"
                  class="w-full pl-10 pr-4 py-2.5 rounded-xl border border-slate-300 dark:border-slate-700 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </div>
              </div>
            </div>

            <!-- Password Field with Show/Hide -->
            <div class="space-y-1">
              <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                Password
              </label>
              <div class="relative">
                <input
                  :type="showPassword ? 'text' : 'password'"
                  v-model="password"
                  required
                  minlength="6"
                  placeholder="Minimal 6 karakter"
                  class="w-full pl-10 pr-11 py-2.5 rounded-xl border border-slate-300 dark:border-slate-700 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                  </svg>
                </div>
                <button
                  type="button"
                  @click="showPassword = !showPassword"
                  class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors"
                  :title="showPassword ? 'Sembunyikan Password' : 'Lihat Password'"
                  tabindex="-1"
                >
                  <svg v-if="!showPassword" xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-brand-600 dark:text-brand-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
                  </svg>
                </button>
              </div>

              <!-- Password Strength Meter -->
              <div v-if="password" class="pt-1.5 space-y-1 animate-fade-in">
                <div class="flex items-center justify-between text-[11px] font-semibold">
                  <span class="text-slate-500 dark:text-slate-400">Kekuatan Password:</span>
                  <span :class="strengthTextColor">{{ strengthLabel }}</span>
                </div>
                <div class="grid grid-cols-4 gap-1.5 h-1.5 w-full bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden p-0.5">
                  <div class="h-full rounded-full transition-all duration-300" :class="passwordScore >= 1 ? strengthBgColor : 'bg-transparent'"></div>
                  <div class="h-full rounded-full transition-all duration-300" :class="passwordScore >= 2 ? strengthBgColor : 'bg-transparent'"></div>
                  <div class="h-full rounded-full transition-all duration-300" :class="passwordScore >= 3 ? strengthBgColor : 'bg-transparent'"></div>
                  <div class="h-full rounded-full transition-all duration-300" :class="passwordScore >= 4 ? strengthBgColor : 'bg-transparent'"></div>
                </div>
              </div>
            </div>

            <!-- Confirm Password Field with Show/Hide -->
            <div class="space-y-1">
              <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                Konfirmasi Password
              </label>
              <div class="relative">
                <input
                  :type="showConfirmPassword ? 'text' : 'password'"
                  v-model="confirmPassword"
                  required
                  placeholder="Ulangi kata sandi Anda"
                  class="w-full pl-10 pr-11 py-2.5 rounded-xl border focus:outline-none focus:ring-2 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                  :class="[
                    confirmPassword && confirmPassword !== password
                      ? 'border-red-300 dark:border-red-800 focus:ring-red-500'
                      : 'border-slate-300 dark:border-slate-700 focus:ring-brand-500'
                  ]"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                  </svg>
                </div>
                <button
                  type="button"
                  @click="showConfirmPassword = !showConfirmPassword"
                  class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-slate-400 hover:text-slate-600 dark:hover:text-slate-200 transition-colors"
                  :title="showConfirmPassword ? 'Sembunyikan Password' : 'Lihat Password'"
                  tabindex="-1"
                >
                  <svg v-if="!showConfirmPassword" xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-brand-600 dark:text-brand-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
                  </svg>
                </button>
              </div>
              <p v-if="confirmPassword && confirmPassword !== password" class="text-[11px] text-red-500 font-medium">
                Kata sandi tidak cocok.
              </p>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="isLoading || isGoogleLoading"
              class="w-full mt-2 py-3.5 rounded-xl bg-gradient-to-r from-brand-500 to-rose-500 hover:from-brand-600 hover:to-rose-600 disabled:opacity-50 text-white font-bold text-sm shadow-md shadow-brand-500/25 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
            >
              <svg v-if="isLoading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path>
              </svg>
              <span>{{ isLoading ? 'Mendaftarkan Akun...' : 'Buat Akun Gratis' }}</span>
            </button>

            <!-- Quick Demo Fill -->
            <div class="pt-1 text-center">
              <button
                type="button"
                @click="quickDemoFill"
                class="text-xs font-semibold text-slate-500 dark:text-slate-400 hover:text-brand-600 dark:hover:text-brand-400 transition-colors inline-flex items-center space-x-1"
              >
                <span>⚡</span>
                <span class="underline">Isi Data Demo Otomatis</span>
              </button>
            </div>
          </form>

          <!-- Switch to Login -->
          <div class="mt-5 pt-5 border-t border-slate-100 dark:border-slate-800 text-center">
            <p class="text-xs text-slate-600 dark:text-slate-400">
              Sudah memiliki akun HyperPDF?
              <button
                type="button"
                @click="$emit('navigate', 'login')"
                class="font-bold text-brand-600 dark:text-brand-400 hover:underline ml-1"
              >
                Masuk di sini
              </button>
            </p>
          </div>
        </div>
      </div>
    </main>

    <!-- Clean Minimal Footer -->
    <footer class="w-full py-6 text-center text-xs text-slate-400 dark:text-slate-600">
      <p>© 2026 HyperPDF. All rights reserved. • Pemrosesan Dokumen Aman & Terenkripsi</p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useAuth } from '../composables/useAuth'

type AppPage = 'home' | 'login' | 'register' | 'forgot-password' | 'reset-password' | 'profile'

const emit = defineEmits<{
  (e: 'navigate', page: AppPage): void
  (e: 'registerSuccess'): void
}>()

const { register, loginWithGoogle } = useAuth()

const name = ref('')
const username = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const isLoading = ref(false)
const isGoogleLoading = ref(false)
const errorMessage = ref<string | null>(null)

// Password strength calculation
const passwordScore = computed(() => {
  const val = password.value
  if (!val) return 0
  let score = 0
  if (val.length >= 6) score += 1
  if (val.length >= 8) score += 1
  if (/[A-Z]/.test(val) && /[a-z]/.test(val)) score += 1
  if (/[0-9]/.test(val) || /[^A-Za-z0-9]/.test(val)) score += 1
  return score
})

const strengthLabel = computed(() => {
  switch (passwordScore.value) {
    case 1:
      return 'Lemah'
    case 2:
      return 'Cukup'
    case 3:
      return 'Kuat'
    case 4:
      return 'Sangat Kuat'
    default:
      return 'Terlalu Pendek'
  }
})

const strengthTextColor = computed(() => {
  switch (passwordScore.value) {
    case 1:
      return 'text-red-500'
    case 2:
      return 'text-amber-500'
    case 3:
      return 'text-blue-500'
    case 4:
      return 'text-emerald-500'
    default:
      return 'text-slate-400'
  }
})

const strengthBgColor = computed(() => {
  switch (passwordScore.value) {
    case 1:
      return 'bg-red-500'
    case 2:
      return 'bg-amber-500'
    case 3:
      return 'bg-blue-500'
    case 4:
      return 'bg-emerald-500'
    default:
      return 'bg-slate-300'
  }
})

function quickDemoFill() {
  name.value = 'Demo Creator'
  username.value = 'democreator'
  email.value = 'creator@hyperpdf.io'
  password.value = 'HyperSecret123!'
  confirmPassword.value = 'HyperSecret123!'
}

async function handleSubmit() {
  errorMessage.value = null

  if (password.value !== confirmPassword.value) {
    errorMessage.value = 'Konfirmasi password tidak cocok dengan password yang dimasukkan.'
    return
  }

  isLoading.value = true

  try {
    await register(name.value, username.value, email.value, password.value)
    emit('registerSuccess')
  } catch (err: any) {
    errorMessage.value = err.message || 'Gagal membuat akun. Pastikan data terisi dengan benar.'
  } finally {
    isLoading.value = false
  }
}

async function handleGoogleSignUp() {
  errorMessage.value = null
  isGoogleLoading.value = true

  try {
    await loginWithGoogle()
    emit('registerSuccess')
  } catch (err: any) {
    errorMessage.value = err.message || 'Gagal mendaftar dengan Google.'
  } finally {
    isGoogleLoading.value = false
  }
}
</script>
