<template>
  <header class="bg-white border-b border-slate-200 sticky top-0 z-50 shadow-sm">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between">
      <!-- Logo -->
      <div class="flex items-center space-x-3">
        <div class="w-10 h-10 rounded-xl bg-brand-500 text-white flex items-center justify-center shadow-md shadow-brand-500/20">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M14 3v4a2 2 0 002 2h4" />
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 13h6m-6 4h4" />
          </svg>
        </div>
        <div>
          <span class="text-xl font-extrabold tracking-tight text-slate-900">
            PDF<span class="text-brand-500">Compressor</span>
          </span>
          <span class="hidden sm:inline-block ml-2 px-2 py-0.5 text-xs font-semibold text-brand-600 bg-brand-50 rounded-full border border-brand-200">
            Go + Ghostscript
          </span>
        </div>
      </div>

      <!-- Navigation & Health status -->
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
