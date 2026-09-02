<template>
  <div class="bg-white border border-slate-200 rounded-3xl p-6 sm:p-8 shadow-sm space-y-6">
    <div>
      <h3 class="text-lg font-bold text-slate-900">Choose Compression Level</h3>
      <p class="text-sm text-slate-500">Select the balance between file size reduction and document visual quality.</p>
    </div>

    <!-- 3 Level Options Grid -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <div
        v-for="opt in options"
        :key="opt.id"
        @click="$emit('selectLevel', opt.id)"
        class="relative border-2 rounded-2xl p-5 cursor-pointer transition-all duration-200 flex flex-col justify-between"
        :class="[
          selectedLevel === opt.id
            ? 'border-brand-500 bg-brand-50/40 shadow-sm ring-2 ring-brand-500/20'
            : 'border-slate-200 hover:border-slate-300 hover:bg-slate-50/60'
        ]"
      >
        <!-- Top Badge -->
        <div class="flex items-center justify-between mb-3">
          <span
            class="px-2.5 py-1 text-xs font-bold rounded-full"
            :class="[
              opt.recommended
                ? 'bg-brand-500 text-white'
                : 'bg-slate-100 text-slate-700'
            ]"
          >
            {{ opt.badge }}
          </span>

          <div
            class="w-5 h-5 rounded-full border-2 flex items-center justify-center transition-colors"
            :class="[
              selectedLevel === opt.id
                ? 'border-brand-500 bg-brand-500 text-white'
                : 'border-slate-300 bg-white'
            ]"
          >
            <div v-if="selectedLevel === opt.id" class="w-2 h-2 bg-white rounded-full"></div>
          </div>
        </div>

        <!-- Name & Description -->
        <div class="space-y-1 mb-4">
          <h4 class="font-bold text-slate-800 text-base">{{ opt.name }}</h4>
          <p class="text-xs text-slate-500 leading-relaxed">{{ opt.description }}</p>
        </div>

        <!-- Specs breakdown -->
        <div class="border-t border-slate-200/80 pt-3 space-y-1 text-xs">
          <div class="flex justify-between text-slate-600">
            <span>Resolution:</span>
            <span class="font-medium text-slate-800">{{ opt.quality }}</span>
          </div>
          <div class="flex justify-between text-slate-600">
            <span>Reduction:</span>
            <span class="font-bold text-brand-600">{{ opt.compression }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Action Button -->
    <div class="pt-2 flex justify-center">
      <button
        type="button"
        @click="$emit('compress')"
        :disabled="isProcessing"
        class="w-full sm:w-auto min-w-[240px] px-8 py-4 rounded-2xl bg-brand-500 hover:bg-brand-600 disabled:bg-slate-300 text-white font-bold text-base shadow-lg shadow-brand-500/30 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
          <path stroke-linecap="round" stroke-linejoin="round" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10" />
        </svg>
        <span>Compress PDF Now</span>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { CompressionLevel, LevelOption } from '../types'

defineProps<{
  options: LevelOption[]
  selectedLevel: CompressionLevel
  isProcessing: boolean
}>()

defineEmits<{
  (e: 'selectLevel', level: CompressionLevel): void
  (e: 'compress'): void
}>()
</script>
