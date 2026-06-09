import { useMutation } from '../../lib/tanstackUtil'
import { instance } from '../../lib/axiosLib'
import type { AxiosResponse } from 'axios'

export interface IProps {
  param: {
    file_name: string
  }
}

export const deleteLog = () =>
  useMutation<AxiosResponse, Error, IProps>(() => ({
    mutationFn: (props: IProps) => instance.delete(`/log/${encodeURIComponent(props.param.file_name)}`),
  }))
