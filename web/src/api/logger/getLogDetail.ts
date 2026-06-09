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

interface IProps {
  param: {
    file_name: string
  }
  query: {
    search: string
    page: string
    limit: string
    levels: string
  }
}

export const getLogDetail = (getProps: () => IProps, options?: Partial<CreateQueryOptions<LogDetailResponse>>) =>
  useQuery<LogDetailResponse>(() => {
    const props = getProps()
    return {
      ...options,
      queryKey: ['log', 'detail', props.param.file_name, props.query.search, props.query.page, props.query.levels],
      queryFn: async () => {
        const res = await instance.get(`/log/${encodeURIComponent(props.param.file_name)}${createQueryStr(props)}`)
        return {
          message: res.data.message ?? '',
          data: res.data.data as LogEntry[],
          meta: res.data.meta as PaginationMeta,
        }
      },
      enabled: !!props.param.file_name,
    }
  })
