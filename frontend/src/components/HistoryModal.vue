<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm animate-fade-in"
    @click.self="close"
  >
    <div class="relative w-full max-w-2xl bg-white dark:bg-slate-900 border border-slate-200 dark:border-slate-800 rounded-3xl shadow-2xl overflow-hidden animate-scale-up max-h-[85vh] flex flex-col">
      <!-- Modal Header -->
      <div class="p-6 border-b border-slate-200 dark:border-slate-800 flex items-center justify-between bg-slate-50/50 dark:bg-slate-800/50">
        <div class="flex items-center space-x-3">
          <div class="w-10 h-10 rounded-xl bg-brand-100 dark:bg-brand-900/40 text-brand-600 dark:text-brand-400 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <div>
            <h3 class="text-lg font-extrabold text-slate-900 dark:text-white">Riwayat Konversi File</h3>
            <p class="text-xs text-slate-500 dark:text-slate-400">File yang baru saja Anda proses (tersimpan di browser ini)</p>
          </div>
        </div>

        <div class="flex items-center space-x-2">
          <button
            v-if="historyItems.length > 0"
            type="button"
            @click="clearHistory"
            class="text-xs font-semibold text-red-600 dark:text-red-400 hover:underline px-2 py-1"
          >
            Hapus Semua
          </button>
          <button
            type="button"
            @click="close"
            class="p-2 rounded-xl text-slate-400 hover:text-slate-700 dark:hover:text-slate-200 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>

      <!-- Modal Body (Scrollable List) -->
      <div class="p-6 overflow-y-auto flex-grow space-y-3">
        <!-- Empty State -->
        <div v-if="historyItems.length === 0" class="py-12 text-center space-y-3">
          <div class="w-16 h-16 mx-auto rounded-2xl bg-slate-100 dark:bg-slate-800 text-slate-400 flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
            </svg>
          </div>
          <h4 class="font-bold text-slate-800 dark:text-slate-200 text-base">Belum Ada Riwayat</h4>
          <p class="text-xs text-slate-500 dark:text-slate-400 max-w-xs mx-auto">
            File yang Anda kompres, gabung, atau pisah akan otomatis tercatat di sini untuk memudahkan unduh ulang.
          </p>
        </div>

        <!-- History Item Cards -->
        <div
          v-for="item in historyItems"
          :key="item.id"
          class="p-4 rounded-2xl border border-slate-200/90 dark:border-slate-800 bg-white dark:bg-slate-800/70 hover:shadow-md transition-all flex items-center justify-between gap-4"
        >
          <div class="flex items-center space-x-3.5 min-w-0">
            <!-- Tool Badge Icon -->
            <div
              class="w-10 h-10 rounded-xl flex items-center justify-center shrink-0 font-bold text-sm"
              :class="getToolBadgeClass(item.tool)"
            >
              <span v-if="item.tool === 'compress'">⚡</span>
              <span v-else-if="item.tool === 'merge'">📑</span>
              <span v-else>✂️</span>
            </div>

            <!-- Item Details -->
            <div class="min-w-0">
              <div class="flex items-center space-x-2">
                <p class="font-bold text-slate-900 dark:text-white text-sm truncate max-w-[220px] sm:max-w-xs">
                  {{ item.filename }}
                </p>
                <span
                  v-if="item.savedPercentage"
                  class="text-[10px] font-extrabold px-1.5 py-0.5 rounded-full bg-emerald-100 dark:bg-emerald-900/40 text-emerald-700 dark:text-emerald-300 shrink-0"
                >
                  -{{ item.savedPercentage }}%
                </span>
              </div>
              <div class="flex items-center space-x-2 text-xs text-slate-500 dark:text-slate-400 mt-0.5">
                <span class="capitalize font-semibold text-brand-600 dark:text-brand-400">{{ item.tool }}</span>
                <span>•</span>
                <span v-if="item.resultSize">{{ formatSize(item.resultSize) }}</span>
                <span v-else>{{ formatSize(item.originalSize) }}</span>
                <span>•</span>
                <span>{{ formatTime(item.timestamp) }}</span>
              </div>
            </div>
          </div>

          <!-- Actions -->
          <div class="flex items-center space-x-2 shrink-0">
            <a
              :href="item.downloadUrl"
              target="_blank"
              download
              class="px-3.5 py-2 rounded-xl bg-brand-50 hover:bg-brand-100 dark:bg-brand-900/30 dark:hover:bg-brand-900/50 text-brand-600 dark:text-brand-300 font-bold text-xs flex items-center space-x-1.5 transition-colors"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
              </svg>
              <span>Unduh</span>
            </a>

            <button
              type="button"
              @click="removeHistoryItem(item.id)"
              class="p-2 text-slate-400 hover:text-red-500 rounded-lg transition-colors"
              title="Hapus dari riwayat"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useHistory } from '../composables/useHistory'

defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
}>()

const { historyItems, removeHistoryItem, clearHistory } = useHistory()

function getToolBadgeClass(tool: string) {
  switch (tool) {
    case 'merge':
      return 'bg-blue-100 dark:bg-blue-900/40 text-blue-600 dark:text-blue-300'
    case 'split':
      return 'bg-emerald-100 dark:bg-emerald-900/40 text-emerald-600 dark:text-emerald-300'
    default:
      return 'bg-red-100 dark:bg-red-900/40 text-brand-600 dark:text-brand-300'
  }
}

function formatSize(bytes: number): string {
  if (!bytes || bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

function formatTime(isoStr: string): string {
  try {
    const date = new Date(isoStr)
    const now = new Date()
    const diffMs = now.getTime() - date.getTime()
    const diffMins = Math.floor(diffMs / (1000 * 60))
    if (diffMins < 1) return 'Baru saja'
    if (diffMins < 60) return `${diffMins} mnt lalu`
    const diffHours = Math.floor(diffMins / 60)
    if (diffHours < 24) return `${diffHours} jam lalu`
    return date.toLocaleDateString('id-ID', { month: 'short', day: 'numeric' })
  } catch {
    return 'Hari ini'
  }
}

function close() {
  emit('close')
}
</script>
