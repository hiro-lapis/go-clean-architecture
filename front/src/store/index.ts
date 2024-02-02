import { create } from "zustand"

// zustand 状態管理

type EditedTask = {
    id: number
    title: string
    description: string
}
// stateで管理する値と更新関数のinterface定義
type State = {
    editedTask: EditedTask
    updateEditedTask: (payload: EditedTask) => void
    resetEditedTask: () => void
}

// zustandのcreate関数を使ってstateを作成
// nuxt同様、関心ごとを管理する状態変数と更新関数を定義
const useStore = create<State>((set) => ({
    editedTask: { id: 0, title: '', description: ''},
    updateEditedTask: (payload) =>
    set({
        editedTask: payload,
    }),
    resetEditedTask: () => set({ editedTask: { id: 0, title: '', description: ''}})
}))

export default useStore
