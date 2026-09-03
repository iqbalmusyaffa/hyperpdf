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

      <!-- Back to Login Link -->
      <div class="flex items-center space-x-3 sm:space-x-4 text-xs sm:text-sm">
        <button
          type="button"
          @click="$emit('navigate', 'login')"
          class="font-bold text-brand-600 dark:text-brand-400 hover:underline inline-flex items-center space-x-1"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          <span>Kembali ke Masuk</span>
        </button>
      </div>
    </header>

    <!-- Main Content: Centered Forgot Password Card -->
    <main class="flex-grow flex items-center justify-center px-4 py-8 sm:py-12">
      <div class="w-full max-w-md bg-white dark:bg-slate-900 border border-slate-200/90 dark:border-slate-800 rounded-3xl shadow-xl overflow-hidden animate-scale-up">
        <!-- Header -->
        <div class="bg-gradient-to-r from-brand-500 via-brand-600 to-rose-500 p-6 sm:p-8 text-white text-center relative">
          <div class="w-12 h-12 mx-auto rounded-2xl bg-white/20 backdrop-blur-md flex items-center justify-center mb-3 shadow-inner">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z" />
            </svg>
          </div>

          <h1 class="text-2xl sm:text-3xl font-black tracking-tight">
            Lupa Kata Sandi?
          </h1>
          <p class="text-xs sm:text-sm text-white/85 mt-1.5 max-w-xs mx-auto">
            Masukkan email atau username Anda untuk menerima instruksi pemulihan kata sandi.
          </p>
        </div>

        <!-- Body -->
        <div class="p-6 sm:p-8">
          <!-- Success State -->
          <div v-if="isSuccess" class="text-center space-y-5 animate-fade-in">
            <div class="w-16 h-16 mx-auto rounded-2xl bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-300 flex items-center justify-center shadow-inner">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
            </div>

            <div class="space-y-1.5">
              <h3 class="text-lg font-bold text-slate-900 dark:text-white">Tautan Pemulihan Terkirim!</h3>
              <p class="text-xs sm:text-sm text-slate-600 dark:text-slate-300 leading-relaxed">
                Kami telah mengirimkan instruksi pemulihan ke <span class="font-bold text-slate-800 dark:text-white">{{ identifier }}</span>. Silakan periksa folder kotak masuk atau spam Anda.
              </p>
            </div>

            <div class="space-y-2">
              <button
                type="button"
                @click="$emit('navigate', 'reset-password')"
                class="w-full py-3.5 rounded-xl bg-gradient-to-r from-brand-500 to-rose-500 text-white font-bold text-sm shadow-md shadow-brand-500/25 transition-all transform active:scale-98"
              >
                Lanjut Buat Kata Sandi Baru
              </button>

              <button
                type="button"
                @click="$emit('navigate', 'login')"
                class="w-full py-3 rounded-xl border border-slate-300 dark:border-slate-700 hover:bg-slate-50 dark:hover:bg-slate-800 text-slate-700 dark:text-slate-200 font-bold text-xs transition-colors"
              >
                Kembali ke Halaman Masuk
              </button>
            </div>
          </div>

          <!-- Form State -->
          <form v-else @submit.prevent="handleSubmit" class="space-y-4">
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

            <!-- Identifier Field -->
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
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                </div>
              </div>
            </div>

            <!-- Submit Button -->
            <button
              type="submit"
              :disabled="isLoading"
              class="w-full mt-2 py-3.5 rounded-xl bg-gradient-to-r from-brand-500 to-rose-500 hover:from-brand-600 hover:to-rose-600 disabled:opacity-50 text-white font-bold text-sm shadow-md shadow-brand-500/25 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
            >
              <svg v-if="isLoading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path>
              </svg>
              <span>{{ isLoading ? 'Mengirim...' : 'Kirim Tautan Pemulihan' }}</span>
            </button>
          </form>

          <!-- Back to Login -->
          <div class="mt-6 pt-5 border-t border-slate-100 dark:border-slate-800 text-center">
            <p class="text-xs text-slate-600 dark:text-slate-400">
              Sudah ingat kata sandi Anda?
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
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'

type AppPage = 'home' | 'login' | 'register' | 'forgot-password' | 'reset-password' | 'profile'

const emit = defineEmits<{
  (e: 'navigate', page: AppPage): void
}>()

const { forgotPassword } = useAuth()

const identifier = ref('')
const isLoading = ref(false)
const isSuccess = ref(false)
const errorMessage = ref<string | null>(null)

async function handleSubmit() {
  errorMessage.value = null
  isLoading.value = true

  try {
    await forgotPassword(identifier.value)
    isSuccess.value = true
  } catch (err: any) {
    errorMessage.value = err.message || 'Gagal mengirim instruksi pemulihan. Pastikan data benar.'
  } finally {
    isLoading.value = false
  }
}
</script>
