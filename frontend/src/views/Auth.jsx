import React from "react";
import {Box, Typography} from "@mui/material";

import Login from "../components/Auth/Login.jsx";
import Register from "../components/Auth/Register.jsx";

const Auth = () => {
    const [internalComponentName, setInternalComponentName] = React.useState("login")
    const spanStyle = {
        color: "blue",
        textDecoration: "underline",
        cursor: "pointer"
    }
    const spanOnClick = (componentName) => {
        setInternalComponentName(componentName)
    }
    return (
        <Box>
            {
                internalComponentName === "login" ? <Login />
                :internalComponentName === "register" ? <Register /> : null
            }
            {/*<Box>*/}
            {
                internalComponentName === "login" ? <Typography variant={"h4"} paragraph >If you don't have an account <span style={spanStyle} onClick={() => spanOnClick("register")} >register!</span></Typography>
                    :internalComponentName === "register" ? <Typography variant={"h4"} paragraph >If you already have an account <span style={spanStyle} onClick={() => spanOnClick("login")}>login!</span></Typography>
                    :null
            }
            {/*</Box>*/}
        </Box>
    )
}

export default Auth