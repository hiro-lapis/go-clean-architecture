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
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
    <div className="flex items-center">
      <CheckBadgeIcon className="h-8 w-8 mr-2 text-blue-500" />
      <span className="text-center text-3xl font-extrabold">
        Todo app by React/Go(Echo)
      </span>
    </div>
    <h2 className="my-6">{isLogin ? 'Login' : 'Create a new account'}</h2>
    <form onSubmit={submit}>
      <div>
        <input
          className="mb-3 px-3 text-sm py-2 border border-gray-300"
          name="email"
          type="email"
          autoFocus
          placeholder="Email address"
          onChange={(e) => setEmail(e.target.value)}
          value={email}
        />
      </div>
      <div>
        <input
          className="mb-3 px-3 text-sm py-2 border border-gray-300"
          name="password"
          type="password"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
          value={pw}
        />
      </div>
      <div className="flex justify-center my-2">
        <button
          className="disabled:opacity-40 py-2 px-4 rounded text-white bg-indigo-600"
          disabled={!email || !password}
          type="submit"
        >
          {isLogin ? 'Login' : 'Sign Up'}
        </button>
      </div>
    </form>
    <ArrowPathIcon
      onClick={() => setIsLogin(!isLogin)}
      className="h-6 w-6 my-2 text-blue-500 cursor-pointer"
    />
  </div>
  )
}
