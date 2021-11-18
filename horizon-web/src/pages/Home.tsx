import React from 'react'
import { Link } from 'react-router-dom'

const Home = (): JSX.Element => {
  return (
    <main>
      <h1>不思議な世界、素敵な世界</h1>
      <section>
        <p>サインアップしてはじめよう</p>
        <Link to="/signup">新規作成</Link>
        <Link to="/signin">すでにアカウントを持っている方</Link>
      </section>
    </main>
  )
}

export default Home
