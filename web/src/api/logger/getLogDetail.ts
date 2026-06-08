import { useQuery, type CreateQueryOptions } from '../../lib/tanstackUtil'
import { createQueryStr, instance } from '../../lib/axiosLib'

type LogEntry = {
  level: string
  time: string
  message: string
}

type PaginationMeta = {
  total: number
  page: number
  limit: number
}

type LogDetailResponse = {
  message: string
  data: LogEntry[]
  meta: PaginationMeta
}

export type { LogEntry, LogDetailResponse, PaginationMeta }

interface IProps {
  param: {
    file_name: string
  }
  query: {
    search: string
    page: string
    limit: string
  }
}

export const getLogDetail = (
  getProps: () => IProps,
  getLevels: () => string,
  options?: Partial<CreateQueryOptions<LogDetailResponse>>
) =>
  useQuery<LogDetailResponse>(() => {
    const props = getProps()
    const levels = getLevels()
    return {
      ...options,
      queryKey: ['log', 'detail', props.param.file_name, props.query.search, props.query.page, levels],
      queryFn: async () => {
        const q = new URLSearchParams()
        if (props.query.search) q.set('search', props.query.search)
        if (props.query.page) q.set('page', props.query.page)
        if (props.query.limit) q.set('limit', props.query.limit)
        if (levels) q.set('levels', levels)
        const qs = q.toString()
        const res = await instance.get(`/log/${props.param.file_name}${qs ? '?' + qs : ''}`)
        return { message: res.data.message ?? '', data: res.data.data as LogEntry[], meta: res.data.meta as PaginationMeta }
      },
      enabled: !!props.param.file_name,
    }
  })
