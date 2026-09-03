<template>
  <div class="bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-3xl p-8 sm:p-12 shadow-sm text-center max-w-2xl mx-auto space-y-8 animate-fade-in">
    <!-- Success Badge / Icon -->
    <div class="w-20 h-20 mx-auto rounded-full bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-300 flex items-center justify-center shadow-lg shadow-emerald-500/10">
      <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
        <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
      </svg>
    </div>

    <!-- Title -->
    <div class="space-y-1">
      <h3 class="text-2xl sm:text-3xl font-extrabold text-slate-900 dark:text-white">
        PDF Anda Telah Siap!
      </h3>
      <p class="text-sm sm:text-base text-slate-500 dark:text-slate-400">
        File dikompresi menggunakan optimasi <span class="font-bold text-slate-700 dark:text-slate-200">{{ result.compression_level }}</span>.
      </p>
    </div>

    <!-- Stats Comparison Card -->
    <div class="bg-slate-50 dark:bg-slate-900/50 border border-slate-200 dark:border-slate-700 rounded-2xl p-6 space-y-5">
      <div class="grid grid-cols-3 gap-4 text-center">
        <!-- Original Size -->
        <div class="space-y-1">
          <p class="text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider">Ukuran Asli</p>
          <p class="text-lg sm:text-xl font-bold text-slate-700 dark:text-slate-200">{{ originalSizeFormatted }}</p>
        </div>

        <!-- Compressed Size -->
        <div class="space-y-1">
          <p class="text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider">Hasil Kompres</p>
          <p class="text-lg sm:text-xl font-extrabold text-emerald-600 dark:text-emerald-400">{{ compressedSizeFormatted }}</p>
        </div>

        <!-- Saved Percentage -->
        <div class="space-y-1">
          <p class="text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase tracking-wider">Hemat</p>
          <p class="text-lg sm:text-xl font-extrabold text-brand-600 dark:text-brand-400">-{{ result.compression_percentage }}%</p>
        </div>
      </div>

      <!-- Reduction bar -->
      <div class="space-y-1.5 pt-2 border-t border-slate-200/60 dark:border-slate-700">
        <div class="w-full bg-slate-200 dark:bg-slate-700 rounded-full h-3 overflow-hidden p-0.5">
          <div
            class="bg-emerald-500 h-2 rounded-full transition-all duration-1000 ease-out"
            :style="{ width: `${Math.max(result.compression_percentage, 5)}%` }"
          ></div>
        </div>
        <p class="text-xs text-slate-500 dark:text-slate-400 font-medium">
          Anda berhasil menghemat <span class="font-bold text-slate-800 dark:text-slate-100">{{ savedBytesFormatted }}</span> ruang penyimpanan.
        </p>
      </div>
    </div>

    <!-- Action Buttons -->
    <div class="flex flex-col sm:flex-row items-center justify-center gap-4 pt-2">
      <button
        type="button"
        @click="$emit('download')"
        class="w-full sm:w-auto px-8 py-4 rounded-2xl bg-brand-500 hover:bg-brand-600 text-white font-bold text-base shadow-lg shadow-brand-500/30 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
        </svg>
        <span>Unduh PDF Kompresi</span>
      </button>

      <button
        type="button"
        @click="$emit('reset')"
        class="w-full sm:w-auto px-6 py-4 rounded-2xl bg-white dark:bg-slate-800 hover:bg-slate-50 dark:hover:bg-slate-700 text-slate-700 dark:text-slate-200 font-semibold text-base border border-slate-200 dark:border-slate-700 shadow-sm transition-all flex items-center justify-center space-x-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
        </svg>
        <span>Kompres File Lain</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { JobResponse } from '../types'

defineProps<{
  result: JobResponse
  originalSizeFormatted: string
  compressedSizeFormatted: string
  savedBytesFormatted: string
}>()

defineEmits<{
  (e: 'download'): void
  (e: 'reset'): void
}>()
</script>
