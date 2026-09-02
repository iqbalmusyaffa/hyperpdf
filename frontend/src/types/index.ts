export type CompressionLevel =
  | 'ULTRA_EXTREME'
  | 'EXTREME'
  | 'HIGH'
  | 'RECOMMENDED'
  | 'MEDIUM'
  | 'HIGH_FIDELITY'
  | 'LOW'
  | 'STUDIO_MASTER'
  | 'CUSTOM_TARGET'

export type JobStatus = 'PENDING' | 'PROCESSING' | 'COMPLETED' | 'FAILED'

export interface JobResponse {
  id: string
  original_filename: string
  original_size: number
  compressed_size: number
  saved_bytes: number
  compression_percentage: number
  compression_level: CompressionLevel
  status: JobStatus
  error_message?: string
  created_at: string
  completed_at?: string
}

export interface APIResponse<T = any> {
  success: boolean
  message: string
  data: T
}

export interface APIErrorResponse {
  success: boolean
  message: string
  errors?: string[]
}

export interface LevelOption {
  id: CompressionLevel
  name: string
  badge: string
  description: string
  quality: string
  compression: string
  recommended?: boolean
  isProOnly?: boolean
}

export type PDFTool = 'compress' | 'merge' | 'split'

export interface MergeResponse {
  id: string
  merged_filename: string
  file_count: number
  total_size: number
  download_url: string
  created_at: string
}

export interface SplitResponse {
  id: string
  original_filename: string
  split_mode: string
  page_ranges?: string
  generated_count: number
  is_zip_archive: boolean
  download_filename: string
  download_url: string
  created_at: string
}
