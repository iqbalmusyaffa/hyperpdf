<template>
  <div class="min-h-screen w-full flex flex-col justify-between bg-slate-50 dark:bg-slate-950 text-slate-800 dark:text-slate-100 transition-colors">
    <!-- Clean Minimal Header -->
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

      <!-- Action Navigation -->
      <div class="flex items-center space-x-3 sm:space-x-4 text-xs sm:text-sm">
        <button
          type="button"
          @click="$emit('navigate', 'home')"
          class="inline-flex items-center space-x-1.5 font-semibold text-slate-600 dark:text-slate-300 hover:text-slate-900 dark:hover:text-white transition-colors"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
          </svg>
          <span>Kembali ke Alat PDF</span>
        </button>

        <span class="text-slate-300 dark:text-slate-700">|</span>

        <button
          type="button"
          @click="handleLogout"
          class="inline-flex items-center space-x-1 font-semibold text-red-600 dark:text-red-400 hover:text-red-700 transition-colors"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
          <span>Keluar</span>
        </button>
      </div>
    </header>

    <!-- Main Content Container -->
    <main class="flex-grow max-w-4xl mx-auto w-full px-4 sm:px-6 lg:px-8 py-6 sm:py-10">
      <!-- Profile Header Banner -->
      <div class="bg-gradient-to-r from-slate-900 via-slate-800 to-brand-900 rounded-3xl p-6 sm:p-8 text-white shadow-xl mb-8 relative overflow-hidden">
        <div class="relative z-10 flex flex-col sm:flex-row items-center sm:items-start space-y-4 sm:space-y-0 sm:space-x-6 text-center sm:text-left">
          <!-- Avatar -->
          <div class="w-20 h-20 rounded-2xl bg-gradient-to-tr from-brand-500 to-rose-400 text-white flex items-center justify-center text-3xl font-black shadow-lg shadow-brand-500/30 shrink-0">
            {{ user?.name?.charAt(0).toUpperCase() || 'U' }}
          </div>

          <!-- User Details -->
          <div class="space-y-1.5 flex-1">
            <div class="flex flex-wrap items-center justify-center sm:justify-start gap-2">
              <h1 class="text-2xl sm:text-3xl font-black tracking-tight">{{ user?.name }}</h1>
              <span
                class="px-2.5 py-0.5 text-xs font-extrabold rounded-full"
                :class="isPro ? 'bg-amber-400 text-slate-950 shadow-xs' : 'bg-white/20 text-white'"
              >
                {{ isPro ? '⭐ PRO MEMBER' : 'FREE MEMBER' }}
              </span>
            </div>
            <p class="text-brand-200 text-sm font-semibold">@{{ user?.username || 'user' }}</p>
            <p class="text-slate-300 text-xs flex items-center justify-center sm:justify-start space-x-1.5">
              <span>📧 {{ user?.email }}</span>
              <span>•</span>
              <span>📅 Bergabung {{ formatDate(user?.joinedAt) }}</span>
            </p>
          </div>

          <!-- Upgrade / Status Button -->
          <button
            v-if="!isPro"
            type="button"
            @click="$emit('openPricing')"
            class="px-5 py-2.5 rounded-xl bg-gradient-to-r from-amber-500 to-rose-500 hover:from-amber-600 hover:to-rose-600 text-white font-extrabold text-xs shadow-md shadow-amber-500/20 transition-all transform active:scale-95 shrink-0"
          >
            ✨ Upgrade ke PRO (Rp 49rb)
          </button>
        </div>

        <!-- Decorative background circle -->
        <div class="absolute -bottom-10 -right-10 w-48 h-48 bg-brand-500/10 rounded-full blur-2xl"></div>
      </div>

      <!-- Navigation Tabs -->
      <div class="flex border-b border-slate-200 dark:border-slate-800 mb-8 space-x-2 sm:space-x-4 overflow-x-auto no-scrollbar">
        <button
          type="button"
          @click="activeTab = 'profile'"
          class="pb-3 px-3 text-xs sm:text-sm font-bold border-b-2 transition-all shrink-0 flex items-center space-x-2"
          :class="activeTab === 'profile' ? 'border-brand-500 text-brand-600 dark:text-brand-400' : 'border-transparent text-slate-500 hover:text-slate-800 dark:hover:text-slate-200'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
          </svg>
          <span>Informasi Profil</span>
        </button>

        <button
          type="button"
          @click="activeTab = 'security'"
          class="pb-3 px-3 text-xs sm:text-sm font-bold border-b-2 transition-all shrink-0 flex items-center space-x-2"
          :class="activeTab === 'security' ? 'border-brand-500 text-brand-600 dark:text-brand-400' : 'border-transparent text-slate-500 hover:text-slate-800 dark:hover:text-slate-200'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
          </svg>
          <span>Keamanan & Kata Sandi</span>
        </button>

        <button
          type="button"
          @click="activeTab = 'history'"
          class="pb-3 px-3 text-xs sm:text-sm font-bold border-b-2 transition-all shrink-0 flex items-center space-x-2"
          :class="activeTab === 'history' ? 'border-brand-500 text-brand-600 dark:text-brand-400' : 'border-transparent text-slate-500 hover:text-slate-800 dark:hover:text-slate-200'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span>Riwayat Aktivitas ({{ historyItems.length }})</span>
        </button>

        <button
          type="button"
          @click="activeTab = 'subscription'"
          class="pb-3 px-3 text-xs sm:text-sm font-bold border-b-2 transition-all shrink-0 flex items-center space-x-2"
          :class="activeTab === 'subscription' ? 'border-brand-500 text-brand-600 dark:text-brand-400' : 'border-transparent text-slate-500 hover:text-slate-800 dark:hover:text-slate-200'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z" />
          </svg>
          <span>Penggunaan & Kuota</span>
        </button>
      </div>

      <!-- TAB 1: EDIT PROFILE -->
      <div v-if="activeTab === 'profile'" class="bg-white dark:bg-slate-800 border border-slate-200/90 dark:border-slate-700 rounded-3xl p-6 sm:p-8 shadow-sm">
        <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-1">Edit Informasi Akun</h3>
        <p class="text-xs text-slate-500 dark:text-slate-400 mb-6">Perbarui nama lengkap dan username Anda.</p>

        <!-- Alert messages -->
        <div v-if="profileSuccess" class="mb-5 p-3.5 rounded-2xl bg-emerald-50 dark:bg-emerald-900/30 border border-emerald-200 dark:border-emerald-800 text-emerald-700 dark:text-emerald-300 text-xs font-semibold flex items-center space-x-2">
          <span>✓</span>
          <span>{{ profileSuccess }}</span>
        </div>
        <div v-if="profileError" class="mb-5 p-3.5 rounded-2xl bg-red-50 dark:bg-red-900/30 border border-red-200 dark:border-red-800 text-red-700 dark:text-red-300 text-xs font-semibold flex items-center space-x-2">
          <span>✕</span>
          <span>{{ profileError }}</span>
        </div>

        <form @submit.prevent="handleSaveProfile" class="space-y-4 max-w-lg">
          <!-- Full Name -->
          <div class="space-y-1.5">
            <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Nama Lengkap</label>
            <input
              type="text"
              v-model="nameInput"
              required
              class="w-full px-4 py-2.5 rounded-xl border border-slate-300 dark:border-slate-600 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-900 text-slate-900 dark:text-white shadow-xs"
            />
          </div>

          <!-- Username -->
          <div class="space-y-1.5">
            <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Username</label>
            <div class="relative">
              <input
                type="text"
                v-model="usernameInput"
                required
                class="w-full pl-8 pr-4 py-2.5 rounded-xl border border-slate-300 dark:border-slate-600 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-900 text-slate-900 dark:text-white shadow-xs"
              />
              <span class="absolute inset-y-0 left-0 pl-3 flex items-center text-slate-400 font-bold text-xs">@</span>
            </div>
          </div>

          <!-- Email (Readonly) -->
          <div class="space-y-1.5">
            <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Email (Terdaftar)</label>
            <input
              type="email"
              :value="user?.email"
              disabled
              class="w-full px-4 py-2.5 rounded-xl border border-slate-200 dark:border-slate-700 text-sm bg-slate-100 dark:bg-slate-900 text-slate-500 dark:text-slate-400 cursor-not-allowed"
            />
            <p class="text-[11px] text-slate-400">Alamat email utama digunakan untuk autentikasi dan pemulihan.</p>
          </div>

          <button
            type="submit"
            :disabled="isProfileSaving"
            class="mt-4 px-6 py-3 rounded-xl bg-slate-900 dark:bg-brand-600 hover:bg-slate-800 dark:hover:bg-brand-700 text-white font-bold text-xs sm:text-sm shadow-md transition-all disabled:opacity-50"
          >
            {{ isProfileSaving ? 'Menyimpan...' : 'Simpan Perubahan' }}
          </button>
        </form>
      </div>

      <!-- TAB 2: SECURITY & PASSWORD -->
      <div v-else-if="activeTab === 'security'" class="bg-white dark:bg-slate-800 border border-slate-200/90 dark:border-slate-700 rounded-3xl p-6 sm:p-8 shadow-sm">
        <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-1">Ganti Kata Sandi</h3>
        <p class="text-xs text-slate-500 dark:text-slate-400 mb-6">Ubah kata sandi akun Anda secara berkala demi keamanan.</p>

        <!-- Alert messages -->
        <div v-if="securitySuccess" class="mb-5 p-3.5 rounded-2xl bg-emerald-50 dark:bg-emerald-900/30 border border-emerald-200 dark:border-emerald-800 text-emerald-700 dark:text-emerald-300 text-xs font-semibold flex items-center space-x-2">
          <span>✓</span>
          <span>{{ securitySuccess }}</span>
        </div>
        <div v-if="securityError" class="mb-5 p-3.5 rounded-2xl bg-red-50 dark:bg-red-900/30 border border-red-200 dark:border-red-800 text-red-700 dark:text-red-300 text-xs font-semibold flex items-center space-x-2">
          <span>✕</span>
          <span>{{ securityError }}</span>
        </div>

        <form @submit.prevent="handleChangePassword" class="space-y-4 max-w-lg">
          <!-- Old Password -->
          <div class="space-y-1.5">
            <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Kata Sandi Saat Ini</label>
            <div class="relative">
              <input
                :type="showOldPass ? 'text' : 'password'"
                v-model="oldPassword"
                required
                placeholder="••••••••"
                class="w-full pl-4 pr-11 py-2.5 rounded-xl border border-slate-300 dark:border-slate-600 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-900 text-slate-900 dark:text-white shadow-xs"
              />
              <button
                type="button"
                @click="showOldPass = !showOldPass"
                class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-slate-400 hover:text-slate-600"
              >
                <svg v-if="!showOldPass" xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-brand-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
                </svg>
              </button>
            </div>
          </div>

          <!-- New Password -->
          <div class="space-y-1.5">
            <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Kata Sandi Baru</label>
            <div class="relative">
              <input
                :type="showNewPass ? 'text' : 'password'"
                v-model="newPassword"
                required
                minlength="6"
                placeholder="Minimal 6 karakter"
                class="w-full pl-4 pr-11 py-2.5 rounded-xl border border-slate-300 dark:border-slate-600 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-900 text-slate-900 dark:text-white shadow-xs"
              />
              <button
                type="button"
                @click="showNewPass = !showNewPass"
                class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-slate-400 hover:text-slate-600"
              >
                <svg v-if="!showNewPass" xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-brand-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
                </svg>
              </button>
            </div>
          </div>

          <!-- Confirm New Password -->
          <div class="space-y-1.5">
            <label class="block text-xs font-bold text-slate-700 dark:text-slate-300 uppercase tracking-wider">Konfirmasi Kata Sandi Baru</label>
            <div class="relative">
              <input
                :type="showNewConfirmPass ? 'text' : 'password'"
                v-model="newConfirmPassword"
                required
                placeholder="Ulangi kata sandi baru"
                class="w-full pl-4 pr-11 py-2.5 rounded-xl border border-slate-300 dark:border-slate-600 focus:outline-none focus:ring-2 focus:ring-brand-500 text-sm bg-white dark:bg-slate-900 text-slate-900 dark:text-white shadow-xs"
              />
              <button
                type="button"
                @click="showNewConfirmPass = !showNewConfirmPass"
                class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-slate-400 hover:text-slate-600"
              >
                <svg v-if="!showNewConfirmPass" xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-brand-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
                </svg>
              </button>
            </div>
          </div>

          <button
            type="submit"
            :disabled="isSecuritySaving"
            class="mt-4 px-6 py-3 rounded-xl bg-slate-900 dark:bg-brand-600 hover:bg-slate-800 dark:hover:bg-brand-700 text-white font-bold text-xs sm:text-sm shadow-md transition-all disabled:opacity-50"
          >
            {{ isSecuritySaving ? 'Memproses...' : 'Ubah Kata Sandi' }}
          </button>
        </form>
      </div>

      <!-- TAB 3: ACTIVITY HISTORY -->
      <div v-else-if="activeTab === 'history'" class="bg-white dark:bg-slate-800 border border-slate-200/90 dark:border-slate-700 rounded-3xl p-6 sm:p-8 shadow-sm space-y-4">
        <div class="flex items-center justify-between pb-3 border-b border-slate-100 dark:border-slate-700">
          <div>
            <h3 class="text-lg font-bold text-slate-900 dark:text-white">Riwayat Konversi Dokumen</h3>
            <p class="text-xs text-slate-500 dark:text-slate-400">File yang baru saja diproses dan siap diunduh ulang.</p>
          </div>
          <button
            v-if="historyItems.length > 0"
            type="button"
            @click="clearHistory"
            class="text-xs font-semibold text-red-600 dark:text-red-400 hover:underline"
          >
            Hapus Semua Riwayat
          </button>
        </div>

        <div v-if="historyItems.length === 0" class="py-10 text-center space-y-2">
          <p class="text-sm font-semibold text-slate-600 dark:text-slate-300">Belum ada riwayat konversi.</p>
          <p class="text-xs text-slate-400">File yang Anda proses akan muncul di sini secara otomatis.</p>
        </div>

        <div v-else class="space-y-2.5">
          <div
            v-for="item in historyItems"
            :key="item.id"
            class="p-4 rounded-2xl bg-slate-50 dark:bg-slate-900/60 border border-slate-200/80 dark:border-slate-700 flex items-center justify-between gap-4"
          >
            <div class="min-w-0 flex items-center space-x-3">
              <span class="text-lg">
                {{ item.tool === 'compress' ? '⚡' : (item.tool === 'merge' ? '📑' : '✂️') }}
              </span>
              <div class="min-w-0">
                <p class="font-bold text-slate-800 dark:text-slate-100 text-sm truncate max-w-xs sm:max-w-md">
                  {{ item.filename }}
                </p>
                <p class="text-xs text-slate-500 dark:text-slate-400">
                  {{ item.tool.toUpperCase() }} • {{ formatFileSize(item.resultSize || item.originalSize) }}
                  <span v-if="item.savedPercentage" class="text-emerald-600 font-bold"> (-{{ item.savedPercentage }}%)</span>
                </p>
              </div>
            </div>

            <div class="flex items-center space-x-2 shrink-0">
              <a
                :href="item.downloadUrl"
                target="_blank"
                download
                class="px-3.5 py-1.5 rounded-xl bg-brand-50 hover:bg-brand-100 dark:bg-brand-900/30 text-brand-600 dark:text-brand-300 font-bold text-xs"
              >
                Unduh
              </a>
              <button
                type="button"
                @click="removeHistoryItem(item.id)"
                class="p-1.5 text-slate-400 hover:text-red-500"
                title="Hapus"
              >
                ✕
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- TAB 4: USAGE & SUBSCRIPTION -->
      <div v-else-if="activeTab === 'subscription'" class="space-y-6">
        <!-- Usage Card -->
        <div class="bg-white dark:bg-slate-800 border border-slate-200/90 dark:border-slate-700 rounded-3xl p-6 sm:p-8 shadow-sm">
          <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-1">Status Kuota Harian</h3>
          <p class="text-xs text-slate-500 dark:text-slate-400 mb-6">Kuota pemrosesan dokumen direset setiap tengah malam (00:00).</p>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div class="p-5 rounded-2xl bg-slate-50 dark:bg-slate-900/50 border border-slate-200 dark:border-slate-700 space-y-2">
              <span class="text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider">Sisa Kuota Hari Ini</span>
              <div class="flex items-baseline space-x-2">
                <span class="text-3xl font-black text-slate-900 dark:text-white">{{ isPro ? 'Tak Terbatas' : remainingQuota }}</span>
                <span v-if="!isPro" class="text-xs text-slate-500 dark:text-slate-400 font-semibold">/ {{ maxDailyLimit }}</span>
              </div>
              <p class="text-[11px] text-slate-500 dark:text-slate-400">
                {{ isPro ? 'Paket PRO memiliki kuota tak terbatas tanpa antrean' : 'Batas free tier hingga 1.000 konversi per hari' }}
              </p>
            </div>

            <div class="p-5 rounded-2xl bg-slate-50 dark:bg-slate-900/50 border border-slate-200 dark:border-slate-700 space-y-2">
              <span class="text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider">Status Layanan</span>
              <div class="flex items-center space-x-2">
                <span class="w-3 h-3 rounded-full" :class="isPro ? 'bg-amber-500' : 'bg-emerald-500'"></span>
                <span class="text-lg font-bold text-slate-900 dark:text-white">{{ isPro ? 'HyperPDF PRO' : 'HyperPDF Starter' }}</span>
              </div>
              <p class="text-[11px] text-slate-500 dark:text-slate-400">
                {{ isPro ? 'Dukungan prioritas 24/7 dan file hingga 500 MB' : 'Dukungan standar dan file hingga 50 MB' }}
              </p>
            </div>
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
import { useHistory } from '../composables/useHistory'

