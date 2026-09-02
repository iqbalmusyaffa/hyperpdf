export type CompressionLevel = 'LOW' | 'MEDIUM' | 'HIGH'
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
}
