import React from 'react'
import { useNavigate } from 'react-router-dom'
import {AuthContext} from "../../context/AuthContext.jsx";

const UserGuard = ({ children, redirection = "/" }) => {
    const { auth } = React.useContext(AuthContext);
    const navigate = useNavigate();
    React.useEffect(() => {
        const validate = () => {
            if (!auth.logged_in) {
                navigate("/")
            }
        }
        validate();

    }, [])
    return (
        <>
            {children}
        </>
    )
}

export default UserGuard