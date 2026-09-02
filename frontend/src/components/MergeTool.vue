<template>
  <div class="space-y-8">
    <!-- Result View -->
    <div
      v-if="mergeResult"
      class="bg-white border border-slate-200 rounded-3xl p-8 sm:p-12 shadow-sm text-center max-w-2xl mx-auto space-y-8 animate-fade-in"
    >
      <div class="w-20 h-20 mx-auto rounded-full bg-emerald-100 text-emerald-600 flex items-center justify-center shadow-lg shadow-emerald-500/10">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
          <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
        </svg>
      </div>

      <div class="space-y-1">
        <h3 class="text-2xl sm:text-3xl font-extrabold text-slate-900">
          PDFs Merged Successfully!
        </h3>
        <p class="text-sm sm:text-base text-slate-500">
          Combined <span class="font-bold text-slate-700">{{ mergeResult.file_count }} PDF files</span> into a single document.
        </p>
      </div>

      <div class="bg-slate-50 border border-slate-200 rounded-2xl p-6 space-y-2">
        <div class="flex justify-between items-center text-sm">
          <span class="text-slate-500">Output Filename:</span>
          <span class="font-bold text-slate-800">{{ mergeResult.merged_filename }}</span>
        </div>
        <div class="flex justify-between items-center text-sm">
          <span class="text-slate-500">Total Size:</span>
          <span class="font-bold text-emerald-600">{{ formatFileSize(mergeResult.total_size) }}</span>
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
          <span>Download Merged PDF</span>
        </button>

        <button
          type="button"
          @click="reset"
          class="w-full sm:w-auto px-6 py-4 rounded-2xl bg-white hover:bg-slate-50 text-slate-700 font-semibold text-base border border-slate-200 shadow-sm transition-all flex items-center justify-center space-x-2"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 text-slate-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
          <span>Merge More Files</span>
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
            <path stroke-linecap="round" stroke-linejoin="round" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
          </svg>
        </div>
      </div>

      <div class="space-y-2">
        <h3 class="text-2xl font-extrabold text-slate-900">Merging PDF Documents...</h3>
        <p class="text-sm text-slate-500">Combining {{ selectedFiles.length }} files in sequential order.</p>
      </div>

      <div class="w-full bg-slate-100 rounded-full h-3 overflow-hidden p-0.5 border border-slate-200">
        <div
          class="bg-gradient-to-r from-brand-500 to-rose-400 h-2 rounded-full transition-all duration-300 ease-out"
          :style="{ width: `${uploadProgress || 90}%` }"
        ></div>
      </div>
    </div>

    <!-- Upload & File Management View -->
    <div v-else class="space-y-8">
      <!-- Drag & Drop Zone -->
      <div
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
          multiple
          class="hidden"
          @change="onFileInputChange"
        />

        <div class="flex flex-col items-center justify-center space-y-4">
          <div
            class="w-20 h-20 rounded-2xl flex items-center justify-center transition-transform group-hover:scale-110 duration-200 shadow-inner"
            :class="isDragging ? 'bg-brand-500 text-white' : 'bg-red-50 text-brand-500'"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.75">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
            </svg>
          </div>

          <div class="space-y-1.5">
            <p class="text-xl font-bold text-slate-800">
              <span class="text-brand-600 group-hover:underline">Select multiple PDF files</span> or drag & drop here
            </p>
            <p class="text-sm text-slate-500">
              Combine 2 or more PDFs into a single polished document.
            </p>
          </div>

          <button
            type="button"
            class="mt-2 inline-flex items-center px-6 py-3 rounded-xl bg-brand-500 hover:bg-brand-600 text-white font-semibold shadow-md shadow-brand-500/25 transition-all active:scale-95"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
            </svg>
            Select PDF Files
          </button>
        </div>
      </div>

      <!-- Selected Files Order List -->
      <div v-if="selectedFiles.length > 0" class="bg-white border border-slate-200 rounded-3xl p-6 sm:p-8 shadow-sm space-y-6">
        <div class="flex items-center justify-between border-b border-slate-100 pb-4">
          <div>
            <h4 class="font-bold text-slate-900 text-lg">Files to Merge ({{ selectedFiles.length }})</h4>
            <p class="text-xs text-slate-500">Reorder files using the up/down arrows to set merge sequence.</p>
          </div>

          <button
            type="button"
            @click="triggerFileInput"
            class="px-4 py-2 rounded-xl text-xs font-bold text-brand-600 bg-brand-50 hover:bg-brand-100 transition-colors flex items-center space-x-1.5"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4" />
            </svg>
            <span>Add More</span>
          </button>
        </div>

        <div class="space-y-3">
          <div
            v-for="(file, index) in selectedFiles"
            :key="index"
            class="flex items-center justify-between p-4 rounded-2xl bg-slate-50 border border-slate-200/80 hover:bg-slate-100/60 transition-colors"
          >
            <div class="flex items-center space-x-4 min-w-0">
              <span class="w-7 h-7 rounded-lg bg-brand-500 text-white text-xs font-extrabold flex items-center justify-center shrink-0">
                {{ index + 1 }}
              </span>
              <div class="truncate">
                <p class="font-bold text-slate-800 text-sm truncate max-w-xs sm:max-w-md">{{ file.name }}</p>
                <p class="text-xs text-slate-500">{{ formatFileSize(file.size) }}</p>
              </div>
            </div>

            <div class="flex items-center space-x-1.5 shrink-0">
              <button
                type="button"
                @click="moveUp(index)"
                :disabled="index === 0"
                class="p-2 rounded-lg text-slate-500 hover:text-slate-800 hover:bg-white disabled:opacity-30 transition-colors"
                title="Move Up"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M5 15l7-7 7 7" />
                </svg>
              </button>

              <button
                type="button"
                @click="moveDown(index)"
                :disabled="index === selectedFiles.length - 1"
                class="p-2 rounded-lg text-slate-500 hover:text-slate-800 hover:bg-white disabled:opacity-30 transition-colors"
                title="Move Down"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
                </svg>
              </button>

              <button
                type="button"
                @click="removeFile(index)"
                class="p-2 rounded-lg text-red-500 hover:bg-red-50 transition-colors"
                title="Remove"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
              </button>
            </div>
          </div>
        </div>

        <div class="pt-4 flex justify-center">
          <button
            type="button"
            @click="startMerge"
            :disabled="selectedFiles.length < 2 || isProcessing"
            class="w-full sm:w-auto min-w-[240px] px-8 py-4 rounded-2xl bg-brand-500 hover:bg-brand-600 disabled:bg-slate-300 text-white font-bold text-base shadow-lg shadow-brand-500/30 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M8 7v8a2 2 0 002 2h6M8 7V5a2 2 0 012-2h4.586a1 1 0 01.707.293l4.414 4.414a1 1 0 01.293.707V15a2 2 0 01-2 2h-2M8 7H6a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2v-2" />
            </svg>
            <span>Merge {{ selectedFiles.length }} PDFs</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { pdfApi } from '../api/pdfApi'
