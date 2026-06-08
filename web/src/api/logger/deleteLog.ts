import { instance } from "../../lib/axiosLib"

export const deleteLog = async (filename: string) => {
  const res = await instance.delete(`/log/${encodeURIComponent(filename)}`)
  return res.data
}
