import { BrowserRouter, Routes, Route } from 'react-router-dom'
import UserContextProvider from "../context/AuthContext.jsx";

import Auth from "../views/Auth.jsx";
import Dashboard from "../views/Dashboard.jsx";

import UserGuard from "./guards/UserGuard.jsx";

const AllRoutes = () => {
    return (
        <BrowserRouter>
            <UserContextProvider>
                <Routes>
                    <Route path="/" element={<Auth/>} />
                    <Route path="/dashboard" element={<UserGuard><Dashboard /></UserGuard>} />
                </Routes>
            </UserContextProvider>
        </BrowserRouter>
    )
}

export default AllRoutes