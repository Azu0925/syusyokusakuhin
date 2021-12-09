import React from 'react'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Home from '../pages/Home'
import NotFound from '../pages/NotFound'
import Setting from '../pages/Setting'
import SignIn from '../pages/SignIn'
import SignUp from '../pages/SignUp'
import TimeLine from '../pages/TimeLine'
import User from '../pages/User'

const Router = (): JSX.Element => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/home" element={<TimeLine />} />
        <Route path="/user/:id" element={<User />} />
        <Route path="/user/:id/settings" element={<Setting />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  )
}
export default Router
