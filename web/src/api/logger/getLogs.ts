import { useQuery, type CreateQueryOptions } from "../../lib/tanstackUtil"
import { instance } from "../../lib/axiosLib"

export type LogFileItem = {
  name: string
  size: number
}

export interface IProps {
  param: Record<string, never>
  query: Record<string, never>
}

export const getLogs = (options?: Partial<CreateQueryOptions<LogFileItem[]>>) =>
  useQuery<LogFileItem[]>(() => ({
    ...options,
    queryKey: ["log"],
    queryFn: async () => {
      const res = await instance.get("/log")
      return res.data.data as LogFileItem[]
    },
  }))
