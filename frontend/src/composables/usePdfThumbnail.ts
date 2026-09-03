import { ref } from 'vue'

let pdfjsLoaded = false
let pdfjsPromise: Promise<any> | null = null

function loadPdfJs(): Promise<any> {
  if (pdfjsLoaded && (window as any).pdfjsLib) {
    return Promise.resolve((window as any).pdfjsLib)
  }
  if (pdfjsPromise) {
    return pdfjsPromise
  }

  pdfjsPromise = new Promise((resolve) => {
    if ((window as any).pdfjsLib) {
      pdfjsLoaded = true
      resolve((window as any).pdfjsLib)
      return
    }

    const script = document.createElement('script')
    script.src = 'https://cdnjs.cloudflare.com/ajax/libs/pdf.js/3.11.174/pdf.min.js'
    script.async = true
    script.onload = () => {
      const pdfjsLib = (window as any).pdfjsLib
      if (pdfjsLib) {
        pdfjsLib.GlobalWorkerOptions.workerSrc =
          'https://cdnjs.cloudflare.com/ajax/libs/pdf.js/3.11.174/pdf.worker.min.js'
        pdfjsLoaded = true
        resolve(pdfjsLib)
      } else {
        resolve(null)
      }
    }
    script.onerror = () => resolve(null)
    document.head.appendChild(script)
  })

  return pdfjsPromise
}

export function usePdfThumbnail() {
  const isGenerating = ref(false)

  async function generateThumbnail(
    file: File,
    scale: number = 0.5
  ): Promise<{ thumbnail: string | null; pageCount: number }> {
    isGenerating.value = true
    try {
      const pdfjs = await loadPdfJs()
      if (!pdfjs) {
        return { thumbnail: null, pageCount: 1 }
      }

      const arrayBuffer = await file.arrayBuffer()
      const loadingTask = pdfjs.getDocument({ data: arrayBuffer })
      const pdf = await loadingTask.promise
      const pageCount = pdf.numPages || 1

      // Render page 1
      const page = await pdf.getPage(1)
      const viewport = page.getViewport({ scale })

      const canvas = document.createElement('canvas')
      const context = canvas.getContext('2d')
      if (!context) {
        return { thumbnail: null, pageCount }
      }

      canvas.height = viewport.height
      canvas.width = viewport.width

      const renderContext = {
        canvasContext: context,
        viewport: viewport,
      }

      await page.render(renderContext).promise
      const thumbnail = canvas.toDataURL('image/jpeg', 0.85)

      return { thumbnail, pageCount }
    } catch (err) {
      console.warn('Could not generate PDF thumbnail, using fallback icon', err)
      return { thumbnail: null, pageCount: 1 }
    } finally {
      isGenerating.value = false
    }
  }

  return {
    isGenerating,
    generateThumbnail,
  }
}
