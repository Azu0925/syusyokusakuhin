import React, { useState } from 'react'
import Footer from '../components/footers/Footer'
import Header from '../components/headers/Header'

const SignUp = (): JSX.Element => {
  const [makeState, setMakeState] = useState()

  return (
    <>
      <Header />
      <main>
        <h2>ようこそ</h2>

      </main>
      <Footer />
    </>
  )
}

export default SignUp
