<template>
  <div class="min-h-screen flex flex-col justify-between bg-slate-50 text-slate-800">
    <Navbar :active-tool="activeTool" @select-tool="setTool" />

    <main class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-12 w-full flex-grow">
      <!-- Tool Navigation Pills -->
      <div class="flex items-center justify-center mb-8">
        <div class="inline-flex p-1.5 rounded-2xl bg-slate-200/80 border border-slate-300/60 shadow-inner">
          <button
            type="button"
            @click="setTool('compress')"
            class="px-5 py-2.5 rounded-xl text-xs sm:text-sm font-bold transition-all duration-200 flex items-center space-x-2"
            :class="[
              activeTool === 'compress'
                ? 'bg-white text-brand-600 shadow-md scale-[1.02]'
                : 'text-slate-600 hover:text-slate-900 hover:bg-slate-200/50'
            ]"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
            </svg>
            <span>Compress PDF</span>
          </button>

          <button
            type="button"
            @click="setTool('merge')"
            class="px-5 py-2.5 rounded-xl text-xs sm:text-sm font-bold transition-all duration-200 flex items-center space-x-2"
            :class="[
              activeTool === 'merge'
                ? 'bg-white text-brand-600 shadow-md scale-[1.02]'
                : 'text-slate-600 hover:text-slate-900 hover:bg-slate-200/50'
            ]"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
            </svg>
            <span>Merge PDF</span>
          </button>

          <button
            type="button"
            @click="setTool('split')"
            class="px-5 py-2.5 rounded-xl text-xs sm:text-sm font-bold transition-all duration-200 flex items-center space-x-2"
            :class="[
              activeTool === 'split'
                ? 'bg-white text-brand-600 shadow-md scale-[1.02]'
                : 'text-slate-600 hover:text-slate-900 hover:bg-slate-200/50'
            ]"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z" />
            </svg>
            <span>Split PDF</span>
          </button>
        </div>
      </div>

      <!-- Hero Header -->
      <div class="text-center space-y-3 mb-10">
        <h1 class="text-3xl sm:text-5xl font-black text-slate-900 tracking-tight">
          {{ toolTitle }}
        </h1>
        <p class="text-base sm:text-lg text-slate-600 max-w-xl mx-auto font-normal">
          {{ toolSubtitle }}
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
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- TOOL 1: COMPRESS PDF -->
      <div v-if="activeTool === 'compress'">
        <!-- Result View -->
        <ResultCard
          v-if="jobResult"
          :result="jobResult"
          :original-size-formatted="formattedOriginalSize"
          :compressed-size-formatted="formattedCompressedSize"
          :saved-bytes-formatted="formattedSavedBytes"
          @download="downloadResult"
          @reset="reset"
        />

        <!-- Processing View -->
        <ProcessingCard
          v-else-if="isProcessing"
          :step="processingStep"
          :upload-progress="uploadProgress"
        />

        <!-- Upload & Configure View -->
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
      </div>

      <!-- TOOL 2: MERGE PDF -->
      <div v-else-if="activeTool === 'merge'">
        <MergeTool @error="errorMessage = $event" />
      </div>

      <!-- TOOL 3: SPLIT PDF -->
      <div v-else-if="activeTool === 'split'">
        <SplitTool @error="errorMessage = $event" />
      </div>

      <!-- Features Cards -->
      <div class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-white border border-slate-200/80 rounded-2xl p-6 space-y-2.5 shadow-sm">
          <div class="w-10 h-10 rounded-xl bg-red-50 text-brand-500 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-sm">Ultra Fast Processing</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            High-performance Go backend with Ghostscript and QPDF optimization pipelines.
          </p>
        </div>

        <div class="bg-white border border-slate-200/80 rounded-2xl p-6 space-y-2.5 shadow-sm">
          <div class="w-10 h-10 rounded-xl bg-emerald-50 text-emerald-600 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-sm">Privacy & Security First</h4>
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
          <h4 class="font-bold text-slate-900 text-sm">All-in-One PDF Suite</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            Compress, merge, and split PDF documents easily directly from your browser.
          </p>
        </div>
      </div>
    </main>

    <Footer />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Navbar from './components/Navbar.vue'
import DropZone from './components/DropZone.vue'
import LevelSelector from './components/LevelSelector.vue'
import ProcessingCard from './components/ProcessingCard.vue'
import ResultCard from './components/ResultCard.vue'
import MergeTool from './components/MergeTool.vue'
import SplitTool from './components/SplitTool.vue'
import Footer from './components/Footer.vue'
import { usePdfCompressor } from './composables/usePdfCompressor'
import type { PDFTool } from './types'

const activeTool = ref<PDFTool>('compress')

function setTool(tool: PDFTool) {
  activeTool.value = tool
  errorMessage.value = null
}

const toolTitle = computed(() => {
  switch (activeTool.value) {
    case 'merge':
      return 'Merge PDF Files'
    case 'split':
      return 'Split PDF Document'
    default:
      return 'Compress PDF Files'
  }
})

const toolSubtitle = computed(() => {
  switch (activeTool.value) {
    case 'merge':
      return 'Combine multiple PDF files into a single organized document in seconds.'
    case 'split':
      return 'Extract selected page ranges or separate every page into distinct files.'
    default:
      return 'Reduce PDF file size drastically while preserving optimal visual quality.'
  }
})

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
