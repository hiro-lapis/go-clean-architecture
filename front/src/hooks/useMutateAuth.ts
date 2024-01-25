import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import useStore from '../store'
import { Credential } from '../types'
import { useError } from '../hooks/useError'

export const useMutateAuth = () => {
  const navigate = useNavigate()
  const resetEditedTask = useStore((state) => state.resetEditedTask)
  const { errorHandling } = useError()
  const loginMutation = useMutation(
    async (user: Credential) =>
      await axios.post<{name: string}>(`${process.env.REACT_APP_API_URL}/login`, user),
    {
      onSuccess: () => {
        navigate('/todo')
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          errorHandling(err.response.data.message)
        } else {
          errorHandling(err.response.data)
        }
      },
    }
  )
  const registerMutation = useMutation(
    async (user: Credential) =>
      await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user),
    {
      // onSuccess: () => {
      //   navigate('/todo')
      // },
      onError: (err: any) => {
        if (err.response.data.message) {
          errorHandling(err.response.data.message)
        } else {
          errorHandling(err.response.data)
        }
      },
    }
  )
  const logoutMutation = useMutation(
    async (user: Credential) =>
      await axios.post(`${process.env.REACT_APP_API_URL}/logout`, user),
    {
      onSuccess: () => {
        resetEditedTask()
        navigate('/')
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          errorHandling(err.response.data.message)
        } else {
          errorHandling(err.response.data)
        }
      },
    }
  )

  return {
    loginMutation,
    registerMutation,
    logoutMutation,
    resetEditedTask,
  }
}
