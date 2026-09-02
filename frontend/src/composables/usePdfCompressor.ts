import { ref, computed } from 'vue'
import { pdfApi } from '../api/pdfApi'
import { useAuth } from './useAuth'
import type { CompressionLevel, JobResponse, LevelOption } from '../types'

export const levelOptions: LevelOption[] = [
  {
    id: 'ULTRA_EXTREME',
    name: 'Ultra Extreme',
    badge: 'Smallest Size',
    description: 'Maximum aggressive shrink for tight upload limits (CPNS/job portals).',
    quality: 'Basic (~50 DPI)',
    compression: 'Massive (~80-95%)',
    isProOnly: false,
  },
  {
    id: 'HIGH',
    name: 'Extreme Compression',
    badge: 'Web & Email',
    description: 'High compression ratio for email attachments and web sharing.',
    quality: 'Screen (~72 DPI)',
    compression: 'High (~70-85%)',
    isProOnly: false,
  },
  {
    id: 'MEDIUM',
    name: 'Balanced Compression',
    badge: 'Recommended',
    description: 'Smart compression balancing crisp text and images with great size savings.',
    quality: 'eBook (~150 DPI)',
    compression: 'Balanced (~50-70%)',
    recommended: true,
    isProOnly: false,
  },
  {
    id: 'LOW',
    name: 'High-Fidelity',
    badge: 'Print Ready',
    description: 'Crisp typography and vector graphics ideal for client presentations.',
    quality: 'Print (~220 DPI)',
    compression: 'Light (~25-45%)',
    isProOnly: true,
  },
  {
    id: 'STUDIO_MASTER',
    name: 'Studio Master Lossless',
    badge: 'Lossless Archive',
    description: 'Preserves 300+ DPI resolution for publishing, blueprints, and legal archives.',
    quality: 'Prepress (~300+ DPI)',
    compression: 'Ultra-Light (~10-25%)',
    isProOnly: true,
  },
]

export function usePdfCompressor() {
  const selectedFile = ref<File | null>(null)
  const compressionLevel = ref<CompressionLevel>('MEDIUM')
  const isProcessing = ref(false)
  const uploadProgress = ref(0)
  const processingStep = ref<'idle' | 'uploading' | 'compressing' | 'finalizing' | 'completed' | 'error'>('idle')
  const jobResult = ref<JobResponse | null>(null)
  const errorMessage = ref<string | null>(null)

  const formattedOriginalSize = computed(() => {
    if (!selectedFile.value) return '0 B'
    return formatFileSize(selectedFile.value.size)
  })

  const formattedCompressedSize = computed(() => {
    if (!jobResult.value) return '0 B'
    return formatFileSize(jobResult.value.compressed_size)
  })

  const formattedSavedBytes = computed(() => {
    if (!jobResult.value) return '0 B'
    return formatFileSize(jobResult.value.saved_bytes)
  })

  function formatFileSize(bytes: number): string {
    if (bytes === 0) return '0 B'
    const k = 1024
    const sizes = ['B', 'KB', 'MB', 'GB']
    const i = Math.floor(Math.log(bytes) / Math.log(k))
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
  }

  function handleFileSelected(file: File) {
    errorMessage.value = null
    if (!file.name.toLowerCase().endsWith('.pdf') && file.type !== 'application/pdf') {
      errorMessage.value = 'Please select a valid PDF document (.pdf)'
      return
    }

    if (file.size > 50 * 1024 * 1024) {
      errorMessage.value = 'File size exceeds maximum 50 MB limit.'
      return
    }

    selectedFile.value = file
    jobResult.value = null
    processingStep.value = 'idle'
  }

  function clearFile() {
    selectedFile.value = null
    jobResult.value = null
    errorMessage.value = null
    processingStep.value = 'idle'
    uploadProgress.value = 0
  }

  async function startCompression() {
    if (!selectedFile.value) {
      errorMessage.value = 'Please select a PDF file first.'
      return
    }

    isProcessing.value = true
    errorMessage.value = null
    uploadProgress.value = 0
    processingStep.value = 'uploading'

    try {
      // Simulate progress progression for smooth UX
      const result = await pdfApi.compressPDF(
        selectedFile.value,
        compressionLevel.value,
        (progress) => {
          uploadProgress.value = progress
          if (progress >= 100) {
            processingStep.value = 'compressing'
          }
        }
      )

      processingStep.value = 'finalizing'
      await new Promise((resolve) => setTimeout(resolve, 500))

      jobResult.value = result
      processingStep.value = 'completed'

      // Record daily usage
      const { recordCompressionUsage } = useAuth()
      recordCompressionUsage()
    } catch (err: any) {
      processingStep.value = 'error'
      if (err.response?.data?.errors?.length) {
        errorMessage.value = err.response.data.errors.join(', ')
      } else if (err.response?.data?.message) {
        errorMessage.value = err.response.data.message
      } else {
        errorMessage.value = err.message || 'An unexpected error occurred during compression'
      }
    } finally {
      isProcessing.value = false
    }
  }

  function downloadResult() {
    if (!jobResult.value) return
    const url = pdfApi.getDownloadUrl(jobResult.value.id)
    const link = document.createElement('a')
    link.href = url
    link.setAttribute('download', `compressed_${jobResult.value.original_filename}`)
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }

  function reset() {
    clearFile()
  }

  return {
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
    formatFileSize,
    handleFileSelected,
    clearFile,
    startCompression,
    downloadResult,
    reset,
  }
}
