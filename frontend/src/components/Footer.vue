<template>
  <footer class="mt-20 border-t border-slate-200 bg-white py-10">
    <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex flex-col md:flex-row items-center justify-between gap-6 text-center md:text-left">
        <!-- Brand & Description -->
        <div class="space-y-1">
          <div class="flex items-center justify-center md:justify-start space-x-2">
            <span class="font-extrabold text-slate-800 text-lg">PDF<span class="text-brand-500">Compressor</span></span>
            <span class="text-xs bg-slate-100 text-slate-600 px-2 py-0.5 rounded font-mono">v1.0.0</span>
          </div>
          <p class="text-xs text-slate-500 max-w-sm">
            High performance, secure PDF optimization tool powered by Go, Fiber, PostgreSQL, and Ghostscript.
          </p>
        </div>

        <!-- Privacy & Security Notice -->
        <div class="flex items-center space-x-2 px-4 py-2 rounded-xl bg-emerald-50 border border-emerald-200 text-emerald-800 text-xs font-medium">
          <svg xmlns="http://www.w3.org/2000/svg" class="w-4 h-4 text-emerald-600 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
          </svg>
          <span>End-to-end security: Temporary files are wiped immediately upon completion</span>
        </div>
      </div>
    </div>
  </footer>
</template>
