import { useState, FormEvent } from 'react'
import { CheckBadgeIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'

// vueで言うところのscript,templateがまとめられた関数
export const Auth = () => {
  // useState: refと同じ状態管理変数
  const [email, setEmail] = useState('')
  const [ password, setPassword] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const { loginMutation, registerMutation, } = useMutateAuth()

  const submit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (isLogin) {
      // mutateFnの実行
      loginMutation.mutate({email, password})
    } else {
      await registerMutation
        .mutateAsync({email, password})
        // 登録成功したらログインAPI実行
        .then(() => loginMutation.mutate({email, password}))
    }
  }
  return (
    <div>Auth</div>
  )
}
