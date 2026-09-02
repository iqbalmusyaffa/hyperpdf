<template>
  <header class="bg-white border-b border-slate-200 sticky top-0 z-50 shadow-sm">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
      <!-- Left: Logo & Brand -->
      <div class="flex items-center space-x-6">
        <div class="flex items-center space-x-3 cursor-pointer" @click="$emit('selectTool', 'compress')">
          <div class="w-10 h-10 rounded-xl bg-brand-500 text-white flex items-center justify-center shadow-md shadow-brand-500/20">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M14 3v4a2 2 0 002 2h4" />
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 13h6m-6 4h4" />
            </svg>
          </div>
          <div>
            <span class="text-xl font-extrabold tracking-tight text-slate-900">
              PDF<span class="text-brand-500">Tools</span>
            </span>
          </div>
        </div>

        <!-- Desktop Tools Dropdown / Quick Links -->
        <div class="relative">
          <button
            type="button"
            @click="isDropdownOpen = !isDropdownOpen"
            class="flex items-center space-x-2 px-3.5 py-2 rounded-xl text-sm font-bold text-slate-700 hover:text-brand-600 hover:bg-slate-50 transition-colors border border-slate-200"
          >
            <span>All PDF Tools</span>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="w-4 h-4 transition-transform duration-200"
              :class="isDropdownOpen ? 'rotate-180' : ''"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
            </svg>
          </button>

          <!-- Dropdown Menu -->
          <div
            v-if="isDropdownOpen"
            @click="isDropdownOpen = false"
            class="absolute left-0 mt-2 w-64 bg-white border border-slate-200 rounded-2xl shadow-xl py-2 z-50 animate-fade-in"
          >
            <div
              @click="$emit('selectTool', 'compress')"
              class="px-4 py-3 hover:bg-slate-50 cursor-pointer flex items-center space-x-3 transition-colors"
              :class="activeTool === 'compress' ? 'bg-brand-50/50' : ''"
            >
              <div class="w-8 h-8 rounded-lg bg-red-100 text-brand-600 flex items-center justify-center shrink-0">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
                </svg>
              </div>
              <div>
                <p class="text-sm font-bold text-slate-800">Compress PDF</p>
                <p class="text-xs text-slate-500">Reduce file size dramatically</p>
              </div>
            </div>

            <div
              @click="$emit('selectTool', 'merge')"
              class="px-4 py-3 hover:bg-slate-50 cursor-pointer flex items-center space-x-3 transition-colors"
              :class="activeTool === 'merge' ? 'bg-brand-50/50' : ''"
            >
              <div class="w-8 h-8 rounded-lg bg-blue-100 text-blue-600 flex items-center justify-center shrink-0">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
                </svg>
              </div>
              <div>
                <p class="text-sm font-bold text-slate-800">Merge PDF</p>
                <p class="text-xs text-slate-500">Combine multiple PDFs in order</p>
              </div>
            </div>

            <div
              @click="$emit('selectTool', 'split')"
              class="px-4 py-3 hover:bg-slate-50 cursor-pointer flex items-center space-x-3 transition-colors"
              :class="activeTool === 'split' ? 'bg-brand-50/50' : ''"
            >
              <div class="w-8 h-8 rounded-lg bg-emerald-100 text-emerald-600 flex items-center justify-center shrink-0">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z" />
                </svg>
              </div>
              <div>
                <p class="text-sm font-bold text-slate-800">Split PDF</p>
                <p class="text-xs text-slate-500">Extract ranges or single pages</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Right: Navigation & Health status -->
      <div class="flex items-center space-x-4">
        <div class="flex items-center space-x-2 text-xs font-medium px-3 py-1.5 rounded-full bg-slate-100 text-slate-700">
          <span class="w-2 h-2 rounded-full" :class="isHealthy ? 'bg-emerald-500 animate-pulse' : 'bg-amber-500'"></span>
          <span>{{ healthText }}</span>
        </div>

        <a
          href="/swagger/index.html"
          target="_blank"
          class="text-xs font-semibold text-slate-600 hover:text-brand-600 transition-colors flex items-center space-x-1"
        >
          <span>API Docs</span>
          <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
          </svg>
        </a>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { pdfApi } from '../api/pdfApi'
import type { PDFTool } from '../types'

defineProps<{
  activeTool: PDFTool
}>()

defineEmits<{
  (e: 'selectTool', tool: PDFTool): void
}>()

const isDropdownOpen = ref(false)
const isHealthy = ref(true)
const healthText = ref('Connecting...')

onMounted(async () => {
  try {
    const health = await pdfApi.checkHealth()
    if (health.success) {
      isHealthy.value = true
      healthText.value = 'Engine Ready'
    } else {
      isHealthy.value = false
      healthText.value = 'Engine Standby'
    }
  } catch {
    isHealthy.value = false
    healthText.value = 'Server Offline'
  }
})
</script>
