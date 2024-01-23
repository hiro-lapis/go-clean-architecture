import { useEffect } from 'react';
// import logo from './logo.svg';
// import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Auth } from './components/Auth';
import { Todo } from './components/Todo';
import axios from 'axios';
import { CsrfToken } from './types/index';


function App() {
  // https://ja.react.dev/reference/react/useEffect
  useEffect(() => {
    axios.defaults.withCredentials = true

    const getCsrfToken = async () => {
      // 環境変数をもとにサーバドメインからCSRFトークンを取得
      // 戻り値の型はCsrfToken
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      // api clientのヘッダーにCSRFトークンを設定
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    // 実行
    getCsrfToken()
  // 値の変更にフックさせたい時は、変数を下記の配列に入れる
  }, [])

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
