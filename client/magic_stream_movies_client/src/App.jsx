
import './App.css'
import Home from './components/home/Home'
import Header from './components/header/Header'
import {Route, Routes, useNavigate} from 'react-router-dom'
import Login from './components/login/Login'
import Register from './components/register/Register'

function App() {
    
    return (
        <>
            <Header />
            <Routes>
                <Route path="/" element={<Home />}></Route>
                <Route path="/register" element={<Register />}></Route>
                <Route path="/login" element={<Login />}></Route>
            </Routes>
        </>
    )
}

export default App 