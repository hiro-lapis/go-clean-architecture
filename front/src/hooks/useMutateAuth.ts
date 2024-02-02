import axios from 'axios'
// router
import { useNavigate } from 'react-router-dom'
// update async state
import { useMutation } from '@tanstack/react-query'
import useStore from '../store'
import { Credential } from '../types'
import { useError } from './useError'

export const useMutateAuth = () => {
  const navigate = useNavigate()
  const resetEditedTask = useStore((state) => state.resetEditedTask)
  const { errorHandling } = useError()

  /**
   * 以下のオプション関数を定義
   * useMutation({
   *   mutationFn　// required: 非同期関数
   *   onMutate,   // optional:mutationFnが実行される前に実行される関数
   *   onSuccess,　// optional:mutationFnが成功した時に実行される関数
   *   onError,　 // optional:mutationFnが失敗した時に実行される関数
   *   onSettled　// optional:mutationFnが成功・失敗した時に実行される関数
   * })
   * @link https://tanstack.com/query/v5/docs/react/reference/useMutation
   *
   * 非同期は使わず、ハンドリング関数を利用するでOK
   * @link https://zenn.dev/taisei_13046/books/133e9995b6aadf/viewer/257b1a#mutate%E3%81%A8mutateasync%E3%81%AE%E4%BD%BF%E3%81%84%E5%88%86%E3%81%91
   *
   * nuxt apollo client composables を使う時の関数を自前で実装するイメージ
   */
  const loginMutation = useMutation({
    mutationFn: (user: Credential) => {
      return axios.post<{name: string}>(`${process.env.REACT_APP_API_URL}/login`, user)
    },
    onSuccess: () => navigate('/todo'),
    onError: (err: any) => {
      if (err.response.data.message) {
        errorHandling(err.response.data.message)
      } else {
        errorHandling(err.response.data)
      }
    },
  })
  const registerMutation = useMutation({
    mutationFn: (user: Credential) => {
      return axios.post<{name: string}>(`${process.env.REACT_APP_API_URL}/signup`, user)
    },
    onSuccess: () => navigate('/todo'),
    onError: (err: any) => {
      if (err.response.data.message) {
        errorHandling(err.response.data.message)
      } else {
        errorHandling(err.response.data)
      }
    },
  })
  const logoutMutation = useMutation({
    mutationFn: () => {
      return axios.post<{name: string}>(`${process.env.REACT_APP_API_URL}/logout`)
    },
    onSuccess: () => {
      resetEditedTask()
      navigate('/todo')
    },
    onError: (err: any) => {
      if (err.response.data.message) {
        errorHandling(err.response.data.message)
      } else {
        errorHandling(err.response.data)
      }
    },
  })

  return {
    loginMutation,
    registerMutation,
    logoutMutation,
    resetEditedTask,
  }
}
