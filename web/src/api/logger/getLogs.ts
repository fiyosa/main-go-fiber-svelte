import { useQuery, type CreateQueryOptions } from "../../lib/tanstackUtil"
import { instance } from "../../lib/axiosLib"

export const getLogs = (options?: Partial<CreateQueryOptions<string[]>>) =>
  useQuery<string[]>(() => ({
    ...options,
    queryKey: ["guest", "logs"],
    queryFn: async () => {
      const res = await instance.get("/guest/logs")
      return res.data.data as string[]
    },
  }))
