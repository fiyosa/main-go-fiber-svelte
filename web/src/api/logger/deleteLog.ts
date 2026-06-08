import { instance } from "../../lib/axiosLib"

export interface IProps {
  param: {
    file_name: string
  }
  query: Record<string, never>
}

export const deleteLog = (props: IProps) =>
  instance.delete(`/log/${encodeURIComponent(props.param.file_name)}`)
