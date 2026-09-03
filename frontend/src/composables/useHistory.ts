import { ref } from 'vue'

export interface HistoryItem {
  id: string
  tool: 'compress' | 'merge' | 'split'
  filename: string
  originalSize: number
  resultSize?: number
  savedPercentage?: number
  pageCount?: number
  fileCount?: number
  downloadUrl: string
  timestamp: string
}

const HISTORY_STORAGE_KEY = 'hyperpdf_conversion_history'

function loadHistory(): HistoryItem[] {
  try {
    const raw = localStorage.getItem(HISTORY_STORAGE_KEY)
    if (raw) {
      return JSON.parse(raw)
    }
  } catch (err) {
    console.error('Failed to load conversion history', err)
  }
  return []
}

const historyItems = ref<HistoryItem[]>(loadHistory())

function saveHistory() {
  try {
    localStorage.setItem(HISTORY_STORAGE_KEY, JSON.stringify(historyItems.value))
  } catch (err) {
    console.error('Failed to save conversion history', err)
  }
}

export function useHistory() {
  function addHistoryItem(item: Omit<HistoryItem, 'id' | 'timestamp'>) {
    const newItem: HistoryItem = {
      ...item,
      id: 'hist_' + Math.random().toString(36).substring(2, 9),
      timestamp: new Date().toISOString(),
    }
    // Prepend new item and keep max 50 items
    historyItems.value = [newItem, ...historyItems.value].slice(0, 50)
    saveHistory()
  }

  function removeHistoryItem(id: string) {
    historyItems.value = historyItems.value.filter((i) => i.id !== id)
    saveHistory()
  }

  function clearHistory() {
    historyItems.value = []
    saveHistory()
  }

  return {
    historyItems,
    addHistoryItem,
    removeHistoryItem,
    clearHistory,
  }
}
