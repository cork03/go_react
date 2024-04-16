import { Route, Routes } from "react-router-dom"
import { Register } from "./pages/Register"
import { Home } from "./pages/Home"


export const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/register" element={<Register />} />
    </Routes>
  )
}
