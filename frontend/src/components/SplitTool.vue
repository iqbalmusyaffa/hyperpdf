<template>
  <div class="space-y-8">
    <!-- Result View -->
    <div
      v-if="splitResult"
      class="bg-white border border-slate-200 rounded-3xl p-8 sm:p-12 shadow-sm text-center max-w-2xl mx-auto space-y-8 animate-fade-in"
    >
      <div class="w-20 h-20 mx-auto rounded-full bg-emerald-100 text-emerald-600 flex items-center justify-center shadow-lg shadow-emerald-500/10">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
        </svg>
      </div>

      <div class="space-y-1">
        <h3 class="text-2xl sm:text-3xl font-extrabold text-slate-900">
          PDF Split Successfully!
        </h3>
        <p class="text-sm sm:text-base text-slate-500">
          Generated <span class="font-bold text-slate-700">{{ splitResult.generated_count }} file(s)</span>
          <span v-if="splitResult.is_zip_archive"> packaged into a clean ZIP archive.</span>
        </p>
      </div>

      <div class="bg-slate-50 border border-slate-200 rounded-2xl p-6 space-y-2 text-left">
        <div class="flex justify-between items-center text-sm">
          <span class="text-slate-500">Output File:</span>
          <span class="font-bold text-slate-800">{{ splitResult.download_filename }}</span>
        </div>
        <div class="flex justify-between items-center text-sm">
          <span class="text-slate-500">Split Mode:</span>
          <span class="font-semibold text-brand-600 uppercase">{{ splitResult.split_mode }}</span>
        </div>
        <div v-if="splitResult.page_ranges" class="flex justify-between items-center text-sm">
          <span class="text-slate-500">Extracted Range:</span>
          <span class="font-mono text-xs bg-slate-200 px-2 py-0.5 rounded">{{ splitResult.page_ranges }}</span>
        </div>
      </div>

      <div class="flex flex-col sm:flex-row items-center justify-center gap-4 pt-2">
        <button
          type="button"
          @click="downloadResult"
          class="w-full sm:w-auto px-8 py-4 rounded-2xl bg-brand-500 hover:bg-brand-600 text-white font-bold text-base shadow-lg shadow-brand-500/30 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
          </svg>
          <span>Download {{ splitResult.is_zip_archive ? 'ZIP Archive' : 'Extracted PDF' }}</span>
        </button>

        <button
          type="button"
          @click="reset"
          class="w-full sm:w-auto px-6 py-4 rounded-2xl bg-white hover:bg-slate-50 text-slate-700 font-semibold text-base border border-slate-200 shadow-sm transition-all flex items-center justify-center space-x-2"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <span>Split Another PDF</span>
        </button>
      </div>
    </div>

    <!-- Processing State -->
    <div
      v-else-if="isProcessing"
      class="bg-white border border-slate-200 rounded-3xl p-8 sm:p-12 shadow-sm text-center max-w-xl mx-auto space-y-8"
    >
      <div class="relative w-24 h-24 mx-auto flex items-center justify-center">
        <div class="absolute inset-0 rounded-full border-4 border-slate-100"></div>
        <div class="absolute inset-0 rounded-full border-4 border-brand-500 border-t-transparent animate-spin"></div>
        <div class="w-16 h-16 rounded-full bg-brand-50 text-brand-500 flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 animate-pulse" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z" />
          </svg>
        </div>
      </div>

      <div class="space-y-2">
        <h3 class="text-2xl font-extrabold text-slate-900">Splitting PDF Document...</h3>
        <p class="text-sm text-slate-500">Extracting requested pages and assembling output files.</p>
      </div>

      <div class="w-full bg-slate-100 rounded-full h-3 overflow-hidden p-0.5 border border-slate-200">
        <div
          class="bg-gradient-to-r from-brand-500 to-rose-400 h-2 rounded-full transition-all duration-300 ease-out"
          :style="{ width: `${uploadProgress || 90}%` }"
        ></div>
      </div>
    </div>

    <!-- Upload & Configuration View -->
    <div v-else class="space-y-8">
      <!-- File Selector Zone -->
      <div
        v-if="!selectedFile"
        @dragover.prevent="isDragging = true"
        @dragleave.prevent="isDragging = false"
        @drop.prevent="onDrop"
        class="border-2 border-dashed rounded-3xl p-8 sm:p-12 text-center transition-all duration-200 cursor-pointer group"
        :class="[
          isDragging
            ? 'border-brand-500 bg-brand-50/70 scale-[1.01]'
            : 'border-slate-300 hover:border-brand-400 bg-white hover:bg-slate-50/50 shadow-sm'
        ]"
        @click="triggerFileInput"
      >
        <input
          ref="fileInput"
          type="file"
          accept="application/pdf,.pdf"
          class="hidden"
          @change="onFileInputChange"
        />

        <div class="flex flex-col items-center justify-center space-y-4">
          <div
            class="w-20 h-20 rounded-2xl flex items-center justify-center transition-transform group-hover:scale-110 duration-200 shadow-inner"
            :class="isDragging ? 'bg-brand-500 text-white' : 'bg-red-50 text-brand-500'"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.75">
              <path stroke-linecap="round" stroke-linejoin="round" d="M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z" />
            </svg>
          </div>

          <div class="space-y-1.5">
            <p class="text-xl font-bold text-slate-800">
              <span class="text-brand-600 group-hover:underline">Select PDF file to split</span> or drag & drop here
            </p>
            <p class="text-sm text-slate-500">
              Extract page ranges or separate all pages into individual files.
            </p>
          </div>

          <button
            type="button"
            class="mt-2 inline-flex items-center px-6 py-3 rounded-xl bg-brand-500 hover:bg-brand-600 text-white font-semibold shadow-md shadow-brand-500/25 transition-all active:scale-95"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
            </svg>
            Select PDF File
          </button>
        </div>
      </div>

      <!-- File Loaded & Split Configuration Card -->
      <div v-else class="space-y-6">
        <!-- Selected File Preview -->
        <div class="bg-white border border-slate-200 rounded-3xl p-6 sm:p-8 shadow-sm flex items-center justify-between gap-4">
          <div class="flex items-center space-x-4 truncate">
            <div class="w-12 h-12 rounded-xl bg-red-100 text-brand-600 flex items-center justify-center shrink-0">
              <svg xmlns="http://www.w3.org/2000/svg" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <div class="truncate">
              <h4 class="font-bold text-slate-800 truncate text-base">{{ selectedFile.name }}</h4>
              <p class="text-xs text-slate-500">{{ formatFileSize(selectedFile.size) }}</p>
            </div>
          </div>

          <button
            type="button"
            @click="clearFile"
            class="px-3 py-1.5 text-xs font-semibold text-red-600 hover:bg-red-50 rounded-lg border border-red-200 transition-colors"
          >
            Change File
          </button>
        </div>

        <!-- Mode Selector -->
        <div class="bg-white border border-slate-200 rounded-3xl p-6 sm:p-8 shadow-sm space-y-6">
          <h4 class="font-bold text-slate-900 text-lg">Split Options</h4>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <!-- Mode 1: Range -->
            <div
              @click="splitMode = 'range'"
              class="border-2 rounded-2xl p-5 cursor-pointer transition-all duration-200 flex flex-col justify-between"
              :class="[
                splitMode === 'range'
                  ? 'border-brand-500 bg-brand-50/40 ring-2 ring-brand-500/20'
                  : 'border-slate-200 hover:border-slate-300'
              ]"
            >
              <div class="flex items-center justify-between mb-2">
                <span class="font-bold text-slate-800 text-sm">Extract Specific Pages</span>
                <div
                  class="w-5 h-5 rounded-full border-2 flex items-center justify-center"
                  :class="splitMode === 'range' ? 'border-brand-500 bg-brand-500' : 'border-slate-300 bg-white'"
                >
                  <div v-if="splitMode === 'range'" class="w-2 h-2 bg-white rounded-full"></div>
                </div>
              </div>
              <p class="text-xs text-slate-500">Extract a continuous page range (e.g. 1-3) into a new PDF document.</p>
            </div>

            <!-- Mode 2: All Pages to ZIP -->
            <div
              @click="splitMode = 'all'"
              class="border-2 rounded-2xl p-5 cursor-pointer transition-all duration-200 flex flex-col justify-between"
              :class="[
                splitMode === 'all'
                  ? 'border-brand-500 bg-brand-50/40 ring-2 ring-brand-500/20'
                  : 'border-slate-200 hover:border-slate-300'
              ]"
            >
              <div class="flex items-center justify-between mb-2">
                <span class="font-bold text-slate-800 text-sm">Split Every Page</span>
                <div
                  class="w-5 h-5 rounded-full border-2 flex items-center justify-center"
                  :class="splitMode === 'all' ? 'border-brand-500 bg-brand-500' : 'border-slate-300 bg-white'"
                >
                  <div v-if="splitMode === 'all'" class="w-2 h-2 bg-white rounded-full"></div>
                </div>
              </div>
              <p class="text-xs text-slate-500">Split every page into separate PDF files and bundle them into a ZIP archive.</p>
            </div>
          </div>

          <!-- Page Range Input (Visible when mode === 'range') -->
          <div v-if="splitMode === 'range'" class="p-4 rounded-2xl bg-slate-50 border border-slate-200 space-y-2">
            <label class="block text-xs font-bold text-slate-700 uppercase tracking-wider">Page Range</label>
            <input
              type="text"
              v-model="pageRanges"
              placeholder="e.g. 1-3 or 2"
              class="w-full px-4 py-2.5 rounded-xl border border-slate-300 focus:outline-none focus:ring-2 focus:ring-brand-500 focus:border-brand-500 text-sm bg-white font-mono"
            />
            <p class="text-xs text-slate-500">
              Specify page range format: <span class="font-semibold text-slate-700">1-5</span> (pages 1 to 5) or <span class="font-semibold text-slate-700">2</span> (page 2 only).
            </p>
          </div>

          <!-- Action Button -->
          <div class="pt-2 flex justify-center">
            <button
              type="button"
              @click="startSplit"
              :disabled="isProcessing"
              class="w-full sm:w-auto min-w-[240px] px-8 py-4 rounded-2xl bg-brand-500 hover:bg-brand-600 disabled:bg-slate-300 text-white font-bold text-base shadow-lg shadow-brand-500/30 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M14.121 14.121L19 19m-7-7l7-7m-7 7l-2.879 2.879M12 12L9.121 9.121m0 5.758a3 3 0 10-4.243 4.243 3 3 0 004.243-4.243zm0-5.758a3 3 0 10-4.243-4.243 3 3 0 004.243 4.243z" />
              </svg>
              <span>Split PDF Document</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { pdfApi } from '../api/pdfApi'
