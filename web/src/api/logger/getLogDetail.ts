import { useQuery, type CreateQueryOptions } from '../../lib/tanstackUtil'
import { instance } from '../../lib/axiosLib'

type LogEntry = {
  level: string
  time: string
  message: string
}

type LogDetail = {
  name: string
  total: number
  logs: LogEntry[]
}

export const getLogDetail = (
  filename: () => string | null,
  options?: Partial<CreateQueryOptions<LogDetail>>,
) =>
  useQuery<LogDetail>(() => ({
    queryKey: ['guest', 'logs', filename()],
    queryFn: async () => {
      const res = await instance.get(`/guest/logs/${encodeURIComponent(filename()!)}`)
      return res.data.data as LogDetail
    },
    enabled: filename() !== null,
    ...options,
  }))
