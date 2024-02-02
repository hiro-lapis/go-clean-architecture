import { ArrowRightCircleIcon, ShieldCheckIcon } from "@heroicons/react/24/solid"
import {
  useMutateAuth,
} from '../hooks/useMutateAuth'

export const Todo = () => {
  const { logoutMutation } = useMutateAuth()
  const logout = async () => {
    // Similar to mutate but returns a promise which can be awaited.
    // https://tanstack.com/query/v5/docs/framework/react/reference/useMutation
    await logoutMutation.mutateAsync()
  }
  return (
    <div>
      <ArrowRightCircleIcon
        onClick={ logout }
        className="h-8 w-8 mr-2 text-blue-500 cursor-pointer"
      />
    </div>
  )
}
