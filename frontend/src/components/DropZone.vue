<template>
  <div class="w-full">
    <!-- Drop Zone Box -->
    <div
      v-if="!file"
      @dragover.prevent="isDragging = true"
      @dragleave.prevent="isDragging = false"
      @drop.prevent="onDrop"
      class="relative border-2 border-dashed rounded-3xl p-8 sm:p-12 text-center transition-all duration-200 cursor-pointer group"
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

      <!-- Icon & Prompt -->
      <div class="flex flex-col items-center justify-center space-y-4">
        <div
          class="w-20 h-20 rounded-2xl flex items-center justify-center transition-transform group-hover:scale-110 duration-200 shadow-inner"
          :class="isDragging ? 'bg-brand-500 text-white' : 'bg-red-50 text-brand-500'"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="1.75">
            <path stroke-linecap="round" stroke-linejoin="round" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
          </svg>
        </div>

        <div class="space-y-1.5">
          <p class="text-xl font-bold text-slate-800">
            <span class="text-brand-600 group-hover:underline">Choose PDF file</span> or drag & drop here
          </p>
          <p class="text-sm text-slate-500">
            Up to 50 MB per file • High-fidelity Ghostscript engine
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

    <!-- File Selected View -->
    <div
      v-else
      class="bg-white border border-slate-200 rounded-3xl p-6 sm:p-8 shadow-sm flex flex-col sm:flex-row items-center justify-between gap-6"
    >
      <div class="flex items-center space-x-5 w-full sm:w-auto">
        <div class="w-16 h-16 rounded-2xl bg-red-100/70 border border-red-200 text-brand-600 flex items-center justify-center shrink-0">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
          </svg>
        </div>
        <div class="truncate">
          <h4 class="text-base sm:text-lg font-bold text-slate-800 truncate max-w-xs sm:max-w-md">
            {{ file.name }}
          </h4>
          <p class="text-xs sm:text-sm text-slate-500 font-medium">
            Original Size: <span class="font-semibold text-slate-700">{{ formattedSize }}</span>
          </p>
        </div>
      </div>

      <div class="flex items-center space-x-3 w-full sm:w-auto justify-end">
        <button
          v-if="!isProcessing"
          type="button"
          @click="$emit('clear')"
          class="px-4 py-2 rounded-xl text-sm font-semibold text-slate-600 hover:text-red-600 hover:bg-red-50 border border-slate-200 transition-colors flex items-center space-x-1.5"
        >
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          <span>Change File</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

defineProps<{
  file: File | null
  formattedSize: string
  isProcessing: boolean
}>()

const emit = defineEmits<{
  (e: 'fileSelected', file: File): void
  (e: 'clear'): void
}>()

const isDragging = ref(false)
const fileInput = ref<HTMLInputElement | null>(null)

function triggerFileInput() {
  fileInput.value?.click()
}

function onFileInputChange(event: Event) {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    emit('fileSelected', target.files[0])
  }
}

function onDrop(event: DragEvent) {
  isDragging.value = false
  if (event.dataTransfer?.files && event.dataTransfer.files.length > 0) {
    emit('fileSelected', event.dataTransfer.files[0])
  }
}
</script>
