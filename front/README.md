## エラー
### user ストアにuserId が undefined
- Login 画面で user　ストアに user をセットし, Home 画面で user の userId を取り出す処理の流れにおいて, userId == undefined になる
    - Login 画面でuserId をセットできていない可能性
        - Login した瞬間に userStore 側の set が正しい値でよばれているかチェック
            - setUser 内の log仕込んだ結果 userid, name が undefined
                - backend から id,name を返していると思ったら Id, Name でオブジェクトをレスポンスしていたことが原因
    - Home　画面でuserId の取り出しに失敗している可能性












# Nuxt Minimal Starter

Look at the [Nuxt documentation](https://nuxt.com/docs/getting-started/introduction) to learn more.

## Setup

Make sure to install dependencies:

```bash
# npm
npm install

# pnpm
pnpm install

# yarn
yarn install

# bun
bun install
```

## Development Server

Start the development server on `http://localhost:3000`:

```bash
# npm
npm run dev

# pnpm
pnpm dev

# yarn
yarn dev

# bun
bun run dev
```

## Production

Build the application for production:

```bash
# npm
npm run build

# pnpm
pnpm build

# yarn
yarn build

# bun
bun run build
```

Locally preview production build:

```bash
# npm
npm run preview

# pnpm
pnpm preview

# yarn
yarn preview

# bun
bun run preview
```

Check out the [deployment documentation](https://nuxt.com/docs/getting-started/deployment) for more information.