type AppPage = 'home' | 'login' | 'register' | 'forgot-password' | 'reset-password' | 'profile'

const emit = defineEmits<{
  (e: 'navigate', page: AppPage): void
  (e: 'openPricing'): void
}>()

const { user, isPro, remainingQuota, maxDailyLimit, updateProfile, changePassword, logout } = useAuth()
const { historyItems, removeHistoryItem, clearHistory } = useHistory()

const activeTab = ref<'profile' | 'security' | 'history' | 'subscription'>('profile')

// Profile form states
const nameInput = ref(user.value?.name || '')
const usernameInput = ref(user.value?.username || '')
const isProfileSaving = ref(false)
const profileSuccess = ref<string | null>(null)
const profileError = ref<string | null>(null)

// Security form states
const oldPassword = ref('')
const newPassword = ref('')
const newConfirmPassword = ref('')
const showOldPass = ref(false)
const showNewPass = ref(false)
const showNewConfirmPass = ref(false)
const isSecuritySaving = ref(false)
const securitySuccess = ref<string | null>(null)
const securityError = ref<string | null>(null)

function formatDate(dateStr?: string) {
  if (!dateStr) return 'Hari ini'
  try {
    const d = new Date(dateStr)
    return d.toLocaleDateString('id-ID', { year: 'numeric', month: 'short', day: 'numeric' })
  } catch {
    return 'Hari ini'
  }
}

function formatFileSize(bytes: number): string {
  if (!bytes || bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

async function handleSaveProfile() {
  profileError.value = null
  profileSuccess.value = null
  isProfileSaving.value = true

  try {
    await updateProfile(nameInput.value, usernameInput.value)
    profileSuccess.value = 'Informasi profil berhasil diperbarui!'
  } catch (err: any) {
    profileError.value = err.message || 'Gagal menyimpan profil.'
  } finally {
    isProfileSaving.value = false
  }
}

async function handleChangePassword() {
  securityError.value = null
  securitySuccess.value = null

  if (newPassword.value !== newConfirmPassword.value) {
    securityError.value = 'Konfirmasi kata sandi baru tidak cocok.'
    return
  }

  isSecuritySaving.value = true

  try {
    await changePassword(oldPassword.value, newPassword.value)
    securitySuccess.value = 'Kata sandi Anda berhasil diperbarui!'
    oldPassword.value = ''
    newPassword.value = ''
    newConfirmPassword.value = ''
  } catch (err: any) {
    securityError.value = err.message || 'Gagal mengganti kata sandi.'
  } finally {
    isSecuritySaving.value = false
  }
}

function handleLogout() {
  logout()
  emit('navigate', 'home')
}
</script>