import type { SplitResponse } from '../types'

const emit = defineEmits<{
  (e: 'error', message: string): void
}>()

const selectedFile = ref<File | null>(null)
const splitMode = ref<'range' | 'all'>('range')
const pageRanges = ref('1-3')
const isDragging = ref(false)
const isProcessing = ref(false)
const uploadProgress = ref(0)
const splitResult = ref<SplitResponse | null>(null)
const fileInput = ref<HTMLInputElement | null>(null)

function triggerFileInput() {
  fileInput.value?.click()
}

function onFileInputChange(event: Event) {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    setFile(target.files[0])
  }
}

function onDrop(event: DragEvent) {
  isDragging.value = false
  if (event.dataTransfer?.files && event.dataTransfer.files.length > 0) {
    setFile(event.dataTransfer.files[0])
  }
}

function setFile(file: File) {
  if (!file.name.toLowerCase().endsWith('.pdf') && file.type !== 'application/pdf') {
    emit('error', 'Please select a valid PDF file (.pdf)')
    return
  }
  selectedFile.value = file
  splitResult.value = null
}

function clearFile() {
  selectedFile.value = null
  splitResult.value = null
  isProcessing.value = false
}

async function startSplit() {
  if (!selectedFile.value) {
    emit('error', 'Please select a PDF file first.')
    return
  }

  if (splitMode.value === 'range' && !pageRanges.value.trim()) {
    emit('error', 'Please specify a page range (e.g. 1-3).')
    return
  }

  isProcessing.value = true
  uploadProgress.value = 0

  try {
    const result = await pdfApi.splitPDF(
      selectedFile.value,
      splitMode.value,
      splitMode.value === 'range' ? pageRanges.value.trim() : undefined,
      (progress) => {
        uploadProgress.value = progress
      }
    )
    splitResult.value = result
  } catch (err: any) {
    emit('error', err.response?.data?.message || err.message || 'Failed to split PDF file')
  } finally {
    isProcessing.value = false
  }
}

function downloadResult() {
  if (!splitResult.value) return
  const link = document.createElement('a')
  link.href = splitResult.value.download_url
  link.setAttribute('download', splitResult.value.download_filename)
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

function reset() {
  clearFile()
}

function formatFileSize(bytes: number): string {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