import type { MergeResponse } from '../types'

const emit = defineEmits<{
  (e: 'error', message: string): void
}>()

const selectedFiles = ref<File[]>([])
const isDragging = ref(false)
const isProcessing = ref(false)
const uploadProgress = ref(0)
const mergeResult = ref<MergeResponse | null>(null)
const fileInput = ref<HTMLInputElement | null>(null)

function triggerFileInput() {
  fileInput.value?.click()
}

function onFileInputChange(event: Event) {
  const target = event.target as HTMLInputElement
  if (target.files) {
    addFiles(Array.from(target.files))
  }
}

function onDrop(event: DragEvent) {
  isDragging.value = false
  if (event.dataTransfer?.files) {
    addFiles(Array.from(event.dataTransfer.files))
  }
}

function addFiles(files: File[]) {
  const pdfs = files.filter((f) => f.name.toLowerCase().endsWith('.pdf') || f.type === 'application/pdf')
  if (pdfs.length < files.length) {
    emit('error', 'Some non-PDF files were skipped. Only .pdf files are supported.')
  }
  selectedFiles.value.push(...pdfs)
}

function removeFile(index: number) {
  selectedFiles.value.splice(index, 1)
}

function moveUp(index: number) {
  if (index > 0) {
    const temp = selectedFiles.value[index]
    selectedFiles.value[index] = selectedFiles.value[index - 1]
    selectedFiles.value[index - 1] = temp
  }
}

function moveDown(index: number) {
  if (index < selectedFiles.value.length - 1) {
    const temp = selectedFiles.value[index]
    selectedFiles.value[index] = selectedFiles.value[index + 1]
    selectedFiles.value[index + 1] = temp
  }
}

async function startMerge() {
  if (selectedFiles.value.length < 2) {
    emit('error', 'Please select at least 2 PDF files to merge.')
    return
  }

  isProcessing.value = true
  uploadProgress.value = 0

  try {
    const result = await pdfApi.mergePDF(selectedFiles.value, (progress) => {
      uploadProgress.value = progress
    })
    mergeResult.value = result
  } catch (err: any) {
    emit('error', err.response?.data?.message || err.message || 'Failed to merge PDF files')
  } finally {
    isProcessing.value = false
  }
}

function downloadResult() {
  if (!mergeResult.value) return
  const link = document.createElement('a')
  link.href = mergeResult.value.download_url
  link.setAttribute('download', mergeResult.value.merged_filename)
  document.body.appendChild(link)
  link.click()
  document.body.removeChild(link)
}

function reset() {
  selectedFiles.value = []
  mergeResult.value = null
  isProcessing.value = false
}

function formatFileSize(bytes: number): string {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}
</script>
