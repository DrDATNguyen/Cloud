import React from 'react'
import Login from './Login'
import Register from './Register';
import ForgetPassword from './ForgetPassword';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
function App() {
  return (
    
    <div>
      <ForgetPassword />
    </div>
  //   <Router>
  //   <Routes>
  //     <Route path="/" element={<Login />} />
  //     <Route path="/register" element={<Register />} />
  //   </Routes>
  // </Router>
  )
}

export default App
