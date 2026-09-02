import axios, { AxiosProgressEvent } from 'axios'
import type { APIResponse, JobResponse, CompressionLevel } from '../types'

const apiClient = axios.create({
  baseURL: '/',
  headers: {
    'Accept': 'application/json',
  },
})

export const pdfApi = {
  async compressPDF(
    file: File,
    level: CompressionLevel,
    onProgress?: (progress: number) => void
  ): Promise<JobResponse> {
    const formData = new FormData()
    formData.append('file', file)
    formData.append('compression_level', level)

    const response = await apiClient.post<APIResponse<JobResponse>>(
      '/api/v1/pdf/compress',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        onUploadProgress: (progressEvent: AxiosProgressEvent) => {
          if (progressEvent.total && onProgress) {
            const percentCompleted = Math.round(
              (progressEvent.loaded * 100) / progressEvent.total
            )
            onProgress(percentCompleted)
          }
        },
      }
    )

    if (!response.data.success || !response.data.data) {
      throw new Error(response.data.message || 'Failed to compress PDF')
    }

    return response.data.data
  },

  async getJob(id: string): Promise<JobResponse> {
    const response = await apiClient.get<APIResponse<JobResponse>>(`/api/v1/pdf/jobs/${id}`)
    if (!response.data.success || !response.data.data) {
      throw new Error(response.data.message || 'Failed to fetch job')
    }
    return response.data.data
  },

  getDownloadUrl(id: string): string {
    return `/api/v1/pdf/jobs/${id}/download`
  },

  async deleteJob(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/pdf/jobs/${id}`)
  },

  async checkHealth(): Promise<any> {
    const response = await apiClient.get('/health')
    return response.data
  }
}
