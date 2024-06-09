import React, {createContext} from "react";
import PropTypes from "prop-types";

export const AuthContext = createContext({
    auth: {
        logged_in: false,
        username: "",
        token: ""
    },
    login: (token, uname) => { return token + uname },
    logout: () => {}
})

const UserContextProvider= ({ children }) => {
    let localAuthData = localStorage.getItem("auth")
    const [auth, setAuth] = React.useState((localAuthData && true) ? JSON.parse(localAuthData) : {
        logged_in: false,
        username: "",
        token: ""
    })

    function login(token, uname) {
        let authData = {
            logged_in: true,
            username: uname,
            token: token
        }
        localStorage.setItem("auth", JSON.stringify(authData))
        setAuth(authData)
    }
    const logout = () => {
        setAuth(prevState => {
            return {
                logged_in: false,
                ...prevState
            }
        })
    }

    const contextValues = {
        auth,
        login,
        logout
    }
    return <AuthContext.Provider value={contextValues}>
        {children}
    </AuthContext.Provider>
}

UserContextProvider.propTypes = {
    children: PropTypes.node.isRequired
}

export default UserContextProvider