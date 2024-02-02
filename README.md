# go-clean-architecture

## structure

```
├── README.md
├── db
├── docker-compose.yml
├── go.mod
├── db // DB接続
├── migraion // マイグレーション実行
└── model // model定義
└── repository // リポジトリ
└── usecase // リポジトリに依存して処理を行うユースケース

```

## library

- FW: [echo](https://github.com/labstack/echo)
- Token management: [echo-jwt](https://github.com/labstack/echo-jwt)
- ORM: [gorm](https://github.com/go-gorm/gorm)
- Validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)



## implement flow

1. repository  

1-1. IRepository 作成  
1-2. IRepository に処理単位で関数定義  
1-3. 実装するrepository 作成   
1-4. 実装repositoryのconstrcutor関数作成  
1-5. 実装repositoryにメソッドを作成、interfaceを満たす実装にする  

* 取得系関数の場合、取得する情報は引数にとった変数に書き込み、戻り値は基本的にerrorのみ  
* repositoryはgormインスタンスのポインタのみをプロパティとして持つ
* 取得系関数でも戻り値は基本的にerrorのみ  

2. usecase  

2-1. IUsecase作成
2-2. IUsecase に処理単位で関数定義
2-3. 実装するusecase 作成
2-4. 実装usecaseのconstrcutor関数作成

3. controller  

3-1. IController作成
3-2. IControllerに処理単位の関数定義
3-3. 実装するcontroller 作成
3-4. 実装controllerのconstrcutor関数作成
3-5. router.goにてIControllerを引数にとりrouting と関数の紐付け
3-6. main.goにてusecaseを渡してのcontroller instanciation

4. validator

4-1. IValidator 作成  
4-2. validator 作成  
4-3. validator のconstrcutor関数作成  
4-4. validator でバリデーション関数実装  
4-5. main.goにてvalidator instancication, constructro引数に渡す  

* 4-2において、基本的に依存は不要。DB利用時のみgormを入れる  

# frontend

## directory

```
└── src
    ├── components
    ├── hooks
    ├── store
    └── types
```

## Setup
1. this project using pnpm, then run below code

```
pnpm create create-react-app@latest front --template typescript
```
*create-react-app* is officially supported way to create single-page React applications. It offers a modern build setup with no configuration.  
if you are interested in options, checck [here](https://create-react-app.dev/docs/getting-started#selecting-a-template)!  

2. reinstall pnpm

pnpm is ignored, installing with generate package-lock.json.  
then, delete package-lock.json and node_modules.  

and, reinstall.

```
pnpm i
```

3. test dev server

```
pnpm start
```

if you encounter error, `Property 'toBeInTheDocument' does not exist on type 'JestMatchers<HTMLElement>'.`  
then, try below.

```
pnpm i @types/testing-library__jest-dom
```
＊install is required above [^6.2.0](https://github.com/testing-library/jest-dom/issues/442#issuecomment-1888145410)

4. add dependencies 

also recommend 
```
pnpm i @tanstack/react-query
pnpm i @tanstack/react-query-devtools
pnpm i zustand
pnpm i @heroicons/react
pnpm i react-router-dom
pnpm i axios
pnpm i -D tailwindcss
npx tailwindcss init
```
tailwindcss required set up. do [this](https://tailwindcss.com/docs/guides/create-react-app)


### library

- [tanstack react query](https://tanstack.com/query/latest/docs/react/overview): data fetching library, also can caching, synchronizing and mutation, updating server state

resource:[TanStack Query　〜プロダクトで採用するための勘所〜](https://zenn.dev/taisei_13046/books/133e9995b6aadf/viewer/257b1a)

## FAQ

Q. How solve `Network Error AxiosError: Network Error` when `pnpm start`?
A. run backend server.
