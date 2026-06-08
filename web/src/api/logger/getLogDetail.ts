import { useQuery, type CreateQueryOptions } from "../../lib/tanstackUtil"
import { instance } from "../../lib/axiosLib"

type LogEntry = {
  level: string
  time: string
  message: string
}

type LogDetailResponse = {
  name: string
  total: number
  logs: LogEntry[]
  page: number
  limit: number
}

export type { LogEntry, LogDetailResponse }

type DetailParams = {
  filename: () => string | null
  levels?: () => string
  search?: () => string
  page?: () => number
}

export const getLogDetail = (
  params: DetailParams,
  options?: Partial<CreateQueryOptions<LogDetailResponse>>,
) =>
  useQuery<LogDetailResponse>(() => ({
    queryKey: ["log", params.filename(), params.levels?.(), params.search?.(), params.page?.()],
    queryFn: async () => {
      const fn = params.filename()
      if (!fn) throw new Error("No filename")
      const q = new URLSearchParams()
      if (params.levels?.()) q.set("levels", params.levels()!)
      if (params.search?.()) q.set("search", params.search()!)
      if (params.page) q.set("page", String(params.page()))
      q.set("limit", "50")
      const qs = q.toString()
      const res = await instance.get(`/log/${encodeURIComponent(fn)}${qs ? "?" + qs : ""}`)
      return res.data.data as LogDetailResponse
    },
    enabled: params.filename() !== null,
    ...options,
  }))
