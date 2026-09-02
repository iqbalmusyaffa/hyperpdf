<template>
  <div class="min-h-screen flex flex-col justify-between bg-slate-50 text-slate-800">
    <Navbar />

    <main class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-10 sm:py-14 w-full flex-grow">
      <!-- Hero Header -->
      <div class="text-center space-y-3 mb-10">
        <h1 class="text-3xl sm:text-5xl font-black text-slate-900 tracking-tight">
          Compress PDF Files
        </h1>
        <p class="text-base sm:text-lg text-slate-600 max-w-xl mx-auto font-normal">
          Reduce PDF file size drastically while preserving optimal visual quality and crisp typography.
        </p>
      </div>

      <!-- Error Alert -->
      <div
        v-if="errorMessage"
        class="mb-8 p-4 rounded-2xl bg-red-50 border border-red-200 text-red-700 flex items-start space-x-3 shadow-sm animate-shake"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 text-red-500 shrink-0 mt-0.5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <div class="flex-1">
          <h5 class="font-bold text-sm">Action Failed</h5>
          <p class="text-xs sm:text-sm mt-0.5">{{ errorMessage }}</p>
        </div>
        <button @click="errorMessage = null" class="text-red-400 hover:text-red-600">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Step 3: Result View -->
      <ResultCard
        v-if="jobResult"
        :result="jobResult"
        :original-size-formatted="formattedOriginalSize"
        :compressed-size-formatted="formattedCompressedSize"
        :saved-bytes-formatted="formattedSavedBytes"
        @download="downloadResult"
        @reset="reset"
      />

      <!-- Step 2: Processing View -->
      <ProcessingCard
        v-else-if="isProcessing"
        :step="processingStep"
        :upload-progress="uploadProgress"
      />

      <!-- Step 1: Upload & Configure View -->
      <div v-else class="space-y-8">
        <DropZone
          :file="selectedFile"
          :formatted-size="formattedOriginalSize"
          :is-processing="isProcessing"
          @file-selected="handleFileSelected"
          @clear="clearFile"
        />

        <LevelSelector
          v-if="selectedFile"
          :options="levelOptions"
          :selected-level="compressionLevel"
          :is-processing="isProcessing"
          @select-level="compressionLevel = $event"
          @compress="startCompression"
        />
      </div>

      <!-- Features Cards -->
      <div v-if="!jobResult && !isProcessing" class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white border border-slate-200/80 rounded-2xl p-6 space-y-2.5 shadow-sm">
          <div class="w-10 h-10 rounded-xl bg-red-50 text-brand-500 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-sm">Fast & Efficient</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            Optimized Go backend executing native Ghostscript rendering pipelines in parallel.
          </p>
        </div>

        <div class="bg-white border border-slate-200/80 rounded-2xl p-6 space-y-2.5 shadow-sm">
          <div class="w-10 h-10 rounded-xl bg-emerald-50 text-emerald-600 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-sm">Secure & Private</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            MIME & Magic byte verification with automatic temporary file wiping.
          </p>
        </div>

        <div class="bg-white border border-slate-200/80 rounded-2xl p-6 space-y-2.5 shadow-sm">
          <div class="w-10 h-10 rounded-xl bg-blue-50 text-blue-600 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-sm">Granular Control</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            Choose between Low, Medium, and High compression based on your file requirements.
          </p>
        </div>
      </div>
    </main>

    <Footer />
  </div>
</template>

<script setup lang="ts">
import Navbar from './components/Navbar.vue'
import DropZone from './components/DropZone.vue'
import LevelSelector from './components/LevelSelector.vue'
import ProcessingCard from './components/ProcessingCard.vue'
import ResultCard from './components/ResultCard.vue'
import Footer from './components/Footer.vue'
import { usePdfCompressor } from './composables/usePdfCompressor'

const {
  selectedFile,
  compressionLevel,
  isProcessing,
  uploadProgress,
  processingStep,
  jobResult,
  errorMessage,
  levelOptions,
  formattedOriginalSize,
  formattedCompressedSize,
  formattedSavedBytes,
  handleFileSelected,
  clearFile,
  startCompression,
  downloadResult,
  reset,
} = usePdfCompressor()
</script>
