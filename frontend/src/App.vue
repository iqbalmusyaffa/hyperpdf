<template>
  <div class="min-h-screen flex flex-col justify-between bg-gradient-to-b from-slate-50 via-white to-slate-50 text-slate-800">
    <Navbar
      :active-tool="activeTool"
      @select-tool="setTool"
      @open-auth="openAuthModal"
      @open-pricing="isPricingModalOpen = true"
    />

    <main class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8 sm:py-12 w-full flex-grow">
      <!-- Hero Header & Badge -->
      <div class="text-center space-y-4 mb-8 sm:mb-10">
        <div class="inline-flex items-center space-x-2 px-3.5 py-1.5 rounded-full bg-brand-50 border border-brand-200/80 text-brand-700 text-xs font-bold shadow-xs">
          <span class="flex h-2 w-2 relative">
            <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-brand-400 opacity-75"></span>
            <span class="relative inline-flex rounded-full h-2 w-2 bg-brand-500"></span>
          </span>
          <span>⚡ Next-Gen HyperPDF Processing Engine</span>
        </div>

        <h1 class="text-3xl sm:text-5xl lg:text-6xl font-black text-slate-900 tracking-tight leading-tight">
          {{ heroTitlePrefix }}
          <span class="text-transparent bg-clip-text bg-gradient-to-r from-brand-600 via-brand-500 to-rose-500">
            {{ heroTitleHighlight }}
          </span>
        </h1>
        <p class="text-sm sm:text-base md:text-lg text-slate-600 max-w-2xl mx-auto font-normal leading-relaxed">
          {{ toolSubtitle }}
        </p>
      </div>

      <!-- Responsive Tool Switcher Tabs -->
      <div class="flex items-center justify-center mb-10">
        <div class="inline-flex p-1.5 rounded-2xl bg-slate-200/70 backdrop-blur-md border border-slate-300/60 shadow-inner max-w-full overflow-x-auto no-scrollbar">
          <button
            type="button"
            @click="setTool('compress')"
            class="px-4 sm:px-6 py-2.5 rounded-xl text-xs sm:text-sm font-bold transition-all duration-200 flex items-center space-x-2 shrink-0 select-none"
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
            class="px-4 sm:px-6 py-2.5 rounded-xl text-xs sm:text-sm font-bold transition-all duration-200 flex items-center space-x-2 shrink-0 select-none"
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
            class="px-4 sm:px-6 py-2.5 rounded-xl text-xs sm:text-sm font-bold transition-all duration-200 flex items-center space-x-2 shrink-0 select-none"
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
            @require-pro="isPricingModalOpen = true"
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

      <!-- Modern Features Grid -->
      <div class="mt-16 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div class="bg-white/80 backdrop-blur-sm border border-slate-200/80 rounded-3xl p-6 space-y-3 shadow-sm hover:shadow-md transition-shadow">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-tr from-brand-600 to-rose-400 text-white flex items-center justify-center shadow-md shadow-brand-500/20">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13 10V3L4 14h7v7l9-11h-7z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-base">Hyper-Speed Engine</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            Compiled Go backend with optimized multi-threaded Ghostscript and QPDF pipelines for instant conversions.
          </p>
        </div>

        <div class="bg-white/80 backdrop-blur-sm border border-slate-200/80 rounded-3xl p-6 space-y-3 shadow-sm hover:shadow-md transition-shadow">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-tr from-emerald-600 to-teal-400 text-white flex items-center justify-center shadow-md shadow-emerald-500/20">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-base">Total Data Privacy</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            Your files are processed in isolated sandboxes and automatically wiped from disk right after download.
          </p>
        </div>

        <div class="bg-white/80 backdrop-blur-sm border border-slate-200/80 rounded-3xl p-6 space-y-3 shadow-sm hover:shadow-md transition-shadow sm:col-span-2 md:col-span-1">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-tr from-blue-600 to-indigo-400 text-white flex items-center justify-center shadow-md shadow-blue-500/20">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-900 text-base">All-in-One PDF Suite</h4>
          <p class="text-xs text-slate-500 leading-relaxed">
            Compress, merge, and split with full device responsiveness across mobile, tablet, and desktop.
          </p>
        </div>
      </div>
    </main>

    <Footer />

    <!-- Auth Modal -->
    <AuthModal
      :is-open="isAuthModalOpen"
      :initial-mode="authModalMode"
      @close="isAuthModalOpen = false"
      @success="isAuthModalOpen = false"
    />

    <!-- Pricing / Paywall Modal -->
    <PricingModal
      :is-open="isPricingModalOpen"
      @close="isPricingModalOpen = false"
      @upgraded="isPricingModalOpen = false"
    />
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
import AuthModal from './components/AuthModal.vue'
import PricingModal from './components/PricingModal.vue'
import { usePdfCompressor } from './composables/usePdfCompressor'
import type { PDFTool } from './types'

const activeTool = ref<PDFTool>('compress')
const isAuthModalOpen = ref(false)
const authModalMode = ref<'login' | 'register'>('login')
const isPricingModalOpen = ref(false)

function setTool(tool: PDFTool) {
  activeTool.value = tool
  errorMessage.value = null
}

function openAuthModal(mode: 'login' | 'register') {
  authModalMode.value = mode
  isAuthModalOpen.value = true
}

const heroTitlePrefix = computed(() => {
  switch (activeTool.value) {
    case 'merge':
      return 'Merge PDF '
    case 'split':
      return 'Split & Extract '
    default:
      return 'Compress PDF '
  }
})

const heroTitleHighlight = computed(() => {
  switch (activeTool.value) {
    case 'merge':
      return 'Documents'
    case 'split':
      return 'Pages'
    default:
      return 'Files Online'
  }
})

const toolSubtitle = computed(() => {
  switch (activeTool.value) {
    case 'merge':
      return 'Combine multiple PDF files in order into a single unified document with lightning speed.'
    case 'split':
      return 'Extract selected page ranges or separate every single page into a convenient ZIP archive.'
    default:
      return 'Reduce PDF file size drastically while preserving optimal visual quality and crisp typography.'
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
