<template>
  <header class="bg-white/85 backdrop-blur-md border-b border-slate-200/80 sticky top-0 z-50 shadow-xs transition-all">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
      <!-- Left: Logo & Tools Menu -->
      <div class="flex items-center space-x-4 sm:space-x-6">
        <!-- Logo -->
        <div
          class="flex items-center space-x-2.5 sm:space-x-3 cursor-pointer group select-none"
          @click="selectAndClose('compress')"
        >
          <div class="w-9 h-9 sm:w-10 sm:h-10 rounded-xl bg-gradient-to-tr from-brand-600 via-brand-500 to-rose-400 text-white flex items-center justify-center shadow-md shadow-brand-500/25 group-hover:scale-105 transition-transform duration-200">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 animate-pulse" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <div class="flex items-baseline space-x-1.5">
            <span class="text-xl sm:text-2xl font-black tracking-tight text-slate-900">
              Hyper<span class="text-transparent bg-clip-text bg-gradient-to-r from-brand-500 to-rose-500">PDF</span>
            </span>
            <span
              class="hidden md:inline-block px-2 py-0.5 text-[10px] font-extrabold uppercase tracking-wider rounded-full border"
              :class="isPro ? 'bg-amber-50 text-amber-700 border-amber-300' : 'bg-brand-50 text-brand-600 border-brand-200'"
            >
              {{ isPro ? '⭐ PRO' : 'FREE' }}
            </span>
          </div>
        </div>

        <!-- Desktop Tools Dropdown -->
        <div class="hidden sm:block relative">
          <button
            type="button"
            @click="isDropdownOpen = !isDropdownOpen"
            class="flex items-center space-x-2 px-3.5 py-1.5 rounded-xl text-sm font-bold transition-all border"
            :class="[
              isDropdownOpen
                ? 'text-brand-600 bg-brand-50/70 border-brand-200 ring-2 ring-brand-500/10'
                : 'text-slate-700 hover:text-brand-600 hover:bg-slate-50 border-slate-200'
            ]"
          >
            <span>All PDF Tools</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-4 h-4 transition-transform duration-200"
              :class="isDropdownOpen ? 'rotate-180 text-brand-500' : 'text-slate-400'"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </button>

          <!-- Desktop Dropdown Menu -->
          <div
            v-if="isDropdownOpen"
            @click="isDropdownOpen = false"
            class="absolute left-0 mt-2 w-72 bg-white/95 backdrop-blur-xl border border-slate-200 rounded-2xl shadow-2xl py-2 z-50 animate-fade-in"
          >
            <div
              @click="$emit('selectTool', 'compress')"
              class="px-4 py-3 hover:bg-slate-50 cursor-pointer flex items-center space-x-3.5 transition-colors"
              :class="activeTool === 'compress' ? 'bg-brand-50/70 border-l-4 border-brand-500' : ''"
            >
              <div class="w-9 h-9 rounded-xl bg-red-100 text-brand-600 flex items-center justify-center shrink-0 shadow-xs">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                </svg>
              </div>
              <div>
                <div class="flex items-center space-x-2">
                  <p class="text-sm font-bold text-slate-800">Compress PDF</p>
                  <span v-if="activeTool === 'compress'" class="text-[10px] font-bold bg-brand-500 text-white px-1.5 py-0.2 rounded-full">ACTIVE</span>
                </div>
                <p class="text-xs text-slate-500">Reduce size up to 80%</p>
              </div>
            </div>

            <div
              @click="$emit('selectTool', 'merge')"
              class="px-4 py-3 hover:bg-slate-50 cursor-pointer flex items-center space-x-3.5 transition-colors"
              :class="activeTool === 'merge' ? 'bg-brand-50/70 border-l-4 border-brand-500' : ''"
            >
              <div class="w-9 h-9 rounded-xl bg-blue-100 text-blue-600 flex items-center justify-center shrink-0 shadow-xs">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
                </svg>
              </div>
              <div>
                <div class="flex items-center space-x-2">
                  <p class="text-sm font-bold text-slate-800">Merge PDF</p>
                  <span v-if="activeTool === 'merge'" class="text-[10px] font-bold bg-brand-500 text-white px-1.5 py-0.2 rounded-full">ACTIVE</span>
                </div>
                <p class="text-xs text-slate-500">Combine multiple documents</p>
              </div>
            </div>

            <div
              @click="$emit('selectTool', 'split')"
              class="px-4 py-3 hover:bg-slate-50 cursor-pointer flex items-center space-x-3.5 transition-colors"
              :class="activeTool === 'split' ? 'bg-brand-50/70 border-l-4 border-brand-500' : ''"
            >
              <div class="w-9 h-9 rounded-xl bg-emerald-100 text-emerald-600 flex items-center justify-center shrink-0 shadow-xs">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z" />
                </svg>
              </div>
              <div>
                <div class="flex items-center space-x-2">
                  <p class="text-sm font-bold text-slate-800">Split PDF</p>
                  <span v-if="activeTool === 'split'" class="text-[10px] font-bold bg-brand-500 text-white px-1.5 py-0.2 rounded-full">ACTIVE</span>
                </div>
                <p class="text-xs text-slate-500">Extract ranges or single pages</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right: Pricing/Upgrade, Auth Actions & Mobile Toggle -->
      <div class="flex items-center space-x-2 sm:space-x-3">
        <!-- Upgrade to PRO Button (Sparkling) -->
        <button
          v-if="!isPro"
          type="button"
          @click="$emit('openPricing')"
          class="inline-flex items-center space-x-1.5 px-3 sm:px-4 py-1.5 rounded-xl bg-gradient-to-r from-amber-500 to-rose-500 hover:from-amber-600 hover:to-rose-600 text-white text-xs font-black shadow-md shadow-amber-500/20 transition-all transform active:scale-95 animate-pulse"
        >
          <span>✨</span>
          <span class="hidden xs:inline">Upgrade to PRO</span>
          <span class="xs:hidden">PRO</span>
        </button>

        <!-- Guest Actions: Log In & Sign Up -->
        <div v-if="!isLoggedIn" class="hidden sm:flex items-center space-x-2">
          <button
            type="button"
            @click="$emit('openAuth', 'login')"
            class="px-3.5 py-1.5 rounded-xl text-xs font-bold text-slate-700 hover:text-slate-900 hover:bg-slate-100 transition-colors"
          >
            Log In
          </button>
          <button
            type="button"
            @click="$emit('openAuth', 'register')"
            class="px-4 py-1.5 rounded-xl text-xs font-bold text-white bg-slate-900 hover:bg-slate-800 shadow-sm transition-all transform active:scale-95"
          >
            Sign Up
          </button>
        </div>

        <!-- Logged In User Avatar & Profile Dropdown -->
        <div v-else class="relative">
          <button
            type="button"
            @click="isUserMenuOpen = !isUserMenuOpen"
            class="flex items-center space-x-2 p-1.5 rounded-xl hover:bg-slate-100 transition-colors border border-slate-200"
          >
            <div class="w-7 h-7 rounded-lg bg-gradient-to-tr from-brand-500 to-rose-500 text-white font-bold text-xs flex items-center justify-center shadow-xs">
              {{ user?.name?.charAt(0).toUpperCase() || 'U' }}
            </div>
            <span class="hidden md:inline text-xs font-bold text-slate-800 truncate max-w-[100px]">{{ user?.name }}</span>
            <span class="text-[10px] px-1.5 py-0.2 rounded-md font-bold" :class="isPro ? 'bg-amber-100 text-amber-800' : 'bg-slate-100 text-slate-600'">
              {{ isPro ? 'PRO' : 'FREE' }}
            </span>
          </button>

          <!-- User Menu Dropdown -->
          <div
            v-if="isUserMenuOpen"
            @click="isUserMenuOpen = false"
            class="absolute right-0 mt-2 w-56 bg-white border border-slate-200 rounded-2xl shadow-xl py-2 z-50 animate-fade-in text-xs"
          >
            <div class="px-4 py-2.5 border-b border-slate-100">
              <p class="font-bold text-slate-800 truncate">{{ user?.name }}</p>
              <p class="text-slate-500 text-[11px] truncate">{{ user?.email }}</p>
            </div>

            <button
              v-if="!isPro"
              type="button"
              @click="$emit('openPricing')"
              class="w-full text-left px-4 py-2.5 font-bold text-amber-600 hover:bg-amber-50 flex items-center space-x-2 transition-colors"
            >
              <span>⭐</span>
              <span>Upgrade Subscription</span>
            </button>

            <button
              type="button"
              @click="handleLogout"
              class="w-full text-left px-4 py-2.5 font-semibold text-red-600 hover:bg-red-50 flex items-center space-x-2 transition-colors"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
              </svg>
              <span>Log Out</span>
            </button>
          </div>
        </div>

        <!-- Mobile Hamburger Toggle -->
        <button
          type="button"
          @click="isMobileMenuOpen = !isMobileMenuOpen"
          class="sm:hidden p-2 rounded-xl text-slate-600 hover:text-slate-900 hover:bg-slate-100 transition-colors"
          aria-label="Toggle Menu"
        >
          <svg v-if="!isMobileMenuOpen" xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Mobile Drawer Menu -->
    <div
      v-if="isMobileMenuOpen"
      class="sm:hidden bg-white/98 backdrop-blur-xl border-b border-slate-200 px-4 pt-2 pb-6 space-y-4 animate-fade-in"
    >
      <!-- Tools Links -->
      <div class="space-y-1.5">
        <p class="text-[11px] font-extrabold uppercase tracking-wider text-slate-400 px-2">PDF Tools</p>
        
        <button
          type="button"
          @click="selectAndClose('compress')"
          class="w-full text-left px-4 py-2.5 rounded-xl flex items-center space-x-3 transition-colors text-xs"
          :class="activeTool === 'compress' ? 'bg-brand-50 text-brand-600 font-bold border border-brand-200' : 'text-slate-700 hover:bg-slate-50 font-medium'"
        >
          <span class="w-6 h-6 rounded-lg bg-red-100 text-brand-600 flex items-center justify-center font-bold">⚡</span>
          <span>Compress PDF</span>
        </button>

        <button
          type="button"
          @click="selectAndClose('merge')"
          class="w-full text-left px-4 py-2.5 rounded-xl flex items-center space-x-3 transition-colors text-xs"
          :class="activeTool === 'merge' ? 'bg-brand-50 text-brand-600 font-bold border border-brand-200' : 'text-slate-700 hover:bg-slate-50 font-medium'"
        >
          <span class="w-6 h-6 rounded-lg bg-blue-100 text-blue-600 flex items-center justify-center font-bold">📑</span>
          <span>Merge PDF</span>
        </button>

        <button
          type="button"
          @click="selectAndClose('split')"
          class="w-full text-left px-4 py-2.5 rounded-xl flex items-center space-x-3 transition-colors text-xs"
          :class="activeTool === 'split' ? 'bg-brand-50 text-brand-600 font-bold border border-brand-200' : 'text-slate-700 hover:bg-slate-50 font-medium'"
        >
          <span class="w-6 h-6 rounded-lg bg-emerald-100 text-emerald-600 flex items-center justify-center font-bold">✂️</span>
          <span>Split PDF</span>
        </button>
      </div>

      <!-- Auth Actions in Mobile -->
      <div class="pt-3 border-t border-slate-100 space-y-2">
        <p class="text-[11px] font-extrabold uppercase tracking-wider text-slate-400 px-2">Account</p>

        <div v-if="!isLoggedIn" class="grid grid-cols-2 gap-2">
          <button
            type="button"
            @click="openAuthAndClose('login')"
            class="py-2.5 text-center text-xs font-bold rounded-xl border border-slate-200 text-slate-700 hover:bg-slate-50"
          >
            Log In
          </button>
          <button
            type="button"
            @click="openAuthAndClose('register')"
            class="py-2.5 text-center text-xs font-bold rounded-xl bg-slate-900 text-white hover:bg-slate-800"
          >
            Sign Up
          </button>
        </div>

        <div v-else class="space-y-2">
          <div class="flex items-center justify-between px-2 py-1 text-xs">
            <span class="font-bold text-slate-800">{{ user?.name }}</span>
            <span class="text-[10px] font-extrabold px-2 py-0.5 rounded-full" :class="isPro ? 'bg-amber-100 text-amber-800' : 'bg-slate-100 text-slate-600'">
              {{ isPro ? '⭐ PRO' : 'FREE' }}
            </span>
          </div>

          <button
            v-if="!isPro"
            type="button"
            @click="openPricingAndClose"
            class="w-full py-2.5 text-center text-xs font-extrabold rounded-xl bg-gradient-to-r from-amber-500 to-rose-500 text-white shadow-sm"
          >
            ✨ Upgrade to PRO (Rp 49.000)
          </button>

          <button
            type="button"
            @click="handleLogout"
            class="w-full py-2 text-center text-xs font-semibold rounded-xl text-red-600 hover:bg-red-50 border border-red-100"
          >
            Log Out
          </button>
        </div>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuth } from '../composables/useAuth'
import type { PDFTool } from '../types'

defineProps<{
  activeTool: PDFTool
}>()

const emit = defineEmits<{
  (e: 'selectTool', tool: PDFTool): void
  (e: 'openAuth', mode: 'login' | 'register'): void
  (e: 'openPricing'): void
}>()

const { user, isLoggedIn, isPro, logout } = useAuth()

const isDropdownOpen = ref(false)
const isUserMenuOpen = ref(false)
const isMobileMenuOpen = ref(false)

function selectAndClose(tool: PDFTool) {
  emit('selectTool', tool)
  isDropdownOpen.value = false
  isMobileMenuOpen.value = false
}

function openAuthAndClose(mode: 'login' | 'register') {
  emit('openAuth', mode)
  isMobileMenuOpen.value = false
}

function openPricingAndClose() {
  emit('openPricing')
  isMobileMenuOpen.value = false
}

function handleLogout() {
  logout()
  isUserMenuOpen.value = false
  isMobileMenuOpen.value = false
}
</script>
