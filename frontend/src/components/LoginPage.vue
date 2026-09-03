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

      <!-- Quick Back & Register Link -->
      <div class="flex items-center space-x-3 sm:space-x-4 text-xs sm:text-sm">
        <span class="hidden sm:inline text-slate-500 dark:text-slate-400">Belum punya akun?</span>
        <button
          type="button"
          @click="$emit('navigate', 'register')"
          class="font-bold text-brand-600 dark:text-brand-400 hover:underline"
        >
          Daftar Gratis
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

    <!-- Main Content: Centered Auth Card -->
    <main class="flex-grow flex items-center justify-center px-4 py-6 sm:py-10">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 border border-slate-200/90 dark:border-slate-800 rounded-3xl shadow-xl overflow-hidden animate-scale-up">
        <!-- Card Header -->
        <div class="bg-gradient-to-r from-brand-500 via-brand-600 to-rose-500 p-6 sm:p-7 text-white text-center relative">
          <div class="w-12 h-12 mx-auto rounded-2xl bg-white/20 backdrop-blur-md flex items-center justify-center mb-2.5 shadow-inner">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
            </svg>
          </div>

          <h1 class="text-2xl sm:text-3xl font-black tracking-tight">
            Masuk ke Akun
          </h1>
          <p class="text-xs sm:text-sm text-white/85 mt-1 max-w-xs mx-auto">
            Akses seluruh alat PDF, kuota konversi, dan riwayat Anda.
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

          <!-- Google Sign In Button -->
          <button
            type="button"
            @click="handleGoogleSignIn"
            :disabled="isGoogleLoading || isLoading"
            class="w-full py-3 px-4 rounded-xl border border-slate-300 dark:border-slate-700 hover:border-slate-400 bg-white dark:bg-slate-800 hover:bg-slate-50 dark:hover:bg-slate-750 text-slate-700 dark:text-slate-200 font-bold text-xs sm:text-sm flex items-center justify-center space-x-3 transition-all transform active:scale-98 shadow-xs"
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
            <span>{{ isGoogleLoading ? 'Menghubungkan Google...' : 'Masuk dengan Google' }}</span>
          </button>

          <!-- Divider -->
          <div class="relative flex items-center justify-center my-2">
            <div class="border-t border-slate-200 dark:border-slate-800 w-full"></div>
            <span class="bg-white dark:bg-slate-900 px-3 text-[11px] font-semibold text-slate-400 uppercase tracking-wider">atau dengan akun</span>
          </div>

          <form @submit.prevent="handleSubmit" class="space-y-4">
            <!-- Identifier Field (Email / Username) -->
            <div class="space-y-1.5">
              <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                Email atau Username
              </label>
              <div class="relative">
                <input
                  type="text"
                  v-model="identifier"
                  required
                  placeholder="nama@email.com atau username"
                  class="w-full pl-10 pr-4 py-3 rounded-xl border border-slate-300 dark:border-slate-700 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                  </svg>
                </div>
              </div>
            </div>

            <!-- Password Field with Show/Hide & Forgot Link -->
            <div class="space-y-1.5">
              <div class="flex items-center justify-between">
                <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">
                  Password
                </label>
                <button
                  type="button"
                  @click="$emit('navigate', 'forgot-password')"
                  class="text-xs font-semibold text-brand-600 dark:text-brand-400 hover:underline"
                >
                  Lupa kata sandi?
                </button>
              </div>
              <div class="relative">
                <input
                  :type="showPassword ? 'text' : 'password'"
                  v-model="password"
                  required
                  placeholder="••••••••"
                  class="w-full pl-10 pr-11 py-3 rounded-xl border border-slate-300 dark:border-slate-700 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-800 text-slate-900 dark:text-white transition-all shadow-xs"
                />
                <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                  <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                  </svg>
                </div>
                <!-- Toggle show/hide -->
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
            </div>

            <!-- Remember Me Checkbox -->
            <div class="flex items-center justify-between pt-0.5">
              <label class="flex items-center space-x-2.5 cursor-pointer select-none">
                <input
                  type="checkbox"
                  v-model="rememberMe"
                  class="w-4 h-4 rounded border-slate-300 dark:border-slate-700 text-brand-600 focus:ring-brand-500 rounded-md cursor-pointer accent-brand-600"
                />
                <span class="text-xs font-semibold text-slate-600 dark:text-slate-400 hover:text-slate-800 dark:hover:text-slate-200">
                  Ingat saya di perangkat ini
                </span>
              </label>
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
              <span>{{ isLoading ? 'Memproses...' : 'Masuk ke Akun' }}</span>
            </button>

            <!-- Quick Demo Fill -->
            <div class="pt-1 text-center">
              <button
                type="button"
                @click="quickDemoFill"
                class="text-xs font-semibold text-slate-500 dark:text-slate-400 hover:text-brand-600 dark:hover:text-brand-400 transition-colors inline-flex items-center space-x-1"
              >
                <span>⚡</span>
                <span class="underline">Isi Akun Demo Otomatis</span>
              </button>
            </div>
          </form>

          <!-- Switch to Register -->
          <div class="mt-5 pt-5 border-t border-slate-100 dark:border-slate-800 text-center">
            <p class="text-xs text-slate-600 dark:text-slate-400">
              Belum memiliki akun HyperPDF?
              <button
                type="button"
                @click="$emit('navigate', 'register')"
                class="font-bold text-brand-600 dark:text-brand-400 hover:underline ml-1"
              >
                Daftar Sekarang
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
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'

type AppPage = 'home' | 'login' | 'register' | 'forgot-password' | 'reset-password' | 'profile'

const emit = defineEmits<{
  (e: 'navigate', page: AppPage): void
  (e: 'loginSuccess'): void
}>()

const { login, loginWithGoogle } = useAuth()

const identifier = ref('')
const password = ref('')
const rememberMe = ref(true)
const showPassword = ref(false)
const isLoading = ref(false)
const isGoogleLoading = ref(false)
const errorMessage = ref<string | null>(null)

function quickDemoFill() {
  identifier.value = 'alex.johnson@hyperpdf.io'
  password.value = 'HyperSecret123!'
  rememberMe.value = true
}

async function handleSubmit() {
  errorMessage.value = null
  isLoading.value = true

  try {
    await login(identifier.value, password.value, rememberMe.value)
    emit('loginSuccess')
  } catch (err: any) {
    errorMessage.value = err.message || 'Gagal masuk. Silakan periksa kembali email/username & password Anda.'
  } finally {
    isLoading.value = false
  }
}

async function handleGoogleSignIn() {
  errorMessage.value = null
  isGoogleLoading.value = true

  try {
    await loginWithGoogle()
    emit('loginSuccess')
  } catch (err: any) {
    errorMessage.value = err.message || 'Gagal masuk dengan Google.'
  } finally {
    isGoogleLoading.value = false
  }
}
</script>
