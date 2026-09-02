import { ref, computed } from 'vue'
import { pdfApi } from '../api/pdfApi'
import type { CompressionLevel, JobResponse, LevelOption } from '../types'

export const levelOptions: LevelOption[] = [
  {
    id: 'LOW',
    name: 'Low Compression',
    badge: 'Best Quality',
    description: 'Minimal reduction, preserves highest image and vector fidelity.',
    quality: 'High (~300 DPI)',
    compression: 'Light (~20-40%)',
  },
  {
    id: 'MEDIUM',
    name: 'Medium Compression',
    badge: 'Recommended',
    description: 'Smart compression balancing sharp text/images with great reduction.',
    quality: 'Medium (~150 DPI)',
    compression: 'Balanced (~50-70%)',
    recommended: true,
  },
  {
    id: 'HIGH',
    name: 'Extreme Compression',
    badge: 'Smallest Size',
    description: 'Maximum compression to reach the smallest possible file footprint.',
    quality: 'Standard (~72 DPI)',
    compression: 'Maximum (~70-90%)',
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
