<template>
  <div class="bg-white border border-slate-200 rounded-3xl p-8 sm:p-12 shadow-sm text-center max-w-xl mx-auto space-y-8">
    <!-- Animated Processing Icon -->
    <div class="relative w-24 h-24 mx-auto flex items-center justify-center">
      <div class="absolute inset-0 rounded-full border-4 border-slate-100"></div>
      <div
        class="absolute inset-0 rounded-full border-4 border-brand-500 border-t-transparent animate-spin"
      ></div>
      <div class="w-16 h-16 rounded-full bg-brand-50 text-brand-500 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 animate-pulse" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
        </svg>
      </div>
    </div>

    <!-- Status Messages -->
    <div class="space-y-2">
      <h3 class="text-2xl font-extrabold text-slate-900">
        {{ currentStepTitle }}
      </h3>
      <p class="text-sm text-slate-500 max-w-md mx-auto">
        {{ currentStepDescription }}
      </p>
    </div>

    <!-- Progress Bar -->
    <div class="space-y-2">
      <div class="w-full bg-slate-100 rounded-full h-3 overflow-hidden p-0.5 border border-slate-200">
        <div
          class="bg-gradient-to-r from-brand-500 to-rose-400 h-2 rounded-full transition-all duration-300 ease-out"
          :style="{ width: `${progressPercentage}%` }"
        ></div>
      </div>
      <div class="flex justify-between text-xs font-semibold text-slate-500">
        <span>Processing</span>
        <span>{{ progressPercentage }}%</span>
      </div>
    </div>

    <!-- Active Step List -->
    <div class="grid grid-cols-3 gap-2 pt-4 border-t border-slate-100 text-xs">
      <div
        class="flex flex-col items-center space-y-1"
        :class="step === 'uploading' || step === 'compressing' || step === 'finalizing' ? 'text-brand-600 font-bold' : 'text-slate-400'"
      >
        <div class="w-2 h-2 rounded-full" :class="step === 'uploading' ? 'bg-brand-500 animate-ping' : 'bg-brand-500'"></div>
        <span>1. Upload</span>
      </div>
      <div
        class="flex flex-col items-center space-y-1"
        :class="step === 'compressing' || step === 'finalizing' ? 'text-brand-600 font-bold' : 'text-slate-400'"
      >
        <div class="w-2 h-2 rounded-full" :class="step === 'compressing' ? 'bg-brand-500 animate-ping' : step === 'finalizing' ? 'bg-brand-500' : 'bg-slate-300'"></div>
        <span>2. Compress</span>
      </div>
      <div
        class="flex flex-col items-center space-y-1"
        :class="step === 'finalizing' ? 'text-brand-600 font-bold' : 'text-slate-400'"
      >
        <div class="w-2 h-2 rounded-full" :class="step === 'finalizing' ? 'bg-brand-500 animate-ping' : 'bg-slate-300'"></div>
        <span>3. Finalize</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  step: 'idle' | 'uploading' | 'compressing' | 'finalizing' | 'completed' | 'error'
  uploadProgress: number
}>()

const progressPercentage = computed(() => {
  if (props.step === 'uploading') return Math.min(props.uploadProgress, 60)
  if (props.step === 'compressing') return 85
  if (props.step === 'finalizing') return 95
  if (props.step === 'completed') return 100
  return 10
})

const currentStepTitle = computed(() => {
  switch (props.step) {
    case 'uploading':
      return 'Uploading Document...'
    case 'compressing':
      return 'Compressing PDF...'
    case 'finalizing':
      return 'Calculating Savings...'
    default:
      return 'Preparing Optimization...'
  }
})

const currentStepDescription = computed(() => {
  switch (props.step) {
    case 'uploading':
      return 'Securely sending your file to the processing engine.'
    case 'compressing':
      return 'Applying Ghostscript compression algorithm and optimizing raster streams.'
    case 'finalizing':
      return 'Generating summary metrics and preparing your download link.'
    default:
      return 'Hold tight, this usually takes only a few seconds.'
  }
})
</script>
