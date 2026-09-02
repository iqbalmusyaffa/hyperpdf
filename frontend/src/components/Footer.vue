<template>
  <footer class="mt-20 border-t border-slate-200/80 bg-white/70 backdrop-blur-md py-10">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex flex-col md:flex-row items-center justify-between gap-6 text-center md:text-left">
        <!-- Brand & Description -->
        <div class="space-y-1.5">
          <div class="flex items-center justify-center md:justify-start space-x-2">
            <span class="font-black text-slate-900 text-xl tracking-tight">
              Hyper<span class="text-transparent bg-clip-text bg-gradient-to-r from-brand-500 to-rose-500">PDF</span>
            </span>
            <span class="text-[10px] bg-slate-100 text-slate-700 px-2 py-0.5 rounded-full font-mono font-bold border border-slate-200">
              v1.0.0
            </span>
          </div>
          <p class="text-xs text-slate-500 max-w-sm leading-relaxed">
            Blazing-fast, private, and secure PDF tools orchestrated with Go, Fiber, PostgreSQL, and Ghostscript.
          </p>
        </div>

        <!-- Privacy & Security Badge -->
        <div class="flex items-center space-x-2.5 px-4 py-2.5 rounded-2xl bg-emerald-50/80 border border-emerald-200/80 text-emerald-800 text-xs font-semibold shadow-xs">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-emerald-600 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
          </svg>
          <span>End-to-End Privacy: Temporary files are wiped immediately upon job completion</span>
        </div>
      </div>
    </div>
  </footer>
</template>
