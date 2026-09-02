<template>
  <div
    v-if="isOpen"
    class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/60 backdrop-blur-sm animate-fade-in"
    @click.self="close"
  >
    <div class="relative w-full max-w-3xl bg-white border border-slate-200 rounded-3xl shadow-2xl overflow-hidden animate-scale-up max-h-[90vh] overflow-y-auto">
      <!-- Close Button -->
      <button
        type="button"
        @click="close"
        class="absolute top-4 right-4 z-10 text-slate-400 hover:text-slate-700 p-2 rounded-full hover:bg-slate-100 transition-colors"
      >
        <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>

      <!-- Header -->
      <div class="pt-8 pb-4 px-6 text-center space-y-2">
        <div class="inline-flex items-center space-x-1.5 px-3 py-1 rounded-full bg-amber-50 border border-amber-200 text-amber-700 text-xs font-extrabold shadow-xs">
          <span>✨</span>
          <span>HYPERPDF PRO SUBSCRIPTION</span>
        </div>
        <h3 class="text-2xl sm:text-4xl font-black text-slate-900 tracking-tight">
          Supercharge Your PDF Workflow
        </h3>
        <p class="text-xs sm:text-sm text-slate-500 max-w-md mx-auto">
          Get ultra-high quality conversions, unlimited merges, 500 MB file sizes, and dedicated high-speed server queues.
        </p>
      </div>

      <!-- Success State (After upgrade) -->
      <div v-if="upgradeSuccess" class="p-8 sm:p-12 text-center space-y-6 animate-fade-in">
        <div class="w-20 h-20 mx-auto rounded-full bg-emerald-100 text-emerald-600 flex items-center justify-center shadow-lg shadow-emerald-500/10 animate-bounce">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-10 h-10" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2.5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M5 13l4 4L19 7" />
          </svg>
        </div>

        <div class="space-y-1">
          <h4 class="text-2xl sm:text-3xl font-extrabold text-slate-900">Welcome to HyperPDF PRO! ⭐</h4>
          <p class="text-sm text-slate-500">Your subscription is active. All premium features are unlocked instantly.</p>
        </div>

        <button
          type="button"
          @click="close"
          class="px-8 py-3.5 rounded-2xl bg-gradient-to-r from-brand-500 to-rose-500 text-white font-bold text-sm shadow-lg shadow-brand-500/30 transition-all transform active:scale-95"
        >
          Start Using PRO Features
        </button>
      </div>

      <!-- Plans Comparison Grid -->
      <div v-else class="p-6 sm:p-8 grid grid-cols-1 md:grid-cols-2 gap-6 items-stretch">
        <!-- Free Plan -->
        <div class="border border-slate-200 rounded-3xl p-6 flex flex-col justify-between bg-slate-50/50">
          <div class="space-y-4">
            <div class="flex justify-between items-center">
              <h4 class="font-extrabold text-slate-800 text-lg">Starter Free</h4>
              <span v-if="!isPro" class="text-[10px] font-bold bg-slate-200 text-slate-700 px-2.5 py-0.5 rounded-full uppercase">Current Plan</span>
            </div>

            <div class="flex items-baseline space-x-1">
              <span class="text-3xl font-black text-slate-900">Rp 0</span>
              <span class="text-xs text-slate-500 font-medium">/ forever</span>
            </div>

            <ul class="space-y-2.5 text-xs text-slate-600 pt-2 border-t border-slate-200/60">
              <li class="flex items-center space-x-2">
                <span class="text-emerald-500 font-bold">✓</span>
                <span>Max 50 MB file size</span>
              </li>
              <li class="flex items-center space-x-2">
                <span class="text-emerald-500 font-bold">✓</span>
                <span>Standard compression quality</span>
              </li>
              <li class="flex items-center space-x-2">
                <span class="text-emerald-500 font-bold">✓</span>
                <span>Basic Merge & Split (up to 5 files)</span>
              </li>
              <li class="flex items-center space-x-2 text-slate-400">
                <span>✕</span>
                <span class="line-through">High DPI 300 Best Quality</span>
              </li>
              <li class="flex items-center space-x-2 text-slate-400">
                <span>✕</span>
                <span class="line-through">Priority Queue</span>
              </li>
            </ul>
          </div>

          <button
            type="button"
            disabled
            class="w-full mt-6 py-3 rounded-2xl bg-slate-200 text-slate-500 font-bold text-xs cursor-default"
          >
            {{ isPro ? 'Included in PRO' : 'Active Plan' }}
          </button>
        </div>

        <!-- PRO Plan -->
        <div class="relative border-2 border-brand-500 rounded-3xl p-6 flex flex-col justify-between bg-gradient-to-b from-brand-50/50 via-white to-brand-50/20 shadow-xl shadow-brand-500/10">
          <div class="absolute -top-3 right-6 bg-gradient-to-r from-amber-500 to-rose-500 text-white text-[10px] font-extrabold uppercase px-3 py-0.5 rounded-full shadow-md">
            POPULAR CHOICE
          </div>

          <div class="space-y-4">
            <div class="flex justify-between items-center">
              <div class="flex items-center space-x-2">
                <h4 class="font-extrabold text-slate-900 text-lg">HyperPDF PRO</h4>
                <span class="text-amber-500">⭐</span>
              </div>
              <span v-if="isPro" class="text-[10px] font-bold bg-emerald-500 text-white px-2.5 py-0.5 rounded-full uppercase">Active Plan</span>
            </div>

            <div class="flex items-baseline space-x-1">
              <span class="text-3xl sm:text-4xl font-black text-slate-900">Rp 49.000</span>
              <span class="text-xs text-slate-500 font-medium">/ bulan</span>
            </div>

            <ul class="space-y-2.5 text-xs text-slate-700 pt-2 border-t border-brand-200/60 font-medium">
              <li class="flex items-center space-x-2">
                <span class="text-brand-600 font-bold">✓</span>
                <span><strong>500 MB</strong> Max Upload File Size</span>
              </li>
              <li class="flex items-center space-x-2">
                <span class="text-brand-600 font-bold">✓</span>
                <span><strong>Ultra High-Fidelity (~300 DPI)</strong> Best Quality</span>
              </li>
              <li class="flex items-center space-x-2">
                <span class="text-brand-600 font-bold">✓</span>
                <span><strong>Unlimited Batch</strong> Merges & Splits</span>
              </li>
              <li class="flex items-center space-x-2">
                <span class="text-brand-600 font-bold">✓</span>
                <span><strong>Zero Limits</strong> & Dedicated Server Queue</span>
              </li>
              <li class="flex items-center space-x-2">
                <span class="text-brand-600 font-bold">✓</span>
                <span>Priority 24/7 Support</span>
              </li>
            </ul>
          </div>

          <button
            type="button"
            @click="handleSubscribe"
            :disabled="isUpgrading || isPro"
            class="w-full mt-6 py-3.5 rounded-2xl bg-gradient-to-r from-brand-500 via-brand-600 to-rose-500 hover:from-brand-600 hover:to-rose-600 disabled:opacity-60 text-white font-extrabold text-sm shadow-lg shadow-brand-500/30 transition-all transform active:scale-98 flex items-center justify-center space-x-2"
          >
            <svg v-if="isUpgrading" class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8z"></path>
            </svg>
            <span>{{ isPro ? 'You are on PRO' : (isUpgrading ? 'Processing Subscription...' : 'Subscribe to HyperPDF PRO') }}</span>
          </button>
        </div>
      </div>

      <!-- Trust Footer -->
      <div class="px-6 pb-6 text-center text-slate-400 text-[11px] flex items-center justify-center space-x-4">
        <span>🔒 Secure 256-bit SSL Checkout</span>
        <span>•</span>
        <span>⚡ Cancel Anytime</span>
        <span>•</span>
        <span>💳 Instant Activation</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useAuth } from '../composables/useAuth'

const props = defineProps<{
  isOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'upgraded'): void
}>()

const { isPro, upgradeToPro } = useAuth()

const isUpgrading = ref(false)
const upgradeSuccess = ref(false)

watch(
  () => props.isOpen,
  (open) => {
    if (open) {
      upgradeSuccess.value = false
    }
  }
)

async function handleSubscribe() {
  isUpgrading.value = true
  try {
    await upgradeToPro()
    upgradeSuccess.value = true
    emit('upgraded')
  } catch (err) {
    console.error(err)
  } finally {
    isUpgrading.value = false
  }
}

function close() {
  emit('close')
}
</script>
